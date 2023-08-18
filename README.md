### chainArt
区块链的开源创作社区
[![Security Status](https://www.murphysec.com/platform3/v31/badge/1692592543501275136.svg)](https://www.murphysec.com/console/report/1692585989234380800/1692592543501275136)
## 安装geth

```text
https://goethereumbook.org/zh/
```
```bash
yum install -y golang git
git clone https://github.com/ethereum/go-ethereum.git
cd go-ethereum
export GOPROXY=https://proxy.golang.com.cn,direct
make all
```

### geth dev

```bash
geth --datadir ./devdata/ --dev --http --http.addr 0.0.0.0 --http.api web3,eth,debug,personal,net --rpc.allow-unprotected-txs console
```

### solidity compile

+ solidity into the abi and bin

    + idea config

      ```bash
      --abi --bin $FileName$ -o ./
      ```

+ abi and bin Into the java

    + web3j-cli download

        + windows

      ```bash
      Set-ExecutionPolicy Bypass -Scope Process -Force; iex ((New-Object System.Net.WebClient).DownloadString('https://raw.githubusercontent.com/web3j/web3j-installer/master/installer.ps1'))
      ```

        + linux

      ```bash
      curl -L get.web3j.io | sh
      source $HOME/.web3j/source.sh
      ```

    + Into the java

  ```bash
  web3j generate solidity --abiFile GoodsStore.abi --binFile GoodsStore.bin -o . -p top.xpit.geth.contract
  ```

  ```bash
  web3j generate solidity --abiFile Escrow.abi --binFile Escrow.bin -o . -p top.xpit.geth.contract  
  ```

```text
Kwei(Babbage) = 10^3 Wei
Mwei(Lovelace) = 10^6 Wei
Gwei(Shannon) = 10^9 Wei
Microether(Szabo) = 10^12 Wei
Milliether(Finney) = 10^15 Wei
Ether = 10^18 Wei
```


```text
eth.sendTransaction({from: eth.accounts[0], to: '0x62E04b39B03c731697c71c9b9A09Ab06f1A0D5c2', value: web3.toWei(1000, 'ether')})
```

```text
"test","1","/profile/upload/2023/03/22/1_20230322194759A001.jpg","<p>test</p>","1679414400","1679673600","5","0"
```

```text
keccak256(abi.encode("15", "test"))
0x7ff1ee2fbf7d5831d8bdf7c2714eef4c4a3f3e225f7e4b80a76b81c3a9f858a7
```

### 项目搭建
+ 获取geth的rpc依赖 通过json-rpc连接区块链
`go get github.com/ethereum/go-ethereum/rpc`
```shell
var name = 'chainArt'
kratos new -- -r https://gitee.com/go-kratos/kratos-layout.git
#生成proto文件
kratos proto add api/${name}/v1/*.proto
#生成pb文件
kratos proto client api/${name}/v1/*.proto
#生成服务端代码
kratos proto server api/${name}/v1/*.proto -t internal/services
#运行程序
kratos run
```
