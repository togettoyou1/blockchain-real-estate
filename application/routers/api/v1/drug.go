/**
 * @Author: CPU
 * @Email: cai_sner@163.com
 * @Date: 2021/10/07 15:41
 * @Description: 药品库存接口
 */
package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spaolacci/murmur3"
	bc "github.com/togettoyou1/blockchain-real-estate/application/blockchain"
	"github.com/togettoyou1/blockchain-real-estate/application/pkg/app"
	"github.com/togettoyou1/blockchain-real-estate/application/routers/tools"
	"github.com/togettoyou1/blockchain-real-estate/chaincode/blockchain-real-estate/lib"
)

//从前端接收到的数据结构
type DrugBody struct {
	YPID               string  `json:"ypid"`
	Downer             string  `json:"downer"`             //数据提供者(用户AccountId)
	DName              string  `json:"dname"`              //产品通用名
	DtradeName         string  `json:"dtradename"`         //商品名
	DosageForm         string  `json:"dosageform"`         //剂型
	DosageClass        string  `json:"dosageclass"`        //剂型分类名称
	DosageCode         string  `json:"dosagecode"`         //剂型分类码
	Specifications     string  `json:"specifications"`     //规格
	SpecificationsCode string  `json:"specificationscode"` //规格分类码
	PharmaClass        string  `json:"pharmaclass`         //药理、药效分类
	PharmaCode         string  `json:"pharmacode`          //药理、药效分类编码
	ConvertRate        int64   `json:"convertrate`         //转换比
	PackageMaterial    string  `json:"packagematerial"`    // 包装材质
	MiniPackageUnit    string  `json:"minipackageunit`     //最小包装单位
	Conpany            string  `json:"conpany"`            //生产企业
	USCC               string  `json:"uscc"`               //统一社会信用代码
	City               string  `json:"city"`               //地市
	ConpanyCode        string  `json:"conpanycode"`        //企业编码
	ApprovalNum        string  `json:"approvalnum"`        //批准文号
	Hospital           string  `json:"hospital"`           //医疗机构
	Rank               string  `json:"rank"`               //医院等级
	Number             int64   `json:"number"`             //采购数量
	Amount             float64 `json:"amount"`             //采购金额
	InsureClass        string  `json:"insureclass"`        //医保甲乙类
	BasicDrug          string  `json:"basicdrug"`          //基药
	Year               int64   `json:"year"`               //采购年份
}

type DrugQueryBody struct {
	downer         string `json:"downer"` //所有者(用户)(用户AccountId)
	YPID           string `json:"ypid"`
	DName          string `json:"dname"`          //产品通用名
	Specifications string `json:"specifications"` //规格
	Year           int64  `json:"year"`           //采购年份
}

// @Summary 公开库存
func CreateDrug(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(DrugBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	if body.Number <= 0 {
		appG.Response(http.StatusBadRequest, "失败", "采购数量必须大于0")
		return
	}
	if body.Amount <= 0 {
		appG.Response(http.StatusBadRequest, "失败", "采购金额必须大于0")
		return
	}
	if body.Year == 0 {
		appG.Response(http.StatusBadRequest, "失败", "采购年份不能为空")
		return
	}
	if body.DName == "" {
		appG.Response(http.StatusBadRequest, "失败", "药品名称不能为空")
		return
	}
	if body.YPID == "" {
		appG.Response(http.StatusBadRequest, "失败", "YPID不能为空")
	}

	//操作数据库
	db := tools.GetDB()
	//创建表
	if !db.HasTable(&lib.Account{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&lib.Drug{}).Error; err != nil {
			panic(err)
		}
	}

	var DrugID string
	DrugID = fmt.Sprintf("%d", time.Now().Local().UnixNano())

	var DrugHash uint64
	DrugHash = murmur3.Sum64([]byte(strings.Join([]string{DrugID, body.YPID, body.Downer, body.DName, body.DtradeName, body.Conpany, body.Hospital, body.DosageClass, body.DosageForm}, "-"))) >> 1

	//插入数据
	drug := &lib.Drug{
		DrugID:             DrugID,
		YPID:               body.YPID,
		Downer:             body.Downer,
		DName:              body.DName,
		DtradeName:         body.DtradeName,
		DosageForm:         body.DosageForm,
		DosageClass:        body.DosageClass,
		DosageCode:         body.DosageCode,
		Specifications:     body.Specifications,
		SpecificationsCode: body.SpecificationsCode,
		PharmaClass:        body.PharmaClass,
		PharmaCode:         body.PharmaCode,
		ConvertRate:        body.ConvertRate,
		PackageMaterial:    body.PackageMaterial,
		MiniPackageUnit:    body.MiniPackageUnit,
		Conpany:            body.Conpany,
		USCC:               body.USCC,
		City:               body.City,
		ConpanyCode:        body.ConpanyCode,
		ApprovalNum:        body.ApprovalNum,
		Hospital:           body.Hospital,
		Rank:               body.Rank,
		Number:             body.Number,
		Amount:             body.Amount,
		InsureClass:        body.InsureClass,
		BasicDrug:          body.BasicDrug,
		Year:               body.Year,
		DrugHash:           DrugHash,
	}
	if err := db.Create(drug).Error; err != nil {
		fmt.Println(err)
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(DrugID))
	bodyBytes = append(bodyBytes, []byte(body.Downer))
	bodyBytes = append(bodyBytes, []byte(strconv.FormatUint(DrugHash, 10)))

	//调用智能合约
	resp, err := bc.ChannelExecute("createDrug", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	var data map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

// @Summary 获取库存信息(空json{}可以查询所有，指定proprietor可以查询指定用户名下库存)
// @Param drugStockQuery body DrugStockQueryRequestBody true "drugStockQuery"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/queryDrugStockList [post]
func QueryDrugList(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(DrugQueryBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	//操作数据库
	db := tools.GetDB()
	var data []lib.Drug
	if body.downer == "" {
		db.Find(&data)
	} else {
		db.Where("downer = ? and drug_name = ? and stock_year >= ? ", body.downer, body.DName, body.Year).Find(&data)
		// db.Find(&data, "proprietor = ? and drug_id = ? and stock_number >= ?", body.Proprietor, body.DrugID, body.Number)
	}

	appG.Response(http.StatusOK, "成功", data)
}
