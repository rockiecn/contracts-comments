//SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/IRole.sol";
import "../interfaces/IAuth.sol";
import "./Owner.sol";
import "./Pool.sol";
import "./Kmanage.sol";

/**
 *@author MemoLabs
 *@title Manage account, roles and groups in the memo system.
 */
contract Role is IRole, Owner {

    struct RoleInfo {
        bool isActive;  // user注册后即生效，k-p注册后还需要质押才能生效
        bool isBanned;  // 是否被admin禁止
        uint8 roleType; // 0 account; 1 user; 2 provider; 3 keeper
        uint64 index;   // 账户序列号，从1开始
        uint64 gIndex;  // 所属group，从1开始
        address payee;  // acoount for receiving income
        bytes extra;    // 0 empty; 1 verify key; 2 ; 3 bls public key
    }

    struct GroupInfo {
        bool isActive;    // true when keeper count >= level 
        bool isBanned;    // 组是否已被禁用
        uint16 level;     // security level
        uint256 kpr;      // keeper pledge require 
        uint256 ppr;      // provider pledge require
        address pool;     // poolAddr
        address kManage;  // keeper manage contract address
        uint64[] keepers;   // 里面有哪些keeper
        uint64[] providers; // 有哪些provider
        uint64[] users;     // 有哪些provider
    }

    uint16 public version = 2;

    address[] addrs; // all roles 序列号即为index,从0开始, addrs[0]为foundation地址
    mapping(address => RoleInfo) info;

    GroupInfo[] groups; // manage group

    event ReAcc(address addr, uint64 index); // to get all registered account by filter logs
    event ReRole(uint8 indexed _rType, uint64 index); // to get all users/keepers/providers by filter logs
    event CreateGroup(uint64 gIndex);

    constructor(address _ctl, address _a, address f) Owner(_ctl, _a) {
        addrs.push(f);
        info[f].payee = f;
        info[f].isActive = true;
        emit ReAcc(f, 0);

        GroupInfo memory g;
        g.isBanned = true;
        groups.push(g);
    }

    // used for keeper
    function activate(uint64 _index, bool _active) external onlyOwner override returns (address) {
        address a = addrs[_index];
        require(!info[a].isBanned, "IB"); // is banned
        if (info[a].roleType == 3 && _active && !info[a].isActive && info[a].gIndex > 0) {
            uint64 gIndex = info[a].gIndex; 
            groups[gIndex].keepers.push(_index);
            if (groups[gIndex].keepers.length >= groups[gIndex].level) {
                groups[gIndex].isActive = true;
            }
            return groups[gIndex].kManage;
        }
        info[a].isActive = _active;
        return address(0);
    }

    function ban(uint64 _index, bool _banned) external onlyOwner override {
        address a = addrs[_index];
        info[a].isBanned = _banned;
    }

    function banG(uint64 _gi, bool _isBan) external onlyOwner override {
        groups[_gi].isBanned = _isBan;
    }

    /// @dev check if 'index' is rType and not in some group
    function checkIR(uint64 _index, uint8 _rType) external view override returns (address) {
        address a = addrs[_index];
        require(info[a].roleType==_rType && !info[a].isActive && !info[a].isBanned, "AE"); // account error
        return a;
    }

    // role is already in some group
    function checkIG(uint64 _index, uint8 _rType) external view override returns (address,address, uint64, uint256) {
        address a = addrs[_index];
        require( !info[a].isBanned, "AE"); // account error
        if (_rType > 0) {
            require(info[a].isActive && info[a].roleType==_rType, "TE"); // Type error
        }

        uint256 lock;
        uint64 gIndex = info[a].gIndex;
        if (gIndex > 0) {
            require(groups[gIndex].isActive&&!groups[gIndex].isBanned, "GE");
            if (_rType == 2) {
                lock = groups[gIndex].ppr;
            }
            if (_rType == 3) {
                lock = groups[gIndex].kpr;
            }
        }
        
        return (a,info[a].payee, gIndex, lock);
    }

    function registerAccount(address a) external onlyOwner override {
        if(info[a].index == 0 && !info[a].isActive) {
            uint64 len = uint64(addrs.length);
            info[a].index = len;
            info[a].payee = a;
            addrs.push(a);
            emit ReAcc(a, len);
        }
    }

    function registerRole(uint64 _index, uint8 _rType, bytes memory _extra) external onlyOwner override {
        address a = this.checkIR(_index, 0);
        info[a].roleType = _rType;
        info[a].extra = _extra;
        emit ReRole(_rType, _index);
    }

    function createGroup(uint16 _level, uint8 mr, uint256 _kr, uint256 _pr) external onlyOwner override {
        uint64 _gIndex = uint64(groups.length);

        // create pool address; force each group has unique pool  
        Pool p = new Pool(instances[1],instances[2]);
        Kmanage k = new Kmanage(instances[1], instances[2], mr);
        
        GroupInfo memory g;
        g.level = _level;
        g.kpr = _kr;
        g.ppr = _pr; 
        g.pool = address(p);
        g.kManage = address(k);
        groups.push(g);

        emit CreateGroup(_gIndex);
    }

    function addToGroup(uint64 _index, uint64 _gIndex, uint256 _pm) external onlyOwner override {
        require(!groups[_gIndex].isBanned, "GB"); // group is banned
    
        address a = addrs[_index];
        require(!info[a].isActive && !info[a].isBanned && info[a].gIndex == 0, "AE"); // account error

        if (info[a].roleType == 1) {
            info[a].isActive = true;
            groups[_gIndex].users.push(_index);
        }

        if (info[a].roleType == 2) {
            require(_pm >= groups[_gIndex].ppr, "KPI"); // pledge insuf
            info[a].isActive = true;
            groups[_gIndex].providers.push(_index);
        }

        if (info[a].roleType == 3) {
            require(_pm >= groups[_gIndex].kpr, "PPI"); // pledge insuf
        }

        info[a].gIndex = _gIndex;
    }

    // ===================get===================
    function getACnt() external view returns (uint64) {
        return uint64(addrs.length);
    }

    function getGCnt() external view returns (uint64) {
        return uint64(groups.length);
    }

    function getAddr(uint64 _i) external view override returns (address) {
        return addrs[_i];
    }

    function getIndex(address _a) external view override returns (uint64) {
        return info[_a].index;
    }

    function getRInfo(address acc) external view override returns (RoleOut memory) {
        RoleOut memory r;
        r.isActive = info[acc].isActive;
        return r;
    }

    function getPInfo(uint64 i) external view override returns (uint256, uint256) {
        return (groups[i].kpr, groups[i].ppr);
    }

    function getGInfo(uint64 i) external view override returns (GroupOut memory) {
        GroupOut memory g;
        g.isActive = groups[i].isActive;
        return g;
    }

    function getKManage(uint64 _i) external view override returns (address) {
        return groups[_i].kManage;
    }

    function getPool(uint64 i) external view override returns (address) {
        return groups[i].pool;
    }

    function getKCnt(uint64 i) external view override returns (uint64) {
        return uint64(groups[i].keepers.length);
    }

    function getGroupK(uint64 _ig, uint64 _ik) external view override returns (uint64) {
        return groups[_ig].keepers[_ik];
    }

    function getGroupP(uint64 _ig, uint64 _ip) external view override returns (uint64) {
        return groups[_ig].providers[_ip];
    }
}