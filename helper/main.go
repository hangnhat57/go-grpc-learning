package helper

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ContextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return LogError(status.Error(codes.Canceled, "request cancelled from client"))
	case context.DeadlineExceeded:
		return LogError(status.Error(codes.DeadlineExceeded, "context deadline exceeded!"))
	default:
		return nil
	}

}

func LogError(err error) error {
	if err != nil {
		log.Print(err)
	}
	return err
}
