// 2017-12-09 00:49:36.585005 -0800 PST m=+20.138860148
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package lb

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type ModifyLoadBalancerListenerResp struct {
	Code      int64  `json:"code"`
	CodeDesc  string `json:"codeDesc"`
	Message   string `json:"message"`
	RequestID int64  `json:"requestId"`
}

// Implement https://cloud.tencent.com/document/api/214/3601
func ModifyLoadBalancerListener(options ...string) (*ModifyLoadBalancerListenerResp, error) {
	resp, err := DoAction("ModifyLoadBalancerListener", options...)
	if err != nil {
		return nil, err
	}
	var s ModifyLoadBalancerListenerResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func (r *ModifyLoadBalancerListenerResp) String(args ...interface{}) (string, error) {
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