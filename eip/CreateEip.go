// 2017-12-08 16:12:00.33067 -0800 PST m=+6.562693086
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package eip

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type CreateEipResp struct {
	Code     int64  `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Data     struct {
		EipIds    []string `json:"eipIds"`
		RequestID int64    `json:"requestId"`
	} `json:"data"`
	Message string `json:"message"`
}

// Implement https://cloud.tencent.com/document/api/213/1381
func CreateEip(options ...string) (*CreateEipResp, error) {
	resp, err := DoAction("CreateEip", options...)
	if err != nil {
		return nil, err
	}
	var s CreateEipResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func (r *CreateEipResp) String(args ...interface{}) (string, error) {
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