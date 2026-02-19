// scripts/mine.js
async function main() {
  const hre = require("hardhat");
  // 挖一个块
  await hre.ethers.provider.send("evm_mine", []);
  console.log("✅ Block mined");
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
