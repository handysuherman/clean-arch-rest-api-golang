package repository

import (
	"context"
	"fmt"

	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/tracing"
	"github.com/opentracing/opentracing-go"
)

type CreateTxParams struct {
	Create CreateParams
}

type CreateTxResult struct {
	ConsumerTransaction *ConsumerTransaction
}

// Create Consumer Transaction with Database Transaction
func (r *Store) CreateTx(ctx context.Context, arg *CreateTxParams) (CreateTxResult, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Repository.CreateTx")
	defer span.Finish()

	var result CreateTxResult

	err := r.execTx(ctx, func(q *Queries) error {
		var err error

		resultId, err := r.Create(ctx, &arg.Create)
		if err != nil {
			return tracing.TraceWithError(span, fmt.Errorf("r.Create.err: %v", err))
		}

		id, err := resultId.LastInsertId()
		if err != nil {
			return tracing.TraceWithError(span, fmt.Errorf("resultId.err: %v", err))
		}

		result.ConsumerTransaction, err = q.FindByID(ctx, id)
		if err != nil {
			return tracing.TraceWithError(span, fmt.Errorf("q.FindByID.err: %v", err))
		}

		return err
	})

	return result, err
}
