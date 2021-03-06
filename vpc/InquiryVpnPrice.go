// 2018-01-09 09:22:52.028789 -0800 PST m=+53.354552430
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package vpc

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type InquiryVpnPriceResp struct {
	Code     int64  `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Message  string `json:"message"`
	Price    int64  `json:"price"`
}

// Implement https://cloud.tencent.com/document/api/215/5104
func (c *VpcClient) InquiryVpnPrice(options ...string) (*InquiryVpnPriceResp, error) {
	resp, err := c.DoAction("InquiryVpnPrice", options...)
	if err != nil {
		return nil, err
	}
	var s InquiryVpnPriceResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func InquiryVpnPrice(options ...string) (*InquiryVpnPriceResp, error) {
	return DefaultClient.InquiryVpnPrice(options...)
}

func (r *InquiryVpnPriceResp) String(args ...interface{}) (string, error) {
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
