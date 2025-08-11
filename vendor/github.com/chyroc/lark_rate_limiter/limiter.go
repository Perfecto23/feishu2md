package lark_rate_limiter

import (
	"context"

	"github.com/chyroc/lark"
	"golang.org/x/time/rate"
)

func Wait(r rate.Limit, b int) func(lark.ApiEndpoint) lark.ApiEndpoint {
	limiter := rate.NewLimiter(r, b)

	return func(next lark.ApiEndpoint) lark.ApiEndpoint {
		return func(ctx context.Context, rawHttpReq *lark.RawRequestReq, resp interface{}) (*lark.Response, error) {
			_ = limiter.Wait(ctx)
			return next(ctx, rawHttpReq, resp)
		}
	}
}

func WaitWithErr(r rate.Limit, b int) func(lark.ApiEndpoint) lark.ApiEndpoint {
	limiter := rate.NewLimiter(r, b)

	return func(next lark.ApiEndpoint) lark.ApiEndpoint {
		return func(ctx context.Context, rawHttpReq *lark.RawRequestReq, resp interface{}) (*lark.Response, error) {
			if err := limiter.Wait(ctx); err != nil {
				return nil, err
			}
			return next(ctx, rawHttpReq, resp)
		}
	}
}
