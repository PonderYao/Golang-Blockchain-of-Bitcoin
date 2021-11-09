# Golang-Blockchain-of-Bitcoin



>  基于Go语言实现的区块链模型之比特币



## Project Structure

```shell
│  README.md
│
├─bin
│      blockchain.db -- 本地持久化存储数据文件
│      coin.exe      -- 可执行文件
│
├─pkg
│   └─windows_amd64  -- Windows下编译
│          coin.a          -- 静态库文件
│
└─src
    ├─github.com     -- 相关github库源码引入
    |   └─boltdb
    |       └─bolt   -- bolt数据库工具
    |
    ├─coin           -- 驱动代码
    │      main.go         -- 程序入口
    │
    └─core           -- 核心代码
           block.go        -- 区块
           blockchain.go   -- 区块链
           client.go       -- 命令解析客户端
           proofOfWork.go  -- 工作量证明
           transaction.go  -- 交易记录
           utils.go        -- 工具方法
```

## Note
### 1. About the Bolt Database
if you would like to test the birth of "Genesis Block", it is suggested to **delete** the file named **blockchain.db** under the package bin by hand.
<br>
如果你想要测试创世纪区块的诞生，建议手动删除`bin`目录下的`blockchain.db`文件