package pkg

import (
	"context"
)

type Mailer interface {
	Send(ctx context.Context, to, subject, html string) error
}
