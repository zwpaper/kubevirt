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
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"

	v1alpha1 "kubevirt.io/client-go/generated/kubevirt/clientset/versioned/typed/flavor/v1alpha1"
)

type FakeFlavorV1alpha1 struct {
	*testing.Fake
}

func (c *FakeFlavorV1alpha1) VirtualMachineClusterFlavors() v1alpha1.VirtualMachineClusterFlavorInterface {
	return &FakeVirtualMachineClusterFlavors{c}
}

func (c *FakeFlavorV1alpha1) VirtualMachineFlavors(namespace string) v1alpha1.VirtualMachineFlavorInterface {
	return &FakeVirtualMachineFlavors{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeFlavorV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
