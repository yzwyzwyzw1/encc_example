package main

import (
	"fmt"
	"github.com/hyperledger/fabric/bccsp"
	"github.com/hyperledger/fabric/bccsp/factory"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/chaincode/shim/ext/entities"
	pb "github.com/hyperledger/fabric/protos/peer"
)

const DECKEY = "DECKEY"
const VERKEY = "VERKEY"
const ENCKEY = "SIGKEY"
const SIGKEY = "SIGKEY"
const IV = "IV"

type Encc struct {

	// 将BCCSP接口类型对象赋值给bccspInst
	bccspInst bccsp.BCCSP // BCCSP is the blockchain cryptographic service provider that offers
	                       // the implementation of cryptographic standards and algorithms.
}

func (t *Encc)Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (t * Encc)Encrypter(stub shim.ChaincodeStubInterface, args []string,encKey,IV []byte)  pb.Response {
	ent,err := entities.NewAES256EncrypterEntity("ID",t.bccspInst,encKey,IV) //创建加密实体
	if err != nil {
		return shim.Error(fmt.Sprintf("entities.NewAES256EncrypterEntity failed, err %s", err))
	}
	if len(args) != 2 {
		return shim.Error("Expected 2 parameters to function Encrypter")
	}
	key := args[0]
	cleartextValue := []byte(args[1])

	err = encryptAndPutState(stub,ent,key,cleartextValue)
	if err != nil {
		return shim.Error (fmt.Sprintf("encryptAndPutState failed err %+v",err))
	}
	return shim.Success(nil)

}

func (t *Encc)Decrypter(stub shim.ChaincodeStubInterface,args []string,decKey,IV []byte) pb.Response  {
	ent,err := entities.NewAES256EncrypterEntity("ID",t.bccspInst,decKey,IV)
	if err != nil {
		return shim.Error(fmt.Sprintf("entities.NewAES256EncrypterEntity failed, err %s",err))
	}
	if len(args) != 1 {
		return shim.Error("Expected 1 parameters to function Decrypter")
	}
	key := args[0]

	cleartextValue,err := getStateAndDecrypt(stub,ent,key)
	if err != nil {
		return shim.Error(fmt.Sprintf("getStateAndDecrypt failed, err %+v", err))
	}
	return shim.Success(cleartextValue)
}


func (t *Encc)EncrypterSigner(stub shim.ChaincodeStubInterface,args []string,encKey,sigKey []byte) pb.Response {
	ent,err := entities.NewAES256EncrypterECDSASignerEntity("ID",t.bccspInst,encKey,sigKey)
	if err != nil {
		return shim.Error(fmt.Sprintf("entities.NewAES256EncrypterEntity failed, err %s", err))
	}
	if len(args) != 2 {
		return shim.Error("Expected 2 parameters to function EncrypterSigner")
	}
	key := args[0]
	cleartextValue := []byte(args[1])

	err = signEncryptAndPutState(stub,ent,key,cleartextValue)
	if err != nil {
		return  shim.Error(fmt.Sprintf("signEncryptAndPutState failed, err %+v", err))
	}
	return shim.Success(nil)
}

func (t *Encc)DecrypterVerify(stub shim.ChaincodeStubInterface,args []string,decKey,verKey []byte) pb.Response {
	ent,err := entities.NewAES256EncrypterECDSASignerEntity("ID",t.bccspInst,decKey,verKey)
	if err != nil {
		return shim.Error(fmt.Sprintf("entities.NewAES256DecrypterEntity failed, err %s", err))
	}
	if len(args) != 1 {
		return shim.Error("Expected 1 parameters to function DecrypterVerify")
	}
	key := args[0]

	cleartextValue,err := getStateDecryptAndVerify(stub,ent,key)
	if err != nil {
		return shim.Error(fmt.Sprintf("getStateDecryptAndVerify failed, err %+v", err))
	}
	return shim.Success(cleartextValue)
}

func (t * Encc) RangeDecrypter(stub shim.ChaincodeStubInterface,decKey []byte)pb.Response {
	ent,err := entities.NewAES256EncrypterEntity("ID",t.bccspInst,decKey,nil)
	if err != nil {
		return shim.Error(fmt.Sprintf("entities.NewAES256EncrypterEntity failed, err %s", err))
	}
	bytes,err := getStateByRangeAndDecrypt(stub,ent,"","")
	if err != nil {
		return shim.Error(fmt.Sprintf("getStateByRangeAndDecrypt failed, err %+v", err))
	}
	return shim.Success(bytes)
}


func (t *Encc)Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	f,args := stub.GetFunctionAndParameters()
	tMap,err := stub.GetTransient() //GetTransient returns the `ChaincodeProposalPayload.Transient
	if err != nil {
		return shim.Error(fmt.Sprintf("Could not retrieve transient,err %s",err))//不能检索传输
	}

	switch f {
	case "ENCRYPT":

		if _,in :=tMap[ENCKEY];!in {
			return shim.Error(fmt.Sprintf("Expected transient encryption key %s",ENCKEY))
		}
		return t.Encrypter(stub,args[0:],tMap[ENCKEY],tMap[IV])
	case "DECRYPT":

		if _,in :=tMap[DECKEY];!in{
			return shim.Error(fmt.Sprintf("Expected transient decryption key %s", DECKEY))

		}
		return t.Decrypter(stub,args[0:],tMap[DECKEY],tMap[IV])
	case "ENCRYPTSIGN":
		// make sure keys are there in the transient map - the assumption is that they
		// are associated to the string "ENCKEY" and "SIGKEY"

		if _,in := tMap[ENCKEY];!in {
			return shim.Error(fmt.Sprintf("Expected transient key %s", DECKEY))
		}else  if _,in := tMap[SIGKEY];!in{
			return shim.Error(fmt.Sprintf("Expected transient key %s", SIGKEY))
		}
		return t.EncrypterSigner(stub,args[0:],tMap[ENCKEY],tMap[SIGKEY])
	case "DECRYPTVERIFY":
		if _,in := tMap[DECKEY];!in {
			return shim.Error(fmt.Sprintf("Expected transient key %s", DECKEY))
		}else if _,in := tMap[VERKEY];!in {
			return shim.Error(fmt.Sprintf("Expected transient key %s", VERKEY))
		}
		return t.DecrypterVerify(stub,args[0:],tMap[DECKEY],tMap[VERKEY])
	case "RANGEQQUERY":
		if _,in := tMap[DECKEY];!in {
			return shim.Error(fmt.Sprintf("Expected transient key %s", DECKEY))
		}
		return t.RangeDecrypter(stub,tMap[DECKEY])
	default:
		return shim.Error(fmt.Sprintf("Unsupported function %s", f))
	}
}


func main() {
	// InitFactories must be called before using factory interfaces
	// It is acceptable to call with config = nil, in which case
	// some defaults will get used
	// Error is returned only if defaultBCCSP cannot be found
	factory.InitFactories(nil)

	err := shim.Start(&Encc{factory.GetDefault()})
	if err != nil {
		fmt.Printf("Error starting EncCC chaincode: %s", err)
	}
}