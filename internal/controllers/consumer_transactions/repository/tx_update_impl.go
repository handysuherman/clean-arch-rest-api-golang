package repository

import (
	"context"
	"fmt"

	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/tracing"
	"github.com/opentracing/opentracing-go"
)

type UpdateTxParams struct {
	Update UpdateParams
}

type UpdateTxResult struct {
	ConsumerTransaction *ConsumerTransaction
}

func (r *Store) UpdateTx(ctx context.Context, arg *UpdateTxParams) (UpdateTxResult, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Repository.CreateTx")
	defer span.Finish()

	var result UpdateTxResult

	err := r.execTx(ctx, func(q *Queries) error {
		var err error

		err = r.Update(ctx, &arg.Update)
		if err != nil {
			return tracing.TraceWithError(span, fmt.Errorf("r.Create.err: %v", err))
		}

		result.ConsumerTransaction, err = q.FindByID(ctx, arg.Update.ID)
		if err != nil {
			return tracing.TraceWithError(span, fmt.Errorf("q.FindByID.err: %v", err))
		}

		return err
	})

	return result, err
}
