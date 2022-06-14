# memo-contracts-v2

`@author:Memolabs`

此仓库包含了 memo 项目中的智能合约文件。

## 环境

**语言** solidity

**编辑器** vscode

**语法及功能测试** solc、remix

**集成调试工具** Hardhat

## 文件

合约文件在 contracts 目录下。role 目录包含了角色相关的合约；fileSystem 目录包含了文件系统与结算相关的合约；pledgePool 目录包含了质押池相关的合约；token 目录包含了代币相关的合约；interfaces 目录包含了合约接口；Owner.sol 主要记录 Role 合约的 owner 信息，包括更换 owner；Recover.sol 是一个根据签名信息恢复出签名者的库文件。

目前由 admin 部署 Role 合约（Owner 合约会被同时自动部署、RToken 合约也会被 Role 合约创建）、FileSys 合约、PledgePool 合约、ERC20 合约（AccessControl 合约会被同时自动部署）、Issuance 合约、RoleFS 合约。

admin 部署合约顺序：ERC20、Role、RoleFS、PledgePool、Issuance、(setPI)、FileSys、(setAddr)

优化 gas 消耗，需考虑：

- 将 msg.sender、msg.value、tx.origin 赋值给局部变量，避免多次访问 msg、tx 信息
- 限制数组的使用
- 限制修改及访问 storage 存储

### Role

```sol
constructor(address f, address t, uint256 pk, uint256 ppro){}
```

需要指定基金会地址、主代币地址以及 pledgeK、pledgeP 参数。其中主代币地址即上述 ERC20 合约地址。
创建后需要根据该合约地址创建 PledgePool 合约、并需要创建 Issuance 合约、RoleFS 合约，然后调用**setPI 函数**将 PledgePool 合约地址、Issuance 合约地址以及 RoleFS 合约地址赋值给该合约。

部署 Role 合约时通过 salt 值创建 RToken 合约。
salt 值设定为：

```sol
bytes32 salt = keccak256(abi.encodePacked(foundation, primaryToken));
```

### FileSys

```sol
/// @dev created by admin; 'r' indicates role-contract address, 'rfs' indicates RoleFS-contract address
constructor(uint64 founder, uint64 _gIndex, address r, address rfs, uint64[] memory _keepers){}
```

需要指定基金会地址、group 索引参数、Role 合约地址、RoleFS 合约地址以及 keepers 的索引。
由 admin 调用 Role 合约中的 createGroup 函数之后，同步部署一个 FileSys 合约（\_gIndex 就是 createGroup 函数的返回值，\_keepers 与调用 createGroup 函数时的第一个参数一致），之后调用 Role 合约中的**setGF 函数**将 FileSys 合约地址赋值给 Role 合约。
注意：每当 CreateGroup 后，需要部署一个 FileSys 合约，并且同样地需要部署一个 RoleFS 合约。

文件系统合约，该合约包含了存储订单及存储支付信息。

### PledgePool

```sol
/// @dev created by admin; 'r' indicates role-contract address
constructor(address primeToken, address _rToken, address r){}
```

需要指定主代币地址、RToken 合约地址（可从 Role 合约中获取）、Role 合约地址。

质押池合约，任何拥有序列号的账户（调用 Role 合约中的 register 函数获得序列号）可以进行质押，从而参与分润。账户在注册角色时质押的金额也是质押到该质押池中。

### ERC20

```sol
constructor(string memory _name, string memory _symbol){}
```

需要指定代币名称及符号。

作为 memo 系统支付的主代币，其在 addOrder 被触发时，进行增发。addOrder 函数在 RoleFS.sol 文件中。

deploy Gas Used: 2212493 gas (根据 Remix 中的测试结果，下同)

### Issuance

```sol
constructor(address rfs){}
```

需要指定 RoleFS 合约地址。

### RoleFS

```sol
constructor(){}
```

不需要指定参数。但是部署 RoleFS 合约之后，需要 admin 调用该合约中的**setAddr 函数**，并指定 Issuance、Role、FileSys、RToken 合约地址。

## 注册角色流程

