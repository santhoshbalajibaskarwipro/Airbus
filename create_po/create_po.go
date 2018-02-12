/***** Project ChainDaaS - Phase0 - SRC - UI1 for Creation Purchase Order ***/
/*** Last updated on 29th Nov 2017  by Santhosh Balaji ***/


package main

import (    
"errors"
"fmt"
"strconv"
"encoding/json"
"github.com/hyperledger/fabric/core/chaincode/shim"
//"github.com/hyperledger/fabric/core/util"
)

// Proposal example simple Chaincode implementation
type ManagePurchaseOrder struct {
}

var approved_purchase_order_entry = "approved_purchase_order_entry"				//name for the key/value that will store a list of all known  Tier3 Form

type purchase_order struct{ 
								// Attributes of a Purchase Form 
								
								
	 
	
	unique_proposal_purchase_id string `json:"unique_proposal_purchase_id"`								
	ProposalID string `json:"ProposalID"`	
	Location string `json:"Location"`	
	City string `json:"City"`
	PartNumber string `json:"PartNumber_code"`
	PartDescription string `json:"PartDescription"`
	WiproOrderReference string `json:"WiproOrderReference"`
	Vendor string `json:"Vendor"`
	VendorSO string `json:"VendorSO"`
	OrderDate string `json:"OrderDate"`
	OrderedQuantity string `json:"OrderedQuantity"`
	Status string `json:"Status"`

	
	
	
	
	
	
}
// ============================================================================================================================
// Main - start the chaincode for Form management
// ============================================================================================================================
func main() {			
	err := shim.Start(new(ManagePurchaseOrder))
	if err != nil {
		fmt.Printf("Error starting Form Purchase order chaincode: %s", err)
	}
}
// ============================================================================================================================
// Init - reset all the things
// ============================================================================================================================
func (t *ManagePurchaseOrder) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var msg string
	var err error
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}
	// Initialize the chaincode
	msg = args[0]
	fmt.Println("ManagePurchaseOrder chaincode is deployed successfully.");
	
	// Write the state to the ledger
	err = stub.PutState("abc", []byte(msg))	//making a test var "abc", I find it handy to read/write to it right away to test the network
	if err != nil {
		return nil, err
	}
	var purchase_order_form_empty []string
	purchase_order_form_empty_json_as_bytes, _ := json.Marshal(purchase_order_form_empty)								//marshal an emtpy array of strings to clear the index
	err = stub.PutState(approved_purchase_order_entry, purchase_order_form_empty_json_as_bytes)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
// ============================================================================================================================
// Run - Our entry Formint for Invocations - [LEGACY] obc-peer 4/25/2016
// ============================================================================================================================
func (t *ManagePurchaseOrder) Run(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("run is running " + function)
	return t.Invoke(stub, function, args)
}
// ============================================================================================================================
// Invoke - Our entry Formint for Invocations
// ============================================================================================================================
func (t *ManagePurchaseOrder) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {													//initialize the chaincode state, used as reset
		return t.Init(stub, "init", args)
	} else if function == "create_purchase_order_id" {											//create a new Form
		return t.create_purchase_order_id(stub, args)
	} 
	fmt.Println("invoke did not find func: " + function)	
	jsonResp := "Error : Received unknown function invocation: "+ function 				//error
	return nil, errors.New(jsonResp)
}

// ============================================================================================================================
// Query - Our entry for Queries
// ============================================================================================================================
func (t *ManagePurchaseOrder) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Query is running " + function)

	// Handle different functions
	if function == "get_all_purchase_order_data" {													//Read all Forms
		return t.get_all_purchase_order_data(stub, args)
	} else if function == "get_all_purchase_order_id" {													//Read all Forms
		return t.get_all_purchase_order_id(stub, args)
	  }
	fmt.Println("query did not find func: " + function)				//error
	jsonResp := "Error : Received unknown function query: "+ function 
	return nil, errors.New(jsonResp)
}


