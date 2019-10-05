
# CLI测试

```bash
docker exec -it cli bash
export CHANNEL_NAME=mychannel
peer channel create -o orderer.example.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/mychannel.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
peer channel join -b mychannel.block
peer chaincode install -n mycc -v 1.0 -p github.com/chaincode/eventsender
peer chaincode instantiate -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n mycc -v 1.0 -c'{"Args":["init"]}' -P "OR ('Org1MSP.peer')"

peer chaincode query -C $CHANNEL_NAME -n mycc -c '{"Args":["query","noevents"]}'
{"NoEvents":"0"}

peer chaincode invoke -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n mycc  -c '{"Args":["invoke","alice"]}'
peer chaincode query -C $CHANNEL_NAME -n mycc -c '{"Args":["query","noevents"]}'
{"NoEvents":"1"}
```


分析：

效果：当被调用的这个交易被committer验证通过写入区块的时候会触发智能合约中所发送的事件。

EventSender链码将每次调用所产生的递增的数字以及传入的参数作为event的内容。evtsenter调用stub.PutState,将会在committer节点写入区块的时候触发event事件。

> 可以使用SDK中封装的方法监听所发出的event事件，也可以使用examples/events/block-listener进行监听事件。

