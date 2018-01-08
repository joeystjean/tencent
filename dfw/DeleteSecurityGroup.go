// 2018-01-08 14:25:12.345767 -0800 PST m=+13.399430705
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package dfw

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type DeleteSecurityGroupResp struct {
	Code     int64  `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Message  string `json:"message"`
}

// Implement https://cloud.tencent.com/document/api/213/1362
func (c *DfwClient) DeleteSecurityGroup(options ...string) (*DeleteSecurityGroupResp, error) {
	resp, err := c.DoAction("DeleteSecurityGroup", options...)
	if err != nil {
		return nil, err
	}
	var s DeleteSecurityGroupResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func DeleteSecurityGroup(options ...string) (*DeleteSecurityGroupResp, error) {
	return DefaultClient.DeleteSecurityGroup(options...)
}

func (r *DeleteSecurityGroupResp) String(args ...interface{}) (string, error) {
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
