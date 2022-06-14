// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/IAuth.sol";

/**
 *@author MemoLabs
 *@title Manage owner and auth in memo system
 */
contract Owner {
    // can support mask if using uint32 instead of uint8 
    mapping(uint8=>address) public instances;

    event Alter(uint8, address from, address to);

    constructor(address _o, address _a) {
        instances[1] = _o;
        instances[2] = _a;
    }

    //函数修饰符，判断是不是owner调用
    modifier onlyOwner(){
        require(msg.sender == instances[1], "N");
        _;
    }

    function alter(uint8 _type, address _a, bytes[] memory signs) external {
        uint size;
        assembly {
            size := extcodesize(_a)
        }
        require(size != 0,"NE"); // need ext addr
        
        address au = instances[2];
        bytes32 h = keccak256(abi.encodePacked(address(this), "alter", _type, _a));
        IAuth(au).perm(h, signs);

        emit Alter(_type, instances[_type], _a);
        instances[_type] = _a;
    }
}