package main 

import (
"fmt"
"errors"
"strconv"
"github.com/hyperledger/fabric/core/chaincode/shim"
)

type SimpleChaincode struct {
}


func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args[]string) ([]byte,error) {

var A, B string
var Aval, Bval int
var err error


if len(args) != 4{
  return nil, errors.New("Incorrect number of arguments")
}

A = args[0]
Aval, err = strconv.Atoi(args[1])

if err != nil {
		return nil, errors.New("Expecting integer value for asset holding")
	}

B = args[2]
	Bval, err = strconv.Atoi(args[3])
	if err != nil {
		return nil, errors.New("Expecting integer value for asset holding")
	}
fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)

err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		return nil, err
	}

	err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
	if err != nil {
		return nil, err
	}
return nil, nil
}


func (t *SimpleChaincode) invoke(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
fmt.Printf("running invoke code")

var A,B string
var Aval,Bval int
var X int
var err error


if len(args) != 3 {
return nil, errors.New("Incorrect number of arguments. Expecting 3")
}

A = args[0]
B = args[1]

Avalbytes, err := stub.GetState(A)

if err != nil {
   return nil, errors.New("Failed to get state")
}

if Avalbytes == nil {
   return nil, errors.New("Entity not found")
}
Aval, _ = strconv.Atoi(string(Avalbytes))


Bvalbytes, err := stub.GetState(B)

if err != nil {
   return nil, errors.New("Failed to get state")
}

if Bvalbytes == nil {
   return nil, errors.New("Entity not found")
}

Bval, _ = strconv.Atoi(string(Bvalbytes))

X, err = strconv.Atoi(args[2])
Aval = Aval - X
Bval = Bval + X

fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)

err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
if err != nil {
		return nil, err
	}
err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
if err != nil{
               return nil, err
}

return nil, nil
}

func (t *SimpleChaincode) delete(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

fmt.Printf("running delete function")

if len(args) != 1 {
    return nil, errors.New("Incorrect number of arguments. Expecting 1")
}

A := args[0]

err := stub.DelState(A)

if err != nil {
		return nil, errors.New("Failed to delete state")
	}

	return nil, nil
}


// Invoke callback representing the invocation of a chaincode
// This chaincode will manage two accounts A and B and will transfer X units from A to B upon invoke
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Invoke called, determining function")
	
	// Handle different functions
	if function == "invoke" {
		// Transaction makes payment of X units from A to B
		fmt.Printf("Function is invoke")
		return t.invoke(stub, args)
	} else if function == "init" {
		fmt.Printf("Function is init")
		return t.Init(stub, function, args)
	} else if function == "delete" {
		// Deletes an entity from its state
		fmt.Printf("Function is delete")
		return t.delete(stub, args)
	}

	return nil, errors.New("Received unknown function invocation")
}

func (t* SimpleChaincode) Run(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Run called, passing through to Invoke (same function)")
	
	// Handle different functions
	if function == "invoke" {
		// Transaction makes payment of X units from A to B
		fmt.Printf("Function is invoke")
		return t.invoke(stub, args)
	} else if function == "init" {
		fmt.Printf("Function is init")
		return t.Init(stub, function, args)
	} else if function == "delete" {
		// Deletes an entity from its state
		fmt.Printf("Function is delete")
		return t.delete(stub, args)
	}

	return nil, errors.New("Received unknown function invocation")
}

// Query callback representing the query of a chaincode
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Query called, determining function")
	
	if function != "query" {
		fmt.Printf("Function is query")
		return nil, errors.New("Invalid query function name. Expecting \"query\"")
	}
	var A string // Entities
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the person to query")
	}

	A = args[0]

	// Get the state from the ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
		return nil, errors.New(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
		return nil, errors.New(jsonResp)
	}

	jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return Avalbytes, nil
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

