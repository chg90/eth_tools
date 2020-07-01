package service

import (
	"fmt"
	"time"

	"github.com/eth_tools/address"
	"github.com/eth_tools/mail"
	"github.com/eth_tools/util"
	"github.com/ethereum/go-ethereum/rpc"
)

func GetBalance() {
	myConfig := new(util.Config)
	myConfig.InitConfig("common.conf")

	url := myConfig.Read("main_net", "url")
	mailto := myConfig.Read("mailto", "to1")
	client, err := rpc.Dial(url)
	if err != nil {
		fmt.Println("错误:", err)
	}
	for {
		priv, addr := address.CreateKey()
		var reply string
		err := client.Call(&reply, "eth_getBalance", addr, "latest") //第一个是用来存放回复数据的格式，第二个是请求方法
		if err != nil {
			fmt.Println("错误:", err)
		}
		//fmt.Println(priv, addr)
		//16进制
		if reply != "0x0" {
			msg := fmt.Sprintf("地址:%v ,私钥:%v ,Wei:%v", addr, priv, reply)
			mail.SendMail([]string{mailto}, "getBalanceResult", msg)
		}
		time.Sleep(time.Millisecond * 700)
	}
	defer client.Close()

}
