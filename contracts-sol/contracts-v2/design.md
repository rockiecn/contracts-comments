# design

## root contract

+ proxy

## contract type

+ 1 owner (caller limited)
+ 2 auth (set authrity) can upgrade
+ 3 access (erc control) can upgrade
+ 4 erc (erc data) non-upgrade
+ 5 pool (pledge data) non-upgrade  
+ 6 role (role and group data) non-upgrade 
+ 7 token (token data) non-upgrade
+ 8 pledge (pledge control) can upgrade
+ 9 issue (issue control) can upgrade
+ 10 fs (fs data) non-upgrade

+ 100 control (interact control) can upgrade
+ 101 getter

kmanage (keeper profit data) has no contract type, get it by role.getKManage()

## base contracts

erc
token
pool
pledge
role
issue
fs
kmanage

## layer

proxy -> control -> base contracts 

getter -> base contracts 


## deploy

### deploy erc

+ deploy access => access address
+ deploy erc(access) => erc address

### deploy role

+ deploy auth => auth address
+ deploy proxy(auth)  => proxy address as root
+ deploy getter(proxy) => getter address
+ deploy control(proxy, auth) => control address 
+ deploy token(control, auth) => token address
+ deploy pool(control, auth) => pool address, deployed when createGroup by Role.sol
+ deploy kmanage(control, auth) => kmanage address, deployed when createGroup by Role.sol
+ deploy pledge(control, auth, token, pool) => pledge address
+ deploy role(control, auth, foundation) => role address
+ deploy fs(control, auth) => fs address
+ deploy issuance(control, auth) => issu address

+ set token, pool, kmanage, pledge, role, fs, issuance to control address

### add token

+ addT(erc address)

### create group

+ createGroup(level, kpledge, ppledge, manageRatio) => create pool(control, auth), create kmanage(control, auth) => g.pool, g.kmanage
 
### register

+ register account => role index
+ register role => rtype
+ add to group

note: keeper need activate

## proxy/control

+ function activate(uint64 _i, bool _active, bytes[] memory signs) external

admin激活账户，用于keeper加入组后激活

+ function ban(uint64 _i, bool _ban, bytes[] memory signs) external;

admin禁止某账户

+ function addT(address _t, bool _ban, bytes[] memory signs) external;

admin添加代币合约地址

+ function banG(uint64 _gi, bool _ban, bytes[] memory signs) external;

admin禁止某组

+ function createGroup(uint16 _level, uint256 _kr, uint256 _pr, uint8 _mr) external;

创建一个新的组，组中会创建对应的kmanage，pool合约；kr: keeper pledge requirement, typically 1 memo; pr: provider pledge requirement, typically 1 memo; mr: manage rate, typically 4;

+ function registerAccount(address _a) external; 

地址注册获得账户号

+ function registerRole(address _a, uint8 _rtype, bytes memory _extra) external;

地址注册角色

+ function addToGroup(address _a, uint64 _gi) external;

加入组

+ function pledge(address _a, uint64 _i, uint256 _money) external;

给账户_i质押， 钱从调用者_a中转出

+ function unpledge(address _a, uint64 _i, uint8 _ti, uint256 _money) external;

从质押池取钱

+ function addOrder(address _a, OrderIn memory _oi) external;

验证账户，token合法性
验证user钱是否足够
验证订单合法


+ function subOrder(address _a, OrderIn memory _oi) external;

验证账户，token合法性
验证订单合法
分润给provider、kmanage

+ function recharge(address _a, uint64 _i, uint8 _ti, uint256 _money, bool isLock) external;

给账户_i充值， 钱从调用者_a中转出； isLock表示钱只能用于fs支付，不能提取出来

+ function withdraw(address _a, uint64 _i, uint8 _ti, uint256 _money) external

从fs池中取钱

+ function proWithdraw(address _a, PWIn memory _ps, uint64[] memory _kis, bytes[] memory ksigns) external;

pro取钱

+ function get(uint8 _type) external view returns(address); 

获取合约地址

## golang

+ go-mefs only need proxy.sol(include Owner.sol, which contains address of other contracts) and getter.sol; 

## upgrade

+ data sol(erc、pool、role、token、fs、kmanage): better no upgrade
+ control sol(control、access): can upgrade 
+ base control sol(pledge、issu): can upgrage
+ auth sol: can upgrade, be caution;
