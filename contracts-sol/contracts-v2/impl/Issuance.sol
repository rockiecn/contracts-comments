// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/IIssuance.sol";
import "./Owner.sol";

/**
 *@author MemoLabs
 *@title mint reward in memo system
 */
contract Issuance is IIssuance, Owner {
    struct MintInfo {
        uint16 ratio;
        uint64 duration;
        uint256 size;
    }

    uint16 public version = 2;

    uint16 public issuRatio; // issuance ratio
    uint16 public minRatio; // minimum issuance ratio

    uint public mintLevel;
    MintInfo[] mint;
    uint256 public lastMint;
    uint256 public sprice; // totalSizePrice
    uint256 public size;  // totalSize
    uint256 public spaceTime; // totalSpaceTime
    uint256 public totalPay;
    uint256 public totalPaid;

    /// @notice Issuance target at the current stage.Initial issuance target is 50 billion,after reaching,the target value of each stage is halved
    uint256 public periodTarget;
    uint256 public periodTotalReward; // total reward at the current stage

    mapping(uint64 => uint256) public subPMap; // 保存sprice
    mapping(uint64 => uint256) public subSMap; // 保存size

    constructor(address _ctl, address _a) Owner(_ctl,_a) {
        uint256 time = block.timestamp; // now time
        lastMint = time;

        mint.push(MintInfo(50,8640000,109951162777600)); // 50%, 100*24*60*60 s = 100 days, 100*1024*1024*1024*1024 B = 100 TB
        mint.push(MintInfo(80, 12960000, 1125899906842624)); // 80%, 150*24*60*60 s = 150 days, 1PB
        mint.push(MintInfo(100, 17280000, 56294995342131200)); // 100%, 200days, 50PB
        mint.push(MintInfo(80, 8640000, 1152921504606846976)); // 80%, 300days, 1EB
        mint.push(MintInfo(50, 63072000, 57646075230342348800)); // 50%, 730 days, 50EB

        minRatio = 3;
        periodTarget = 150000000000000000000000000; // initialTarget 150 million * token_decimal
        issuRatio = 50; // mint[0].ratio
    }

    // 主代币激励增加量计算
    function issu(IssuIn memory ps) external onlyOwner returns (uint256) {
        // add sub sprice and size
        subPMap[ps.end] += ps.sPrice;
        subSMap[ps.end] += uint256(ps.size);

        // old
        uint256 nowTime = block.timestamp;
        if(nowTime - lastMint > 86400){
            nowTime = lastMint + 86400;
        }

        uint256 dur = nowTime - uint256(lastMint);
        uint256 paid = sprice * dur;
        // 至少一天后调用
        if(lastMint/86400 < nowTime/86400){
            uint256 midTime = nowTime/86400 * 86400; // 取整天数时间值
            uint256 sp = subPMap[uint64(midTime)];
            if(sp!=0){
                paid -= sp * uint256(nowTime - midTime);
                sprice -= sp;
            } 
            size -= subSMap[uint64(midTime)];
        }
        totalPaid += paid;

        // update info
		// update spacetime for reward ratio
        spaceTime += ps.size * uint256(ps.end - ps.start);

        // update total pay
        totalPay += ps.sPrice * uint256(ps.end - ps.start);

        size += uint256(ps.size);
        sprice += ps.sPrice;

        lastMint = nowTime;

        if(paid<=0){
            return 0;
        }

        // 计算 mintLevel 及 reward
        for(uint256 i=mintLevel; i<mint.length; i++) {
            uint256 esize = mint[i].size;
            if(esize < size) {
                esize = size;
            }
            uint256 du = mint[i].duration;
            if(spaceTime/esize >= du) { // 进入下一阶段
                mintLevel = i+1;
                issuRatio = mint[mintLevel].ratio;
                periodTarget = periodTarget / 2;
                periodTotalReward = 0;
            }else {
                break;
            }
        }

        uint256 reward = paid * uint256(issuRatio) / uint256(100);

        // 如果当前阶段累积reward超出了目标激励增加值，则超出部分触发减半，ratio也需要除以2，直到ratio到达最小值minRatio，将按照最小ratio激励增加。
        uint256 _periodTotalReward = periodTotalReward + reward;
        if(_periodTotalReward>periodTarget && issuRatio>minRatio){
            uint256 left = periodTarget - periodTotalReward;
            uint256 overflow = (reward - left)/2;
            reward = reward - overflow;
            periodTotalReward += reward;
            issuRatio = issuRatio / 2;
            if(issuRatio<minRatio){
                issuRatio = minRatio;
            }
        }

        return reward;
    }
}