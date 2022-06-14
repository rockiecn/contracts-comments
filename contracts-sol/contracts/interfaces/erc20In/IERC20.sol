// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

interface IERC20 {
    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(
        address indexed owner,
        address indexed spender,
        uint256 value
    );

    function getTotalSupply() external view returns (uint256);

    function balanceOf(address account) external view returns (uint256);

    function transfer(address recipient, uint256 amount)
        external
        returns (bool);

    function allowance(address owner, address spender)
        external
        view
        returns (uint256);

    function approve(address spender, uint256 amount) external returns (bool);

    // can remove these two; use approve instead
    function increaseAllowance(address spender, uint256 addedValue) external returns (bool);
    function decreaseAllowance(address spender, uint256 subtractedValue) external returns (bool);

    function transferFrom(
        address sender,
        address recipient,
        uint256 amount
    ) external returns (bool);

    // 代币发行相关
    function mint(address target, uint256 mintedAmount) external returns (bool);

    // 烧毁代币
    function burn(uint256 burnAmount) external returns (bool);

    // 代币基本信息
    function getName() external view returns (string memory);

    function getSymbol() external view returns (string memory);

    function getDecimals() external view returns (uint8);
}
