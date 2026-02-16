// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package compute

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

/*
IsSelfLinkEqual Terraform and mock uses the /beta/ endpoints, while direct controller uses /v1/.
Compute resources(i.e. ComputeServiceAttachment) might be managed by legacy controller and still use beta endpoint.

(Might be a bug/intended behavior in Compute service)
When v1 resource references to a beta resource, after creation the version in selfLink of the referenced resource changed from beta to v1.

Compare resource selfLink by eliminating the version to avoid version mismatching.
todo(yuhou): Should direct controller use the same version that TF uses to avoid this mixed version issue in Compute?
*/

func IsSelfLinkEqual(a, b *string) bool {
	if reflect.DeepEqual(a, b) {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	aVal := *a
	bVal := *b

	for _, basePath := range []string{"https://compute.googleapis.com/compute", "https://www.googleapis.com/compute"} {
		for _, version := range []string{"/beta/", "/v1/"} {
			prefix := basePath + version
			if strings.HasPrefix(aVal, prefix) {
				aVal = strings.TrimPrefix(aVal, prefix)
			}
			if strings.HasPrefix(bVal, prefix) {
				bVal = strings.TrimPrefix(bVal, prefix)
			}
		}
	}
	return aVal == bVal
}

func resolveResourceName(ctx context.Context, reader client.Reader, key client.ObjectKey, gvk schema.GroupVersionKind) (*unstructured.Unstructured, error) {
	resource := &unstructured.Unstructured{}
	resource.SetGroupVersionKind(gvk)
	if err := reader.Get(ctx, key, resource); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, k8s.NewReferenceNotFoundError(resource.GroupVersionKind(), key)
		}
		return nil, fmt.Errorf("error reading referenced %v %v: %w", gvk.Kind, key, err)
	}

	return resource, nil
}

func int64Ptr(v *int32) *int64 {
	if v == nil {
		return nil
	}
	val := int64(*v)
	return &val
}

func int32Ptr(v *int64) *int32 {
	if v == nil {
		return nil
	}
	val := int32(*v)
	return &val
}

func int64PtrFromUint32Ptr(v *uint32) *int64 {
	if v == nil {
		return nil
	}
	val := int64(*v)
	return &val
}

func uint32PtrFromInt64Ptr(v *int64) *uint32 {
	if v == nil {
		return nil
	}
	val := uint32(*v)
	return &val
}

func int64PtrFromUint64Ptr(v *uint64) *int64 {
	if v == nil {
		return nil
	}
	val := int64(*v)
	return &val
}

func uint64PtrFromInt64Ptr(v *int64) *uint64 {
	if v == nil {
		return nil
	}
	val := uint64(*v)
	return &val
}

func resolveResourceRef(ctx context.Context, reader client.Reader, obj client.Object, ref *v1alpha1.ResourceRef, gvk schema.GroupVersionKind, targetField string) error {
	if ref == nil {
		return nil
	}

	if ref.External != "" {
		if ref.Name != "" {
			return fmt.Errorf("cannot specify both name and external on reference")
		}
		return nil
	}

	if ref.Name == "" {
		return fmt.Errorf("must specify either name or external on reference")
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = obj.GetNamespace()
	}

	resource, err := resolveResourceName(ctx, reader, key, gvk)
	if err != nil {
		return err
	}

	val, _, err := unstructured.NestedString(resource.Object, "status", targetField)
	if err != nil || val == "" {
		return fmt.Errorf("cannot get %s for referenced %s %v (status.%s is empty)", targetField, resource.GetKind(), resource.GetNamespace(), targetField)
	}
	ref.External = val
	return nil
}