在 Role 合约中，账户首先调用 register 函数获得序列号，注册 User 的账户可以直接调用 registerUser 函数，注册 Keeper 和 Provider 的账户，需要先调用 PledgePool 合约中的 pledge 函数进行质押，之后再调用 Role 合约中的 registerKeeper 和 registerProvider 函数。

## 先决条件

账户调用 RoleFS 合约中的 addOrder 函数前，需要 admin 调用 AccessControl 合约的 setRole 函数，为 RoleFS 合约赋予 MINTER_ROLE 权限。

## TODO list

### FileSys.sol

1. 针对 ChannelInfo 数据结构：
   在每次 addOrder 时，会根据 end 值更新 ChannelInfo.expire，但是另外两个字段：amount、nonce 目前暂未用到；
2. period 状态变量可考虑做成可修改的，目前是一个常量；
3. addOrder 操作中，应该相对应地也更新 fs[_uIndex].ao[_pIndex].sInfo[_tIndex].time，目前该值一直未被更新；（已处理）
4. 1%金额在结束时支付给 keeper，在 subOrder 操作中，记录到 proInfo[\_pIndex][_tokenindex].endPaid 中；另外 3%线性支付的部分，在 proWithdraw 时处理。但是处理逻辑有问题,可以做的更精简。（已处理）
5. proWithdraw 中，取回余额要加上 balances[pindex][tokenindex]，但其实 balances[pindex][tokenindex]一直是 0，所以这个操作比较多余，需去掉。（已处理）
6. 首次 addOrder 时，调用\_settlementAdd 函数更新 Provider 的存储结算相关信息时，se.time 是 0，所以此时计算的 proInfo[\_pIndex][_tokenindex].canPay = se.price \* (start - 0)，这是错误的，需要判断 se.time 是否为 0，如果是 0，则不更新 canPay。（已处理）
7. fs[uIndex].tokenIndex 没有用，因为每次 addOrder 时可以指定任意 Memo 支持的 token，所以应该去掉。（已处理）
8. 合约状态变量 penalty（记录惩罚信息）并未处理。

### PledgePool.sol

1. pledge 函数中，签名信息的 hash 值暂定为 keccak256(abi.encodePacked(caller, money, "pledge"))，若是用户是由其他账户（caller）代为调用 pledge 函数进行质押，这条签名信息，理论上可以被 caller 无限利用，使得用户多次重复质押。如何避免该签名被重复利用？
   可以在链下维护一个 nonce 值，记录每次签名的 nonce，nonce 值依次累加。
2. withdraw 函数中，也面临上述问题；

### Role.sol

1. recharge 函数中，如何避免签名被重复利用；
2. withdrawFromFs 函数中，如何避免签名被重复利用；
3. 暂未添加修改 RoleInfo 和 GroupInfo 结构中`isBanned`参数的函数；
4. 暂未实现 Keeper、Provider 取回角色申请金、取消自己的角色这一功能；

### RoleFS.sol

1. addOrder 函数中，尚未实现对 uSign、pSign、kSigns 的签名验证；
2. subOrder 函数中，尚未实现对 uSign、pSign、kSigns 的签名验证；
3. proWithdraw 函数中，尚未实现对 kSigns 的签名验证；

## 交接待办

代币随着新增订单进行增发外，还需要随着时间增发。需要加入一个增发因子、最小增发比例。如果累积奖励 reward+totalreward 超出了目标奖励，且增发比例没有达到最小增发比例时，则超出部分需要除以 2.
Memo 的增发将会在触发目标增发值后减半，所以早期的生态参与方可以获得更多红利。
激励增发分配的首次目标释放为创世发行的一半，并在达到目标后减半，直至到达设定的最低增发率，保持稳定。
也就是说：mintLevel 内应该加入增发数量的属性，限制每个阶段可以增发的代币上限，超过后，增发率减半（但不低于最小值），直至订单总量达到下一阶段要求后，跳至下一阶段。（已解决）

此外，账户质押金额获得分润与账户质押量所占比例决定，也应与时长有关，质押时间长，就获得分润更多。这一点可通过分润时机实现。比如每 3 天触发一次分润，那么质押的越早的账户可获得分润的次数就越多。
目前是当有账户质押或者取款时，就触发分润。也能实现质押的越久，分润值越多。
