
该链码的主要用于测试chaincode.executimeout属性而设计的合约，能够根据传入的sleeptime来验证超时时间。
sleeptime---。超时操作。

# CLI


```bash
docker exec -it cli bash

export CHANNEL_NAME=mychannel
peer channel create -o orderer.example.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/mychannel.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
peer channel join -b mychannel.block
peer chaincode install -n mycc -v 1.0 -p github.com/chaincode/sleeper

peer chaincode instantiate -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n mycc -v 1.0 -c '{"Args":["init","2000"]}' -P "OR ('Org1MSP.peer')"
Error: could not assemble transaction, err proposal response was not successful, error code 500, msg transaction returned with failure: Incorrect number of arguments. Expecting 1



peer chaincode query -C $CHANNEL_NAME -n mycc -c '{"Args":["query","a","10"]}'

peer chaincode invoke -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n mycc  -c '{"Args":["invoke","a","10"]}'

```