package rest

import (
	"github.com/go-ginseng/ginseng"
	"github.com/go-ginseng/ginseng-plugin-core/error_map"
)

type Response struct {
	Success    bool        `json:"success"`
	Error      *Error      `json:"error,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

type Error struct {
	Code    interface{} `json:"code"`
	Message string      `json:"message"`
}

type Pagination struct {
	Page  int   `json:"page" form:"page"`
	Size  int   `json:"size" form:"size"`
	Total int64 `json:"total" form:"-"`
}

func OK[T any](ctx *ginseng.Context[T], data interface{}, p *Pagination) {
	ctx.Response = &Response{
		Success:    true,
		Data:       data,
		Pagination: p,
	}
	ctx.GinCtx().JSON(200, ctx.Response)
}

func ERR[T any](ctx *ginseng.Context[T], code interface{}, data interface{}) {
	ctx.Response = &Response{
		Success: false,
		Error: &Error{
			Code:    code,
			Message: error_map.Message(code),
		},
		Data: data,
	}
	ctx.GinCtx().JSON(200, ctx.Response)
}
