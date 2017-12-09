// 2017-12-08 16:14:16.163327 -0800 PST m=+11.143191655
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package eip

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type DeleteEipResp struct {
	Code     int64  `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Data     struct {
		RequestID int64 `json:"requestId"`
	} `json:"data"`
	Message string `json:"message"`
}

// Implement https://cloud.tencent.com/document/api/213/1380
func DeleteEip(options ...string) (*DeleteEipResp, error) {
	resp, err := DoAction("DeleteEip", options...)
	if err != nil {
		return nil, err
	}
	var s DeleteEipResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func (r *DeleteEipResp) String(args ...interface{}) (string, error) {
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
