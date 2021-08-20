// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "github.com/openshift/cluster-kube-descheduler-operator/pkg/apis/descheduler/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKubeDeschedulers implements KubeDeschedulerInterface
type FakeKubeDeschedulers struct {
	Fake *FakeKubedeschedulersV1beta1
	ns   string
}

var kubedeschedulersResource = schema.GroupVersionResource{Group: "operator.openshift.io", Version: "v1beta1", Resource: "kubedeschedulers"}

var kubedeschedulersKind = schema.GroupVersionKind{Group: "operator.openshift.io", Version: "v1beta1", Kind: "KubeDescheduler"}

// Get takes name of the kubeDescheduler, and returns the corresponding kubeDescheduler object, and an error if there is any.
func (c *FakeKubeDeschedulers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.KubeDescheduler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(kubedeschedulersResource, c.ns, name), &v1beta1.KubeDescheduler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.KubeDescheduler), err
}

// List takes label and field selectors, and returns the list of KubeDeschedulers that match those selectors.
func (c *FakeKubeDeschedulers) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.KubeDeschedulerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(kubedeschedulersResource, kubedeschedulersKind, c.ns, opts), &v1beta1.KubeDeschedulerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.KubeDeschedulerList{ListMeta: obj.(*v1beta1.KubeDeschedulerList).ListMeta}
	for _, item := range obj.(*v1beta1.KubeDeschedulerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kubeDeschedulers.
func (c *FakeKubeDeschedulers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(kubedeschedulersResource, c.ns, opts))

}

// Create takes the representation of a kubeDescheduler and creates it.  Returns the server's representation of the kubeDescheduler, and an error, if there is any.
func (c *FakeKubeDeschedulers) Create(ctx context.Context, kubeDescheduler *v1beta1.KubeDescheduler, opts v1.CreateOptions) (result *v1beta1.KubeDescheduler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(kubedeschedulersResource, c.ns, kubeDescheduler), &v1beta1.KubeDescheduler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.KubeDescheduler), err
}

// Update takes the representation of a kubeDescheduler and updates it. Returns the server's representation of the kubeDescheduler, and an error, if there is any.
func (c *FakeKubeDeschedulers) Update(ctx context.Context, kubeDescheduler *v1beta1.KubeDescheduler, opts v1.UpdateOptions) (result *v1beta1.KubeDescheduler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(kubedeschedulersResource, c.ns, kubeDescheduler), &v1beta1.KubeDescheduler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.KubeDescheduler), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKubeDeschedulers) UpdateStatus(ctx context.Context, kubeDescheduler *v1beta1.KubeDescheduler, opts v1.UpdateOptions) (*v1beta1.KubeDescheduler, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(kubedeschedulersResource, "status", c.ns, kubeDescheduler), &v1beta1.KubeDescheduler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.KubeDescheduler), err
}

// Delete takes name of the kubeDescheduler and deletes it. Returns an error if one occurs.
func (c *FakeKubeDeschedulers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(kubedeschedulersResource, c.ns, name), &v1beta1.KubeDescheduler{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKubeDeschedulers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(kubedeschedulersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.KubeDeschedulerList{})
	return err
}

// Patch applies the patch and returns the patched kubeDescheduler.
func (c *FakeKubeDeschedulers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.KubeDescheduler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(kubedeschedulersResource, c.ns, name, pt, data, subresources...), &v1beta1.KubeDescheduler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.KubeDescheduler), err
}