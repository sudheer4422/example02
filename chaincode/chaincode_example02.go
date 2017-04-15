package main

import (
"fmt"
"strconv"
"github.com/hyperledger/fabric/core/chaincode/shim"
"encoding/json"	
"errors"
)


type MasChaincode struct {
	
}

func main() {
	err := shim.Start(new(MasChaincode))
	if err != nil {
		
	}
}

func (t *MasChaincode) Init (stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var err error
	columnDefinations := []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name:"DistributorName",Type: shim.ColumnDefinition_STRING ,Key: true},
		&shim.ColumnDefinition{Name:"DistAddress",Type: shim.ColumnDefinition_STRING,Key: false},
		&shim.ColumnDefinition{Name:"Distcode",Type: shim.ColumnDefinition_INT64,Key: true},
		&shim.ColumnDefinition{Name:"Diststate",Type: shim.ColumnDefinition_STRING,Key: true},
	}
	
	columnDefinationOfShopKeeper := []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name:"ShopKeeperName",Type: shim.ColumnDefinition_STRING ,Key: true},
		&shim.ColumnDefinition{Name:"ShopKeeperAddress",Type: shim.ColumnDefinition_STRING,Key: false},
		&shim.ColumnDefinition{Name:"ShopKeepercode",Type: shim.ColumnDefinition_INT64,Key: true},
		&shim.ColumnDefinition{Name:"ShopKeeperState",Type: shim.ColumnDefinition_STRING,Key: true},
	}
	
	
	columnDefinationOfProducts := []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name:"ProductName", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name:"ProductId",Type: shim.ColumnDefinition_INT64, Key: true},
		&shim.ColumnDefinition{Name:"ProductType",Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name:"NumofProduct",Type: shim.ColumnDefinition_INT64, Key: true},
	}
	
	
	columnDefinationOfInvoice := []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name:"InvoiceNumber",Type:shim.ColumnDefinition_INT64,Key: true},
		&shim.ColumnDefinition{Name:"InvoiceAmount",Type:shim.ColumnDefinition_INT64,Key: true},
		&shim.ColumnDefinition{Name:"InvoiceDate",Type:shim.ColumnDefinition_INT64,Key: false},
		&shim.ColumnDefinition{Name:"InvoceState",Type:shim.ColumnDefinition_STRING,Key: false},
	}
	
	columnDefinationOfOrder := []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name:"OrderNumber",Type:shim.ColumnDefinition_INT64,Key:true},
		&shim.ColumnDefinition{Name:"ProductName",Type:shim.ColumnDefinition_STRING,Key:true},
		&shim.ColumnDefinition{Name:"ProductQuantity",Type:shim.ColumnDefinition_INT64,Key:true},
		&shim.ColumnDefinition{Name:"ProductValue",Type:shim.ColumnDefinition_INT64,Key:true},
	}
	
	
	err = stub.CreateTable("DistributerTable", columnDefinations)
	
	if err != nil {
		fmt.Println("error while creating the table")
	}
	
	err = stub.CreateTable("ShopKeeperTable", columnDefinationOfShopKeeper)
	
	if err != nil {
	 fmt.Println("error while creating the table")
	}
	err = stub.CreateTable("ProductTable", columnDefinationOfProducts)
	if err != nil {
		fmt.Println("error while creating the table")
	}
	err = stub.CreateTable("InvoiceTable", columnDefinationOfInvoice)
	if err != nil {
		fmt.Println("error while creating the table")
	}
	err = stub.CreateTable("OrderTable", columnDefinationOfOrder)
	if err != nil {
		fmt.Println("error while creating the table")
	}
	return nil, nil
}


func (t *MasChaincode) Invoke (stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	
	if function == "placeOrder" {
		return t.placeOrder(stub, args)
	}
		
	 if function == "insertProducts" {
		return t.insertProducts(stub, args)
	}	
	
	 if function == "insertNewShopKeeper" {
		return t.insertNewShopKeeper(stub, args)
	}
	
	 if function == "insertDistributor" {
		return t.insertDistributor(stub, args)
	}
	return nil, nil
}

func (t *MasChaincode) Query (stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	
	if function == "getOrderDetails" {
		return t.getOrderDetails(stub, args)
	}
	 if function == "getShopkeeperDetails" {
		return t.getShopkeeperDetails(stub, args)
	}
	 return nil, nil
}

func (t *MasChaincode) placeOrder(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	
	if len(args) != 3 {
		
	}
	
	var productName string
	var productQuantity, productValue, OrderNum int64
	
	
	OrderNumval, err := strconv.Atoi(args[3])
	
	if err != nil {
		
	}
	OrderNum = int64(OrderNumval)
	productName = args[0]
	productQuantityval, err := strconv.Atoi(args[1])
	productQuantity = int64(productQuantityval)
	
	productValueval, err := strconv.Atoi(args[2])
	productValue = int64(productValueval)
	
	Validated , err := stub.InsertRow("OrderTable", shim.Row{
			Columns: []*shim.Column {
				&shim.Column {Value: &shim.Column_String_{productName}},
				&shim.Column {Value: &shim.Column_Int64{productQuantity}},
				&shim.Column {Value: &shim.Column_Int64{productValue}},
				&shim.Column {Value: &shim.Column_Int64{OrderNum}},
			}})
	if Validated != true {
		
	}
	return nil, nil
}

