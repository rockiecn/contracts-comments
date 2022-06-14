# 合约设计文档

## 通用

**代为调用**

合约中 external 类的涉及到交易的函数都做了可由其他人代为调用的功能。比如账户 A 想要调用 register 函数获得角色索引号，可以给账户 B 一个相关的签名信息，由账户 B 帮 A 调用 register 函数，也即发起交易。

## 角色注册

角色分类：User、Keeper、Provider
角色类型：`uint8 roleType; // 0 ordinary account; 1 user; 2 provider; 3 keeper`
admin 创建 group，并随机将一些 Provider 加入该 group 中，创建 group 时设置一个`level`等级，当该 group 中 Provider 数量达到`level`值时，该 group 才生效（由`isActive`参数表示）。每个 group 代表一个服务组，每个角色只能存在于一个 group 中。Keeper 和 User 可自己选择加入哪一个 group。每个 group 都有一个 FileSys 合约地址，也即当 admin 创建 group 时，需要同时部署一个 FileSys 合约。
用户调用 Role 合约中的 register 函数获得账户序列号`index`（从 1 开始，依次累加），账户序列号用于组成 group;
通过调用 Role 合约中的 registerUser、registerKeeper、registerProvider 函数依次注册成为 User、Keeper、Provider；
Keeper、Provider 进行注册前，需要通过调用 PledgePool 合约中的 pledge 函数，往 PledgePool 合约中质押指定金额（即角色申请金），质押额足够时，角色才生效。指定金额由 admin 设置在 Role 合约中。
角色`isBanned`参数：表示该账户是否被 admin 禁止，用于惩罚机制。

## 代币支持

RToken 合约用于存放 MEMO 支持的代币信息。包含代币地址、代币序列号、该币种是否被禁止。
MEMO 支持的主代币由 MEMO 部署，合约代码在 contracts/token 目录下。也支持其他 erc20 协议的子代币。
支持方式：admin 部署 Role 合约时，将主代币地址添加到 RToken 合约中，主代币序列号为 0。其他子代币由 admin 调用 Role 合约中的 registerToken 函数，添加支持的代币。

## 质押机制

提供一个质押池，由 PledgePool 合约实现。
用户往质押池中质押主代币，可以根据质押量获得增发的主代币以及质押池中支持的其他代币（这些被支持的代币存放在 RToken 合约中），我们把这种获利叫做分润；

- tInfo[i].lastReward 记录 PledgePool 合约地址在代币 i 上的余额；
- tInfo[i].rewardAccu 记录每个质押主代币应得的代币 i 分润值的累积总和，该值总是最新的；
- allAmount[rindex][i].lastReward 记录账户 rindex 在质押池中代币 i 上的余额。如果是主代币（i 为 0），则不仅包括质押值，还包括分润值；如果非主代币，则即为在代币 i 上的分润值；
- allAmount[rindex][i].rewardAccu 记录账户 rindex 上一次更新分润值时，每质押主代币应得的代币 i 分润值的累积总和；
- tInfo[i].rewardAccu - allAmount[rindex][i].rewardAccu 即为账户 index 在此次代币增发时，应得的 每质押主代币应得的代币 i 的分润值；
- totalPledge 代表账户质押主代币总量

1. 主代币增发
   addOrder 触发 mintToken 进行代币增发，代币增发到 PledgePool 合约地址，增发量 tv 除以总质押量 totalPledge 即为均分到每个质押主代币上的分润值 tv，将该分润值加到 tInfo[0].rewardAccu 中。tv 乘以账户此时的质押量，即为该账户在此次主代币增发时获得的主代币奖励。

2. 质押主代币
   账户调用 PledgePool 合约中的 pledge 函数进行质押。
   账户在质押主代币时，会首先更新该账户的 allAmount 信息以及 tInfo 信息，然后执行质押操作（即账户转账到 PledgePool 合约地址），之后更新账户的 allAmount[index][0].lastReward 以及 tInfo[0].lastReward、totalPledge。
   **注意**：质押量过小时（具体值不能确定，但小于 1e9），由于合约整除，可能前几次的分润值为 0，随着代币的多次分发，后面才能获得分润值。

3. 查询账户余额
   账户调用 PledgePool 合约中的 getBalance 函数查询余额。
   账户可以查询在质押池中，对于指定代币的余额，包括账户本身的质押金额和分润值。
   用户查询余额时，会计算分润值，但不会更新合约状态变量。

