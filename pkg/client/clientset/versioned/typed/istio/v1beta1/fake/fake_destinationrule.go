/*
Copyright 2020 The Flux authors

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

	v1beta1 "github.com/fluxcd/flagger/pkg/apis/istio/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDestinationRules implements DestinationRuleInterface
type FakeDestinationRules struct {
	Fake *FakeNetworkingV1beta1
	ns   string
}

var destinationrulesResource = v1beta1.SchemeGroupVersion.WithResource("destinationrules")

var destinationrulesKind = v1beta1.SchemeGroupVersion.WithKind("DestinationRule")

// Get takes name of the destinationRule, and returns the corresponding destinationRule object, and an error if there is any.
func (c *FakeDestinationRules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.DestinationRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(destinationrulesResource, c.ns, name), &v1beta1.DestinationRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DestinationRule), err
}

// List takes label and field selectors, and returns the list of DestinationRules that match those selectors.
func (c *FakeDestinationRules) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.DestinationRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(destinationrulesResource, destinationrulesKind, c.ns, opts), &v1beta1.DestinationRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.DestinationRuleList{ListMeta: obj.(*v1beta1.DestinationRuleList).ListMeta}
	for _, item := range obj.(*v1beta1.DestinationRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested destinationRules.
func (c *FakeDestinationRules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(destinationrulesResource, c.ns, opts))

}

// Create takes the representation of a destinationRule and creates it.  Returns the server's representation of the destinationRule, and an error, if there is any.
func (c *FakeDestinationRules) Create(ctx context.Context, destinationRule *v1beta1.DestinationRule, opts v1.CreateOptions) (result *v1beta1.DestinationRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(destinationrulesResource, c.ns, destinationRule), &v1beta1.DestinationRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DestinationRule), err
}

// Update takes the representation of a destinationRule and updates it. Returns the server's representation of the destinationRule, and an error, if there is any.
func (c *FakeDestinationRules) Update(ctx context.Context, destinationRule *v1beta1.DestinationRule, opts v1.UpdateOptions) (result *v1beta1.DestinationRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(destinationrulesResource, c.ns, destinationRule), &v1beta1.DestinationRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DestinationRule), err
}

// Delete takes name of the destinationRule and deletes it. Returns an error if one occurs.
func (c *FakeDestinationRules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(destinationrulesResource, c.ns, name, opts), &v1beta1.DestinationRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDestinationRules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(destinationrulesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.DestinationRuleList{})
	return err
}

// Patch applies the patch and returns the patched destinationRule.
func (c *FakeDestinationRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.DestinationRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(destinationrulesResource, c.ns, name, pt, data, subresources...), &v1beta1.DestinationRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DestinationRule), err
}