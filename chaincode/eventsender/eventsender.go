package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"

	"strconv"
)
type EventSender struct {

}

func (t *EventSender)Init(stub shim.ChaincodeStubInterface)pb.Response {
	err := stub.PutState("noevents",[]byte("0"))
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (t *EventSender)query(stub shim.ChaincodeStubInterface,args []string) pb.Response {
	b,err := stub.GetState("noevents")
	if err != nil {
		return shim.Error("Failed to get state")
	}
	jsonResp := "{\"NoEvents\":\""+string(b)+"\"}"
	return shim.Success([]byte(jsonResp))
}

func (t *EventSender)invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	b,err := stub.GetState("noevents")
	if err != nil {
		return shim.Error("Failed to get state")
	}
	noevts,_ := strconv.Atoi(string(b))//获取noevents的整型值

	tosend := "Event" + string(b)
	for _, s := range args {
		tosend = tosend + "," + s
	}
	err = stub.PutState("noevents", []byte(strconv.Itoa(noevts+1)))//对noevents的值加一
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.SetEvent("evtsender", []byte(tosend))//此处evtsenter的的作用是记录事件的发送者
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)

}
func (t *EventSender)Invoke(stub shim.ChaincodeStubInterface)pb.Response {
	function,args := stub.GetFunctionAndParameters()
	if function == "invoke" {
		return t.invoke(stub,args)
	}else if function == "query" {
		return t.query(stub,args)
	}
	return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"query\"")
}

func main(){
	err := shim.Start(new(EventSender))
	if err != nil {
		fmt.Printf("Error starting EventSender chaincode: %s", err)
	}
}