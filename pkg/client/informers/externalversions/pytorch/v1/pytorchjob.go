// Copyright 2019 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	pytorchv1 "github.com/kubeflow/pytorch-operator/pkg/apis/pytorch/v1"
	versioned "github.com/kubeflow/pytorch-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/kubeflow/pytorch-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/kubeflow/pytorch-operator/pkg/client/listers/pytorch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PyTorchJobInformer provides access to a shared informer and lister for
// PyTorchJobs.
type PyTorchJobInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.PyTorchJobLister
}

type pyTorchJobInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPyTorchJobInformer constructs a new informer for PyTorchJob type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPyTorchJobInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPyTorchJobInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPyTorchJobInformer constructs a new informer for PyTorchJob type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPyTorchJobInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeflowV1().PyTorchJobs(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeflowV1().PyTorchJobs(namespace).Watch(options)
			},
		},
		&pytorchv1.PyTorchJob{},
		resyncPeriod,
		indexers,
	)
}

func (f *pyTorchJobInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPyTorchJobInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *pyTorchJobInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&pytorchv1.PyTorchJob{}, f.defaultInformer)
}

func (f *pyTorchJobInformer) Lister() v1.PyTorchJobLister {
	return v1.NewPyTorchJobLister(f.Informer().GetIndexer())
}