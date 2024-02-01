package types

import (
	"errors"
	"strings"
)

const (
	FailedVerify = iota
	Expired
	ValueRequestFailed
	EnoughBalance
	ZeroValueTransfer
)

var VerifyTokenErrMap = map[int]string{
	FailedVerify:       "Not Existed token at server",
	Expired:            "Expired Token Login Again",
	ValueRequestFailed: "Failed To Send Value Check Request",
	EnoughBalance:      "Enough From Balance Under Value Transfer",
	ZeroValueTransfer:  "Try Zero Value Transfer",
}

const (
	InvalidNodeInfo = iota
)

var NodeErrMap = map[int]string{
	InvalidNodeInfo: "Not Valid Chain Name",
}

func ToNodeErrType(i int) error {
	return errors.New(NodeErrMap[i])
}

type header struct {
	Result int    `json:"result"`
	Data   string `json:"data"`
}

func newHeader(result int, data ...string) *header {
	return &header{
		Result: result,
		Data:   strings.Join(data, ","),
	}
}

type response struct {
	*header
	Result interface{} `json:"result"`
}

func NewRes(result int, res interface{}, data ...string) *response {
	return &response{
		header: newHeader(result, data...),
		Result: res,
	}
}
