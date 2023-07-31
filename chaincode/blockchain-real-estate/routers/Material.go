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

type MaterialObj struct {
	MaterialID   string `json:"materialId"`             //药品ID
	Mowner       string `json:"mowner"`                 //数据提供者(用户AccountId)
	MaterialHash uint64 `json:"unique_index:hash_idx;"` //该药品的唯一hash值
}

//公开库存
func CreateMaterial(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// 验证参数
	if len(args) != 3 {
		return shim.Error("参数个数不满足")
	}
	MaterialID := args[0]
	Mowner := args[1]
	MaterialHash := args[2]
	if MaterialID == "" || Mowner == "" || MaterialHash == "" {
		return shim.Error("参数存在空值")
	}

	// 参数数据格式转换
	var formattedMaterialHash uint64
	if val, err := strconv.ParseUint(MaterialHash, 10, 64); err != nil {
		return shim.Error(fmt.Sprintf("RealEstate参数格式转换出错: %s", err))
	} else {
		formattedMaterialHash = val
	}

	resultsMowner, err := utils.GetStateByPartialCompositeKeys(stub, lib.AccountKey, []string{Mowner})
	if err != nil || len(resultsMowner) != 1 {
		return shim.Error(fmt.Sprintf("用户owner信息验证失败%s", err))
	}
	Material := &MaterialObj{
		MaterialID:   MaterialID,
		Mowner:       Mowner,
		MaterialHash: formattedMaterialHash,
	}

	// 写入账本
	if err := utils.WriteLedger(Material, stub, lib.MaterialKey, []string{Material.Mowner, Material.MaterialID}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	//将成功创建的信息返回
	MaterialByte, err := json.Marshal(Material)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}
	// 成功返回
	return shim.Success(MaterialByte)

}

//查询房地产(可查询所有，也可根据所有人查询名下房产)
func QueryMaterialList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
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

//申请医疗耗材信息
func ApplyMaterial(stub shim.ChaincodeStubInterface, args []string) pb.Response {
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