4. 取回余额
   账户调用 PledgePool 合约中的 withdraw 函数取回余额。
   Keeper、Provider 账户的角色申请金额被锁定，不能被取回。
   账户取回余额时，指定代币序列号 i 和取回金额 money，会先更新该账户在主代币和代币 i 上的 allAmount 信息以及 tInfo 信息，之后根据账户余额和角色申请金确定可以取回的金额\_money，取回余额\_money 后，更新账户 allAmount 信息和 totalPledge.

5. 查询质押池在代币序列号 i 上的余额
   账户调用 PledgePool 合约中的 getPledge 函数，查询质押池在代币 i 上的余额。

## 代币增发机制

User 增加存储订单（调用 addOrder）且指定使用的代币为主代币时，触发主代币增发。订单中的 start、end、sPrice、\_size 决定了代币增发量。具体机制在 Issuance 合约中实现。

Issuance 合约中的 subPMap 记录 end 时刻的累积 sPrice，即每笔 end 值相同的存储订单的 sPrice 之和；subSMap 记录 end 时刻的累积 \_size，即每笔 end 值相同的存储订单的 \_size 之和；lastMint 记录上次触发增发的时间；price 记录所有触发 issu 函数的 sPrice 之和；size 记录所有触发 issu 函数的\_size 之和；totalPaid 记录所有订单中此时应付的金额；spaceTime 记录所有订单的累积时空值（时间乘以空间）；

增发逻辑：

如果此次增发时间 ntime 距离上次增发时间 lastmint 超出了一天，那么将两次增发的时间间隔 dur 设为 86400（即一天），否则正常计算 dur，应支付 paid=price _ dur。如果 ntime 距离 lastmint 超出了一天，那么获得整天数时间 midTime=nowTime/86400 _ 86400，从 subPMap 得出这个时刻的累积 sp。如果 sp 非 0，paid 应该减去 subPay,subPay=sp*(ntime-midTime),price 也应该减去 sp。从 subSMap 得出这个 midTime 时刻的累积 ssize，如果非 0，则 size 减去 ssize。最终 totalPaid 加上此次的 paid 值，spaceTime 加上此次的 st(st=\_size*(end-start))。

更新 totalPay，totalPay 加上此次订单的支付金额 pay(pay=sPrice\*(end-start))。更新 size，size 加上此次订单的\_size。更新 price，price 加上此次订单的 sPrice。

此时，如果 paid 为 0，更新 lastMint 值为 ntime，之后直接返回 0（返回增发量）。

接下来计算 mintLevel 和 reward。memo 维护 4 个 mintLevel，mintLevel 初始值为 0。如果当前全网的时空值除以下一个 mintLevel 规定的 size(或者 Issuance 合约维护的 size，取两者中的最大值)大于等于下一个 mintLevel 的 duration，那么就进入下一个 mintLevel，否则，继续停留在当前 mintLevel 上。增发值 reward 等于 paid 乘以当前 mintLevel 规定的百分比，之后更新 lastMint 值为 ntime，并返回 reward.

即：
根据当前市场需要支付的额度计算增发额度；根据总时空值确定增发系数。
激励系数 A，初始为 1；当前的时空总量为 st，当前的存储总量 s；若 st/max(s, 100TB) > 1 年；解锁 A=1.2；若 st/max(s, 1PB) > 1 年；解锁 A=1.5；若 st/max(s, 10PB) > 1 年；解锁 A=2； 鼓励大家增加存储来解锁更高的比例，以此来提高奖励上限；
激励金额，当前存储需要支付的总额度为 m，当前存储的平均存储时长为 T，当前有 t 时间未分配奖励，则激励为 A*m*t/T。

## 文件系统

admin 创建 group，随机将任意多个 Provider 加到某 group 中。Keeper 自主选择加入某个 group，User 也自主选择加入某 group.每个角色只能存在于一个 group 中，且一旦选择，不可更改。每个 group 都配置了一个文件系统组（FileSys 合约），文件系统组管理着该 group 中所有 User 的文件系统。每个 User 有它自己的文件系统，该系统中，User 付费寻求存储服务，该 group 中的 Provider 可为其提供存储服务、Keeper 为其提供存储管理服务。

Role合约中的GroupInfo.size存储该group中所有订单目前有效的累积存储空间，GroupInfo.price存储该group中所有订单目前有效的累积sPrice(空间单价)。

FileSys 合约中的 balances 记录账户在 FileSys 合约中的余额。

FileSys 合约中的 count 会记录 Keeper 触发 subOrder、addRepair、subRepair 的次数，totalCount 会记录触发的总次数，tAcc 记录 Keeper 应得的收益（即存储费用的 4%）。

