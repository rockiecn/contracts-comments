// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/IProxy.sol";
import "../interfaces/IControl.sol";
import "./Owner.sol";

/**
 *@author MemoLabs
 *@title Operate Interface in the memo system.
 */
contract Proxy is IProxy, Owner {
    uint16 public version = 2;

    constructor(address _o,address _a) Owner(_o, _a) {
    }

    receive() external payable {}

    function activate(uint64 _i, bool _active, bytes[] memory signs) external override {
        IControl(instances[100]).activate(_i, _active, signs);
    }

    function ban(uint64 _i, bool _ban, bytes[] memory signs) external override {
        IControl(instances[100]).ban(_i, _ban, signs);
    }

    function addT(address _t,bool _ban, bytes[] memory signs) external override {
        IControl(instances[100]).addT(_t, _ban, signs);
    }

    function banG(uint64 _gi, bool _ban, bytes[] memory signs) external override {
        IControl(instances[100]).banG(_gi, _ban, signs);
    }

    function createGroup(uint16 _level, uint8 _mr, uint256 _kr, uint256 _pr) external override {
        IControl(instances[100]).createGroup(_level, _mr, _kr, _pr);
    }

    function registerAccount() external override {
        IControl(instances[100]).registerAccount(msg.sender);
    }
    
    function registerRole(uint8 _rtype, bytes memory _extra) external override {
        IControl(instances[100]).registerRole(msg.sender, _rtype, _extra);
    }

    function addToGroup(uint64 _gi) external override {
        IControl(instances[100]).addToGroup(msg.sender, _gi);
    }
    
    function pledge(uint64 _i, uint256 _money) external override {
        IControl(instances[100]).pledge(msg.sender, _i, _money);
    }

    function unpledge(uint64 _i, uint8 _ti, uint256 _money) external override {
        IControl(instances[100]).unpledge(msg.sender, _i, _ti, _money);
    }

    // called by user
    function addOrder( OrderIn memory _oi) external override {
        IControl(instances[100]).addOrder(msg.sender, _oi);
    }

    // called by user or keeper
    function subOrder(OrderIn memory _oi) external override {
        IControl(instances[100]).subOrder(msg.sender, _oi);
    }

    function recharge(uint64 _i, uint8 _ti, uint256 _money, bool isLock) external override {
        IControl(instances[100]).recharge(msg.sender, _i, _ti, _money, isLock);
    }

    function withdraw(uint64 _i, uint8 _ti, uint256 _money) external override {
        IControl(instances[100]).withdraw(msg.sender, _i, _ti, _money);
    }

    function proWithdraw(PWIn memory _ps, uint64[] memory _kis, bytes[] memory ksigns) external override {
        IControl(instances[100]).proWithdraw(msg.sender, _ps, _kis, ksigns);
    }

    function get(uint8 _type) external view override returns(address) {
        return IControl(instances[100]).get(_type);
    }
}