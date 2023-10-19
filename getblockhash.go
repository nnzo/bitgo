package bitgo

func (c *Client) GetBlockHash(blockNumber int64) (string, error) {
	response, err := c.sendRequest("getblockhash", []interface{}{blockNumber})
	if err != nil {
		return "", err
	}

	blockHash := response.(string)
	return blockHash, nil
}
