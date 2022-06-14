const hre = require("hardhat");
const ethers = hre.ethers;

async function main() {
    //查询过去的日志
    const logs = await hre.web3.eth.getPastLogs({
      address: "0xa51c1fc2f0d1a1b8494ed1fe312d7c3a78ed91c0", // 指定合约地址
      topics: ['0x0361d74fc16827043a0abbc44931d101a36f8a404257aff6dc5c6f07c46c9e25']  //根据getTopics.js计算得出
    });
    console.log(logs);
  }
  

main()
.argument()
.then(() => process.exit(0))
.catch((error) => {
    console.error(error);
    process.exit(1);
  });