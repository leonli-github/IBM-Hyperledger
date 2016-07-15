package main
 
import(
      "errors"
      "fmt"
     //"encoding/json"
      "github.com/hyperledger/fabric/core/chaincode/shim"
      //"strings"
      "strconv"
)


type Chaincode struct{
     
}

// init
func (t *Chaincode) Init(stub *shim.ChaincodeStub, function string,args []string)([]byte,error){
    
   return nil,nil
}

func(t *Chaincode) Invoke(stub *shim.ChaincodeStub, function string,args []string)([]byte,error){
     
     //method := args[0]
     
     fmt.Println("Step1 : get args")
     voter := args[0]
     item:=args[1]
     
    //test
    
    //todo : check if v is valid
    err := stub.PutState(voter,[]byte(item))
    fmt.Println("Step2 : put vote")
    if err != nil{
    	 return nil,err
    }

    bcount,err := stub.GetState(item)
    if err != nil {

        stub.PutState(item,[]byte(strconv.Itoa(1)))
     

    } else{
    count,_:=strconv.Atoi(string(bcount))
    count =count + 1
    stub.PutState(item,[]byte(strconv.Itoa(count)))
    fmt.Println("Step3 : put stat")
    }

    return nil,nil
     
}

func(t *Chaincode) Query(stub *shim.ChaincodeStub, function string,args []string)([]byte,error){
    
      if len(args) !=1 {
     	return nil, errors.New("Invalid Input")
     } 
      

	A:=args[0]
	key := A
	value,err := stub.GetState(key)
	if err != nil {
	    return nil,errors.New("Can't find Voter")	
	}
	jsonresponse := "{\"Vote_to_count\":\"" + string(value) +"\"}"
	 
	return []byte(jsonresponse),nil
   
   
}
func main(){
   err := shim.Start(new(Chaincode))
   if err !=nil{
   	fmt.Printf("Error Starting cc_voting: %s",err)
   }
}
