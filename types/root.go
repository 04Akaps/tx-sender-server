package types

import (
	"errors"
	"strings"
)

const (
	FailedVerify = iota
	Expired
)

var VerifyTokenErrMap = map[int]string{
	FailedVerify: "Not Existed token at server",
	Expired:      "Expired Token Login Again",
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
