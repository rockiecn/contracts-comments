const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("Role", function () {
  it("Should return the right value", async function () {
    const signer0 = await ethers.getSigner(0);

    // 部署ERC20合约
    const ERC20 = await ethers.getContractFactory("ERC20");
    const erc20 = await ERC20.deploy("MEMO", "m");
    await erc20.deployed();

    // // 部署Recover合约
    // const Recover = await ethers.getContractFactory("Recover");
    // const recover = await Recover.deploy();
    // await recover.deployed();

    // 需要链接库
    const Role = await this.env.ethers.getContractFactory("Role", {
      libraries: {
        Recover: "0x2279B7A0a67DB372996a5FaB50D91eAA73d2eBe6",
      },
    });
    const role = await Role.deploy(0x1c111472F298E4119150850c198C657DA1F8a368,erc20.address,9000000000000000000, 10000000000000000000);
    await role.deployed();

    expect(await role.getOwner()).to.equal(signer0.address); //验证owner
    console.log("pledgePool: ", await role.pledgePool()); // 打印pledgePool合约地址

    const signer1 = await ethers.getSigner(1);
    const registerTx = await role.connect(signer1).register(signer1.address, 0x0);
    const index = await registerTx.wait();
    console.log("index: ", index);
    expect(index.to.equal(1));

    // let result = await user.getUserInfo(signer1.address); //直接用result接收所有值，验证getUserInfo
    // expect(result[0]).to.equal(false);
    // expect(result[1]).to.equal(0);  //不需要处理这个大数
    // expect(result[2]).to.equal(verifyPk);
  });
});
