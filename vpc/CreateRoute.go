// 2017-11-17 04:09:18.122307 -0800 PST m=+43.445143932
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package vpc

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type CreateRouteResp struct {
	Code     int64         `json:"code"`
	CodeDesc string        `json:"codeDesc"`
	Data     []interface{} `json:"data"`
	Message  string        `json:"message"`
}

// Implement https://cloud.tencent.com/document/api/215/5688
func CreateRoute(options ...string) (*CreateRouteResp, error) {
	resp, err := DoAction("CreateRoute", options...)
	if err != nil {
		return nil, err
	}
	var s CreateRouteResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func (r *CreateRouteResp) String(args ...interface{}) (string, error) {
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