在一笔存储订单中，每个 Provider 根据提供的存储空间和时间进行收费，这笔订单中所有 Provider 的收益记为 PsIncome。此外，4% _ PsIncome 的费用还需要付给 Keepers（其中的 1%在订单结束(调用 subOrder 和 subRepair)后发放（endPaid），剩余的 3%在 Provider 执行取回收益操作时线性发放（linearPaid）），1% _ PsIncome 的费用付给 Memo 基金会。因此，在一笔存储订单中，User 需要付费 105% \* PsIncome。

nonce 值，用来记录 User 针对 Provider 增加订单的次数；subNonce 用来记录 User 针对 Provider 结束订单的次数。

FileSys 合约中的 proInfo 保存 Provider 的存储收益相关信息，包括对该 group 中所有 User 提供的存储服务。其中 Settlement.maxPay 表示该 Provider 在代币 i 上的订单累积最大收益；Settlement.hasPaid 表示已经支付给 Provider 的金额；Settlement.canPay 表示目前这个时间可以支付给 Provider 的金额；Settlement.time 记录的是store状态改变或者支付的时间（每一次增加的订单（包括修复订单）的 start 时间，每一次结束订单或结束修复订单时的end时间，以及每一次Provider取回余额的时间）；Settlement.price 记录的是该 Provider 对所有 User 累积的有效 sprice(即每次 addOrder、addRepair 的 sprice 累积之和，在订单到期结束后，还需减去 subOrder 和 subRepair 对应的 sprice)；Settlement.size 记录的是该 Provider 对所有 User 累积的有效存储空间（即每次 addOrder、addRepair 的 size 之和，在订单到期结束后，还需减去 subOrder 和 subRepair 对应的 size）；Settlement.managePay 记录的是与该 Provider 接收的所有存储订单相关联的 Keeper 管理收益。

FileSys 合约中的 fs 记录 User 的文件系统信息。其中 ChannelInfo 记录 User 读取文件数据的相关信息，StoreInfo 记录 User 存储文件的相关信息。sInfo[tindex].time 记录每次 addOrder 时订单的 end 值，每次 addOrder 时，该笔订单的 end 值需要大于等于上笔订单的 end 值；sInfo[tIndex].price 记录每次 addOrder 的 sprice 累积之和，即所有订单的总 sprice，当 subOrder 时，需要减掉相应的 sprice；sInfo[tIndex].size 记录每次 addOrder 的 size 累积之和，即所有订单的总 size，当 subOrder 时，需要减掉相应的 size.

FileSys 合约中的 repairFs 记录所有修复 Provider 的存储服务信息。sInfo[tIndex].price 记录每次 addRepair 的 sprice 累积之和，即所有修复订单的总 sprice，当 subRepair 时，需要减掉相应的 sprice；sInfo[tIndex].size 记录每次 addRepair 的 size 累积之和，即所有修复订单的总 size，当 subRepair 时，需要减掉相应的 size.

### User 充值

User 调用 Role 合约中的 recharge 函数进行充值。
User 往它对应的文件系统中充值，用于存储付费。可充值系统支持的代币，即 RToken 合约中保存的代币。
充值操作即将账户上的金额转移到 FileSys 合约账户，并且更新 FileSys 合约中的 balances[uIndex][tindex]。

### 增加存储订单

User 调用 RoleFS 合约中的 addOrder 函数，指定 pindex、订单开始时间 start（时间戳）、订单结束时间 end（时间戳）、订单存储空间 size、订单使用代币 tindex 以及订单的单价 sPrice(每秒钟的价格)，提交存储订单。

要求 end 是整天数时间；User 和 Provider 在同一个 group 中；tindex 需要是 MEMO 系统支持的代币；此外维护一个 nonce 值，记录增加订单的次数，防止同一订单多次增加；订单 end 不小于 fs[_uIndex].ao[_pIndex].sInfo[_tIndex].time，即上笔订单的 end 时间，从而保证每笔订单到期的时间是有序的；

首先判断 User 余额 balances[uIndex][tindex]是否大于该订单需支付的金额 payAndTax，其中 payAndTax=105% _ sprice _ (end - start)，因此 User 在增加存储订单前，需要往 FileSys 中充值足够的金额。其中付给 Provider 的费用 pay=sprice \* (end - start).

