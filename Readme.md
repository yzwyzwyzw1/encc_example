[TOC]

切记：不支持fabric-1.4.4版本

此文件夹中的链码程序均为fabric源码中的example程序！更加详细的链码介绍，请看源码！

# 生成组织结构和证书文件
```
cd fixtures/
mkdir channel-artifacts
export FABRIC_CFG_PATH=$PWD
export CHANNEL_NAME=mychannel
configtxgen  -profile TwoOrgsOrdererGenesis  -outputBlock ./channel-artifacts/genesis.block
configtxgen  -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/mychannel.tx -channelID $CHANNEL_NAME
```
# 更换CA的私钥

修改docker-compose-base.yaml 中的ca.org1.example.com的CA私钥：
- crypto-config/peerOrganizations/org1.example.com/ca/f8a583bc75b8f56d083491ad7222699c1e9372d2c7ac3bb7c409eed31209dd24_sk

修改docker-compose-base.yaml 中的ca.org2.example.com的CA私钥：
- crypto-config/peerOrganizations/org2.example.com/ca/f12aced81a48fb4209c9a2ac36028e718fe9f7f9f0570b58798ae231c0d61b30_sk

# 启动网络

```
docker-compose -f docker-compose-cli.yaml up -d
```
如果需要关闭网络的话可以执行命令：
```
docker-compose -f docker-compose-cli.yaml down
```

# 其他

```
docker volume prune -f  # 清理挂载卷
docker network prune -f # 来清理没有再被任何容器引用的networks
```

```
rm -f ./fixtures/channel-artifacts/*
rm -rf ./fixtures/crypto-config
rm -rf  /var/hyperledger/production


cd fixtures
cryptogen generate  --config ./crypto-config.yaml  --output  crypto-config
export FABRIC_CFG_PATH=$PWD
export CHANNEL_NAME=mychannel
configtxgen  -profile TwoOrgsOrdererGenesis  -outputBlock ./channel-artifacts/genesis.block
configtxgen  -profile TwoOrgsChannel -outputCreateChannelTx  ./channel-artifacts/mychannel.tx  -channelID $CHANNEL_NAME
configtxgen  -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org1MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org1MSP
configtxgen  -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org2MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org2MSP

```

```
cd fixtures
cryptogen generate  --config ./crypto-config.yaml  --output  crypto-config
export FABRIC_CFG_PATH=$PWD
export CHANNEL_NAME=mychannel
configtxgen  -profile TwoOrgsOrdererGenesis  -outputBlock ./channel-artifacts/genesis.block
configtxgen  -profile TwoOrgsChannel -outputCreateChannelTx  ./channel-artifacts/mychannel.tx  -channelID $CHANNEL_NAME
configtxgen  -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org1MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org1MSP
configtxgen  -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org2MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org2MSP

```