module.exports = {
  networks: {
    localhost: {
      url: "http://127.0.0.1:8545",
      mining: {
        auto: true,
        interval: 0  // 立即挖块
      }
    }
  }
};
