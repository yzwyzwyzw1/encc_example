marble02是一个大理石资产管理的案例，介绍了如何在智能合约中定义资产，在区块链上进行资产的创建、转移、查询等操作。



# CLI

链码的主要功能是：
- （initMarble）创建一个新的大理石信息
- （transferMarble）转让大理石
- （transferMarblesBasedOnColor）根据某个颜色转让所有大理石
- （delete）删除大理石信息
- （readMarble）读取大理石信息
- （queryMarblesByOwner）根据大理石的拥有者查询大理石的信息
- （queryMarbles）基于ad hoc查询大理石信息
- （getHistoryForMarble）获取大理石的历史交易信息
- （getMarblesByRange）基于范围查询获取大理石信息
- （getMarblesByRangeWithPagination）
- （queryMarblesWithPagination）使用查询字段，页大小和书签执行查询操作。


```bash
docker exec -it cli bash

export CHANNEL_NAME=mychannel
peer channel create -o orderer.example.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/mychannel.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
peer channel join -b mychannel.block
peer chaincode install -n mycc -v 1.0 -p github.com/chaincode/marbles02

//实例化链码
peer chaincode instantiate -c '{"Args":["init"]}' -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n mycc -v 1.0  -P "OR ('Org1MSP.peer')"

//创建大理石信息
peer chaincode invoke -C $CHANNEL_NAME -n mycc -c '{"Args":["initMarble","marble1","blue","35","tom"]}' -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
peer chaincode invoke -C $CHANNEL_NAME -n mycc -c '{"Args":["initMarble","marble2","red","50","tom"]}' -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
peer chaincode invoke -C $CHANNEL_NAME -n mycc -c '{"Args":["initMarble","marble3","blue","70","tom"]}' -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
```

1. 转让大理石信息
```
peer chaincode invoke -C $CHANNEL_NAME -n mycc -c '{"Args":["transferMarble","marble2","jerry"]}' -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
```
查询安装结果
```
peer chaincode query -C $CHANNEL_NAME -n mycc -c '{"Args":["readMarble","marble2"]}' -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem

输出结果：
{"color":"red","docType":"marble","name":"marble2","owner":"jerry","size":50}
```
获取大理石的交易信息
```
peer chaincode query -C $CHANNEL_NAME -n mycc -c '{"Args":["getHistoryForMarble","marble2"]}' -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem

输出结果：
[{"TxId":"8eae83f00b4ffa171c0959a81b683b9b48b4f828dd56121ad52fdaf29c48cbe7", "Value":{"docType":"marble","name":"marble2","color":"red","size":50,"owner":"tom"}, "Timestamp":"2019-06-11 19:07:16.687743516 +0000 UTC", "IsDelete":"false"},{"TxId":"08b3f3996dc9e26ba342176d8ac2046ff3324ff589337a08e893708718bebe74", "Value":{"docType":"marble","name":"marble2","color":"red","size":50,"owner":"jerry"}, "Timestamp":"2019-06-11 19:11:52.079682154 +0000 UTC", "IsDelete":"false"}]
```
根据大理石的拥有者查询大理石的信息
```
peer chaincode query -C $CHANNEL_NAME -n mycc -c '{"Args":["queryMarblesByOwner","tom"]}'

输出结果：
[{"Key":"marble1", "Record":{"color":"blue","docType":"marble","name":"marble1","owner":"tom","size":35}},{"Key":"marble3", "Record":{"color":"blue","docType":"marble","name":"marble3","owner":"tom","size":70}}]
```

2. 根据某个颜色转让所有大理石
```
peer chaincode invoke -C $CHANNEL_NAME -n mycc -c '{"Args":["transferMarblesBasedOnColor","blue","jerry"]}' -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
```
根据大理石的拥有者查询大理石的信息
```
peer chaincode query -C $CHANNEL_NAME -n mycc -c '{"Args":["queryMarblesByOwner","jerry"]}'

输出结果：
[{"Key":"marble1", "Record":{"color":"blue","docType":"marble","name":"marble1","owner":"tom","size":35}},{"Key":"marble3", "Record":{"color":"blue","docType":"marble","name":"marble3","owner":"tom","size":70}}]
```


3. 删除大理石信息
```
peer chaincode invoke -C $CHANNEL_NAME -n mycc -c '{"Args":["delete","marble1"]}' -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
```
查询测试：
```
peer chaincode query -C $CHANNEL_NAME -n mycc -c '{"Args":["readMarble","marble1"]}'

输出结果：
Error: endorsement failure during query. response: status:500 message:"{\"Error\":\"Marble does not exist: marble1\"}"

```


3. 基于ad hoc查询大理石信息
```
peer chaincode query -C $CHANNEL_NAME -n mycc -c '{"Args":["queryMarbles","{\"selector\":{\"owner\":\"tom\"}}"]}'

输出结果：
[{"Key":"marble1", "Record":{"color":"blue","docType":"marble","name":"marble1","owner":"tom","size":35}},{"Key":"marble3", "Record":{"color":"blue","docType":"marble","name":"marble3","owner":"tom","size":70}}]
```


4. 基于范围查询marbles信息
```
peer chaincode invoke -C $CHANNEL_NAME -n mycc -c '{"Args":["initMarble","marble4","red","70","tom"]}' -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem


peer chaincode query -C $CHANNEL_NAME -n mycc -c '{"Args":["getMarblesByRange","marble1","marble4"]}' -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem

输出结果[marble1,marble4)：
[{"Key":"marble1", "Record":{"color":"blue","docType":"marble","name":"marble1","owner":"tom","size":35}},{"Key":"marble2", "Record":{"color":"red","docType":"marble","name":"marble2","owner":"tom","size":50}},{"Key":"marble3", "Record":{"color":"blue","docType":"marble","name":"marble3","owner":"tom","size":70}}]
```

5. 使用查询字段，页大小和书签执行查询操作。
```
peer chaincode query -C $CHANNEL_NAME -n mycc -c '{"Args":["queryMarblesWithPagination","{\"selector\":{\"owner\":\"tom\"}}","3",""]}'

输出结果:
[{"Key":"marble1", "Record":{"color":"blue","docType":"marble","name":"marble1","owner":"tom","size":35}},{"Key":"marble2", "Record":{"color":"red","docType":"marble","name":"marble2","owner":"tom","size":50}},{"Key":"marble3", "Record":{"color":"blue","docType":"marble","name":"marble3","owner":"tom","size":70}}][{"ResponseMetadata":{"RecordsCount":"3", "Bookmark":"g1AAAAA-eJzLYWBgYMpgSmHgKy5JLCrJTq2MT8lPzkzJBYqz5yYWJeWkGoOkOWDSyBJZABygEg8"}}]
```

CouchDB访问网址：http://127.0.0.1:5984/_utils/#login

6. 有范围的，使用查询字段，页大小和书签执行查询操作。
```
peer chaincode query -C $CHANNEL_NAME -n mycc -c '{"Args":["queryMarblesWithPagination","{\"selector\":{\"name\":\"marble1\"}}","3",""]}'

```