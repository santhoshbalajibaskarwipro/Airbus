/*/*
-a-Licensed to the Apache Software Foundation (ASF) under one
or more contributor license Forms.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at
  http://www.apache.org/licenses/LICENSE-2.0 
Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

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
type Manage_po_order struct {
}

var approved_po_order_entry = "approved_po_order_entry"	
type proposal struct{
								// Attributes of a Form 
	proposal_id string `json:"proposal_id"`	
	region string `json:"region"`
	country string `json:"country"`
	proposal_type string `json:"proposal_type"`
	proposal_date string `json:"proposal_date"`
	approval_date string `json:"approval_date"`
	shared_with_procurement_team_on string `json:"shared_with_procurement_team_on"`
	approver string `json:"approver"`
	number_of_tasks_covered string `json:"number_of_tasks_covered"`
	device_qty string `json:"device_qty"`
	accessary_periperal_qty string `json:"accessary_periperal_qty"`
	total_qty string `json:"total_qty"`
	status string `json:"status"`
	
}





// ============================================================================================================================
// Main - start the chaincode for Form management
// ============================================================================================================================
func main() {			
	err := shim.Start(new(Manage_po_order))
	if err != nil {
		fmt.Printf("Error starting Form management of po order chaincode: %s", err)
	}
}





// ============================================================================================================================
// Init - reset all the things
// ============================================================================================================================
func (t *Manage_po_order) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var msg string
	var err error
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}
	// Initialize the chaincode
	msg = args[0]
	fmt.Println("Manage Po Order chaincode is deployed successfully.");
	
	// Write the state to the ledger
	err = stub.PutState("abc", []byte(msg))	//making a test var "abc", I find it handy to read/write to it right away to test the network
	if err != nil {
		return nil, err
	}
	var po_order_form_empty []string
	po_order_form_empty_json_as_bytes, _ := json.Marshal(po_order_form_empty)								//marshal an emtpy array of strings to clear the index
	err = stub.PutState(approved_po_order_entry, po_order_form_empty_json_as_bytes)
	if err != nil {
		return nil, err
	}
	return nil, nil
}




// ============================================================================================================================
// Run - Our entry Formint for Invocations - [LEGACY] obc-peer 4/25/2016
// ============================================================================================================================
func (t *Manage_po_order) Run(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("run is running " + function)
	return t.Invoke(stub, function, args)
}




// ============================================================================================================================
// Invoke - Our entry Formint for Invocations
// ============================================================================================================================
func (t *Manage_po_order) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {													//initialize the chaincode state, used as reset
		return t.Init(stub, "init", args)
	} else if function == "create_po_order_id" {											//create a new Form
		return t.create_po_order_id(stub, args)
	} 
	fmt.Println("invoke did not find func: " + function)	
	jsonResp := "Error : Received unknown function invocation: "+ function 				//error
	return nil, errors.New(jsonResp)
}



// ============================================================================================================================
// Query - Our entry for Queries
// ============================================================================================================================
func (t *Manage_po_order) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Query is running " + function)

	// Handle different functions
	if function == "get_all_po_order" {													//Read all Forms
		return t.get_all_proposal(stub, args)
	} 

	fmt.Println("query did not find func: " + function)				//error
	jsonResp := "Error : Received unknown function query: "+ function 
	return nil, errors.New(jsonResp)
}













