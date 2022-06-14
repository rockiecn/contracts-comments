// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/IControl.sol";
import "../interfaces/IRole.sol";
import "../interfaces/IFileSys.sol";
import "../interfaces/IIssuance.sol";
import "../interfaces/IToken.sol";
import "../interfaces/IERC20.sol";
import "../interfaces/IAuth.sol";
import "../interfaces/IPledge.sol";
import "../interfaces/IPool.sol";
import "../interfaces/IKmanage.sol";
import "./Owner.sol";
import "../Recover.sol";

/**
 *@author MemoLabs
 *@title controlling Role, Pledge, FileSys, Issue, Token, erc20, Pool, Kmanage in memo system
 */
contract Control is IControl, Owner {
    using Recover for bytes32;
    uint16 public version = 2;

    // deployed by admin
    constructor(address _o,address _a) Owner(_o, _a) {
    }

    receive() external payable {}

    function activate(uint64 _i, bool _active, bytes[] memory signs) external override {
        IAuth ia = IAuth(instances[2]);
        bytes32 h = keccak256(abi.encodePacked(address(this), "activate", _i, _active));
        ia.perm(h, signs);

        address kmanage = IRoleSetter(instances[6]).activate(_i, _active);
        if (kmanage != address(0)) {
            IKmanageSetter(kmanage).addKeeper(_i);
        }
    }

    function ban(uint64 _i, bool _ban, bytes[] memory signs) external override {
        IAuth ia = IAuth(instances[2]);
        bytes32 h = keccak256(abi.encodePacked(address(this), "ban", _i, _ban));
        ia.perm(h, signs);

        IRoleSetter(instances[6]).ban(_i, _ban);
    } 

    function addT(address _t, bool _ban, bytes[] memory signs) external override {
        IAuth ia = IAuth(instances[2]);
        bytes32 h = keccak256(abi.encodePacked(address(this), "addT", _t, _ban));
        ia.perm(h, signs);
        if (_ban) {
            ITokenSetter(instances[7]).banT(_t);
        } else {
            uint8 ti = ITokenSetter(instances[7]).addT(_t);
            IPledgeSetter(instances[8]).addT(ti);
        }
    }

    function banG(uint64 _gi, bool _isBan, bytes[] memory signs) external override {
        IAuth ia = IAuth(instances[2]);
        bytes32 h = keccak256(abi.encodePacked(address(this), "banG", _gi, _isBan));
        ia.perm(h, signs);

        IRoleSetter(instances[6]).banG(_gi, _isBan);
    }

    function createGroup(uint16 _level, uint8 _mr, uint256 _k, uint256 _p) external override {
        require(_mr > 1 && _mr % 4 == 0, "MRL");
        IRoleSetter(instances[6]).createGroup(_level, _mr, _k, _p);
    }

    function registerAccount(address _a) external onlyOwner override {
        IRoleSetter r = IRoleSetter(instances[6]);
        return r.registerAccount(_a);
    }

    function registerRole(address _a, uint8 _rtype, bytes memory extra) external onlyOwner override {
        uint64 _i = IRoleGetter(instances[6]).getIndex(_a);
        return IRoleSetter(instances[6]).registerRole(_i, _rtype, extra);
    }
    
    function addToGroup(address _a, uint64 _gi) external onlyOwner override {
        uint64 _i = IRoleGetter(instances[6]).getIndex(_a);
        uint256 bal = IPledgeGetter(instances[8]).balanceOf(_i, 0);
        return IRoleSetter(instances[6]).addToGroup(_i, _gi, bal);
    }

    // anyone can pledge using its money, use a's money, pledge for i
    function pledge(address _a ,uint64 _i, uint256 money) external override {
        (address _t, bool _v) = ITokenGetter(instances[7]).getTA(0);
        require(!_v, "TB"); // token banned

        IPool(instances[5]).inflow(_t, _a, money);
        IPledgeSetter(instances[8]).pledge(_i, money);
    }

    // cancle pledge
    function unpledge(address _a, uint64 _i, uint8 _ti, uint256 money) external onlyOwner override {
        (address _t, bool _v) = ITokenGetter(instances[7]).getTA(_ti);
        require(!_v, "TB"); // token banned

        (,address _re, , uint256 _lock) = IRoleGetter(instances[6]).checkIG(_i, 0);
        require(_re == _a, "IC");

        // todo: need lock more from fs

        uint256 _money = IPledgeSetter(instances[8]).withdraw(_i, _ti, money, _lock);
        IPool(instances[5]).outflow(_t, _re, _money);
    }

    // use a's money, recharge for ui
    function recharge(address _a, uint64 _ui, uint8 _ti, uint256 money, bool isLock) external override {
        (address _t, bool _v) = ITokenGetter(instances[7]).getTA(0);
        require(!_v, "TB"); // token banned
        
        IRoleGetter r = IRoleGetter(instances[6]);
        (,,uint64 _gi,)  = r.checkIG(_ui, 0);
        
        IPool(r.getPool(_gi)).inflow(_t, _a, money);
        
        IFileSys(instances[10]).recharge(_ui, _ti, money, isLock); 
    }

    // called by Proxy.sol, tx.origin = a, need a = addr[i] or a = addr[i].payee
    function withdraw(address _a, uint64 _i, uint8 _ti, uint256 money) external onlyOwner override {
        (address _t, bool _v) = ITokenGetter(instances[7]).getTA(_ti);
        require(!_v, "TB"); // token banned

        IRoleGetter r = IRoleGetter(instances[6]);
        (address _w, address _re, uint64 _gi, ) = r.checkIG(_i, 0);
        require((_w == _a ||_re == _a), "IC");

        uint256 _money = IFileSys(instances[10]).withdraw(_i, _ti, money); // reduce balances[i] in fs
        if (money > _money) {
            money -= _money;
            uint256 _m = IKmanage(r.getKManage(_gi)).withdraw(_i, _ti, money);
            _money += _m;
        } 
        IPool(r.getPool(_gi)).outflow(_t, _re, _money);
    }

    function checkParam(OrderIn memory ps) internal view returns (address, uint64) {
        IRoleGetter r = IRoleGetter(instances[6]);

        require(ps.size > 0, "sz"); // size zero
        require(ps.sPrice > 0, "pz" ); // sizePrice zero
        require(ps.end > ps.start + 8640000, "es"); // end more than start+100 day
        require(ps.end < ps.start + 86400000, "el"); // end no more than start+1000 day
        require(ps.end%86400 == 0, "te"); // end time error; align to day

        // 检查uIndex、pIndex角色是否正确
        (address uAddr,,uint64 _gi,) = r.checkIG(ps.uIndex, 1);
        (address pAddr,,uint64 pGIndex,) = r.checkIG(ps.pIndex, 2);
        require(_gi == pGIndex, "GD"); // group different

        //// verify signatures 
        // calc hash
        bytes32 h = keccak256(abi.encodePacked(
            ps.uIndex,
            ps.pIndex,
            ps.nonce,
            ps.start,
            ps.end,
            ps.size, 
            ps.tIndex, 
            ps.sPrice)
            );

        // verify user sig
        require(h.recover(ps.usign) == uAddr, "IUS"); // illegal user sign

        // verify provider sig
        require(h.recover(ps.psign) == pAddr, "IPS"); // illegal pro sign
        return (uAddr, _gi);
    }

    // order duration >= 100 days & <= 1000 days, a = tx.origin, a is user
    function addOrder(address _a, OrderIn memory ps) external onlyOwner override {
        (address _t, bool _v) = ITokenGetter(instances[7]).getTA(ps.tIndex);
        require(!_v, "TB"); //
        
        (address _u, uint64 _gi) = checkParam(ps);
        require(_a == _u, "IC"); // illegal caller
        
        IRoleGetter r = IRoleGetter(instances[6]);
        IKmanage ikm =  IKmanage(r.getKManage(_gi));
        IFileSys(instances[10]).addOrder(ps, uint256(ikm.getRate()));        
        
        // 产生奖励
        if (ps.tIndex == 0) { 
            IssuIn memory _is;
            _is.start = ps.start;
            _is.end = ps.end;
            _is.size = ps.size;
            _is.sPrice = ps.sPrice;

            uint256 reward = IIssuance(instances[9]).issu(_is);
            if(reward>0){
                IERC20(_t).mint(instances[5], reward);
            }
        }
        ikm.addSP(ps.tIndex, ps.size, ps.sPrice, true);  
    }

    function subOrder(address _a, OrderIn memory ps) external onlyOwner override {
        (, bool _v) = ITokenGetter(instances[7]).getTA(ps.tIndex);
        require(!_v, "TB"); //
        
        (address uAddr, uint64 _gi) = checkParam(ps);
        require(ps.end <= block.timestamp, "ET"); // time error

        IRoleGetter r = IRoleGetter(instances[6]);
        IKmanage ikm =  IKmanage(r.getKManage(_gi));

        uint256 ep = IFileSys(instances[10]).subOrder(ps, uint256(ikm.getRate()));

        uint64 kIndex;
        if (_a != uAddr) {
            kIndex = r.getIndex(_a);
            // not user, should be keeper
            (,,uint64 _ngi,) = r.checkIG(kIndex, 3);
            require(_gi == _ngi, "GD");
            
            ikm.addCnt(kIndex, 1);
            ikm.addProfit(ps.tIndex, ep);
        }   

        ikm.addSP(ps.tIndex, ps.size, ps.sPrice, false);  
    }

    // kIndexes is incremental, a = tx.origin, a should be ps.pIndex or ps.pIndex.payee
    function proWithdraw(address _a, PWIn memory ps, uint64[] memory kIndexes, bytes[] memory ksigns) external onlyOwner override {
        (address _t, bool _v) = ITokenGetter(instances[7]).getTA(ps.tIndex);
        require(!_v, "TB"); 

        IRoleGetter r = IRoleGetter(instances[6]);  

        (address _w, address pOwner, uint64 _gi, ) = r.checkIG(ps.pIndex, 2);
        require((_a == _w || _a == pOwner), "IC");

        uint256 indexNum = kIndexes.length;
        require(indexNum <= ksigns.length, "LS"); // lack sign

        uint8 sigCnt;
        uint64 _start;
        bytes32 h = keccak256(abi.encodePacked(ps.pIndex,ps.tIndex,ps.pay,ps.lost));
        for(uint64 i = 0; i < indexNum; i++) {
            require(kIndexes[i] > _start, "DI"); // larger than previous
            (address kAddr,, uint64 gIndex,) = r.checkIG(kIndexes[i], 3);
            if (gIndex == _gi && h.recover(ksigns[i]) == kAddr) {
                sigCnt += 1;
            }
            _start = kIndexes[i];           
        }
        // valid sig should not less than 2*(N+1)/3, N: kNum of group
        require(sigCnt >= 2 * (r.getKCnt(_gi) + 1) / 3, "KSE"); // kSigns error

        IKmanage ikm =  IKmanage(r.getKManage(_gi));
        
        (uint256 _money, uint256 _pr) = IFileSys(instances[10]).proWithdraw(ps, uint256(ikm.getRate()));
        IPool(r.getPool(_gi)).outflow(_t, pOwner, _money);

        ikm.addProfit(ps.tIndex, _pr);
    }

    function get(uint8 _type) external override view returns(address) {
        return instances[_type];
    } 
}