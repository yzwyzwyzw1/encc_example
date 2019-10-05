package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"strconv"
	"time"
)

type SleeperChaincode struct {
}


// Init initializes chaincode...all it does is sleep a bi
func (t *SleeperChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	args := stub.GetStringArgs()

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	sleepTime := args[0] //获取睡眠时间

	t.sleep(sleepTime)   //睡眠sleepTime时间长度

	return shim.Success(nil)
}

//链码睡眠函数
func (t *SleeperChaincode) sleep(sleepTime string) {
	st, _ := strconv.Atoi(sleepTime)
	if st >= 0 {
		time.Sleep(time.Duration(st) * time.Millisecond)
	}
}

// Transaction makes payment of X units from A to B
func (t *SleeperChaincode) invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// set state
	key := args[0]
	val := args[1]

	err := stub.PutState(key, []byte(val))
	if err != nil {
		return shim.Error(err.Error())
	}

	sleepTime := args[2]

	//sleep for a bit
	t.sleep(sleepTime)

	return shim.Success([]byte("OK"))
}

// query callback representing the query of a chaincode
func (t *SleeperChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	key := args[0]

	// Get the state from the ledger
	val, err := stub.GetState(key)
	if err != nil {
		return shim.Error(err.Error())
	}

	sleepTime := args[1]

	//sleep for a bit
	t.sleep(sleepTime)

	return shim.Success(val)
}


// Invoke sets key/value and sleeps a bit
func (t *SleeperChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	if function == "put" {
		if len(args) != 3 {
			return shim.Error("Incorrect number of arguments. Expecting 3")
		}

		// Make payment of X units from A to B
		return t.invoke(stub, args)
	} else if function == "get" {
		if len(args) != 2 {
			return shim.Error("Incorrect number of arguments. Expecting 2")
		}

		// the old "Query" is now implemtned in invoke
		return t.query(stub, args)
	}

	return shim.Error("Invalid invoke function name. Expecting \"put\" or \"get\"")
}

func main() {
	err := shim.Start(new(SleeperChaincode))
	if err != nil {
		fmt.Printf("Error starting Sleeper chaincode: %s", err)
	}
}
