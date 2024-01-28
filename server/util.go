package server

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strings"
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

func (s *Server) VerifyLogined() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := s.GetAuthorizationToken(c)
		if t == "" {
			// 로그인이 안되어 있는 경우
		} else {
			// call to gRPC
		}
		fmt.Println(t)
	}
}

func (s *Server) GetAuthorizationToken(c *gin.Context) string {
	var token string

	authorization := c.Request.Header.Get("Authorization")
	authSlided := strings.Split(authorization, " ")
	if len(authSlided) > 1 {
		token = authSlided[1]
	}

	return token
}
