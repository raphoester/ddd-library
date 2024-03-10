package tokens

import (
	"context"
)

type Storage interface {
	SaveToken(ctx context.Context, token *Token) error
}
