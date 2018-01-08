// 2018-01-08 14:25:10.571692 -0800 PST m=+11.625379349
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package dfw

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type DescribeInstancesOfSecurityGroupResp struct {
	Code       int64         `json:"code"`
	CodeDesc   string        `json:"codeDesc"`
	Data       []interface{} `json:"data"`
	Message    string        `json:"message"`
	TotalCount int64         `json:"totalCount"`
}

// Implement https://cloud.tencent.com/document/api/213/1366
func (c *DfwClient) DescribeInstancesOfSecurityGroup(options ...string) (*DescribeInstancesOfSecurityGroupResp, error) {
	resp, err := c.DoAction("DescribeInstancesOfSecurityGroup", options...)
	if err != nil {
		return nil, err
	}
	var s DescribeInstancesOfSecurityGroupResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func DescribeInstancesOfSecurityGroup(options ...string) (*DescribeInstancesOfSecurityGroupResp, error) {
	return DefaultClient.DescribeInstancesOfSecurityGroup(options...)
}

func (r *DescribeInstancesOfSecurityGroupResp) String(args ...interface{}) (string, error) {
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
