const { task } = require('hardhat/config');

/**
 * @type import('hardhat/config').HardhatUserConfig
 */
require("@nomiclabs/hardhat-waffle");
require("@nomiclabs/hardhat-web3");
require("@nomiclabs/hardhat-ethers");
require("hardhat-contract-sizer");

task("accounts", "Prints the list of accounts", async (taskArgs, hre) => {
  const accounts = await hre.ethers.getSigners();

  for (const account of accounts) {
    console.log(account.address);
  }
});

task("balance", "Prints the balance of inputed account")
.addParam("account", "the account's address")
.setAction(async (taskArgs) => {
  const account = hre.web3.utils.toChecksumAddress(taskArgs.account);
  const balance = await hre.web3.eth.getBalance(account);
  console.log(hre.web3.utils.fromWei(balance, "ether"), "ETH");
});

// task("logs", "Prints the logs filtered by params")
// .addOptionalParam("from", "Account that generated the log, is a contract address")
// .addParam("topic", "The topic of log name")
// .setAction(async ({from}, taskArgs) => {
//   //const account = web3.utils.toChecksumAddress(from);
//   const result = await hre.web3.eth.subscribe('logs', {
//     address: from,
//     topics: [taskArgs.topic]
//   });
//   console.log(result);
// })

module.exports = {
  solidity: "0.8.10",
};
