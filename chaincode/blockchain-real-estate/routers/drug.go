/**
 * @Author: 夜央 Oh oh oh oh oh oh (https://github.com/togettoyou1)
 * @Email: zoujh99@qq.com
 * @Date: 2020/3/10 12:33 上午
 * @Description: 房地产相关合约路由
 */
package routers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
	"github.com/togettoyou1/blockchain-real-estate/chaincode/blockchain-real-estate/lib"
	"github.com/togettoyou1/blockchain-real-estate/chaincode/blockchain-real-estate/utils"
)

type DrugObj struct {
	DrugID   string `json:"drugId"`                 //药品ID
	Downer   string `json:"downer"`                 //数据提供者(用户AccountId)
	DrugHash uint64 `json:"unique_index:hash_idx;"` //该药品的唯一hash值
}

//创建药物
func CreateDrug(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// 验证参数
	if len(args) != 3 {
		return shim.Error("参数个数不满足")
	}
	drugID := args[0]
	downer := args[1]
	drugHash := args[2]
	if drugID == "" || downer == "" || drugHash == "" {
		return shim.Error("参数存在空值")
	}

	// 参数数据格式转换
	var formattedDrugHash uint64
	if val, err := strconv.ParseUint(drugHash, 10, 64); err != nil {
		return shim.Error(fmt.Sprintf("RealEstate参数格式转换出错: %s", err))
	} else {
		formattedDrugHash = val
	}

	resultsDowner, err := utils.GetStateByPartialCompositeKeys(stub, lib.AccountKey, []string{downer})
	if err != nil || len(resultsDowner) != 1 {
		return shim.Error(fmt.Sprintf("用户owner信息验证失败%s", err))
	}
	drug := &DrugObj{
		DrugID:   drugID,
		Downer:   downer,
		DrugHash: formattedDrugHash,
	}

	// 写入账本
	if err := utils.WriteLedger(drug, stub, lib.DrugKey, []string{drug.Downer, drug.DrugID}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	//将成功创建的信息返回
	drugByte, err := json.Marshal(drug)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}
	// 成功返回
	return shim.Success(drugByte)

}

//查询药物(可查询所有，也可根据药物名称、提供者)
func QueryDrugList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var drugList []lib.Drug
	results, err := utils.GetStateByPartialCompositeKeys2(stub, lib.DrugKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var drug lib.Drug
			err := json.Unmarshal(v, &drug)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryDrugStockList-反序列化出错: %s", err))
			}
			drugList = append(drugList, drug)
		}
	}
	drugListByte, err := json.Marshal(drugList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryDrugList-序列化出错: %s", err))
	}
	return shim.Success(drugListByte)
}

//申请药物信息
func ApplyDrug(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var MaterialList []lib.Material
	results, err := utils.GetStateByPartialCompositeKeys2(stub, lib.MaterialKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var Material lib.Material
			err := json.Unmarshal(v, &Material)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryMaterialStockList-反序列化出错: %s", err))
			}
			MaterialList = append(MaterialList, Material)
		}
	}
	MaterialListByte, err := json.Marshal(MaterialList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryMaterialList-序列化出错: %s", err))
	}
	return shim.Success(MaterialListByte)
}
