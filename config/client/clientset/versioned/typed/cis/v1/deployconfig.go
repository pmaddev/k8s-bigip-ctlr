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
	"context"
	"time"

	v1 "github.com/F5Networks/k8s-bigip-ctlr/v3/config/apis/cis/v1"
	scheme "github.com/F5Networks/k8s-bigip-ctlr/v3/config/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DeployConfigsGetter has a method to return a DeployConfigInterface.
// A group's client should implement this interface.
type DeployConfigsGetter interface {
	DeployConfigs(namespace string) DeployConfigInterface
}

// DeployConfigInterface has methods to work with DeployConfig resources.
type DeployConfigInterface interface {
	Create(ctx context.Context, deployConfig *v1.DeployConfig, opts metav1.CreateOptions) (*v1.DeployConfig, error)
	Update(ctx context.Context, deployConfig *v1.DeployConfig, opts metav1.UpdateOptions) (*v1.DeployConfig, error)
	UpdateStatus(ctx context.Context, deployConfig *v1.DeployConfig, opts metav1.UpdateOptions) (*v1.DeployConfig, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.DeployConfig, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.DeployConfigList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.DeployConfig, err error)
	DeployConfigExpansion
}

// deployConfigs implements DeployConfigInterface
type deployConfigs struct {
	client rest.Interface
	ns     string
}

// newDeployConfigs returns a DeployConfigs
func newDeployConfigs(c *CisV1Client, namespace string) *deployConfigs {
	return &deployConfigs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the deployConfig, and returns the corresponding deployConfig object, and an error if there is any.
func (c *deployConfigs) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.DeployConfig, err error) {
	result = &v1.DeployConfig{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("deployconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DeployConfigs that match those selectors.
func (c *deployConfigs) List(ctx context.Context, opts metav1.ListOptions) (result *v1.DeployConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.DeployConfigList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("deployconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested deployConfigs.
func (c *deployConfigs) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("deployconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a deployConfig and creates it.  Returns the server's representation of the deployConfig, and an error, if there is any.
func (c *deployConfigs) Create(ctx context.Context, deployConfig *v1.DeployConfig, opts metav1.CreateOptions) (result *v1.DeployConfig, err error) {
	result = &v1.DeployConfig{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("deployconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(deployConfig).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a deployConfig and updates it. Returns the server's representation of the deployConfig, and an error, if there is any.
func (c *deployConfigs) Update(ctx context.Context, deployConfig *v1.DeployConfig, opts metav1.UpdateOptions) (result *v1.DeployConfig, err error) {
	result = &v1.DeployConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("deployconfigs").
		Name(deployConfig.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(deployConfig).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *deployConfigs) UpdateStatus(ctx context.Context, deployConfig *v1.DeployConfig, opts metav1.UpdateOptions) (result *v1.DeployConfig, err error) {
	result = &v1.DeployConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("deployconfigs").
		Name(deployConfig.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(deployConfig).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the deployConfig and deletes it. Returns an error if one occurs.
func (c *deployConfigs) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("deployconfigs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *deployConfigs) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("deployconfigs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched deployConfig.
func (c *deployConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.DeployConfig, err error) {
	result = &v1.DeployConfig{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("deployconfigs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
