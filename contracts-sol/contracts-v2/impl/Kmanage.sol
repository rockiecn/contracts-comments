// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/IKmanage.sol";
import "./Owner.sol";

/**
 *@author MemoLabs
 *@title keeper management in memo system, per group
 */
contract Kmanage is IKmanage, Owner {

    struct StoreInfo {
        uint64 size; 
        uint256 sprice; 
    }

    uint16 public version = 2;

    uint8 public manageRate = 4; // group分得资金的百分比；4% for group, 其中3% linear and 1% at end;

    uint64 period = 7*86400;     // keeper根据比例获取收益的时间间隔, one week
    uint64 lastTime;             // 上次分利润时间

    uint64 totalCount; // 记录所有keeper触发order相关函数的总次数

    uint64[] keepers;
    uint8[] tokens;
    mapping(uint8 => uint256) tAcc; // 记录分润值
    
    mapping(uint64 => uint64) count; // 记录keeper触发Order相关函数的次数，用于分润

    mapping(uint64 => mapping(uint8 => uint256)) balances; // 账户可用的余额
    mapping(uint8 => StoreInfo) sinfo;
    
    // owner: Control-contract-address
    constructor(address _ctl, address _a, uint8 mr) Owner(_ctl, _a) {
        manageRate = mr; 
        lastTime = uint64(block.timestamp);
    }

    // judge if tokens has the _tokenIndex
    function _hasToken(uint8 _ti) internal view returns (bool) {
        for(uint8 i = 0; i<tokens.length; i++){
            if(tokens[i]==_ti){
                return true;
            }
        }
        return false;
    }

    // after add keeper to group, called by Control.sol
    function addKeeper(uint64 _ki) external override onlyOwner {
        keepers.push(_ki);
        count[_ki] = 1;
        totalCount++;
    }

    // after sub order, called by Control.sol
    function addCnt(uint64 _ki, uint64 cnt) external override onlyOwner {
        require(count[_ki] > 0, "IC");

        count[_ki]+= cnt;
        totalCount+= cnt;
    }

    // after pro withdraw and subOrder, called by Control.sol
    function addProfit(uint8 _ti, uint256 _money) external override onlyOwner {
        tAcc[_ti] += _money;
    }

    // after addOrder and subOrder, called by Control.sol
    function addSP(uint8 _ti, uint64 _size, uint256 _sprice, bool isAdded) external override onlyOwner {
        if (isAdded) {
            sinfo[_ti].size += _size;
            sinfo[_ti].sprice += _sprice;
        } else {
           sinfo[_ti].size -= _size;
            sinfo[_ti].sprice -= _sprice; 
        }
    }

    // after withdraw, called by Control.sol
    function withdraw(uint64 _ki, uint8 _ti, uint256 amount) external override onlyOwner returns(uint256){
        if (count[_ki] > 0) {
            uint64 ntime = uint64(block.timestamp);
            if(ntime-lastTime > period){
                if(totalCount != 0){
                    for(uint i = 0; i< tokens.length; i++){
                        uint256 value = tAcc[tokens[i]]; // 分润值
                        if(value==0){
                            continue;
                        }
                        uint256 per = value / uint256(totalCount);
                        for(uint j =0; j<keepers.length; j++){
                            uint256 pro = per * count[keepers[j]];
                            balances[keepers[j]][tokens[i]] += pro;
                        }
                        tAcc[tokens[i]] = 0;
                    }
                }
                lastTime = ntime;
            }
        }
    
        uint256 bal = balances[_ki][_ti];
        if(amount>bal){
            amount=bal;
        }

        balances[_ki][_ti] -= amount;
        return amount;
    }   

    // ================get=============
    function getRate() external view override returns (uint8) {
        return manageRate;
    }

    function getSP(uint8 _ti) external view override returns (uint64, uint256) {
        return (sinfo[_ti].size, sinfo[_ti].sprice);
    }

    function balanceOf(uint64 _ki, uint8 _ti) external view override returns(uint256, uint256){
        uint256 avail = balances[_ki][_ti];
        uint256 tmp = 0;

        if(totalCount == 0){
            return (avail, tmp);
        }

        if(count[_ki]!=0){
            uint256 sum = tAcc[_ti] * count[_ki];
            uint256 pro = sum / totalCount;
            if((block.timestamp-lastTime)>=period){
                avail += pro;
            }else{
                tmp+=pro;
            }
        }

        return (avail, tmp);
    }
}