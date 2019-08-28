package binance

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"gitlab.com/thorchain/bepswap/common"

	"gitlab.com/thorchain/bepswap/chain-service/config"
)

// Creating this binance client because the official go-sdk doesn't support
// these endpoints it seems

var netClient = &http.Client{
	Timeout: time.Second * 10,
}

type Binance interface {
	GetTx(txID common.TxID) (time.Time, error)
}

type BinanceClient struct {
	BaseURL string
}

// NewBinanceClient create a new instance of BinanceClient
func NewBinanceClient(cfg config.BinanceConfiguration) (*BinanceClient, error) {
	if len(cfg.DEXHost) == 0 {
		return nil, errors.New("DEXHost is empty")
	}
	return &BinanceClient{BaseURL: fmt.Sprintf("https:%s", cfg.DEXHost)}, nil
}

type httpRespGetTx struct {
	Height string `json:"height"`
}

type TxDetail struct {
	TxHash      string    `json:"txHash"`
	ToAddress   string    `json:"toAddr"`
	FromAddress string    `json:"fromAddr"`
	Timestamp   time.Time `json:"timeStamp"`
}

type httpRespGetBlock struct {
	Height int64      `json:"blockHeight"`
	Tx     []TxDetail `json:"tx"`
}

// TODO update it to get tx from binnance node as we are going to run  binane full node
func (bnb BinanceClient) GetTx(txID common.TxID) (TxDetail, error) {
	noTx := TxDetail{}
	// Rate Limit: 10 requests per IP per second.
	uri := fmt.Sprintf("%s/api/v1/tx/%s", bnb.BaseURL, txID.String())
	resp, err := netClient.Get(uri)
	if err != nil {
		return noTx, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return noTx, err
	}
	resp.Body.Close()

	var tx httpRespGetTx
	err = json.Unmarshal(body, &tx)
	if err != nil {
		return noTx, err
	}

	// Rate Limit: 60 requests per IP per minute.
	uri = fmt.Sprintf("%s/api/v1/transactions-in-block/%s", bnb.BaseURL, tx.Height)
	resp, err = netClient.Get(uri)
	if err != nil {
		return noTx, err
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return noTx, err
	}
	resp.Body.Close()

	var block httpRespGetBlock
	err = json.Unmarshal(body, &block)
	if err != nil {
		return noTx, err
	}

	for _, transaction := range block.Tx {
		if transaction.TxHash == txID.String() {
			return transaction, nil
		}
	}

	return noTx, nil
}