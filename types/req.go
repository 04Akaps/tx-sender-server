package types

type LoginReq struct {
	Address string `form:"address" binding:"required"`
}
