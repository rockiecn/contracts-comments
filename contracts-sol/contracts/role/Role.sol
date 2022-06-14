//SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/fileSysIn/IFileSys.sol";
import "../interfaces/pledgePoolIn/IPledgePool.sol";
import "../interfaces/erc20In/IERC20.sol";
import "../interfaces/roleIn/IIssuance.sol";
import "./RToken.sol";
import "../Recover.sol";
import "../Owner.sol";
import "../interfaces/roleIn/IRole.sol";

/**
 *@author MemoLabs
 *@title Used to register account roles for people participating in the memo system
 */
contract Role is IRole, Owner {
    using Recover for bytes32;

    struct RoleInfo {
        bool isActive;  // user注册后即生效，k-p注册后还需要质押才能生效
        bool isBanned;  // 是否被admin禁止
        uint8 roleType; // 0 account; 1 user; 2 provider; 3 keeper
        uint64 index;   // 账户序列号，从1开始
        uint64 gIndex;  // 所属group，从1开始
        bytes extra;    // 用于存储额外的信息，0 empty; 1 verify key; 2 proxy address?; 3 bls public key 
    }

    struct GroupInfo {
        bool isActive; // true when keeper count >= level 
        bool isBanned; // 组是否已被禁用
        bool isReady;  // 是否已在线下成组; 由签名触发
        uint16 level;  // security level
        uint64[] keepers;   // 里面有哪些keeper
        uint64[] providers; // 有哪些provider
        uint64[] users;     // 有哪些provider
        uint256 size;   // storeSize
        uint256 price;  // storePrice
        address fsAddr; // fs contract address
    }

    uint16 public version;

    uint256 public pledgeK; // keeper质押的最小金额
    uint256 public pledgeP; // provider质押的最小金额
    
    address public foundation; // foundation address 基金会账户地址，索引默认为0,不需要进行register

    address[] addrs; // all roles 序列号即为index,从1开始
    mapping(address => RoleInfo) info;

    // manage group
    GroupInfo[] groups;

    address public rToken;
    address public issuance;
    address public rolefs;
    address public pledgePool; // pledge pool 合约地址

    event RUser(uint64 index, address addr);
    event RKeeper(uint64 index, address addr);
    event RProvider(uint64 index, address addr);
    event CreateGroup(uint64 gIndex);

    // called by admin, specify foundation、primaryToken、pledgeK、pledgeP, and create RToken-contract
    // 创建后需要根据该合约地址创建PledgePool合约及Issuance合约，然后调用setPI函数将PledgePool合约地址及Issuance合约地址赋值给该合约
    constructor(address f, address t, uint256 pk, uint256 pp, uint16 _version) {
        foundation = f;
        pledgeK = pk;
        pledgeP = pp;
        version = _version;

        // create RToken-contract by "create2"
        bytes32 salt = keccak256(abi.encodePacked(f, t));
        address rtoken;
        bytes memory b = type(RToken).creationCode;
        assembly {
            rtoken := create2(0, add(b, 0x20), mload(b), salt)
        }
        rToken = rtoken;
        RToken r = RToken(rtoken);
        r.addT(t);
    }

    receive() external payable {}

    /// @dev set pledgePool and issuance and rolefs
    function setPI(address _p, address i, address rfs) external onlyOwner {
        pledgePool = _p;
        issuance = i;
        rolefs = rfs;
    }

    // 账户本身调用register()获得序号
    function register() external override returns (uint64) {
        if(info[msg.sender].index != 0) {
            return info[msg.sender].index;
        } else {
            addrs.push(msg.sender);
            uint64 len = uint64(addrs.length);
            info[msg.sender].index = len;
            return len;
        }
    }

    /// @dev check if 'index' indicates user(roleType:1) or provider(roleType:2) or keeper(roleType:3); and not in some group
    // for register
    function checkIR(uint64 _index, uint8 _rType) external view override returns (address) {
        address a = addrs[_index-1];
        require(info[a].roleType==_rType && !info[a].isActive && !info[a].isBanned, "AE"); // account error
        return a;
    }

    // role is already in some group
    function checkIG(uint64 _index, uint8 _rType) private view returns (address, uint64) {
        address a = addrs[_index-1];
        require(info[a].roleType==_rType && info[a].isActive && !info[a].isBanned, "AE"); // account error
        return (a, info[a].gIndex);
    }

    /// @dev check whether 'tIndex' is valid
    function checkT(uint32 tIndex) public view override returns (address) {
        RToken rt = RToken(rToken);
        require(rt.isValid(tIndex), "TE"); // payToken error
        return rt.getTA(tIndex);
    }

    function registerKeeper(uint64 _index, bytes memory _blsKey) external override {
        address addr = checkIR(_index, 0);
        require(msg.sender == addr, "IC");

        // 检查质押金额是否满足要求
        uint256 pledgedMoney = IPledgePool(payable(pledgePool)).getBalance(_index,0); // 获得_index账户在主代币上的质押金额
        require(pledgedMoney >= pledgeK, "NE"); // pledged money is not enough
        
        info[addr].roleType = 3;
        info[addr].extra = _blsKey;

        emit RKeeper(_index, addr);
    }

    function registerProvider(uint64 _index) external override {
        address addr = checkIR(_index, 0);
        require(msg.sender == addr, "IC");

        // 检查质押金额是否满足要求
        uint256 pledgedMoney = IPledgePool(payable(pledgePool)).getBalance(_index, 0);
        require(pledgedMoney >= pledgeP, "NE"); // pledged money is not enough

        info[addr].roleType = 2;

        emit RProvider(_index, addr);
    }

    function registerUser(uint64 _index, uint64 _gIndex, bytes memory blsKey) external override {
        // check params
        address addr = checkIR(_index, 0);
        require(msg.sender == addr, "IC");

        require(groups[_gIndex-1].isActive && !groups[_gIndex-1].isBanned, "GB");;

        // create fs
        IFileSys(payable(groups[_gIndex-1].fsAddr)).createFs(_index);

        groups[_gIndex-1].users.push(_index);

        info[addr].roleType = 1;
        info[addr].gIndex = _gIndex;
        info[addr].extra = blsKey;
        info[addr].isActive = true;

        emit RUser(_index, addr);
    }

    // called by owner
    // 前提是：PledgePool合约被部署并将地址告知Role合约
    function registerToken(address tAddr) external override onlyOwner {
        // check if exist
        RToken rt = RToken(rToken);
        uint32 tIndex = rt.addT(tAddr);

        // update PledgePool.sol
        IPledgePool(payable(pledgePool)).addToken(tAddr, tIndex);
    }

    // called by owner
    // indexes是keepers的索引
    // 调用该函数时，需要同步部署一个FileSys合约，并将FileSys合约地址赋值给该group（通过调用setGF）
    function createGroup(uint64[] memory indexes, uint16 _level) external override onlyOwner {
        GroupInfo memory g;
        groups.push(g);
        uint64 _gIndex = uint64(groups.length);

        for(uint8 i = 0; i<indexes.length; i++) {
            address addr = checkIR(indexes[i], 3);
            info[addr].gIndex = _gIndex;
            info[addr].isActive = true;
        }

        groups[_gIndex-1].level = _level;
        groups[_gIndex-1].keepers = indexes;

        if(indexes.length >= uint(_level)) {
            groups[_gIndex-1].isActive = true;
        }
        emit CreateGroup(_gIndex);
    }

    // 给group赋值fsAddr,由admin createGroup之后，部署一个FileSys合约，之后调用setGF函数将FileSys合约地址赋值给Role合约; 只能设置一次
    function setGF(uint64 _gIndex, address _fsAddr) external onlyOwner {
        // require fsAddr in group is empty; only set once 
        require(groups[_gIndex-1].fsAddr == address(0), "NE"); // not empty
        groups[_gIndex-1].fsAddr = _fsAddr;
    }

    // called by owner
    function addKeeperToGroup(uint64 _index, uint64 _gIndex) external override onlyOwner {
        require(!groups[_gIndex-1].isBanned, "GB"); // group is banned
        
        address addr = checkIR(_index, 3);

        IFileSys(payable(groups[_gIndex-1].fsAddr)).addKeeper(_index);

        groups[_gIndex-1].keepers.push(_index);
        info[addr].gIndex = _gIndex;
        info[addr].isActive = true;
        if(groups[_gIndex-1].keepers.length >= groups[_gIndex-1].level) {
            groups[_gIndex-1].isActive = true;
        }
    }

    function addProviderToGroup(uint64 _index, uint64 _gIndex) external override {
        // check params
        address addr = checkIR(_index, 2);
        require(msg.sender == addr, "IC");

        require(groups[_gIndex-1].isActive && !groups[_gIndex-1].isBanned, "GB");

        groups[_gIndex-1].providers.push(_index);
        info[addr].gIndex = _gIndex;
        info[addr].isActive = true;
    }

    function setPledgeMoney(uint256 kPledge, uint256 pPledge) external override onlyOwner {
        pledgeK = kPledge;
        pledgeP = pPledge;
    }


    // 可由账户本身调用或由其他账户代为调用
    // 调用前需要uAddr先approve合约地址（FileSys.sol）账户指定金额
    // hash(caller, uIndex, tIndex, money)
    // user往文件系统FileSys合约中充值
    function recharge(uint64 uIndex, uint32 tIndex, uint256 money) external payable override {
        checkIR(uIndex, 1);

        address tAddr = checkT(tIndex);
        
        IFileSys(payable(groups[info[addr].gIndex-1].fsAddr)).recharge(msg.sender, uIndex, tIndex, tAddr, money); 
    }

    // 可由账户本身调用或由其他账户代为调用
    // hash(caller, tIndex, amount)
    // user、keeper、foundation从FileSys合约中取回余额(caller为foundation时，即为foundation取余额)
    function withdrawFromFs(uint64 index, uint32 tIndex, uint256 amount, bytes memory sign) external override {
        // check params
        address tAddr = checkT(tIndex); // tIndex error
        require(amount > 0); // should not be zero
        require(index>0);

        address acc = addrs[index-1];
        address fsAddr = groups[info[acc].gIndex-1].fsAddr;

        uint8 rtype = info[acc].roleType;
        address caller = msg.sender;
        if(caller == foundation){
            index = 0;
            rtype = 0;
            acc = foundation;
        }else{
            checkIR(index, 2);
        }

        // TODO：如何防止该签名被重复利用

        IFileSys(payable(fsAddr)).withdraw(index, tIndex, rtype, tAddr, acc, amount);
    }

    /// @dev 'index' indicates index of 'groups' rather than gIndex
    function setGInfo(SGParams memory ps) external {
        require(msg.sender == rolefs, "N");
        if(ps.isAdd){
            groups[ps.index].size += ps._size;
            groups[ps.index].price += ps.sPrice;
        }else{
            groups[ps.index].size -= ps._size; // 隐含 require(groups[index].size >= _size)
            groups[ps.index].price -= ps.sPrice; // 隐含 require(groups[index].price >= sPrice)
        }
    }

    // ===================get===================

    function getAddrsNum() external view returns (uint64) {
        return uint64(addrs.length);
    }

    // 根据数组索引值（不是账户index）获取相应的账户地址，超出数组范围将revert
    function getAddr(uint64 i) external view override returns (address) {
        return addrs[i];
    }

    function getRoleIndex(address acc) external view returns (uint64) {
        return info[acc].index;
    }

    function getRoleInfo(address acc) external view override returns (bool, bool, uint8, uint64, uint64, bytes memory) {
        return (info[acc].isActive, info[acc].isBanned, info[acc].roleType, info[acc].index, info[acc].gIndex, info[acc].extra);
    }

    function getGroupsNum() external view returns (uint64) {
        return uint64(groups.length);
    }

    function getGroupInfo(uint64 i) external view returns (bool, bool, bool, uint16, uint256,uint256, address) {
        return (groups[i].isActive, groups[i].isBanned, groups[i].isReady, groups[i].level, groups[i].size, groups[i].price, groups[i].fsAddr);
    }

    function getFsAddr(uint64 i) external view override returns (address) {
        return groups[i-1].fsAddr;
    }

    function getAddrGindex(uint64 i) external view override returns (address,uint64) {
        address acc = addrs[i-1];
        return (acc,info[acc].gIndex);
    }

    function getGKNum(uint64 i) external view override returns (uint64) {
        return uint64(groups[i].keepers.length);
    }

    function getGUPNum(uint64 i) external view returns (uint64, uint64) {
        return (uint64(groups[i].users.length),uint64(groups[i].providers.length));
    }

    // 根据groups的数组索引（不是group-index）、以及keepers的数组索引查找相应的keeper-index
    function getGroupK(uint64 ig, uint64 ik) external view returns (uint64) {
        return groups[ig].keepers[ik];
    }

    function getGroupP(uint64 ig, uint64 ip) external view returns (uint64) {
        return groups[ig].providers[ip];
    }

    function getGU(uint64 ig, uint64 iu) external view returns (uint64) {
        return groups[ig].users[iu];
    }
}