const hre = require("hardhat");

async function main() {
  // We get the contract to deploy
  const Recover = await hre.ethers.getContractFactory("Recover");
  const recover = await Recover.deploy();

  await recover.deployed();

  console.log("Recover deployed to:", recover.address);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
