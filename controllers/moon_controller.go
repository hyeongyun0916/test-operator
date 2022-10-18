/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/util/retry"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"test-operator/api/v1alpha1"
	cachev1alpha1 "test-operator/api/v1alpha1"
)

// MoonReconciler reconciles a Moon object
type MoonReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=cache.example.com,resources=moons,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=cache.example.com,resources=moons/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=cache.example.com,resources=moons/finalizers,verbs=update

// SetupWithManager sets up the controller with the Manager.
func (r *MoonReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&cachev1alpha1.Moon{}).
		Complete(r)
}

func (r *MoonReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	l := log.FromContext(ctx)
	l.Info("start")

	moon, found, err := r.getMoon(ctx, req.NamespacedName)
	if !found {
		return ctrl.Result{}, nil
	}
	if err != nil {
		return ctrl.Result{}, errors.Wrapf(err, "can't get Moon %s", req.NamespacedName.Name)
	}
	copiedMoon := moon.DeepCopy()

	result := &multierror.Error{}
	if err := r.syncFoo(ctx, copiedMoon); err != nil {
		result = multierror.Append(result, errors.Wrap(err, "can't syncFoo"))
	}
	if err := r.syncBar(ctx, copiedMoon); err != nil {
		result = multierror.Append(result, errors.Wrap(err, "can't syncBar"))
	}

	if diff := cmp.Diff(moon.Status, copiedMoon.Status); diff != "" {
		l.Info(diff)
		if err := r.Status().Patch(ctx, copiedMoon, client.MergeFrom(moon.DeepCopy())); err != nil {
			return ctrl.Result{}, errors.Wrap(err, "can't patch")
		}
	}

	return ctrl.Result{}, result.ErrorOrNil()
}

func (r *MoonReconciler) updateWithRetryOnConflict(ctx context.Context, namespacedName types.NamespacedName,
	fn func(ctx context.Context, moon *v1alpha1.Moon) error) error {
	l := Logger(ctx, "updateWithRetryOnConflict")

	if err := RetryOnConflictByChainedError(retry.DefaultBackoff, func() error {
		moon, found, err := r.getMoon(ctx, namespacedName)
		if err != nil {
			return errors.Wrap(err, "can't get Moon")
		}
		if !found {
			l.Info("can't find Moon. may be deleted.")
			return nil
		}
		return fn(ctx, moon)
	}); err != nil {
		return errors.Wrap(err, "can't retryOnConflict")
	}
	return nil
}

func (r *MoonReconciler) getMoon(ctx context.Context, namespacedName types.NamespacedName) (moon *v1alpha1.Moon, found bool, err error) {
	moon = &v1alpha1.Moon{}
	if err := r.Get(ctx, namespacedName, moon); err != nil {
		if k8serrors.IsNotFound(err) {
			return nil, false, nil
		}
		return nil, false, errors.Wrap(err, "can't get Moon")
	}
	return moon, true, nil
}
