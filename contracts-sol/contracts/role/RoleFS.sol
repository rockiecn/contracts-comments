// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/roleIn/IRole.sol";
import "../interfaces/fileSysIn/IFileSys.sol";
import "../interfaces/roleIn/IIssuance.sol";
import "../interfaces/roleIn/IRToken.sol";
import "../interfaces/erc20In/IERC20.sol";
import "../Recover.sol";

/// @dev This contract contains functions involving Role and FileSys contract
contract RoleFS {
    using Recover for bytes32;

    address role;
    address issuance;
    address owner;
    address rToken;
    
    /// @dev created by admin
    constructor() {
        owner = msg.sender;
    }

    // todo: need handle it security
    /// @dev set issuance-contract、role-contract and RToken-contract address, called by owner
    function setAddr(address iss,address r, address rt) external {
        require(msg.sender == owner, "N");
        issuance = iss;
        role = r;
        rToken = rt;
    }

    function checkParam(address caller, AOParams memory ps) internal view returns (address, uint64) {
        IRole r = IRole(role);
        r.checkT(ps.tIndex); // tIndex error

        // 检查参数大小
        require(ps._size > 0, "sz"); // size zero
        require(ps.sPrice > 0, "pz" ); // price zero
        require(ps._end > ps._start + 8640000, "es"); // end should more than start + 100 day
        require(ps._end < ps._start + 86400000, "el"); // end should no more than start + 1000 day
        require(ps._end%86400 == 0, "te"); // end time error; align to day

        // 检查uIndex、pIndex角色是否正确
        (address uAddr, uint64 _gIndex) = r.checkIG(ps.uIndex, 1);
        (address pAddr, uint64 pGIndex) = r.checkIG(ps.pIndex, 2);
        require(_gIndex == pGIndex, "GD"); // group different

        //// verify signatures 
        // calc hash
        bytes32 h = keccak256(abi.encodePacked(
            ps.uIndex,
            ps.pIndex,
            ps.nonce,
            ps._start,
            ps._end,
            ps._size, 
            ps.tIndex, 
            ps.sPrice)
            );

        // verify user sig
        require(h.recover(ps.usign) == uAddr, "IUS"); // illegal user sign

        // verify provider sig
        require(h.recover(ps.psign) == pAddr, "IPS"); // illegal pro sign
        return (uAddr, _gIndex);
    }

    // called by user
    // 拥有user、provider的签名信息，以及其他keeper的聚合签名信息，然后调用此函数
    // hash(uIndex, pIndex, nonce, _start, end, _size, tIndex, sPrice)
    // 调用该函数前需要admin调用AccessControl合约的setUpRole函数，为该合约赋予MINTER_ROLE权限
    function addOrder(AOParams memory ps) external {
        (address uaddr, uint64 _gIndex) = checkParam(ps);
        require(uAddr == msg.sender, "IC"); // illegal caller
        
        IRole r = IRole(role);
        IFileSys(payable(r.getFsAddr(_gIndex))).addOrder(ps);

        // for calling issu()
        IssuParams memory issuPs;
        issuPs._start = ps._start;
        issuPs._end = ps._end;
        issuPs._size = ps._size;
        issuPs.sPrice = ps.sPrice;


        SGParams memory sgPs;
        sgPs.index = _gIndex;
        sgPs.isAdd = true;
        sgPs._size = ps._size;
        sgPs.sPrice = ps.sPrice;

        // update info in group
        r.setGInfo(sgPs);
        
        // 主代币增发
        if (ps.tIndex == 0) { // 0 为主代币
            IIssuance iss = IIssuance(issuance);
            uint256 reward = iss.issu(issuPs);
            
            
   
            // Role合约往PledgePool合约中转入reward数量的tokens[0]代币，即主代币
            // 此处应该为mint，而不是transfer，但是这样的话，需要为RoleFS合约账户赋予mint权限。
            if(reward>0){
                IERC20(IRToken(rToken).getTA(0)).mint(r.pledgePool(), reward);
            }
        }
    }

    // called by keeper or user
    // 拥有user、provider的签名信息，以及其他keeper的聚合签名信息，然后调用此函数
    // hash(uIndex, pIndex, nonce, _start, end, _size, tIndex, sPrice)
    function subOrder(AOParams memory ps) external {
        (address uAddr, uint64 _gIndex) = checkParam(ps);
        require(ps._end <= block.timestamp, "ET"); // time error

        IRole r = IRole(role);
        uint64 kIndex = ps.uIndex;
        if (msg.sender != uAddr) {
            // not user, should be keeper
            (,,,,kIndex,) = r.getRoleInfo(s);
            (bool isActive, bool isBanned, uint8 roleType, kIndex, uint64 gIndex, ) = r.getRoleInfo(msg.sender);
            require(isActive && !isBanned && roleType==3 && gIndex == _gIndex, "IP");
        }    
        
        IFileSys(payable(r.getFsAddr(_gIndex))).subOrder(kIndex,ps);

 
        SGParams memory sgPs;
        sgPs.index = _gIndex;
        sgPs.isAdd = false;
        sgPs._size = ps._size;
        sgPs.sPrice = ps.sPrice;
        r.setGInfo(sgPs);
    }

    // hash（ps.pIndex,ps.tIndex,ps.pay,ps.lost）
    // provider从FileSys合约中取回余额
    function proWithdraw(PWParams memory ps, uint64[] memory kIndexes, bytes[] memory ksigns) external {
        //// prepare params for verify sig
        IRole r = IRole(role);  
        address tAddr = r.checkT(ps.tIndex)

        // get provider address and gIndex from provider location in array
        (address pAddr, uint64 gIndex) = r.checkIG(ps.pIndex);

        uint256 indexNum = kIndexes.length;
        require(indexNum <= ps.ksigns.length, "LS"); // lack sign

        uint8 sigCnt;
        // calc hash
        bytes32 h = keccak256(abi.encodePacked(ps.pIndex,ps.tIndex,ps.pay,ps.lost));
        for(uint64 i = 0; i < indexNum; i++) {
            (address kAddr, uint64 _gIndex) = r.checkIG(kIndexes[i]);
            if (gIndex == _gIndex && h.recover(ksigns[i]) == kAddr) {
                sigCnt += 1;
            }           
        }
        // valid sig should not less than 2*(N+1)/3, N: kNum of group
        require(sigCnt >= 2 * (IRole(role).getGKNum(gIndex-1) + 1) / 3, "KSE"); // kSigns error

        // call fs.proWithdraw
        address fsAddr = r.getFsAddr(gIndex);
        IFileSys(payable(fsAddr)).proWithdraw(pAddr, tAddr, ps);

        // call issuance.setTP, add lost to paid
        if(ps.tIndex==0){
            (,,,,,,uint256 lost,,,,) = IFileSys(payable(fsAddr)).getSettleInfo(ps.pIndex,ps.tIndex); 
            if ps.lost > lost {
                lost = ps.lost - lost
                IIssuance(issuance).setTP(ps.lost);
            }
        }
    }

    // called by keeper
    // hash(pIndex, _start, end, _size, nonce, tIndex, sPrice, "a")
    // signed by newProvider、at least two thirds of the keeper
    function addRepair(uint64 pIndex, uint64 nPIndex, uint64 _start, uint64 end, uint64 _size, uint64 nonce, uint32 tIndex, uint256 sprice, bytes memory pSign, bytes[] memory kSigns) external {
        // todo
    }

    // keeper 调用
    // hash(pIndex, _start, end, _size, nonce, tIndex, sPrice, "s")
    // signed by newProvider、at least two thirds of the keeper
    function subRepair(uint64 pIndex, uint64 nPIndex, uint64 _start, uint64 end, uint64 _size, uint64 nonce, uint32 tIndex, uint256 sPrice, bytes memory pSign, bytes[] memory kSigns) external {
        // todo
    }
}