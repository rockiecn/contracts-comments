// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

// contain destory method
// can be once
contract Mint {
    mintMax, minted public uint256
    erc20Addr, mintTarget public address   

    constructor(address _erc20Addr , _mintTarget, uint256 _mintMax) {
        erc20Addr = _erc20Addr;
        mintTarget = _mintTarget;
        mintMax = _mintMax;
    }

    function mint(mintValue uint256) onlyOwner public {
        require(minted + mintValue <= mintMax, "EMM") // exceed mint max
        IERC20 e = IERC20(erc20Addr);
        e.mint(mintTarget, mintValue);
        minted += mintValue;
    }

    function destroy() onlyOwner public {
        selfdestruct(owner);
    }
}