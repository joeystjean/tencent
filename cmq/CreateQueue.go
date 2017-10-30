// 2017-10-04 18:44:26.846092503 -0700 PDT
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at

package cmq
import (
	"encoding/json"
)


type CreateQueueResp struct {
	Code      int64  `json:"code"`
	Message   string `json:"message"`
	QueueID   string `json:"queueId"`
	RequestID string `json:"requestId"`
}


// Implement https://cloud.tencent.com/document/api/431/5832
func CreateQueue(options ...string) (*CreateQueueResp, error) {
	resp, err := DoAction("CreateQueue", options...)
	if err != nil {
		return nil, err
	}
	var s CreateQueueResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}
