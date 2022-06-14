// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

struct RoleOut {
    bool isActive;  // user注册后即生效，k-p注册后还需要质押才能生效
    bool isBanned;  // 是否被admin禁止
    uint8 roleType; // 0 account; 1 user; 2 provider; 3 keeper
    uint64 index;   // 账户序列号，从1开始
    uint64 gIndex;  // 所属group，从1开始
    address owner;  // acoount for receiving income
    bytes extra;    // 用于存储额外的信息，0 empty; 1 verify key; 2 proxy address?; 3 bls public key
}

struct GroupOut {
    bool isActive; // true when keeper count >= level 
    bool isBanned; // 组是否已被禁用
    bool isReady;  // 是否已在线下成组; 由签名触发
    uint16 level;  // security level
    uint64 kCnt;   // keeper count
    uint64 pCnt;   // keeper count
    uint64 uCnt;   // keeper count
    uint256 size;   // storeSize
    uint256 price;  // storePrice
    address pool;
    address fsAddr; // fs contract address
}

struct PledgeOut {
    uint256 kpledge;
    uint256 ppledge;
    address at;
}

// re:register; ac: account; t: token; p: provider; k: keeper; u: user; pp: pledgepool
interface IRoleSetter {
    // called by admin
    function activate(uint64 _i, bool _active) external returns (address);
    function ban(uint64 _i, bool _ban) external;
    function banG(uint64 _gi, bool _isBan) external;

    function createGroup(uint16 _level, uint8 _mr, uint256 _kr, uint256 _pr) external;
    function registerAccount(address _a) external; 
    function registerRole(uint64 _i, uint8 _rtype, bytes memory extra) external;
    function addToGroup(uint64 _i, uint64 _gi, uint256 money) external;
}

interface IRoleGetter {
    // index is rType and not in some group
    function checkIR(uint64 _i, uint8 _rType) external view returns (address);
    // index is rType and in some group
    function checkIG(uint64 _i, uint8 _rType) external view returns (address, address, uint64, uint256);
    
    function getIndex(address _a) external view returns (uint64);
    function getAddr(uint64 _i) external view returns (address);
    function getRInfo(address _a) external view returns (RoleOut memory);
    function getGInfo(uint64 _i) external view returns (GroupOut memory);
    function getPInfo(uint64 _i) external view returns (uint256, uint256);

    function getKManage(uint64 _i) external view returns (address);
    function getPool(uint64 _i) external view returns (address);
    function getKCnt(uint64 _i) external view returns (uint64);
    function getGroupK(uint64 ig, uint64 ik) external view returns (uint64);
    function getGroupP(uint64 ig, uint64 ip) external view returns (uint64);
}

interface IRole is IRoleSetter,IRoleGetter {}