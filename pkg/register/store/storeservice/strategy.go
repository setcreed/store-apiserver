package storeservice

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/apiserver/pkg/storage/names"

	"github.com/setcreed/store-apiserver/pkg/apis/store"
)

type storeServiceStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

func NewStrategy(typer runtime.ObjectTyper) storeServiceStrategy {
	return storeServiceStrategy{typer, names.SimpleNameGenerator}
}

func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, error) {
	object, ok := obj.(*store.StoreService)
	if !ok {
		return nil, nil, fmt.Errorf("the object isn't a StoreService")
	}
	fs := generic.ObjectMetaFieldsSet(&object.ObjectMeta, true)
	return labels.Set(object.ObjectMeta.Labels), fs, nil
}

func MatchJenkinsService(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	return storage.SelectionPredicate{
		Label:    label,
		Field:    field,
		GetAttrs: GetAttrs,
	}
}

// CreateStrategy接口定义的方法
func (storeServiceStrategy) AllowCreateOnUpdate() bool {
	return false
}

func (storeServiceStrategy) Canonicalize(obj runtime.Object) {

}

func (storeServiceStrategy) NamespaceScoped() bool {
	return true
}

func (storeServiceStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {

}

func (storeServiceStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	errs := field.ErrorList{} //承载发现的错误

	st := obj.(*store.StoreService)
	if st.Spec.DbConfig.Replicas > 10 {
		errs = append(errs, field.TooMany(field.NewPath("spec").Key("dbconfig.replicas"), st.Spec.DbConfig.Replicas, 10))
	}
	if len(errs) > 0 {
		return errs
	} else {
		return nil
	}
}
func (storeServiceStrategy) WarningsOnCreate(ctx context.Context, obj runtime.Object) []string {
	return []string{}
}

// UpdateStrategy接口定义的方法
func (storeServiceStrategy) AllowUnconditionalUpdate() bool {
	return false
}

func (storeServiceStrategy) PrepareForUpdate(ctx context.Context, obj runtime.Object, old runtime.Object) {

}

func (storeServiceStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
}

func (storeServiceStrategy) WarningsOnUpdate(ctx context.Context, obj runtime.Object, old runtime.Object) []string {
	return []string{}
}
