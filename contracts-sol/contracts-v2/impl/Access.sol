// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/IAccess.sol";
import "../Recover.sol";

/**
 *@author MemoLabs
 *@title manage erc20
 */
contract Access is IAccess {
    using Recover for bytes32;

    uint8 public constant floor = 3;
    uint16 public version = 2;
    
    address[] public controls; // five address
    uint256 public nonce;
    
    mapping(address => bool) private access;

    // how to contruct? deployed by admin?
    constructor(address[] memory _addrs) {
        controls = _addrs;
    }

    function set(address account, bool isSet, bytes[] memory signs) external override {
        uint size;
        assembly {
            size := extcodesize(account)
        }
        require(size != 0,"NE"); // need ext addr

        bytes32 h = keccak256(abi.encodePacked(address(this), nonce, "set", account, isSet));
        uint8 valid = 0;
        for(uint256 i=0;i<signs.length;i++){
            if(h.recover(signs[i])==controls[i]){
                valid+=1;
            }
        }
        require(valid>=floor,"SNE"); // sign not enough

        access[account] = isSet;
        nonce+=1;
    }

    function can(address account) external view override returns (bool) {
        return access[account];
    }
}