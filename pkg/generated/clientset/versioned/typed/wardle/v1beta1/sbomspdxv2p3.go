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

package v1beta1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1beta1 "k8s.io/sample-apiserver/pkg/apis/wardle/v1beta1"
	scheme "k8s.io/sample-apiserver/pkg/generated/clientset/versioned/scheme"
)

// SBOMSPDXv2p3sGetter has a method to return a SBOMSPDXv2p3Interface.
// A group's client should implement this interface.
type SBOMSPDXv2p3sGetter interface {
	SBOMSPDXv2p3s(namespace string) SBOMSPDXv2p3Interface
}

// SBOMSPDXv2p3Interface has methods to work with SBOMSPDXv2p3 resources.
type SBOMSPDXv2p3Interface interface {
	Create(ctx context.Context, sBOMSPDXv2p3 *v1beta1.SBOMSPDXv2p3, opts v1.CreateOptions) (*v1beta1.SBOMSPDXv2p3, error)
	Update(ctx context.Context, sBOMSPDXv2p3 *v1beta1.SBOMSPDXv2p3, opts v1.UpdateOptions) (*v1beta1.SBOMSPDXv2p3, error)
	UpdateStatus(ctx context.Context, sBOMSPDXv2p3 *v1beta1.SBOMSPDXv2p3, opts v1.UpdateOptions) (*v1beta1.SBOMSPDXv2p3, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.SBOMSPDXv2p3, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.SBOMSPDXv2p3List, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.SBOMSPDXv2p3, err error)
	SBOMSPDXv2p3Expansion
}

// sBOMSPDXv2p3s implements SBOMSPDXv2p3Interface
type sBOMSPDXv2p3s struct {
	client rest.Interface
	ns     string
}

// newSBOMSPDXv2p3s returns a SBOMSPDXv2p3s
func newSBOMSPDXv2p3s(c *SpdxV1beta1Client, namespace string) *sBOMSPDXv2p3s {
	return &sBOMSPDXv2p3s{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sBOMSPDXv2p3, and returns the corresponding sBOMSPDXv2p3 object, and an error if there is any.
func (c *sBOMSPDXv2p3s) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.SBOMSPDXv2p3, err error) {
	result = &v1beta1.SBOMSPDXv2p3{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sbomspdxv2p3s").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SBOMSPDXv2p3s that match those selectors.
func (c *sBOMSPDXv2p3s) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.SBOMSPDXv2p3List, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.SBOMSPDXv2p3List{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sbomspdxv2p3s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sBOMSPDXv2p3s.
func (c *sBOMSPDXv2p3s) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sbomspdxv2p3s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a sBOMSPDXv2p3 and creates it.  Returns the server's representation of the sBOMSPDXv2p3, and an error, if there is any.
func (c *sBOMSPDXv2p3s) Create(ctx context.Context, sBOMSPDXv2p3 *v1beta1.SBOMSPDXv2p3, opts v1.CreateOptions) (result *v1beta1.SBOMSPDXv2p3, err error) {
	result = &v1beta1.SBOMSPDXv2p3{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sbomspdxv2p3s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sBOMSPDXv2p3).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a sBOMSPDXv2p3 and updates it. Returns the server's representation of the sBOMSPDXv2p3, and an error, if there is any.
func (c *sBOMSPDXv2p3s) Update(ctx context.Context, sBOMSPDXv2p3 *v1beta1.SBOMSPDXv2p3, opts v1.UpdateOptions) (result *v1beta1.SBOMSPDXv2p3, err error) {
	result = &v1beta1.SBOMSPDXv2p3{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sbomspdxv2p3s").
		Name(sBOMSPDXv2p3.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sBOMSPDXv2p3).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *sBOMSPDXv2p3s) UpdateStatus(ctx context.Context, sBOMSPDXv2p3 *v1beta1.SBOMSPDXv2p3, opts v1.UpdateOptions) (result *v1beta1.SBOMSPDXv2p3, err error) {
	result = &v1beta1.SBOMSPDXv2p3{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sbomspdxv2p3s").
		Name(sBOMSPDXv2p3.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sBOMSPDXv2p3).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the sBOMSPDXv2p3 and deletes it. Returns an error if one occurs.
func (c *sBOMSPDXv2p3s) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sbomspdxv2p3s").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sBOMSPDXv2p3s) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sbomspdxv2p3s").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched sBOMSPDXv2p3.
func (c *sBOMSPDXv2p3s) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.SBOMSPDXv2p3, err error) {
	result = &v1beta1.SBOMSPDXv2p3{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sbomspdxv2p3s").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
