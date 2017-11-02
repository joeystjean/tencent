// 2017-11-02 18:40:30.51998075 +0000 UTC m=+50.072266625
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package cdn
import (
	"encoding/json"
	"fmt"
	"github.com/tencentcloudplatform/tcpicli/core"
)


type RefreshCdnUrlResp struct {
	Code     int64  `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Data     struct {
		Count  int64  `json:"count"`
		TaskID string `json:"task_id"`
	} `json:"data"`
	Message string `json:"message"`
}


// Implement https://cloud.tencent.com/document/api/228/3946
func RefreshCdnUrl(urls ...string) (*RefreshCdnUrlResp, error) { 
	var options []string
	for k, v := range urls {
		options = append(options, fmt.Sprintf("urls.%v=%v", k, v))
	}
	resp, err := DoAction("RefreshCdnUrl", options...)
	if err != nil {
		return nil, err
	}
	var s RefreshCdnUrlResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func (r *RefreshCdnUrlResp) String(args ...interface{}) (string, error) {
	var b []byte
	var err error
	if len(args) == 0 {
		b, err = json.MarshalIndent(r, "", "  ")
	} else if len(args) == 1 {
		if val, ok := args[0].(string); ok {
			if len(val) == 0 {
				b, err = json.MarshalIndent(r, "", "  ")
			} else {
				b, err = core.FmtOutput(val, r)
			}
		}
	}
	return string(b), err
}
