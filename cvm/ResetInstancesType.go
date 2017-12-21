// 2017-12-19 13:21:51.486779 -0800 PST m=+422.036161807
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package cvm

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type ResetInstancesTypeResp struct {
	Response struct {
		RequestID string      `json:"RequestId"`
		Error     interface{} `json:"Error,omitempty"`
	} `json:"Response"`
}

// Implement https://cloud.tencent.com/document/api/213/9394
func (c *CvmClient) ResetInstancesType(options ...string) (*ResetInstancesTypeResp, error) {
	resp, err := c.DoAction("ResetInstancesType", options...)
	if err != nil {
		return nil, err
	}
	var s ResetInstancesTypeResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func ResetInstancesType(options ...string) (*ResetInstancesTypeResp, error) {
	return DefaultClient.ResetInstancesType(options...)
}

func (r *ResetInstancesTypeResp) String(args ...interface{}) (string, error) {
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