// ============================================================================================================================
// create Form - create a new Form for purchase id, store into chaincode state
// ============================================================================================================================
func (t *ManagePurchaseOrder) create_purchase_order_id(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error
	if len(args) != 12 {
		return nil, errors.New("Incorrect number of arguments. Expecting 12 ")
	}
	fmt.Println("Creating a new Form for proposal id ")
	if len(args[0]) <= 0 {
		return nil, errors.New("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return nil, errors.New("2nd argument must be a non-empty string")
	}
	
	
	
	
	
	
	
	
	
	unique_proposal_purchase_id :=args[0]
	ProposalID := args[1]
	Location := args[2]
	City := args[3]
	PartNumber := args[4]
	PartDescription := args[5]
	WiproOrderReference := args[6]
	Vendor := args[7]
	VendorSO := args[8]
	OrderDate :=args[9]
	OrderedQuantity :=args[10]
	Status := args[11]
	
	
	//build the Form json string manually
	input := 	`{`+
		`"unique_proposal_purchase_id": "` + unique_proposal_purchase_id + `" , `+
		`"ProposalID": "` + ProposalID + `" , `+
		`"Location": "` + Location + `" , `+ 
		`"City": "` + City + `" , `+ 
		`"PartNumber": "` + PartNumber + `" , `+ 
		`"PartDescription": "` + PartDescription + `" , `+ 
		`"WiproOrderReference": "` + WiproOrderReference + `" , `+
		`"Vendor": "` + Vendor + `" , `+ 
		`"VendorSO": "` + VendorSO + `" , `+ 
		`"OrderDate": "` + OrderDate + `" , `+
		`"OrderedQuantity": "` + OrderedQuantity + `" , `+
		`"Status": "` + Status + `"` +	
		`}`
	
		fmt.Println("input: " + input)
		fmt.Print("input in bytes array: ")
		fmt.Println([]byte(input))
	err = stub.PutState(unique_proposal_purchase_id, []byte(input))					//store Form with unique_proposal_purchase_id as key
	if err != nil {
		return nil, err
	}
	

	
	purchase_order_id_FormIndexAsBytes, err := stub.GetState(approved_purchase_order_entry)
	if err != nil {
		return nil, errors.New("Failed to get proposal id  Form index")
	}
	var purchase_order_id_FormIndex []string
	fmt.Print("purchase_order_id_FormIndexAsBytes: ")
	fmt.Println(purchase_order_id_FormIndexAsBytes)
	
	json.Unmarshal(purchase_order_id_FormIndexAsBytes, &purchase_order_id_FormIndex)							//un stringify it aka JSON.parse()
	fmt.Print("purchase_order_id_FormIndex after unmarshal..before append: ")
	fmt.Println(purchase_order_id_FormIndex)
	//append
	purchase_order_id_FormIndex = append(purchase_order_id_FormIndex, unique_proposal_purchase_id)									//add Form transID to index list
	fmt.Println("! Purchase Order  Form index after appending po id: ", purchase_order_id_FormIndex)
	jsonAsBytes, _ := json.Marshal(purchase_order_id_FormIndex)
	fmt.Print("jsonAsBytes: ")
	fmt.Println(jsonAsBytes)
	err = stub.PutState(approved_purchase_order_entry, jsonAsBytes)						//store name of Form
	if err != nil {
		return nil, err
	}

	fmt.Println("Purchase Order Form created successfully.")
	return nil, nil
	
	
}



func (t *ManagePurchaseOrder) get_all_purchase_order_data(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	
	var jsonPurchaseResp,errResp string
	var purchase_order_id_FormIndex []string
	fmt.Println("Fetching All Purchase Order")
	var err error
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting single space as an argument")
	}
	// fetching all Purchase ORder
	purchase_order_id_FormIndexAsBytes, err := stub.GetState(approved_purchase_order_entry)
	if err != nil {
		return nil, errors.New("Failed to get all Purchase Order")
	}
	fmt.Print("purchase_order_id_FormIndexAsBytes : ")
	fmt.Println(purchase_order_id_FormIndexAsBytes)
	json.Unmarshal(purchase_order_id_FormIndexAsBytes, &purchase_order_id_FormIndex)								//un stringify it aka JSON.parse()
	fmt.Print("purchase_order_id_FormIndex : ")
	fmt.Println(purchase_order_id_FormIndex)
	// Purchase Order Data
	jsonPurchaseResp = "{"
	for i,val := range purchase_order_id_FormIndex{
		fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Purchase")
		valueAsBytes, err := stub.GetState(val)
		if err != nil {
			errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
			return nil, errors.New(errResp)
		}
		fmt.Print("valueAsBytes : ")
		fmt.Println(valueAsBytes)
		jsonPurchaseResp = jsonPurchaseResp + "\""+ val + "\":" + string(valueAsBytes[:])
		if i < len(purchase_order_id_FormIndex)-1 {
			jsonPurchaseResp = jsonPurchaseResp + ","
		}
	}
	fmt.Println("len(purchase_order_id_FormIndex) : ")
	fmt.Println(len(purchase_order_id_FormIndex))

	jsonPurchaseResp = jsonPurchaseResp + "}"
	fmt.Println([]byte(jsonPurchaseResp))
	fmt.Println("Fetched All Proposal Data successfully.")
	return []byte(jsonPurchaseResp), nil
}


func (t *ManagePurchaseOrder) get_all_purchase_order_id(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	
	var jsonPurchaseResp,errResp string
	var purchase_order_id_FormIndex []string
	fmt.Println("Fetching All Purchase Order")
	var err error
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting single space as an argument")
	}
	// fetching all Purchase 
	purchase_order_id_FormIndexAsBytes, err := stub.GetState(approved_purchase_order_entry)
	if err != nil {
		return nil, errors.New("Failed to get all Purchase Orders")
	}
	fmt.Print("purchase_order_id_FormIndexAsBytes : ")
	fmt.Println(purchase_order_id_FormIndexAsBytes)
	json.Unmarshal(purchase_order_id_FormIndexAsBytes, &purchase_order_id_FormIndex)								//un stringify it aka JSON.parse()
	fmt.Print("purchase_order_id_FormIndex : ")
	fmt.Println(purchase_order_id_FormIndex)
	//  Proposal Data
	jsonPurchaseResp = "{"
	for i,val := range purchase_order_id_FormIndex{
		fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Purchase Order")
		valueAsBytes, err := stub.GetState(val)
		if err != nil {
			errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
			return nil, errors.New(errResp)
		}
		fmt.Print("valueAsBytes : ")
		fmt.Println(valueAsBytes)
		jsonPurchaseResp = jsonPurchaseResp + "\""+ val + "\""
		if i < len(purchase_order_id_FormIndex)-1 {
			jsonPurchaseResp = jsonPurchaseResp + ","
		}
	}
	fmt.Println("len(purchase_order_id_FormIndex) : ")
	fmt.Println(len(purchase_order_id_FormIndex))

	jsonPurchaseResp = jsonPurchaseResp + "}"
	fmt.Println([]byte(jsonPurchaseResp))
	fmt.Println("Fetched All PO ID successfully.")
	return []byte(jsonPurchaseResp), nil
}

