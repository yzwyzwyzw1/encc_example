
example04是一个跨chaincode调用的example02的案例，用于展示Fabric中跨智能合约调用场景。该chaincode主要针对example02进行A，B转账时进行跨chaincdoe的调用以及跨chaincode的查询。

# CLI


```bash
docker exec -it cli bash

export CHANNEL_NAME=mychannel
peer channel create -o orderer.example.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/mychannel.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
peer channel join -b mychannel.block

//安装链码example02
peer chaincode install -n mycc02 -v 1.0 -p github.com/chaincode/example0
peer chaincode instantiate -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n mycc02 -v 1.0 -c '{"Args":["init","a","100","b","200"]}' -P "OR ('Org1MSP.peer')"
peer chaincode query -C $CHANNEL_NAME -n mycc02 -c '{"Args":["query","a"]}'
//测试结果：
100


//安装链码example04
peer chaincode install -n mycc04 -v 1.0 -p github.com/chaincode/example04
peer chaincode instantiate -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n mycc04 -v 1.0 -c '{"Args":["init","event","1"]}' -P "OR ('Org1MSP.peer')"

//跨链码调用
peer chaincode invoke -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n mycc04  -c '{"Args":["invoke","mycc02","event","1","mychannel"]}'
peer chaincode query -C $CHANNEL_NAME -n mycc04 -c '{"Args":["query","event","mycc02","a","mychannel"]}'

//测试结果：
90
```