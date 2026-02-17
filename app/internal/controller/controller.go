package controller

import "context"

type Controller interface {
	IsExists(ctx context.Context, idempotencyKey string) bool
}
