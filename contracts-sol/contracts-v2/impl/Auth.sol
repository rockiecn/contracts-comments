// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/IAuth.sol";
import "../Recover.sol";

/**
 *@author MemoLabs
 *@title auth in memo system
 */
contract Auth is IAuth {
    using Recover for bytes32;
    
    uint8 public constant floor = 3;
    uint16 public version = 2;
    
    uint256 public nonce;
    address[] public controls; // five address

    // deployed by admin?
    constructor(address[] memory _addrs) {
        controls = _addrs;
    }

    // multi-signature verification succeeded, then nonce + 1, called by Control.sol and Owner.sol
    function perm(bytes32 ha, bytes[] memory signs) external override {
        // hash(contract address, nonce, method name, params)
        bytes32 h = keccak256(abi.encodePacked(address(this), nonce, "perm", ha));
        uint8 valid = 0;
        for(uint8 i=0;i<signs.length;i++){
            if(h.recover(signs[i])==controls[i]){
                valid+=1;
            }
        }
        require(valid>=floor,"SNE"); // sign not enough
        nonce+=1;
    }
}