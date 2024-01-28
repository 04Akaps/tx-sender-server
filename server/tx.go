package server

import (
	"github.com/04Akaps/tx-sender-server/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

type tx struct {
	server *Server
}

func newTx(server *Server) {
	t := &tx{server: server}

	baseURI := "/tx"

	server.register(baseURI+"/login", GET, t.login)
	server.register(baseURI+"/unSign", POST, server.verifyLogin(), t.unSign)
	server.register(baseURI+"/sendTx", POST, server.verifyLogin(), t.sendTx)
}

func (t *tx) login(c *gin.Context) {
	var req types.LoginReq

	if err := c.BindQuery(&req); err != nil {
		response(c, http.StatusUnprocessableEntity, nil, err.Error())
	} else if res, err := t.server.gRPCClient.CreateAuth(req.Address); err != nil {
		response(c, http.StatusInternalServerError, nil, err.Error())
	} else {
		response(c, http.StatusOK, res, "")
	}
}

func (t *tx) unSign(c *gin.Context) {}

func (t *tx) sendTx(c *gin.Context) {}
