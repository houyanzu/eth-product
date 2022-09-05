package dingtalk

import (
	"github.com/houyanzu/eth-product/config"
	"github.com/houyanzu/eth-product/lib/httptool"
)

func Push(msg string) {
	conf := config.GetConfig()
	js := `{"msgtype":"text","text": {"content": ` + msg + `"}}`

	_, _, _ = httptool.PostJSON(conf.Extra.DingTalkURL, []byte(js))
}
