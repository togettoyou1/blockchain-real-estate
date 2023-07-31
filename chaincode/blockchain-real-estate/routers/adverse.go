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

type AdverseObj struct {
	AdverseID   string `json:"adverseId"`              //药品ID
	Aowner      string `json:"aowner"`                 //数据提供者(用户AccountId)
	AdverseHash uint64 `json:"unique_index:hash_idx;"` //该药品的唯一hash值
}

//创建不良反应
func CreateAdverse(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// 验证参数
	if len(args) != 3 {
		return shim.Error("参数个数不满足")
	}
	adverseID := args[0]
	aowner := args[1]
	adverseHash := args[2]
	if adverseID == "" || aowner == "" || adverseHash == "" {
		return shim.Error("参数存在空值")
	}

	// 参数数据格式转换
	var formattedAdverseHash uint64
	if val, err := strconv.ParseUint(adverseHash, 10, 64); err != nil {
		return shim.Error(fmt.Sprintf("RealEstate参数格式转换出错: %s", err))
	} else {
		formattedAdverseHash = val
	}

	resultsAowner, err := utils.GetStateByPartialCompositeKeys(stub, lib.AccountKey, []string{aowner})
	if err != nil || len(resultsAowner) != 1 {
		return shim.Error(fmt.Sprintf("用户owner信息验证失败%s", err))
	}
	adverse := &AdverseObj{
		AdverseID:   adverseID,
		Aowner:      aowner,
		AdverseHash: formattedAdverseHash,
	}

	// 写入账本
	if err := utils.WriteLedger(adverse, stub, lib.AdverseKey, []string{adverse.Aowner, adverse.AdverseID}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	//将成功创建的信息返回
	AdverseByte, err := json.Marshal(adverse)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}
	// 成功返回
	return shim.Success(AdverseByte)

}

//查询不良反应(可查询所有，也可根据药物名称、年份查询)
func QueryAdverseList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var adverseList []lib.Adverse
	results, err := utils.GetStateByPartialCompositeKeys2(stub, lib.AdverseKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var Adverse lib.Adverse
			err := json.Unmarshal(v, &Adverse)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryAdverseStockList-反序列化出错: %s", err))
			}
			adverseList = append(adverseList, Adverse)
		}
	}
	adverseListByte, err := json.Marshal(adverseList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryAdverseList-序列化出错: %s", err))
	}
	return shim.Success(adverseListByte)
}

//申请不良反应信息
func ApplyAdverse(stub shim.ChaincodeStubInterface, args []string) pb.Response {
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
