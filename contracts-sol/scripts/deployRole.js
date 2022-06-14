const hre = require("hardhat");

async function main() {
  // We get the contract to deploy
  const Role = await hre.ethers.getContractFactory("Role", {
    libraries: {
      Recover: "0x0DCd1Bf9A1b36cE34237eEaFef220932846BCD82",
    },
  });
  // 包含bigNumber数字转化
  const role = await Role.deploy("0x1c111472F298E4119150850c198C657DA1F8a368","0x0B306BF915C4d645ff596e518fAf3F9669b97016", hre.ethers.BigNumber.from("9000000000000000000"), hre.ethers.BigNumber.from("10000000000000000000"));

  await role.deployed();

  console.log("Role deployed to:", role.address);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
