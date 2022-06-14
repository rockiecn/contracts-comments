// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

interface IAccessControl {
    function setRole(uint8 role, address account) external returns (bool);

    function hasRole(uint8 role, address account) external view returns (bool);

    function revokeRole(uint8 role, address account) external returns (bool);

    // 禁止一切转账行为
    function pause() external returns (bool);

    function unpause() external returns (bool);

    function getPause() external view returns (bool);
}