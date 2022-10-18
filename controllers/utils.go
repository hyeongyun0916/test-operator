package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/util/retry"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

func Logger(ctx context.Context, functionName string) logr.Logger {
	return log.FromContext(ctx, "func", functionName)
}

func RetryOnConflictByChainedError(backoff wait.Backoff, fn func() error) error {
	isChainedErrorHasConflict := func(err error) bool {
		for {
			if k8serrors.IsConflict(err) {
				return true
			}
			if err = errors.Unwrap(err); err == nil {
				return false
			}
		}
	}
	return retry.OnError(backoff, isChainedErrorHasConflict, fn)
}