func (t *MasChaincode) insertProducts(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 4 {
		
	}
	var productName, productType string
	var productId, NumofProduct int64
	productName = args[0]
	productType = args[2]
	productIdval, err := strconv.Atoi(args[1])
	if err != nil {
		
	} 
	productId = int64(productIdval)
	NumofProductval, err := strconv.Atoi(args[3])
	NumofProduct = int64(NumofProductval)
	Validated, err := stub.InsertRow("ProductTable", shim.Row {
	   Columns : []*shim.Column {
	   	&shim.Column{Value: &shim.Column_String_{String_: productName}},
	   	&shim.Column{Value: &shim.Column_Int64{Int64: productId}},
	   	&shim.Column{Value: &shim.Column_String_{String_: productType}},
	   	&shim.Column{Value: &shim.Column_Int64{Int64: NumofProduct}},
	   }})
	  if Validated != true {
      	
      } 
	return nil, nil
} 


func (t *MasChaincode) insertNewShopKeeper (stub shim.ChaincodeStubInterface, args []string)  ([]byte, error) {
	if len(args) != 4 {
		
	}
	var ShopKeeperName, ShopKeeperAddress ,ShopKeeperState string
	var ShopKeepercode int64
	var err error 
	ShopKeeperName = args[0]
	ShopKeeperAddress = args[1]
	ShopKeepercodeval, err := strconv.Atoi(args[2])
	if err != nil {
		
	}
	ShopKeepercode = int64(ShopKeepercodeval)
	ShopKeeperState = args[3]	
   	Validated, err := stub.InsertRow("ShopKeeperTable", shim.Row {
      	Columns : []*shim.Column{
      		&shim.Column{Value: &shim.Column_String_{String_: ShopKeeperName}},
      		&shim.Column{Value: &shim.Column_String_{String_: ShopKeeperAddress}},
      		&shim.Column{Value: &shim.Column_Int64{Int64: ShopKeepercode}},
      		&shim.Column{Value: &shim.Column_String_{String_: ShopKeeperState}},
      	}})
      if Validated != true {
      	
      } 
      return nil, nil    
}
   		
   		
   		
   		
func (t *MasChaincode) insertDistributor (stub shim.ChaincodeStubInterface, args []string)  ([]byte, error) {
	
	if len(args) != 4 {
		
	}
	
	 var DistributorName, DistAddress, Diststate string
	 var Distcode int64
      
      DistributorName = args[0]
      DistAddress = args[1]
      Diststate = args[3]
      
      Distcodeval, err := strconv.Atoi(args[2])
      if err != nil {
      	
      }
      Distcode = int64(Distcodeval)
      
      Validated, err := stub.InsertRow("DistributerTable", shim.Row {
      	Columns : []*shim.Column{
      		&shim.Column{Value: &shim.Column_String_{String_: DistributorName}},
      		&shim.Column{Value: &shim.Column_String_{String_: DistAddress}},
      		&shim.Column{Value: &shim.Column_Int64{Int64: Distcode}},
      		&shim.Column{Value: &shim.Column_String_{String_: Diststate}},
      	}})
      if Validated != true {
      	
      } 
      return nil, nil    
}

func (t *MasChaincode) getShopkeeperDetails(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var columns []shim.Column
	if len(args) != 1{
		
	}
	
	     colint, err := strconv.ParseInt(args[0], 10, 64)
	     if err != nil {
				return nil, errors.New("getRowsTableTwo failed. arg[1] must be convertable to int32")
			}
	     ShopKeepercode := int64(colint)
	     col2 := shim.Column{Value: &shim.Column_Int64{Int64: ShopKeepercode}}
	     columns = append(columns, col2)
	     
	   rowChannel, err :=  stub.GetRows("ShopKeeperTable", columns)
	   
	   if err != nil {
	   	
	   }
	  var rows []shim.Row
	   
	   for {
	   	select {
	   		case row, ok := <-rowChannel:
	   		if !ok {
  				rowChannel = nil
  			}else{
  				rows = append(rows,row)
  			}
  			if rowChannel == nil {
				break
			}
	   	}
	   }
	   
	   jsonRows, err := json.Marshal(rows)
	   if err != nil {
  			
  		}
  	return jsonRows, nil	
}

func (t *MasChaincode) getOrderDetails(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    
    var columns []shim.Column
    
       col2Int, err :=  strconv.ParseInt(args[0], 10, 64)
    
    if len(args) != 1{
    	
    }
   
    orderNum := int64(col2Int)
  	col2 := shim.Column {Value: &shim.Column_Int64{Int64: orderNum}}
  	columns = append(columns, col2)
  	rowChannel, err :=stub.GetRows("OrderTable", columns)
  	if err != nil {
			return nil, fmt.Errorf("getRowsTableFour operation failed. %s", err)
		}
  	 var rows []shim.Row
  	for{
  		select {
  			case row, ok := <-rowChannel:
  			if !ok {
  				rowChannel = nil
  			}else{
  				rows = append(rows,row)
  			}
  		}
  		if rowChannel == nil {
				break
			}
  	}
  	
  	jsonRows, err := json.Marshal(rows)
  		
  		if err != nil {
  			
  		}
  	return jsonRows, nil
}
