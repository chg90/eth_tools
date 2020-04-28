package transaction

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/big"
	"time"

	"github.com/eth_tools/util"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// SendTransaction
type SendTransaction struct {
	Client *ethclient.Client
	//From   common.Address
	To common.Address

	FromKeyStoreFilePath string
	FromPwd              string

	FromPrivateKey string
	EthAmount      float64
}

/*

 */
func (s *SendTransaction) SendRawTransaction() (err error) {
	// 交易发送方
	// 获取私钥方式一，通过keystore文件
	fromKeystore, err := ioutil.ReadFile(s.FromKeyStoreFilePath)
	if err != nil {
		return err
	}
	fromKey, err := keystore.DecryptKey(fromKeystore, s.FromPwd)
	if err != nil {
		return err
	}
	fromPrivateKey := fromKey.PrivateKey
	fromPublicKey := fromPrivateKey.PublicKey
	fromAddr := crypto.PubkeyToAddress(fromPublicKey)

	// 获取私钥方式二，通过私钥字符串
	//privateKey, err := crypto.HexToECDSA("s.FromPrivateKey")

	// 数量
	amount := util.ToWei(s.EthAmount, 18)

	// gasLimit
	var gasLimit uint64 = 300000

	// gasPrice
	gasPrice, err := s.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	// nonce获取
	nonce, err := s.Client.PendingNonceAt(context.Background(), fromAddr)
	if err != nil {
		return err
	}

	// 认证信息组装
	auth := bind.NewKeyedTransactor(fromPrivateKey)
	//auth,err := bind.NewTransactor(strings.NewReader(mykey),"111")
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = amount      // wei
	auth.GasLimit = gasLimit // in units
	auth.GasPrice = gasPrice
	auth.From = fromAddr

	// 交易创建
	tx := types.NewTransaction(nonce, s.To, amount, gasLimit, gasPrice, []byte{})
	// 签名
	signedTx, err := auth.Signer(types.HomesteadSigner{}, auth.From, tx)
	//signedTx ,err := types.SignTx(tx,types.HomesteadSigner{},fromPrivkey)

	// 交易发送
	err = s.Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return err
	}
	txhash := signedTx.Hash().Hex()
	now := time.Now()
	fmt.Printf("交易提交时间: %s ,交易hash: %s ", now.Format("2006-01-02 15:04:05"), txhash)
	// 等待挖矿完成
	h, err := bind.WaitMined(context.Background(), s.Client, signedTx)
	if err != nil {
		return err
	}
	blockNumber := h.BlockNumber
	header, err := s.Client.HeaderByNumber(context.Background(), blockNumber)
	if err != nil {
		return err
	}
	tm := time.Unix(int64(header.Time), 0)
	fmt.Printf("区块号: %v ,区块时间: %v ,用时 %v 秒\n",
		blockNumber, tm.Format("2006-01-02 15:04:05"), int64(header.Time)-now.Unix())
	return
}
