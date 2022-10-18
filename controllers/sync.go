package controllers

import (
	"context"
	"test-operator/api/v1alpha1"
)

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
