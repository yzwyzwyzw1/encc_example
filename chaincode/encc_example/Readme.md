
#
```bash
govendor init
govendor add +external

#init     Create the "vendor" folder and the "vendor.json" file.
#+external (e) referenced packages in GOPATH but not in current project.
```
解释：
govendor只是用来管理项目的依赖包，如果GOPATH中本身没有项目的依赖包，则需要通过go get先下载到GOPATH中，再通过govendor add +external拷贝到vendor目录中。

创建所需要的密钥信息：

ENCKEY=`openssl rand 32 | openssl base64` && DECKEY=$ENCKEY

IV=`openssl rand 16 | openssl base64`

# CLI测试

```bash
docker exec -it cli bash
export CHANNEL_NAME=mychannel
peer channel create -o orderer.example.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/mychannel.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
peer channel join -b mychannel.block
peer chaincode install -n mycc -v 1.0 -p github.com/chaincode/encc_example


```