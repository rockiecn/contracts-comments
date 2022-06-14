const hre = require("hardhat");

async function main() {
    const topic0 = hre.Web3.utils.keccak256("RegisterUser(uint64,address)");
    const topic1 = hre.Web3.utils.keccak256("Pledge(address,uint256)");
    const topic2 = hre.Web3.utils.keccak256("Withdraw(address,uint256)")
    console.log("RegisterUser(uint64,address) topic ", topic0);
    console.log("Pledge(address,uint256) topic ", topic1);
    console.log("Withdraw(address,uint256) topic ", topic2);
}

main()
.then(() => process.exit(0))
.catch((error) => {
    console.error(error);
    process.exit(1);
  });