与example04的区别在于查询获取到的A,B的值被记录了下来。

# CLI


```bash
docker exec -it cli bash

export CHANNEL_NAME=mychannel
peer channel create -o orderer.example.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/mychannel.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
peer channel join -b mychannel.block

//安装链码example02
peer chaincode install -n mycc02 -v 1.0 -p github.com/chaincode/example02
peer chaincode instantiate -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n mycc02 -v 1.0 -c '{"Args":["init","a","100","b","200"]}' -P "OR ('Org1MSP.peer')"
peer chaincode query -C $CHANNEL_NAME -n mycc02 -c '{"Args":["query","a"]}'
//测试结果：
100


//安装链码example05
peer chaincode install -n mycc05 -v 1.0 -p github.com/chaincode/example05
peer chaincode instantiate -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n mycc05 -v 1.0 -c '{"Args":["init","event","1"]}' -P "OR ('Org1MSP.peer')"

//跨链码调用
peer chaincode invoke -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n mycc05  -c '{"Args":["invoke","mycc02","c","mychannel"]}'
peer chaincode query -C $CHANNEL_NAME -n mycc05 -c '{"Args":["query","mycc02","c","mychannel"]}'

//测试结果：
300
```