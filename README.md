# memo-go-contracts-v2

## contracts

合约文件通过`abigen`转换为 go 文件，保存在该目录中。

## interface

包含了所有合约的访问接口。

## callcontracts

定义了一个`ContractModule`的结构体，在该目录下，实现了 ContractModule 对上述所有合约接口的调用。

## cmd

该目录下提供了查询余额、转账（erc20 token 以及 eth token）的命令行操作。以及 admin 部署合约、调用合约的命令行操作。

## 注意事项

在 callcontracts/common.go 中定义了 ERC20Addr、RoleAddr、RoleFSAddr、RTokenAddr、FileSysAddr、PledgePoolAddr、IssuanceAddr 以及 admin 的账户地址和私钥信息，以供全局使用。

这些合约地址由 admin 部署合约从而得到，admin 部署合约步骤：

1. DeployERC20 => erc20Addr
2. DeployRole => roleAddr、rTokenAddr
3. DeployRoleFS => rolefsAddr
4. DeployPledgePool => pledgePoolAddr
5. DeployIssuance => issuAddr

recompile memoc

6. role.SetPI 
7. role.CreateGroup => fsAddr
8. rolefs.SetAddr
9. erc20.SetUpRole(rolefsAddr, 1)

admin 选取一部分 keeper(也可以为空，因为后续可以调用 addKeeperToGroup 增添 keeper 到 group 中)创建 group(createGroup)，每个 group 对应一个 filesys。所以在 admin 调用 createGroup 时，会同时部署一个 FileSys 合约。

## index

系统中存在三种 index：角色索引（rIndex）、group 索引（gIndex）、代币索引（tIndex）
rIndex 从 1 开始，不能为 0；
gIndex 也从 1 开始，不能为 0；
tIndex 从 0 开始，tIndex=0 表示主代币；

## nonce

涉及到的 nonce 值都需要从 0 开始依次累加，从而与合约中的 nonce 值匹配；

## 交易返回错误

**ErrTxExecu("transaction mined but execution failed")**
当返回这种错误，可能的情况是：

1. 交易超出设定的 gasLimit
2. 交易的输入信息有问题，导致合约里面执行出现 revert 错误

**ErrTxFail("transaction not packaged")**
出现这种错误，表明链上交易量较多，在给定的时间内交易还未被打包。代码会尝试加大 gasPrice 重新发起交易。

## foundation 取回 fs 中的收益

foundation 在所有的 filesystem 中都有收益。
首先选择一个 userIndex（用来确认哪一个 filesystem ），然后由 foundation 自己调用 Role 合约中的 withdrawFromFs 函数，就可以取得在该 filesystem 中的收益了。
