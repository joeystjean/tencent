// 2017-12-19 12:51:35.704808 -0800 PST m=+12.921231998
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package cvm

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type RunInstancesResp struct {
	Response struct {
		InstanceIDSet []string    `json:"InstanceIdSet,omitempty"`
		Error         interface{} `json:"Error,omitempty"`
		RequestID     string      `json:"RequestId"`
	} `json:"Response"`
}

// Implement https://cloud.tencent.com/document/api/213/9384
func (c *CvmClient) RunInstances(options ...string) (*RunInstancesResp, error) {
	resp, err := c.DoAction("RunInstances", options...)
	if err != nil {
		return nil, err
	}
	var s RunInstancesResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func RunInstances(options ...string) (*RunInstancesResp, error) {
	return DefaultClient.RunInstances(options...)
}

func (r *RunInstancesResp) String(args ...interface{}) (string, error) {
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
