// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "./IFileSys.sol";

interface IControl {
    // called by admin 
    function activate(uint64 _i, bool _active, bytes[] memory signs) external;
    function ban(uint64 _i, bool _ban, bytes[] memory signs) external;
    function addT(address _t, bool _ban, bytes[] memory signs) external;
    function banG(uint64 _gi, bool _ban, bytes[] memory signs) external;

    function createGroup(uint16 _level, uint8 _mr, uint256 _kr, uint256 _pr) external;
    function registerAccount(address _a) external; 
    function registerRole(address _a, uint8 _rtype, bytes memory _extra) external;
    function addToGroup(address _a, uint64 _gi) external;
    
    function pledge(address _a, uint64 _i, uint256 _money) external;
    function unpledge(address _a, uint64 _i, uint8 _ti, uint256 _money) external;

    function addOrder(address _a, OrderIn memory _oi) external;
    function subOrder(address _a, OrderIn memory _oi) external;

    function recharge(address _a, uint64 _i, uint8 _ti, uint256 _money, bool isLock) external;
    function withdraw(address _a, uint64 _i, uint8 _ti, uint256 _money) external;
    function proWithdraw(address _a, PWIn memory _ps, uint64[] memory _kis, bytes[] memory ksigns) external;

    function get(uint8 _type) external view returns(address); 
}