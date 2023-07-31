package main

import (
	"crypto/md5"
	"fmt"
	"time"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
	"github.com/togettoyou1/blockchain-real-estate/chaincode/blockchain-real-estate/lib"
	"github.com/togettoyou1/blockchain-real-estate/chaincode/blockchain-real-estate/routers"
	"github.com/togettoyou1/blockchain-real-estate/chaincode/blockchain-real-estate/utils"
)

type BlockChainRealEstate struct {
}

//链码初始化
func (t *BlockChainRealEstate) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("链码初始化")
	timeLocal, err := time.LoadLocation("Asia/Chongqing")
	if err != nil {
		return shim.Error(fmt.Sprintf("时区设置失败%s", err))
	}
	time.Local = timeLocal
	//初始化默认数据
	var accountIds = [7]string{
		"adminA",
		"njglyy",
		"njsetyy",
		"jqzyy",
		"njyy",
	}
	var userNames = [7]string{"管理员", "南京鼓楼医院", "江苏省阳关招采平台", "江苏省不良反应监测中心", "江苏省医疗保障局"}
	var identitys = [7]string{"卫健委", "医疗机构", "招采机构", "不良反应监测中心", "医保局"}
	//定义密码及加密虚拟机中
	var rowpassWds = [7]string{"adminA", "njglyy", "njsetyy", "jqzyy", "njyy"}
	var passWds []string
	for _, pswd := range rowpassWds {
		data := []byte(pswd)
		has := md5.Sum(data)
		md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制
		passWds = append(passWds, md5str1)
	}
	//初始化账号数据
	for i, val := range accountIds {
		account := &lib.Account{
			AccountId: val,
			UserName:  userNames[i],
			Identity:  identitys[i],
			PassWd:    passWds[i],
		}
		// 写入账本
		if err := utils.WriteLedger(account, stub, lib.AccountKey, []string{val}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
	}
	return shim.Success(nil)
}

//实现Invoke接口调用智能合约
func (t *BlockChainRealEstate) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	funcName, args := stub.GetFunctionAndParameters()
	switch funcName {
	case "queryAccountList":
		return routers.QueryAccountList(stub, args)
	case "queryAccess":
		return routers.QueryAccess(stub, args)
	case "createDrug":
		return routers.CreateDrug(stub, args)
	case "queryDrugList":
		return routers.QueryDrugList(stub, args)
	case "createAdverse":
		return routers.CreateAdverse(stub, args)
	case "queryAdverseList":
		return routers.QueryAdverseList(stub, args)
	case "createMaterial":
		return routers.CreateMaterial(stub, args)
	case "queryMaterialList":
		return routers.QueryMaterialList(stub, args)
	case "applyMaterial":
		return routers.ApplyMaterial(stub, args)
	case "applyDrug":
		return routers.ApplyDrug(stub, args)
	case "applyAdverse":
		return routers.ApplyAdverse(stub, args)

	default:
		return shim.Error(fmt.Sprintf("没有该功能: %s", funcName))
	}
}

func main() {
	err := shim.Start(new(BlockChainRealEstate))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
