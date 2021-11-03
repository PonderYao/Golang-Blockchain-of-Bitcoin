# Golang-Blockchain-of-Bitcoin



>  基于Go语言实现的区块链模型之比特币



## Project Structure

```shell
│  README.md
│
├─bin
│      coin.exe      -- 可执行文件
│
├─pkg
│   └─windows_amd64  -- Windows下编译
│          coin.a          -- 静态库文件
│
└─src
    ├─coin           -- 驱动代码
    │      main.go         -- 程序入口
    │
    └─core           -- 核心代码
           block.go        -- 区块
           blockchain.go   -- 区块链
           proofOfWork.go  -- 工作量证明
           utils.go        -- 工具方法
```

