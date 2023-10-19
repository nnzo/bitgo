package bitgo

import "log"

func (c *Client) GetBlockCount() (int64, error) {
	response, err := c.sendRequest("getblockcount", nil)
	if err != nil {
		log.Println("Please make sure to include the http:// prefix at the beginning of the URL.")
		return 0, err
	}

	blockCount := int64(response.(float64))
	return blockCount, nil
}
