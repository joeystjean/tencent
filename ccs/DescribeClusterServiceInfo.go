// 2018-01-06 16:08:40.644913 -0800 PST m=+298.943314366
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package ccs

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type DescribeClusterServiceInfoResp struct {
	Code     int64  `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Data     struct {
		Service struct {
			AccessType string `json:"accessType"`
			Containers []struct {
				Arguments     interface{} `json:"arguments"`
				Command       string      `json:"command"`
				ContainerName string      `json:"containerName"`
				CPU           int64       `json:"cpu"`
				Envs          interface{} `json:"envs"`
				Image         string      `json:"image"`
				LiveProbe     interface{} `json:"liveProbe"`
				Memory        int64       `json:"memory"`
				Privileged    bool        `json:"privileged"`
				ReadyProbe    interface{} `json:"readyProbe"`
				UnHubID       string      `json:"unHubId"`
				VolumeMounts  interface{} `json:"volumeMounts"`
				WorkingDir    string      `json:"workingDir"`
			} `json:"containers"`
			CreatedAt       string `json:"createdAt"`
			CurrentReplicas int64  `json:"currentReplicas"`
			DesiredReplicas int64  `json:"desiredReplicas"`
			ExternalIP      string `json:"externalIp"`
			Labels          struct {
				Qcloud_app string `json:"qcloud-app"`
			} `json:"labels"`
			LbID         string `json:"lbId"`
			LbStatus     string `json:"lbStatus"`
			Namespace    string `json:"namespace"`
			PortMappings []struct {
				ContainerPort int64  `json:"containerPort"`
				LbPort        int64  `json:"lbPort"`
				NodePort      int64  `json:"nodePort"`
				Protocol      string `json:"protocol"`
			} `json:"portMappings"`
			ReasonMap struct {
				容器未启动 int64 `json:"容器未启动"`
			} `json:"reasonMap"`
			RegionID int64 `json:"regionId"`
			Selector struct {
				Qcloud_app string `json:"qcloud-app"`
			} `json:"selector"`
			ServiceDesc  string `json:"serviceDesc"`
			ServiceIP    string `json:"serviceIp"`
			ServiceName  string `json:"serviceName"`
			SrcReasonMap struct {
				Container_not_started int64 `json:"container not started"`
			} `json:"srcReasonMap"`
			Status    string `json:"status"`
			Strategy  string `json:"strategy"`
			SubnetID  string `json:"subnetId"`
			SysLabels struct {
				Qcloud_app string `json:"qcloud-app"`
			} `json:"sysLabels"`
			UserLabels struct{}      `json:"userLabels"`
			Volumes    []interface{} `json:"volumes"`
		} `json:"service"`
	} `json:"data"`
	Message string `json:"message"`
}

// Implement https://cloud.tencent.com/document/api/457/9441
func (c *CcsClient) DescribeClusterServiceInfo(options ...string) (*DescribeClusterServiceInfoResp, error) {
	resp, err := c.DoAction("DescribeClusterServiceInfo", options...)
	if err != nil {
		return nil, err
	}
	var s DescribeClusterServiceInfoResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func DescribeClusterServiceInfo(options ...string) (*DescribeClusterServiceInfoResp, error) {
	return DefaultClient.DescribeClusterServiceInfo(options...)
}

func (r *DescribeClusterServiceInfoResp) String(args ...interface{}) (string, error) {
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
