package bitgo

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Client struct {
	URL      string
	Username string
	Password string
}

type RPCRequest struct {
	Jsonrpc string        `json:"jsonrpc"`
	ID      string        `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

func NewClient(url, username, password string) *Client {
	return &Client{
		URL:      url,
		Username: username,
		Password: password,
	}
}

func (c *Client) sendRequest(method string, params ...interface{}) ([]byte, error) {
	request := RPCRequest{
		Jsonrpc: "1.0",
		ID:      "curltest",
		Method:  method,
		Params:  params,
	}

	payload, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.URL, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.Username, c.Password)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (c *Client) GetBlockCount() (int64, error) {
	response, err := c.sendRequest("getblockcount")
	if err != nil {
		return 0, err
	}

	var result int64
	err = json.Unmarshal(response, &result)
	return result, err
}

func (c *Client) GetBlockHash(blockHeight int64) (string, error) {
	response, err := c.sendRequest("getblockhash", blockHeight)
	if err != nil {
		return "", err
	}

	var result string
	err = json.Unmarshal(response, &result)
	return result, err
}

func (c *Client) GetBlock(blockHash string) (map[string]interface{}, error) {
	response, err := c.sendRequest("getblock", blockHash)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(response, &result)
	return result, err
}

func (c *Client) GetRawTransaction(txid string) (string, error) {
	response, err := c.sendRequest("getrawtransaction", txid)
	if err != nil {
		return "", err
	}

	var result string
	err = json.Unmarshal(response, &result)
	return result, err
}

func (c *Client) DecodeRawTransaction(hexString string) (map[string]interface{}, error) {
	response, err := c.sendRequest("decoderawtransaction", hexString)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(response, &result)
	return result, err
}
