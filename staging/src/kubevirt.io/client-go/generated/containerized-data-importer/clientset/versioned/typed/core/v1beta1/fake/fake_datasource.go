/*
Copyright 2023 The KubeVirt Authors.

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

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"

	v1beta1 "kubevirt.io/containerized-data-importer-api/pkg/apis/core/v1beta1"
)

// FakeDataSources implements DataSourceInterface
type FakeDataSources struct {
	Fake *FakeCdiV1beta1
	ns   string
}

var datasourcesResource = schema.GroupVersionResource{Group: "cdi.kubevirt.io", Version: "v1beta1", Resource: "datasources"}

var datasourcesKind = schema.GroupVersionKind{Group: "cdi.kubevirt.io", Version: "v1beta1", Kind: "DataSource"}

// Get takes name of the dataSource, and returns the corresponding dataSource object, and an error if there is any.
func (c *FakeDataSources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.DataSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(datasourcesResource, c.ns, name), &v1beta1.DataSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DataSource), err
}

// List takes label and field selectors, and returns the list of DataSources that match those selectors.
func (c *FakeDataSources) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.DataSourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(datasourcesResource, datasourcesKind, c.ns, opts), &v1beta1.DataSourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.DataSourceList{ListMeta: obj.(*v1beta1.DataSourceList).ListMeta}
	for _, item := range obj.(*v1beta1.DataSourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dataSources.
func (c *FakeDataSources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(datasourcesResource, c.ns, opts))

}

// Create takes the representation of a dataSource and creates it.  Returns the server's representation of the dataSource, and an error, if there is any.
func (c *FakeDataSources) Create(ctx context.Context, dataSource *v1beta1.DataSource, opts v1.CreateOptions) (result *v1beta1.DataSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(datasourcesResource, c.ns, dataSource), &v1beta1.DataSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DataSource), err
}

// Update takes the representation of a dataSource and updates it. Returns the server's representation of the dataSource, and an error, if there is any.
func (c *FakeDataSources) Update(ctx context.Context, dataSource *v1beta1.DataSource, opts v1.UpdateOptions) (result *v1beta1.DataSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(datasourcesResource, c.ns, dataSource), &v1beta1.DataSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DataSource), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDataSources) UpdateStatus(ctx context.Context, dataSource *v1beta1.DataSource, opts v1.UpdateOptions) (*v1beta1.DataSource, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(datasourcesResource, "status", c.ns, dataSource), &v1beta1.DataSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DataSource), err
}

// Delete takes name of the dataSource and deletes it. Returns an error if one occurs.
func (c *FakeDataSources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(datasourcesResource, c.ns, name), &v1beta1.DataSource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDataSources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(datasourcesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.DataSourceList{})
	return err
}

// Patch applies the patch and returns the patched dataSource.
func (c *FakeDataSources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.DataSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(datasourcesResource, c.ns, name, pt, data, subresources...), &v1beta1.DataSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DataSource), err
}
