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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	context "context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	applyconfigurationscorev1 "k8s.io/client-go/applyconfigurations/core/v1"
	gentype "k8s.io/client-go/gentype"
	scheme "k8s.io/client-go/kubernetes/scheme"
)

// ConfigMapsGetter has a method to return a ConfigMapInterface.
// A group's client should implement this interface.
type ConfigMapsGetter interface {
	ConfigMaps(namespace string) ConfigMapInterface
}

// ConfigMapInterface has methods to work with ConfigMap resources.
type ConfigMapInterface interface {
	Create(ctx context.Context, configMap *corev1.ConfigMap, opts metav1.CreateOptions) (*corev1.ConfigMap, error)
	Update(ctx context.Context, configMap *corev1.ConfigMap, opts metav1.UpdateOptions) (*corev1.ConfigMap, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*corev1.ConfigMap, error)
	List(ctx context.Context, opts metav1.ListOptions) (*corev1.ConfigMapList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *corev1.ConfigMap, err error)
	Apply(ctx context.Context, configMap *applyconfigurationscorev1.ConfigMapApplyConfiguration, opts metav1.ApplyOptions) (result *corev1.ConfigMap, err error)
	ConfigMapExpansion
}

// configMaps implements ConfigMapInterface
type configMaps struct {
	*gentype.ClientWithListAndApply[*corev1.ConfigMap, *corev1.ConfigMapList, *applyconfigurationscorev1.ConfigMapApplyConfiguration]
}

// newConfigMaps returns a ConfigMaps
func newConfigMaps(c *CoreV1Client, namespace string) *configMaps {
	return &configMaps{
		gentype.NewClientWithListAndApply[*corev1.ConfigMap, *corev1.ConfigMapList, *applyconfigurationscorev1.ConfigMapApplyConfiguration](
			"configmaps",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *corev1.ConfigMap { return &corev1.ConfigMap{} },
			func() *corev1.ConfigMapList { return &corev1.ConfigMapList{} },
			gentype.PrefersProtobuf[*corev1.ConfigMap](),
		),
	}
}
