/*
2020.11.30
*/
package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	bc "github.com/togettoyou1/blockchain-real-estate/application/blockchain"
	"github.com/togettoyou1/blockchain-real-estate/application/pkg/app"
)

type AccessRequestBody struct {
	AccountId string `json:"accountId"` //账户ID
	PassWd    string `json:"passWd"`    //密码
}

func QueryAccess(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(AccessRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	if body.AccountId == "" {
		appG.Response(http.StatusBadRequest, "失败", "帐号不能为空")
		return
	}
	if body.PassWd == "" {
		appG.Response(http.StatusBadRequest, "失败", "密码不能为空")
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.AccountId))
	bodyBytes = append(bodyBytes, []byte(body.PassWd))
	//调用智能合约
	resp, err := bc.ChannelQuery("queryAccess", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	//反序列化json
	var data map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}

	appG.Response(http.StatusOK, "成功", data)
}
