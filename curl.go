package request

import (
	"fmt"

	"github.com/fmonod/request/curl"
)

func (c *Client) PrintCURL() {
	ctx := c.initContext()
	req := ctx.GetRequest()
	cmd, err := curl.GetCommand(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(cmd)
}
