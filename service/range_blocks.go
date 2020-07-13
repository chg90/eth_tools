package service

import (
	"context"
	"fmt"
	"github.com/eth_tools/util"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func RangeBlocks() {

	myConfig := new(util.Config)
	myConfig.InitConfig("common.conf")

	url := myConfig.Read("interface_test", "url")
	client, err := ethclient.Dial(url)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	var startBlockNumber int64 = 2200000
	var endBlockNumber int64 = 2210000
	for {
		blockNumber := big.NewInt(startBlockNumber)
		startBlockNumber++
		if endBlockNumber < startBlockNumber {
			break
		}

		block, err := client.BlockByNumber(context.Background(), blockNumber)
		if err != nil {
			fmt.Println("error: ", err)
		}
		fmt.Printf("blockNumber: %v , transactionCount: %v \n", blockNumber.String(), block.Transactions().Len())
		for _, tx := range block.Transactions() {
			chainID, err := client.NetworkID(context.Background())
			if err != nil {
				fmt.Println("error: ", err)
			}
			msg, err := tx.AsMessage(types.NewEIP155Signer(chainID))
			if err != nil {
				fmt.Println("error: ", err)
			}
			fmt.Printf(" transactionHash: %v , from: %v , to: %v \n",
				tx.Hash().String(), msg.From().String(), msg.To().String())
		}
	}
}
