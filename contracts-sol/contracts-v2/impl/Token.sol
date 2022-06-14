// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/IToken.sol";
import "./Owner.sol";

/**
 *@author MemoLabs
 *@title Manage token addresses that memo supported.
 */
contract Token is IToken, Owner {
    struct TokenInfo {
        uint8 index;  // 代币序列号
        bool isBanned; // 该币种是否被禁止
    }

    uint16 public version = 2;

    address[] tokens;
    mapping(address => TokenInfo) tInfo;

    event AddT(address t, uint8 tIndex);

    constructor(address _ctl, address _a) Owner(_ctl, _a) {
    }

    function addT(address t) external override onlyOwner returns(uint8) {
        require(!_hasToken(t), "TE"); // token exist 

        uint8 tIndex = uint8(tokens.length);
        tokens.push(t);
        tInfo[t].index = tIndex;

        emit AddT(t, tIndex);
        return tIndex;
    }

    // need banT 
    function banT(address tAddr) external override onlyOwner {
        tInfo[tAddr].isBanned = true;
    }

    function _hasToken(address tAddr) internal view returns (bool) {
        if (tokens.length == 0) {
            return false;
        } 

        if (tInfo[tAddr].index > 0) {
            return true;
        }

        if (tokens[0] == tAddr) {
            return true;
        }

        return false;
    }

    // =========get==========
    function getTA(uint8 tIndex) external view override returns (address, bool) {
        return (tokens[tIndex],tInfo[tokens[tIndex]].isBanned);
    }

    function getTI(address t) external view override returns (uint8, bool) {
        uint8 tIndex = tInfo[t].index;
        if(t==tokens[tIndex]){ // 注册过
            return (tIndex, tInfo[t].isBanned);
        }
        return (tIndex, true);
    }

    function getTNum() external view returns (uint8) {
        return uint8(tokens.length);
    }
}