/**
 * @Author: 夜央 Oh oh oh oh oh oh (https://github.com/togettoyou1)
 * @Email: zoujh99@qq.com
 * @Date: 2020/3/4 5:14 下午
 * @Description: sdk的测试
 */
package main_test

import (
	"fmt"
	"testing"

	"github.com/togettoyou1/blockchain-real-estate/application/blockchain"
)

func TestInvoke_QueryAccountList(t *testing.T) {
	blockchain.Init()
	response, e := blockchain.ChannelQuery("queryAccountList", [][]byte{})
	if e != nil {
		fmt.Println(e.Error())
		t.FailNow()
	}
	fmt.Println(string(response.Payload))
}