之后更新 Provider 的 channel.expire 信息为 end 值、将订单中的 sprice 累加到 fs[_uIndex].ao[_pIndex].sInfo[_tIndex].price、将订单 size 累加到 fs[_uIndex].ao[_pIndex].sInfo[_tIndex].size、将 fs[_uIndex].ao[_pIndex].sInfo[_tIndex].time 赋值为 end（从而记录此次订单的 end 时间）、将 fs[_uIndex].ao[_pIndex].nonce 加 1、更新 foundation 的余额(balances[foundation][_tindex],加上 1%\*pay)和 User 的余额（balances[\_uIndex][_tindex],扣除掉 payAndTax）。

最后，需要更新 Provider 的存储结算信息（即 settlement 结构,proInfo[pIndex][tokenindex]）。其中 se.time 记录的是每一次增加的订单（包括修复订单）的 start 时间，此次订单的 start 时间与上一次订单的 start 时间（即 se.time）之间的差值乘以 se.price 得到 hp，即为这两次订单之间可以支付给 Provider 的费用，se.canPay 需要加上 hp 计算累计值，之后，将 se.time 更新为此次订单的 start。并且 se.price 需要加上 sprice、se.size 需要加上 size.此外还需将 se.maxPay 加上 pay 值、se.managePay 加上 4% \* pay.

如果增加的订单使用的支付代币是主代币（tindex=0)，那么就会触发主代币增发机制。根据增加订单的start、end、sPrice、\_size值确定增发量，从而增发代币到PledgePool合约账户上。

### 结束存储订单

Keeper或者User可以调用RoleFS合约中的subOrder函数，从而结束订单。指定订单的start、end、_size、nonce、tIndex、sPrice.
要求_size大于0；end大于start且end不大于目前时间；nonce等于当前User和Provider对应的subNonce（fs[uIndex].ao[_pIndex].subNonce）；_size小于等于当前User和Provider对应的size（fs[uIndex].ao[_pIndex].sInfo[_tokenIndex].size）；
更新UP的存储信息中的price（fs[uIndex].ao[_pIndex].sInfo[_tokenIndex].price -= sprice）；以及UP的存储信息中的size（fs[uIndex].ao[_pIndex].sInfo[_tokenIndex].size -= _size）。
更新Provider的结算信息settlement，计算se.time与订单end之间的时间间隔hp，如果se.time小于当前结束订单的end，那么将se.time更新为end。hp乘以se.time得到这段时间可以支付的金额，将该金额加到se.canPay中。之后更新se.price，se.price需要减掉此次结束订单的sPrice，还需更新se.size，se.size需要减去此次订单的_size.
订单结束，需要将应支付给Keepers的金额的1%(_endPaid)加到se.endPaid以及tAcc中，_endPaid = 1% * sPrice * (end-start).剩下的3%在Provider取回余额时结算。
最后，更新subNonce以及Keeper调用订单相关函数的次数（count[kIndex]以及totalCount）。
如果此次订单指定的代币为主代币，则还需更新Role合约中的groups信息，将该订单所属的group中的size和price依次减去此次订单的_size、sPrice.

### 增加修复订单

Keeper调用RoleFS合约中的addRepair函数，从而增加修复数据的存储订单。需指定旧Provider(即故障Provider)、新Provider、订单开始时间start、订单结束时间end、订单空间_size、订单_nonce、订单sPrice、订单使用的代币tIndex。调用该函数需要新Provider的签名以及至少三分之二Keepers的签名。
要求_size大于0、end大于start、Keeper与新旧Provider同属于一个group、该订单nonce值与repairFs.ao[newProvider].nonce一致。
首先需要验证旧Provider的lost余额（由于未能响应挑战而应付的金额记为lost，由于修复订单而支付的金额记为lostPaid，这里的余额即指se.lost-se.lostPaid）是否足够支付此次订单新Provider应得的金额pay(pay=sPrice*(end-start))。RepairFs.providers记录所有修复订单中的新Provider，所以需更新repairFs.providers，并且需更新repairFs.ao[newProvider].sInfo[tIndex]中的price(累积有效sPrice)和size(累积有效size)。与addOrder逻辑一致，需更新新Provider（newProvider）的settlement信息。
之后将repairFs中newProvider的nonce值加1，旧Provider的lostPaid值加上此次修复订单应付修复金额pay，并且需增加调用者Keeper的调用次数count[kindex]与totalCount.
如果该订单指定的是主代币，那么将需要更新Issuance合约中的数据。修复订单代表旧Provider上有数据丢失，那么该Provider将不会得到该段时间内的存储支付，则需要更新Issuance合约中代表全网已支付金额的变量totalPaid,将totalPaid减去此次订单的sPrice*(end-start).

