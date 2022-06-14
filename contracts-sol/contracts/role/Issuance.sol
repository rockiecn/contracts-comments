// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/roleIn/IIssuance.sol";

/// @dev This contract is about the issuance of token
contract Issuance {
    struct MintInfo {
        uint16 ratio;
        uint256 size;
        uint64 duration;
    }

    uint public mintLevel;
    MintInfo[] mint;
    uint256 public lastMint;
    uint256 public price; // totalSizePrice
    uint256 public size;  // totalSize
    uint256 public spaceTime; // totalSpaceTime
    uint256 public totalPay;
    uint256 public totalPaid;

    /// @notice Issuance target at the current stage.Initial issuance target is 50 billion,after reaching,the target value of each stage is halved
    uint256 public periodTarget;
    uint256 public periodTotalReward; // total reward at the current stage
    uint16 public issuRatio; // issuance ratio
    uint16 public minRatio; // minimum issuance ratio

    mapping(uint64 => uint256) public subPMap; // 保存price
    mapping(uint64 => uint256) public subSMap; // 保存size
    address roleFs;

    /// @dev specify RoleFS-contract address to control caller; created by admin
    constructor(address rfs) {
        uint256 time = block.timestamp; // now time
        lastMint = time;

        mint.push(MintInfo(50,109951162777600,8640000)); // 50%, 100*1024*1024*1024*1024 B = 100 TB, 100*24*60*60 s = 100 days
        mint.push(MintInfo(80, 1125899906842624, 12960000)); // 80%, 1PB, 150*24*60*60 s = 150 days
        mint.push(MintInfo(100, 56294995342131200, 17280000)); // 100%, 50PB, 200days
        mint.push(MintInfo(80, 1152921504606846976, 8640000)); // 80%, 1EB, 300days
        mint.push(MintInfo(50, 57646075230342348800, 63072000)); // 50%, 50EB, 730 days
        roleFs = rfs;

        minRatio = 3;
        periodTarget = 150000000000000000000000000; // initialTarget 150 million * token_decimal
        issuRatio = 50; // mint[0].ratio
    }

    modifier onlyRoleFS() {
        require(msg.sender == roleFs, "N");
        _;
    }

    // 主代币激励增加量计算
    function issu(IssuParams memory ps) external onlyRoleFS returns (uint256) {
        // add sub price and size
        subPMap[ps._end] += ps.sPrice;
        subSMap[ps._end] += uint256(ps._size);

        // old
        uint256 nowTime = block.timestamp;
        if(nowTime - lastMint > 86400){
            nowTime = lastMint + 86400;
        }

        uint256 dur = nowTime - uint256(lastMint);
        uint256 paid = price * dur;
        // 至少一天后调用
        if(lastMint/86400 < nowTime/86400){
            uint256 midTime = nowTime/86400 * 86400; // 取整天数时间值
            uint256 sp = subPMap[uint64(midTime)];
            if(sp!=0){
                paid -= sp * uint256(nowTime - midTime);
                price -= sp;
            } 
            size -= subSMap[uint64(midTime)];
        }
        totalPaid += paid;

        // update info
		// update spacetime for reward ratio
        spaceTime += ps._size * uint256(ps._end - ps._start);

        // update total pay
        totalPay += ps.sPrice * uint256(ps._end - ps._start);

        size += uint256(ps._size);
        price += ps.sPrice;

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

    /// @dev set totalPaid
    function setTP(uint256 _sub) external onlyRoleFS {
        if(totalPay >= _sub){
            totalPay -= _sub;
        }
    }
}