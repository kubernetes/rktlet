/*
Copyright 2017 The Kubernetes Authors.

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

// This file was automatically generated by informer-gen with arguments: --input-dirs=[k8s.io/kubernetes/cmd/kube-aggregator/pkg/apis/apiregistration,k8s.io/kubernetes/cmd/kube-aggregator/pkg/apis/apiregistration/v1alpha1] --internal-clientset-package=k8s.io/kubernetes/cmd/kube-aggregator/pkg/client/clientset_generated/internalclientset --listers-package=k8s.io/kubernetes/cmd/kube-aggregator/pkg/client/listers --output-package=k8s.io/kubernetes/cmd/kube-aggregator/pkg/client/informers --versioned-clientset-package=k8s.io/kubernetes/cmd/kube-aggregator/pkg/client/clientset_generated/clientset

package v1alpha1

import (
	internalinterfaces "k8s.io/kubernetes/cmd/kube-aggregator/pkg/client/informers/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// APIServices returns a APIServiceInformer.
	APIServices() APIServiceInformer
}

type version struct {
	internalinterfaces.SharedInformerFactory
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory) Interface {
	return &version{f}
}

// APIServices returns a APIServiceInformer.
func (v *version) APIServices() APIServiceInformer {
	return &aPIServiceInformer{factory: v.SharedInformerFactory}
}
