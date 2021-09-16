/*
Copyright 2021 Cisco Systems, Inc. and/or its affiliates.

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

package util

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/banzaicloud/istio-operator/v2/api/v1alpha1"
	"github.com/banzaicloud/operator-tools/pkg/utils"
)

func SetICPMetadataOnObject(object metav1.Object, icp *v1alpha1.IstioControlPlane) {
	object.SetOwnerReferences([]metav1.OwnerReference{
		{
			APIVersion:         icp.GroupVersionKind().GroupVersion().String(),
			Kind:               icp.GroupVersionKind().Kind,
			Name:               icp.GetName(),
			UID:                icp.GetUID(),
			Controller:         utils.BoolPointer(true),
			BlockOwnerDeletion: utils.BoolPointer(true),
		},
	})
	object.SetLabels(icp.RevisionLabels())
}