### 结束修复订单

由Keeper调用RoleFS合约中的subRepair函数，从而结束修复订单。需指定旧Provider、新Provider、订单开始时间start、订单结束时间end、订单空间_size、订单_nonce、订单代币tIndex、订单空间单价sPrice。调用该函数需要新Provider的签名以及至少三分之二Keepers的签名。
要求_size大于0、end大于start、end不大于当前时间、Keeper与新旧Provider同属于一个group、该订单nonce值与repairFs.ao[newProvider].subNonce一致、该订单_size不大于repairFs.ao[newProvider].sInfo[tIndex].size。
更新FileSys合约中repairFs变量里面newProvider的累积price信息（减去sPrice）、累积size信息（减去_size）。与subOrder逻辑一致，需更新新Provider（newProvider）的settlement信息。
修复订单结束，触发对Keepers的1%支付，计算_endPaid=1% * sPrice * (end-start)，proInfo[newProvider][tIndex].endPaid加上_endPaid，tAcc[tIndex]也需加上_endPaid。
之后需更新subNonce、count[kIndex]以及totalCount。

### 取回收益

**Provider**

Provider 调用 RoleFS 合约中的 proWithdraw 函数（指定想要取回的余额值 pay）取回在 FileSys 中的余额，pay 需要大于 Settlement.hasPaid 且小于等于 Settlement.canPay，根据调用的时间更新 Settlement.canPay 和 Settlement.time，本次可以取回的存储服务费 thisPay 即为 pay - Settlement.hasPaid , 之后将 Settlement.hasPaid 设置为 pay 值。再根据 thisPay 相应更新 keeper 的 linearPay 值，linearPay 加上 3%的 thisPay，tAcc[tokenIndex] 加上 3%的 thisPay 。Provider 本次真实取回的金额值即为：thisPay。
求出 thisPay 值后，将 FileSys 合约账户上的 thisPay 金额转给 Provider，即完成 Provider 取回收益操作。

如何更新 Settlement.canPay 和 Settlement.time？
Provider 调用 proWithdraw 函数时，当前时间 ntime 与 Settlement.time 之间的时间间隔 hp，乘以 Settlement.price，即为这段时间的存储服务费用 money，则将 Settlement.canPay 加上 money，实现 Settlement.canPay 的更新，之后将 Settlement.time 更新为 ntime.

**Keeper**

Keeper 调用 Role 合约中的 withdrawFromFs 函数，取回自己在文件系统中的收益。Keeper 取回收益时，如果距离上次分发利润的时间（lastTime）超出了系统设置的分利润时间间隔（period），则触发分润。

当触发分润时，根据 Keeper 触发次数占比（count[kindex]/totalCount）,划分应得收益 pro（count[kindex]/totalCount \* tAcc[tindex]），将应得收益 pro（即分润值）加到 balances[kindex]][tindex]上，从而更新 FileSys 合约中 Keeper 的 balance。之后 Keeper 就可以取回收益了，当然 Keeper 指定的取回收益值不能超出其在 FileSys 合约中的 balance.

- Keeper 分润值 tAcc 更新: 订单到期结束订单时（即触发 subOrder、subRepair 时），会将该 Provider 的收益的 1%加到 tAcc[tindex]中；当 Provider 调用 proWithdraw 取回收益时，会将 Provider 此次能取走的收益的 3%加到 tAcc[tindex]中。

**foundation**

基金会在每个文件系统中均能分到收益，即在每个 FileSys 合约中都有余额，foundation 指定 uindex、tindex、amount 调用 Role 合约中的 withdrawFromFs 函数，从 uindex 代表的 FileSys 中取回收益。
foundation 在 FileSys 中的收益金额记录在 FileSys 合约中的 balances[foundation]中，每当 User 增加存储订单（即调用 addOrder 函数）时，更新 balances[foundation]。根据订单中指定的开始结束时间（end、start）和空间价格（sizePrice，简写成 sprice）得出该笔订单在存储服务上需给 Provider 的付费值 pay(pay = sprice _ (end - start))，其中 1% _ pay 即为 foundation 在这笔订单中的收益。

### User 取回充值

User 可以通过调用 Role 合约中的 withdrawFromFs 函数，取回自己在文件系统中充值的金额。直接将 FileSys 合约地址中的金额转移到 User 账户，并更新 FileSys 合约中的 balances[uIndex][tindex]。
