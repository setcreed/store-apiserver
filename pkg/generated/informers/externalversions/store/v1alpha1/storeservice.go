/*
Copyright The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	storev1alpha1 "github.com/setcreed/store-apiserver/pkg/apis/store/v1alpha1"
	versioned "github.com/setcreed/store-apiserver/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/setcreed/store-apiserver/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/setcreed/store-apiserver/pkg/generated/listers/store/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// StoreServiceInformer provides access to a shared informer and lister for
// StoreServices.
type StoreServiceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.StoreServiceLister
}

type storeServiceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewStoreServiceInformer constructs a new informer for StoreService type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewStoreServiceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredStoreServiceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredStoreServiceInformer constructs a new informer for StoreService type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredStoreServiceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SetcreedV1alpha1().StoreServices(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SetcreedV1alpha1().StoreServices(namespace).Watch(context.TODO(), options)
			},
		},
		&storev1alpha1.StoreService{},
		resyncPeriod,
		indexers,
	)
}

func (f *storeServiceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredStoreServiceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *storeServiceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&storev1alpha1.StoreService{}, f.defaultInformer)
}

func (f *storeServiceInformer) Lister() v1alpha1.StoreServiceLister {
	return v1alpha1.NewStoreServiceLister(f.Informer().GetIndexer())
}
