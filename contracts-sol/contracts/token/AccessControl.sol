// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/erc20In/IAccessControl.sol";

// 本合约中的role指代的是下面三个constant值，表示访问控制的3种角色等级
contract AccessControl is IAccessControl {
    address[5] public controls; // five address
    signNonce public uint256
    mapping(uint8 => mapping(address => bool)) accessRoles;
    bool public pause; // 为true表示禁止转账操作

    uint8 public constant MINTER_ROLE = 1; // can mint

    // how to contruct?
    constructor(address[5] memory _addrs) {
        controls = _addrs
    }

    function setRole(uint8 role, address account, bytes[5] memory signs) external override returns (bool) {
        bytes32 h = keccak256(abi.encodePacked(address(this), signNonce, "method", account, role));
        uint8 valid = 0;
        for(uint256 i=0;i<5;i++){
            if(h.recover(signs[i])==controls[i]){
                valid++;
            }
        }
        require(valid>=3,"SNE"); // sign not enough

        accessRoles[role][account] = true;
        return true;
    }

    function hasRole(uint8 role, address account) public view override returns (bool) {
        return accessRoles[role][account];
    }

    // admin revoke other account's role
    function revokeRole(uint8 role, address account, bytes[5] memory signs) external override returns (bool) {
        bytes32 h = keccak256(abi.encodePacked(address(this), signNonce, "method", account, role));
        uint8 valid = 0;
        for(uint256 i=0;i<5;i++){
            if(h.recover(signs[i])==controls[i]){
                valid++;
            }
        }
        require(valid>=3,"SNE"); // sign not enough

        accessRoles[role][account] = false;
        return true;
    }

    function pause() external override returns (bool) {
        bytes32 h = keccak256(abi.encodePacked(address(this), signNonce, "method"));
        uint8 valid = 0;
        for(uint256 i=0;i<5;i++){
            if(h.recover(signs[i])==controls[i]){
                valid++;
            }
        }
        require(valid>=3,"SNE"); // sign not enough
        pause = true;
        return true;
    }

    function unpause() external override returns (bool) {
        bytes32 h = keccak256(abi.encodePacked(address(this), signNonce, "method"));
        uint8 valid = 0;
        for(uint256 i=0;i<5;i++){
            if(h.recover(signs[i])==controls[i]){
                valid++;
            }
        }
        require(valid>=3,"SNE"); // sign not enough
        pause = false;
        return true;
    }

    function getPause() public view override returns (bool) { //可见性如果是external，则继承该合约的ERC20合约将不能直接调用该函数
        return pause;
    }
}