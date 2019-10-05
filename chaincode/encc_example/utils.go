package main

import (
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/chaincode/shim/ext/entities"
	"github.com/pkg/errors"
)

/*
type Encrypter interface {
	// Encrypt returns the ciphertext for the supplied plaintext message
	Encrypt(plaintext []byte) (ciphertext []byte, err error)

	// Decrypt returns the plaintext for the supplied ciphertext message
	Decrypt(ciphertext []byte) (plaintext []byte, err error)
}
 */
func encryptAndPutState(stub shim.ChaincodeStubInterface,ent entities.Encrypter,key string ,value []byte) error {
	ciphertext,err := ent.Encrypt(value)  //谁实现了这个接口
	if err != nil {
		return err
	}
	return stub.PutState(key,ciphertext)
}

func getStateAndDecrypt(stub shim.ChaincodeStubInterface,ent entities.Encrypter,key string) ([]byte,error) {
	ciphertext,err := stub.GetState(key)
	if err != nil {
		return nil,err
	}
	if len(ciphertext) == 0 {
		return nil,errors.New("no ciphertext to decrypt")
	}
	return ent.Decrypt(ciphertext)//返回解密原文
}

func getStateDecryptAndVerify(stub shim.ChaincodeStubInterface,ent entities.EncrypterSignerEntity,key string) ([]byte,error) {
	val,err := getStateAndDecrypt(stub,ent,key)
	if err != nil {
		return  nil,err
	}

	msg := &entities.SignedMessage{}
	err = msg.FromBytes(val)  // FromBytes populates the instance from the supplied byte array
	if err != nil {
		return nil,err
	}
	ok,err := msg.Verify(ent)
	if err != nil {
		return nil,err
	}else if !ok {
		return nil ,errors.New("invalid signature")
	}
	return msg.Payload,nil
	/*
	// Payload contains the message that is signed
	Payload []byte `json:"payload"`
	 */
}

func signEncryptAndPutState(stub shim.ChaincodeStubInterface,ent entities.EncrypterSignerEntity,key string,value []byte) error {
	msg := &entities.SignedMessage{Payload:value,ID:[]byte(ent.ID())} //得到签名消息
	err := msg.Sign(ent)
	if err != nil {
		return err
	}

	b,err := msg.ToBytes()
	if err != nil {
		return err
	}
	return encryptAndPutState(stub,ent,key,b)
}

type keyValuePair struct {
	Key string `json:"key"`
	Value string `json:"value"`
}

//getStateByRangeAndDecrypt从分类账中检索一组KVS对，并使用提供的实体解密每个值;它返回一个json编组的keyValuePair切片
func getStateByRangeAndDecrypt(stub shim.ChaincodeStubInterface,ent entities.Encrypter,startKey, endKey string ) ([]byte,error) {
	iterator,err := stub.GetStateByRange(startKey,endKey) 	// we call get state by range to go through the entire range
	if err != nil {
		return nil,err
	}
	defer iterator.Close()

	keyvalueset := []keyValuePair{}
	if iterator.HasNext() {
		el,err := iterator.Next()
		if err != nil {
			return nil,err
		}
		v,err := ent.Decrypt(el.Value)
		if err != nil {
			return nil,err
		}
		keyvalueset = append(keyvalueset,keyValuePair{el.Key,string(v)})
	}

	bytes,err := json.Marshal(keyvalueset)
	if err != nil {
		return nil,err
	}
	return bytes,nil

}