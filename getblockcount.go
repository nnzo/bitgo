package bitgo

func (c *Client) GetBlockCount() (int64, error) {
	response, err := c.sendRequest("getblockcount", nil)
	if err != nil {
		return 0, err
	}

	blockCount := int64(response.(float64))
	return blockCount, nil
}
