package shared

import (
	"context"
	"errors"
)

func CheckExistingRecord[T any, R any](
	ctx context.Context,
	data T,
	method func(context.Context, T) (*R, error),
	existingErr error,
) error {
	record, err := method(ctx, data)
	if !errors.Is(err, ErrRecordNotFound) && err != nil {
		return err
	}

	if record != nil {
		return existingErr
	}

	return nil
}
