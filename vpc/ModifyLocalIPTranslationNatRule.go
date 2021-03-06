// 2018-01-09 09:23:12.98516 -0800 PST m=+74.310962411
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package vpc

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type ModifyLocalIPTranslationNatRuleResp struct {
	Code     int64  `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Data     struct {
		TaskID int64 `json:"taskId"`
	} `json:"data"`
	Message string `json:"message"`
}

// Implement https://cloud.tencent.com/document/api/215/5187
func (c *VpcClient) ModifyLocalIPTranslationNatRule(options ...string) (*ModifyLocalIPTranslationNatRuleResp, error) {
	resp, err := c.DoAction("ModifyLocalIPTranslationNatRule", options...)
	if err != nil {
		return nil, err
	}
	var s ModifyLocalIPTranslationNatRuleResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func ModifyLocalIPTranslationNatRule(options ...string) (*ModifyLocalIPTranslationNatRuleResp, error) {
	return DefaultClient.ModifyLocalIPTranslationNatRule(options...)
}

func (r *ModifyLocalIPTranslationNatRuleResp) String(args ...interface{}) (string, error) {
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
