package rest

import (
	"encoding/json"

	"github.com/go-ginseng/ginseng"
	"github.com/go-ginseng/sql"
)

func PaginationRequest[T any](ctx *ginseng.Context[T]) *sql.Pagination {
	var p sql.Pagination
	if err := ctx.GinCtx().ShouldBindQuery(&p); err != nil {
		return nil
	}
	if p.Page == 0 && p.Size == 0 {
		return nil
	}
	return &p
}

func SortRequest[T any](ctx *ginseng.Context[T]) *sql.Sort {
	var s sql.Sort
	if err := ctx.GinCtx().ShouldBindQuery(&s); err != nil {
		return nil
	}
	if s.By == "" && !s.Asc {
		return nil
	}
	return &s
}

type clauseRequest struct {
	Clause string `form:"clause"`
}

func ClauseRequest[T any](ctx *ginseng.Context[T]) *sql.Clause {
	var f clauseRequest
	if err := ctx.GinCtx().ShouldBindQuery(&f); err != nil {
		return nil
	}

	if f.Clause == "" {
		return nil
	}

	var clause sql.Clause
	if err := json.Unmarshal([]byte(f.Clause), &clause); err != nil {
		return nil
	}

	return &clause
}
