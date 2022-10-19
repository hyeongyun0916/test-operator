package controllers

import (
	"context"
	"test-operator/api/v1alpha1"

	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (r *MoonReconciler) patchTest(ctx context.Context, moon *v1alpha1.Moon) error {
	moon1 := moon.DeepCopy()
	moon2 := moon.DeepCopy()

	moon.ObjectMeta.ManagedFields = nil
	moon.ObjectMeta.ResourceVersion = ""
	moon.Spec = v1alpha1.MoonSpec{Bar: "a"}
	if err := r.Patch(ctx, moon, client.Apply, client.FieldOwner("controller"), client.ForceOwnership); err != nil {
		return errors.Wrap(err, "can't server-side patch")
	}

	moon1.ObjectMeta.ManagedFields = nil
	moon1.Spec = v1alpha1.MoonSpec{Bar: "x"}
	if err := r.Patch(ctx, moon1, client.Apply, client.FieldOwner("controller"), client.ForceOwnership); err != nil {
		return errors.Wrap(err, "can't server-side patch")
	}

	moon2.ObjectMeta.ManagedFields = nil
	moon2.ObjectMeta.ResourceVersion = ""
	moon2.Spec = v1alpha1.MoonSpec{Bar: "y"}
	if err := r.Patch(ctx, moon2, client.Apply, client.FieldOwner("controller"), client.ForceOwnership); err != nil {
		return errors.Wrap(err, "can't server-side patch")
	}

	return nil
}

func (r *MoonReconciler) syncFoo(ctx context.Context, moon *v1alpha1.Moon) error {
	l := Logger(ctx, "syncFoo")
	l.Info("sync foo")
	moon.Status.Foo = moon.Spec.Foo

	return nil
}

func (r *MoonReconciler) syncBar(ctx context.Context, moon *v1alpha1.Moon) error {
	l := Logger(ctx, "syncBar")
	l.Info("sync bar")
	moon.Status.Bar = moon.Spec.Bar

	return nil
}
