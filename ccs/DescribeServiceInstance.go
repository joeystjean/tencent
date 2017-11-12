// 2017-11-12 20:56:03.095522 +0800 CST m=+8.752934625
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package ccs

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type DescribeServiceInstanceResp struct {
	Code     int64  `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Data     struct {
		Instances []struct {
			Containers []struct {
				ContainerID string `json:"containerId"`
				Image       string `json:"image"`
				Name        string `json:"name"`
				Reason      string `json:"reason"`
				Status      string `json:"status"`
			} `json:"containers"`
			CreatedAt    string `json:"createdAt"`
			IP           string `json:"ip"`
			Name         string `json:"name"`
			NodeIP       string `json:"nodeIp"`
			NodeName     string `json:"nodeName"`
			ReadyCount   int64  `json:"readyCount"`
			Reason       string `json:"reason"`
			RestartCount int64  `json:"restartCount"`
			SrcReason    string `json:"srcReason"`
			Status       string `json:"status"`
		} `json:"instances"`
		TotalCount int64 `json:"totalCount"`
	} `json:"data"`
	Message string `json:"message"`
}

// Implement https://cloud.tencent.com/document/api/457/9433
func DescribeServiceInstance(options ...string) (*DescribeServiceInstanceResp, error) {
	resp, err := DoAction("DescribeServiceInstance", options...)
	if err != nil {
		return nil, err
	}
	var s DescribeServiceInstanceResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func (r *DescribeServiceInstanceResp) String(args ...interface{}) (string, error) {
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
