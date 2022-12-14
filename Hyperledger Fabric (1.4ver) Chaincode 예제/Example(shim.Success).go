package main

import (
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}
// Init ν¨μ
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	// π μ΄λ²μ shim.Success(nil) μ λ°λ‘ λ¦¬ν΄νμ§λ§κ³  νμΈνλ μμμ κ±°μΉλ€
	fmt.Println("ex02 Init")
	_, args := stub.GetFunctionAndParameters() // π Func & Para
	// κ΄λ¦¬μκ° μ²΄μΈμ½λλ₯Ό λ°°ν¬νλ©΄μ μ²μλΆν° key κ°, value κ°μ λ£μ΄μ λ°°ν¬μμ 
	// ν¨μ μ΄λ¦κ³Ό νλΌλ―Έν°λ₯Ό μλΌμ μ¬μ©ν μ μλλ‘ μ²λ¦¬
	var A, B string    // Entities
	var Aval, Bval int // Asset holdings
	var err error

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}
 
	// Initialize the chaincode
	A = args[0] // λ°μ΄ν°λ₯Ό μ½μ΄μμ λ°°μ΄μ λ£λλ€
	Aval, err = strconv.Atoi(args[1])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	B = args[2]
	Bval, err = strconv.Atoi(args[3])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)
	// λ°°ν¬ν λ λ€μ΄μ¨ λ°μ΄ν°λ₯Ό κ°μ§κ³  μ΄κΈ° μλμ€νμ΄νΈμ λ£μ΄μ€λ€

	// Write the state to the ledger
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
	if err != nil {
		return shim.Error(err.Error())
	}
	// π μλμ€νμ΄νΈμ A κ°κ³Ό B κ°μ κ°κ° μ μ₯

	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	// shim μΈν°νμ΄μ€λ₯Ό μ¬μ©ν  κ²½μ° Init κ³Ό κ°μ΄ β­ Invoke λ νμμμμ΄λ€

	// μ²΄μΈμ½λκ° λ°°ν¬κ° λλκ³  μ΄ν, μΉμλ²κ° λμμνκ² λλ©΄
	// μ¬μ©μκ° μμ μ λ¨λ§κΈ°(μ»΄ν¨ν°, νΈλν°, IOT λ±)λ‘ μΈν°λ·μ ν΅ν΄ νΈλμ­μμ μΌμΌν€λ©΄ 
	// νΈλμ­μ λ΄μ©μ΄ ν¨λΈλ¦­ λ€νΈμν¬λ΄μ peer μκ² μ λ¬
	// peer λ νΈλμ­μμ λ°μμ 
	// μμ μ΄ λ³΄μ νκ³  μλ μ²΄μΈμ½λ μ»¨νμ΄λμͺ½μ κ·Έ νΈλμ­μμ λ³΄λΈλ€
	// κ·Έλ¦¬κ³  μλ?¬λ μ΄μμ λλ¦Ό
	// π κ·Έλ νΈμΆλλκ² Invoke ν¨μμ΄λ€
	
	// μ¬μ©μκ° λΈλ‘μ²΄μΈμ νΈλμ­μμ μΌμΌμΌ°μλ
	// λ°μ΄ν°λ₯Ό μ¨ λ£λλ€κ±°λ μ½μ΄μ€κ±°λ μμ²­νκ±°λ λ°μ΄ν°λ₯Ό μ μ₯ν λ
	// λ°λμ νΈμΆλλκ² Invoke ν¨μμ΄λ€
	// Invoke ν¨μκ° νΈμΆλμ΄ λμνκ²λλ©΄ peerλ νΈλμ­μμ λ°μκ²μ΄λ€
	// 
	fmt.Println("ex02 Invoke")
	function, args := stub.GetFunctionAndParameters() // π Func & Para
	// λ€μ΄μ¨ λ°μ΄ν°λ₯Ό μ€λ§νΈμ»¨νΈλνΈ ν¨μ μ΄λ¦κ³Ό 
	// μ€λ§νΈμ»¨νΈλνΈκ° λμνκΈ° μν νλΌλ―Έν°λ€μ΄ λ°°μ΄νμμΌλ‘ λ€μ΄μ¨λ€.
	// ν¨μ μ΄λ¦κ³Ό νλΌλ―Έν°λ₯Ό μλΌμ μ¬μ©ν μ μλλ‘ μ²λ¦¬
	if function == "invoke" {
		// Make payment of X units from A to B
		return t.invoke(stub, args)
	} else if function == "delete" {
		// Deletes an entity from its state
		return t.delete(stub, args)
	} else if function == "query" {
		// the old "Query" is now implemtned in invoke
		return t.query(stub, args)
	}

	return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"delete\" \"query\"")
}

// Transaction makes payment of X units from A to B
func (t *SimpleChaincode) invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var A, B string    // Entities
	var Aval, Bval int // Asset holdings
	var X int          // Transaction value
	var err error

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	A = args[0]
	B = args[1]

	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if Avalbytes == nil {
		return shim.Error("Entity not found")
	}
	Aval, _ = strconv.Atoi(string(Avalbytes))

	Bvalbytes, err := stub.GetState(B)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if Bvalbytes == nil {
		return shim.Error("Entity not found")
	}
	Bval, _ = strconv.Atoi(string(Bvalbytes))

	// Perform the execution
	X, err = strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("Invalid transaction amount, expecting a integer value")
	}
	Aval = Aval - X
	Bval = Bval + X
	fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)

	// Write the state back to the ledger
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

// Deletes an entity from state
func (t *SimpleChaincode) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	A := args[0]

	// Delete the key from the state in ledger
	err := stub.DelState(A)
	if err != nil {
		return shim.Error("Failed to delete state")
	}

	return shim.Success(nil)
}

// query callback representing the query of a chaincode
func (t *SimpleChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var A string // Entities
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}

	A = args[0]

	// Get the state from the ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return shim.Success(Avalbytes)
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
