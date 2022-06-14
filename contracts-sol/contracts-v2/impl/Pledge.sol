//SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/IPledge.sol";
import "../interfaces/IToken.sol";
import "./Owner.sol";

/**
 *@author MemoLabs
 *@title Manage pledge in memo system
 */
contract Pledge is IPledge, Owner {

    struct RewardInfo {
        uint256 rewardAccu; // 本代币的accumulator, multi by 10^18
        uint256 lastReward; // 上一次变更时的本代币奖励总量。等于余额加奖励
    }

    bytes4 private constant SELECTOR0 = bytes4(keccak256(bytes('balanceOf(address)')));

    uint16 public version = 2;

    uint8 private cnt;       

    uint256 public totalPledge; // 账户质押量

    mapping(uint8 => RewardInfo) tInfo;  // 每种代币的信息, tokenIndex=>RewardInfo
    mapping(uint64 => mapping(uint8 => RewardInfo)) allAmount; // 所有质押的人的信息，(roleIndex => tokenIndex => RewardInfo)

    constructor(address _ctl, address _a, address _token, address _pool) Owner(_ctl, _a){
        instances[5] = _pool;
        instances[7] = _token;
    }

    function _updateReward(uint256 amount, uint8 i, uint64 index) internal {
        uint256 bal = _getBalance(i); // 获取质押池在该代币上的余额

        // 代币i有增发，需更新tInfo[i].rewardAccu
        if (bal > tInfo[i].lastReward && totalPledge>0 ){
            tInfo[i].rewardAccu += (bal - tInfo[i].lastReward) * 1e18 / totalPledge; // 此刻每质押主代币应分得的利润
            tInfo[i].lastReward = bal; // update to latest

            // 将账户余额加上应得的分润值
            // res表示：距离上次更新分润值期间，每质押主代币应得的代币i分润值的累积总和
            uint256 res = tInfo[i].rewardAccu - allAmount[index][i].rewardAccu;
            res = res * amount / 1e18; // amount应大于1e18
            allAmount[index][i].lastReward += res; // 添加分润值
            allAmount[index][i].rewardAccu = tInfo[i].rewardAccu; // 更新accu
        }
    }
    
    function pledge(uint64 index, uint256 money) external onlyOwner override {
        RewardInfo memory reward = allAmount[index][0];
        uint256 amount = reward.lastReward;

        // 更新结算奖励
        for(uint8 i=0; i<cnt; i++){
            reward = allAmount[index][i];
            if(reward.rewardAccu==0){
                // init 
                allAmount[index][i].rewardAccu = tInfo[i].rewardAccu;
            }
            _updateReward(amount, i, index);
        }
        
        // update
        tInfo[0].lastReward += money;
        allAmount[index][0].lastReward += money;
        totalPledge += money;
    }

    function withdraw(uint64 index, uint8 tIndex, uint256 money, uint256 lock) external onlyOwner override returns (uint256) {
        RewardInfo memory reward = allAmount[index][0];
        uint256 amount = reward.lastReward;

        // no pledge
        if(amount==0 || totalPledge == 0){
            return 0;
        }
     
        // 更新结算奖励
        if (tIndex == 0) {
            // update all tokens due to pledge change
            for(uint8 i=0; i<cnt; i++) {
                _updateReward(amount, i, index);
            }
        } else {
            // only update i
            _updateReward(amount, tIndex, index);
        }
    
        // 确定value
        reward = allAmount[index][tIndex];
        uint256 rw = reward.lastReward;
        if(tIndex==0){
            if (rw > lock){
                rw -= lock;  // lock minimum pledge value
            }else {
                rw = 0;
            }
        }

        if(money<rw){
            rw = money;
        }

        if(rw>0){
            // update token
            if (tInfo[tIndex].lastReward > rw){
                tInfo[tIndex].lastReward -= rw;
            }else {
                tInfo[tIndex].lastReward = 0;
            }

            // update account value
            if (allAmount[index][tIndex].lastReward > rw){
                allAmount[index][tIndex].lastReward -= rw;
            }else {
                allAmount[index][tIndex].lastReward = 0;
            }

            // update totalPledge
            if(tIndex==0){
                totalPledge -= rw;
            }
        }

        return rw;
    }

    function addT(uint8 tIndex) external onlyOwner override {
        tInfo[tIndex].rewardAccu = 0;
        tInfo[tIndex].lastReward = _getBalance(tIndex);
        if (tIndex +1 > cnt) {
            cnt = tIndex + 1;
        }
    }

    // ========== get ===========
    // get balance of Pool-address with tIndex 
    function _getBalance(uint8 tIndex) internal view returns (uint256) {
        (address tAddr, ) = ITokenGetter(instances[7]).getTA(tIndex);
        (bool success, bytes memory data) = tAddr.staticcall(abi.encodeWithSelector(SELECTOR0, instances[5]));
        require(success && data.length >= 32, "CE"); // staticcall error
        return abi.decode(data, (uint256));
    }

    // pledge + mint, balance of pool
    function getPledge(uint8 tIndex) external view returns (uint256) {
        return _getBalance(tIndex);
    }

    // get balance of index, only calculate the reward, don't update the reward
    function balanceOf(uint64 index, uint8 tIndex) external view override returns (uint256) {
        RewardInfo memory reward0 = allAmount[index][0];
        if (reward0.lastReward==0) {
            return 0;
        }
        uint256 amount = reward0.lastReward;

        RewardInfo memory rewardt = tInfo[tIndex];
        uint256 val = rewardt.rewardAccu;

        uint256 bal = _getBalance(tIndex);
        if(bal > rewardt.lastReward && totalPledge > 0) {
            bal = (bal - rewardt.lastReward) * 1e18 / totalPledge;
            val = val + bal;
        }

        RewardInfo memory rewardi = allAmount[index][tIndex];
        val = rewardi.lastReward + (val - rewardi.rewardAccu)*amount/1e18;

        return val;
    }
}