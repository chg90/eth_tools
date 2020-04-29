package service

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/eth_tools/transaction"
	"github.com/eth_tools/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func NewInterface() {
	myConfig := new(util.Config)
	myConfig.InitConfig("common.conf")

	url := myConfig.Read("interface_test", "url")
	to := myConfig.Read("interface_test", "to")
	fromKeyStoreFilePath := myConfig.Read("interface_test", "fromKeyStoreFilePath")
	fromPwd := myConfig.Read("interface_test", "fromPwd")

	//创建客户端
	client, err := ethclient.Dial(url)
	if err != nil {
		panic(err)
	}

	tx := transaction.SendTransaction{
		Client:               client,
		To:                   common.HexToAddress(to),
		FromKeyStoreFilePath: fromKeyStoreFilePath,
		FromPwd:              fromPwd,
		//FromPrivateKey: "",
		//EthAmount: etherAmount,
	}
	i := 0
	for {
		rand.Seed(time.Now().Unix())
		s, _ := strconv.ParseFloat(strconv.Itoa(rand.Intn(3000)), 32) //3以内的小数
		tx.EthAmount = s / 1000

		fmt.Print(i, ",")
		if err = tx.SendRawTransaction(); err != nil {
			fmt.Println(err)
		}
		i++
		time.Sleep(time.Hour * 2)
	}

}
