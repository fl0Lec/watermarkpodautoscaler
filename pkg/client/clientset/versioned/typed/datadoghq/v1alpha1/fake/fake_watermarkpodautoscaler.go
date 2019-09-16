/*
// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-2019 Datadog, Inc.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/DataDog/watermarkpodautoscaler/pkg/apis/datadoghq/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeWatermarkPodAutoscalers implements WatermarkPodAutoscalerInterface
type FakeWatermarkPodAutoscalers struct {
	Fake *FakeDatadoghqV1alpha1
	ns   string
}

var watermarkpodautoscalersResource = schema.GroupVersionResource{Group: "datadoghq.com", Version: "v1alpha1", Resource: "watermarkpodautoscalers"}

var watermarkpodautoscalersKind = schema.GroupVersionKind{Group: "datadoghq.com", Version: "v1alpha1", Kind: "WatermarkPodAutoscaler"}

// Get takes name of the watermarkPodAutoscaler, and returns the corresponding watermarkPodAutoscaler object, and an error if there is any.
func (c *FakeWatermarkPodAutoscalers) Get(name string, options v1.GetOptions) (result *v1alpha1.WatermarkPodAutoscaler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(watermarkpodautoscalersResource, c.ns, name), &v1alpha1.WatermarkPodAutoscaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WatermarkPodAutoscaler), err
}

// List takes label and field selectors, and returns the list of WatermarkPodAutoscalers that match those selectors.
func (c *FakeWatermarkPodAutoscalers) List(opts v1.ListOptions) (result *v1alpha1.WatermarkPodAutoscalerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(watermarkpodautoscalersResource, watermarkpodautoscalersKind, c.ns, opts), &v1alpha1.WatermarkPodAutoscalerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.WatermarkPodAutoscalerList{ListMeta: obj.(*v1alpha1.WatermarkPodAutoscalerList).ListMeta}
	for _, item := range obj.(*v1alpha1.WatermarkPodAutoscalerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested watermarkPodAutoscalers.
func (c *FakeWatermarkPodAutoscalers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(watermarkpodautoscalersResource, c.ns, opts))

}

// Create takes the representation of a watermarkPodAutoscaler and creates it.  Returns the server's representation of the watermarkPodAutoscaler, and an error, if there is any.
func (c *FakeWatermarkPodAutoscalers) Create(watermarkPodAutoscaler *v1alpha1.WatermarkPodAutoscaler) (result *v1alpha1.WatermarkPodAutoscaler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(watermarkpodautoscalersResource, c.ns, watermarkPodAutoscaler), &v1alpha1.WatermarkPodAutoscaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WatermarkPodAutoscaler), err
}

// Update takes the representation of a watermarkPodAutoscaler and updates it. Returns the server's representation of the watermarkPodAutoscaler, and an error, if there is any.
func (c *FakeWatermarkPodAutoscalers) Update(watermarkPodAutoscaler *v1alpha1.WatermarkPodAutoscaler) (result *v1alpha1.WatermarkPodAutoscaler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(watermarkpodautoscalersResource, c.ns, watermarkPodAutoscaler), &v1alpha1.WatermarkPodAutoscaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WatermarkPodAutoscaler), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeWatermarkPodAutoscalers) UpdateStatus(watermarkPodAutoscaler *v1alpha1.WatermarkPodAutoscaler) (*v1alpha1.WatermarkPodAutoscaler, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(watermarkpodautoscalersResource, "status", c.ns, watermarkPodAutoscaler), &v1alpha1.WatermarkPodAutoscaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WatermarkPodAutoscaler), err
}

// Delete takes name of the watermarkPodAutoscaler and deletes it. Returns an error if one occurs.
func (c *FakeWatermarkPodAutoscalers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(watermarkpodautoscalersResource, c.ns, name), &v1alpha1.WatermarkPodAutoscaler{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeWatermarkPodAutoscalers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(watermarkpodautoscalersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.WatermarkPodAutoscalerList{})
	return err
}

// Patch applies the patch and returns the patched watermarkPodAutoscaler.
func (c *FakeWatermarkPodAutoscalers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.WatermarkPodAutoscaler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(watermarkpodautoscalersResource, c.ns, name, pt, data, subresources...), &v1alpha1.WatermarkPodAutoscaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WatermarkPodAutoscaler), err
}