package storeservice

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	gRegistry "k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"

	"github.com/setcreed/store-apiserver/pkg/apis/store"
	registry "github.com/setcreed/store-apiserver/pkg/register"
)

func NewREST(scheme *runtime.Scheme, optsGetter generic.RESTOptionsGetter) (*registry.REST, error) {
	strategy := NewStrategy(scheme)

	st := &gRegistry.Store{
		NewFunc:                  func() runtime.Object { return &store.StoreService{} },
		NewListFunc:              func() runtime.Object { return &store.StoreServiceList{} },
		PredicateFunc:            MatchJenkinsService,
		DefaultQualifiedResource: store.Resource("storeservices"),

		CreateStrategy: strategy,
		UpdateStrategy: strategy,
		DeleteStrategy: strategy,

		TableConvertor: rest.NewDefaultTableConvertor(store.Resource("storeservices")),
	}
	options := &generic.StoreOptions{RESTOptions: optsGetter, AttrFunc: GetAttrs}
	if err := st.CompleteWithOptions(options); err != nil {
		return nil, err
	}
	return &registry.REST{Store: st}, nil
}
