package storage

import (
	"context"

	"github.com/mailslurper/mailslurper/pkg/mailslurper"
)

// Config
type Config struct {
	DSN string //
}

// Connect
func Connect(ctx context.Context, cfg Config) (mailslurper.IStorage, error) {

	return nil, nil
}
