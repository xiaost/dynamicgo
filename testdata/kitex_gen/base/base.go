// Code generated by thriftgo (0.3.15). DO NOT EDIT.

package base

import (
	"fmt"
)

type TrafficEnv struct {
	Open bool   `thrift:"Open,1" json:"Open"`
	Env  string `thrift:"Env,2" json:"Env"`
}

func NewTrafficEnv() *TrafficEnv {
	return &TrafficEnv{

		Open: false,
		Env:  "",
	}
}

func (p *TrafficEnv) InitDefault() {
	p.Open = false
	p.Env = ""
}

func (p *TrafficEnv) GetOpen() (v bool) {
	return p.Open
}

func (p *TrafficEnv) GetEnv() (v string) {
	return p.Env
}

func (p *TrafficEnv) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TrafficEnv(%+v)", *p)
}

var fieldIDToName_TrafficEnv = map[int16]string{
	1: "Open",
	2: "Env",
}

type Base struct {
	LogID      string            `thrift:"LogID,1" json:"LogID"`
	Caller     string            `thrift:"Caller,2" json:"Caller"`
	Addr       string            `thrift:"Addr,3" json:"Addr"`
	Client     string            `thrift:"Client,4" json:"Client"`
	TrafficEnv *TrafficEnv       `thrift:"TrafficEnv,5,optional" json:"TrafficEnv,omitempty"`
	Extra      map[string]string `thrift:"Extra,6,optional" json:"Extra,omitempty"`
}

func NewBase() *Base {
	return &Base{

		LogID:  "",
		Caller: "",
		Addr:   "",
		Client: "",
	}
}

func (p *Base) InitDefault() {
	p.LogID = ""
	p.Caller = ""
	p.Addr = ""
	p.Client = ""
}

func (p *Base) GetLogID() (v string) {
	return p.LogID
}

func (p *Base) GetCaller() (v string) {
	return p.Caller
}

func (p *Base) GetAddr() (v string) {
	return p.Addr
}

func (p *Base) GetClient() (v string) {
	return p.Client
}

var Base_TrafficEnv_DEFAULT *TrafficEnv

func (p *Base) GetTrafficEnv() (v *TrafficEnv) {
	if !p.IsSetTrafficEnv() {
		return Base_TrafficEnv_DEFAULT
	}
	return p.TrafficEnv
}

var Base_Extra_DEFAULT map[string]string

func (p *Base) GetExtra() (v map[string]string) {
	if !p.IsSetExtra() {
		return Base_Extra_DEFAULT
	}
	return p.Extra
}

func (p *Base) IsSetTrafficEnv() bool {
	return p.TrafficEnv != nil
}

func (p *Base) IsSetExtra() bool {
	return p.Extra != nil
}

func (p *Base) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Base(%+v)", *p)
}

var fieldIDToName_Base = map[int16]string{
	1: "LogID",
	2: "Caller",
	3: "Addr",
	4: "Client",
	5: "TrafficEnv",
	6: "Extra",
}

type BaseResp struct {
	StatusMessage string            `thrift:"StatusMessage,1" json:"StatusMessage"`
	StatusCode    int32             `thrift:"StatusCode,2" json:"StatusCode"`
	Extra         map[string]string `thrift:"Extra,3,optional" json:"Extra,omitempty"`
}

func NewBaseResp() *BaseResp {
	return &BaseResp{

		StatusMessage: "",
		StatusCode:    0,
	}
}

func (p *BaseResp) InitDefault() {
	p.StatusMessage = ""
	p.StatusCode = 0
}

func (p *BaseResp) GetStatusMessage() (v string) {
	return p.StatusMessage
}

func (p *BaseResp) GetStatusCode() (v int32) {
	return p.StatusCode
}

var BaseResp_Extra_DEFAULT map[string]string

func (p *BaseResp) GetExtra() (v map[string]string) {
	if !p.IsSetExtra() {
		return BaseResp_Extra_DEFAULT
	}
	return p.Extra
}

func (p *BaseResp) IsSetExtra() bool {
	return p.Extra != nil
}

func (p *BaseResp) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("BaseResp(%+v)", *p)
}

var fieldIDToName_BaseResp = map[int16]string{
	1: "StatusMessage",
	2: "StatusCode",
	3: "Extra",
}
