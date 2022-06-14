// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

struct SGParams {
        uint64 index;
        bool isAdd;
        uint256 _size;
        uint256 sPrice;
    }

interface IRole {
    // 更新Role合约及 PledgePool 合约
    function registerToken(address taddr) external; // called by admin, add a erc20 token, so the system can use the token

    function register(address addr, bytes memory sign) external returns (uint64); // register address to get index
    function registerKeeper(uint64 index, bytes memory blsKey, bytes memory sign) external; // register Keeper，需要从 PledgePool 合约中查询其质押量
    function registerProvider(uint64 index, bytes memory sign) external; // register Provider，同上
    function registerUser(uint64 index,uint64 _gIndex, bytes memory blsKey, bytes memory sign) external; // register User,需要调用 FileSys 合约创建一个fs
    function recharge(uint64 uIndex, uint32 tokenIndex, uint256 money, bytes memory sign) external payable; //user往 FileSys 合约中充值
    //function proWithdraw(uint64 pIndex, uint32 tIndex, uint256 pay, uint256 lost, bytes[] memory ksigns) external; //provider从 FileSys 合约中取回余额
    function withdrawFromFs(uint64 index, uint32 tIndex, uint256 amount, bytes memory sign) external; // 账户从 FileSys 合约中取回余额

    function checkIR(uint64 _index, uint8 _rType) external view returns (address);
    function checkT(uint32 tIndex) external view returns (address);
    //function setGInfo(uint64 index, bool isAdd, uint256 _size, uint256 sPrice) external;
    function setGInfo(SGParams memory ps) external;

    // // 往 FileSys 合约中增加一个订单，并触发 ERC20 合约中的代币增发
    // function addOrder(uint64 uIndex, uint64 pIndex, uint64 _start, uint64 end, uint64 _size, uint64 nonce, uint32 tIndex, uint256 sPrice) external;
    // // 从 FileSys 合约中减少一个订单
    // function subOrder(uint64 uIndex, uint64 pIndex, uint64 _start, uint64 end, uint64 _size, uint64 nonce, uint32 tIndex, uint256 sPrice) external;
    // // 往 FileSys 合约中增添一个repair
    // function addRepair(uint64 pIndex, uint64 newPIndex, uint64 _start, uint64 end, uint64 _size, uint64 nonce, uint32 tIndex, uint256 sprice, bytes memory pSign, bytes[] memory kSigns) external;
    // // 往 FileSys 合约中减少一个repair
    // function subRepair(uint64 pIndex, uint64 newPIndex, uint64 _start, uint64 end, uint64 _size, uint64 nonce, uint32 tIndex, uint256 sPrice, bytes memory pSign, bytes[] memory kSigns) external;

    // owner新建一个 FileSys 合约
    function createGroup(uint64[] memory indexes, uint16 _level) external returns (uint64); // called by admin, create a group
    // owner往 FileSys 合约中新加一个keeper
    function addKeeperToGroup(uint64 _index, uint64 _gIndex) external;
    // 账户把自己添加进Role合约中的某个group中
    function addProviderToGroup(uint64 _index, uint64 _gIndex, bytes memory sign) external;
    // 设置Role合约中的某些参数
    function setPledgeMoney(uint256 kPledge, uint256 pPledge) external;

    function getAddr(uint64 i) external view returns (address);
    function getRoleInfo(address acc) external view returns (bool, bool, uint8, uint64, uint64, bytes memory);
    function getGKNum(uint64 i) external view returns (uint64);
    function getFsAddr(uint64 i) external view returns (address);
    function getAddrGindex(uint64 i) external view returns (address,uint64);
    function getGroupK(uint64 ig, uint64 ik) external view returns (uint64);
    function pledgeK() external view returns (uint256);
    function pledgeP() external view returns (uint256);
    function pledgePool() external view returns (address);
    function getGroupInfo(uint64 i) external view returns (bool, bool, bool, uint16, uint256, uint256, address);
}