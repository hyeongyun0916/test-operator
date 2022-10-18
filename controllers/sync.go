package controllers

import (
	"context"
	"test-operator/api/v1alpha1"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (r *MoonReconciler) syncFoo(ctx context.Context, moon *v1alpha1.Moon) error {
	l := Logger(ctx, "syncFoo")
	l.Info("sync foo")
	if moon.Spec.Foo != moon.Status.Foo {
		copiedMoon := moon.DeepCopy()
		copiedMoon.Status.Foo = moon.Spec.Foo
		// copiedMoon.ObjectMeta.ResourceVersion = ""
		return r.Status().Patch(ctx, copiedMoon, client.MergeFrom(moon.DeepCopy()))
	}

	return nil
}

func (r *MoonReconciler) syncBar(ctx context.Context, moon *v1alpha1.Moon) error {
	l := Logger(ctx, "syncBar")
	l.Info("sync bar")
	if moon.Spec.Bar != moon.Status.Bar {
		copiedMoon := moon.DeepCopy()
		copiedMoon.Status.Bar = moon.Spec.Bar
		// copiedMoon.ObjectMeta.ResourceVersion = ""
		return r.Status().Patch(ctx, copiedMoon, client.MergeFrom(moon.DeepCopy()))
	}

	return nil
}
