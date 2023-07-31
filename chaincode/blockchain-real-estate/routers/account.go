package routers

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
	"github.com/togettoyou1/blockchain-real-estate/chaincode/blockchain-real-estate/lib"
	"github.com/togettoyou1/blockchain-real-estate/chaincode/blockchain-real-estate/utils"
)

//查询账户列表
func QueryAccountList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var accountList []lib.Account
	results, err := utils.GetStateByPartialCompositeKeys(stub, lib.AccountKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var account lib.Account
			err := json.Unmarshal(v, &account)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryAccountList-反序列化出错: %s", err))
			}
			accountList = append(accountList, account)
		}
	}
	accountListByte, err := json.Marshal(accountList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryAccountList-序列化出错: %s", err))
	}
	return shim.Success(accountListByte)
}

//查询帐号密码是否正确
func QueryAccess(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	accountId := args[0]
	passWd := args[1]
	var args1 []string
	var account lib.Account
	access := &lib.Access{
		Access: false,
	}
	results, err := utils.GetStateByPartialCompositeKeys(stub, lib.AccountKey, args1)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			err := json.Unmarshal(v, &account)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryAccountList-反序列化出错: %s", err))
			}
			if account.AccountId == accountId && account.PassWd == passWd {
				access.Access = true
			}
		}
	}
	accessByte, err := json.Marshal(access)
	if err != nil {
		return shim.Error(fmt.Sprintf("queryAccess-序列化出错: %s", err))
	}
	return shim.Success(accessByte)
}
