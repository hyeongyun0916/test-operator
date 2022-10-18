package controllers

import (
	"context"
	"test-operator/api/v1alpha1"
)

func (r *MoonReconciler) syncFoo(ctx context.Context, moon *v1alpha1.Moon) error {
	l := Logger(ctx, "syncFoo")
	l.Info("sync foo")
	if moon.Spec.Foo != moon.Status.Foo {
		copiedMoon := moon.DeepCopy()
		copiedMoon.Status.Foo = moon.Spec.Foo
		return r.Status().Update(ctx, copiedMoon)
	}

	return nil
}

func (r *MoonReconciler) syncBar(ctx context.Context, moon *v1alpha1.Moon) error {
	l := Logger(ctx, "syncBar")
	l.Info("sync bar")
	if moon.Spec.Bar != moon.Status.Bar {
		copiedMoon := moon.DeepCopy()
		copiedMoon.Status.Bar = moon.Spec.Bar
		return r.Status().Update(ctx, copiedMoon)
	}

	return nil
}
