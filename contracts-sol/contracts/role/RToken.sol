// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;
import "../interfaces/roleIn/IRToken.sol";

/// @dev This contract is about the token addresses that memo supported.
contract RToken is IRToken {
    struct TokenInfo {
        uint32 index;  // 代币序列号
        bool isBanned; // 该币种是否被禁止
    }

    // 代币信息,0为主代币的代币地址
    address[] tokens;
    mapping(address => TokenInfo) tInfo;
    address private owner;

    /// @dev created by role-contract
    constructor() {
        owner = msg.sender;
    }

    event AddT(address t, uint32 tIndex);

    // called by owner only; why not use onlyOwner?
    function addT(address t) external override returns (uint32 index) {
        require(msg.sender == owner, "NO"); // not owner. 如果revert了会传递到上层调用吗？就是Role合约中也会revert吗？答案是：会
        // check if exist
        if(_hasToken(t)){
            return tInfo[t].index;
        }

        // update storage
        uint32 tIndex = uint32(tokens.length);
        tokens.push(t);
        tInfo[t].index = tIndex;

        emit AddT(t, tIndex);

        return tIndex;
    }

    // need banT 
    function banT(address t) external {

    }

    // =========get==========

    // 检查是否tokenAddress已经存在
    function _hasToken(address tAddr) internal view returns (bool) {
        if tokens.length == 0 {
            return false
        } 

        if tInfo[tAddr].index > 0 {
            return true
        }

        if tokens[0] == tAddr {
            return true
        }

        return false;
    }

    function isValid(uint32 tIndex) external view override returns (bool) {
        if(tIndex < tokens.length && !tInfo[tokens[tIndex]].isBanned){
            return true;
        }
        return false;
    }

    // 访问无效的索引将导致revert
    function getTA(uint32 tIndex) external view override returns (address) {
        if(tIndex < tokens.length){
            return tokens[tIndex];
        }
        return address(0);
    }

    // 返回tIndex以及该代币地址是否有效
    function getTI(address t) external view override returns (uint32, bool) {
        uint32 tIndex = tInfo[t].index;
        if(this.isValid(tIndex)){
            if(t==tokens[tIndex]){ // 注册过
                return (tIndex, true);
            }
        }
        return (tIndex, false);
    }

    function getTNum() external view returns (uint32) {
        return uint32(tokens.length);
    }
}