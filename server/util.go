package server

import (
	"errors"
	"github.com/04Akaps/tx-sender-server/types"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Request uint8

const (
	GET Request = iota
	POST
	PUT
	DELETE
)

func (s *Server) setCors() {
	s.engine.Use(gin.Logger())
	s.engine.Use(gin.Recovery())
	s.engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"ORIGIN", "Content-Length", "Content-Type", "Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "Authorization", "X-Requested-With", "expires"},
		ExposeHeaders:    []string{"ORIGIN", "Content-Length", "Content-Type", "Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "Authorization", "X-Requested-With", "expires"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	}))
}

func (s *Server) register(path string, t Request, h ...gin.HandlerFunc) gin.IRoutes {
	switch t {
	case GET:
		return s.engine.GET(path, h...)
	case POST:
		return s.engine.POST(path, h...)
	case PUT:
		return s.engine.PUT(path, h...)
	case DELETE:
		return s.engine.DELETE(path, h...)
	default:
		return nil
	}
}

func response(c *gin.Context, s int, res interface{}, data ...string) {
	c.JSON(s, types.NewRes(s, res, data...))
}

func (s *Server) verifyLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := s.getAuthorizationToken(c)
		if t == "" {
			// 로그인이 안되어 있는 경우
			response(c, http.StatusUnauthorized, nil, errors.New("auth token need").Error())
			c.Abort()
		} else {
			// call to gRPC
			if _, err := s.gRPCClient.VerifyAuth(t); err != nil {
				response(c, http.StatusUnauthorized, nil, err.Error())
				c.Abort()
			} else {
				c.Next()
			}
		}
	}
}

func (s *Server) getAddressByToken(c *gin.Context) (string, error) {
	t := s.getAuthorizationToken(c)
	if res, err := s.gRPCClient.VerifyAuth(t); err != nil {
		return "", err
	} else {
		return res.Auth.Address, nil
	}
}

func (s *Server) getAuthorizationToken(c *gin.Context) string {
	var token string

	authorization := c.Request.Header.Get("Authorization")
	authSlided := strings.Split(authorization, " ")
	if len(authSlided) > 1 {
		token = authSlided[1]
	}

	return token
}
