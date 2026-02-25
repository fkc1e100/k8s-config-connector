// Copyright 2026 Google LLC
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

package v1alpha1

import (
	"context"
	"fmt"
	"strings"

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// SchemaBundleIdentity defines the resource reference to BigtableSchemaBundle, which "External" field
// holds the GCP identifier for the KRM object.
type SchemaBundleIdentity struct {
	parent *v1beta1.TableIdentity
	id     string
}

func (i *SchemaBundleIdentity) String() string {
	return i.ParentString() + "/schemaBundles/" + i.id
}

func (i *SchemaBundleIdentity) ID() string {
	return i.id
}

func (i *SchemaBundleIdentity) Parent() *v1beta1.TableIdentity {
	return i.parent
}

func (i *SchemaBundleIdentity) ParentString() string {
	return i.parent.String()
}

func (i *SchemaBundleIdentity) ParentTableIdString() string {
	return i.parent.Id
}

// New builds a SchemaBundleIdentity from the Config Connector SchemaBundle object.
func NewSchemaBundleIdentity(ctx context.Context, reader client.Reader, obj *BigtableSchemaBundle) (*SchemaBundleIdentity, error) {

	// Get Parent
	tableRef, err := obj.Spec.TableRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
	if err != nil {
		return nil, err
	}
	instanceParent, tableID, err := v1beta1.ParseTableExternal(tableRef)
	if err != nil {
		return nil, err
	}

	// Get desired ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve schema bundle name")
	}

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		actualParent, actualResourceID, err := ParseSchemaBundleExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.Parent.Parent.ProjectID != instanceParent.Parent.ProjectID {
			return nil, fmt.Errorf("ProjectID changed, expect %s, got %s", actualParent.Parent.Parent.ProjectID, instanceParent.Parent.ProjectID)
		}
		if actualParent.Parent.Id != instanceParent.Id {
			return nil, fmt.Errorf("InstanceID changed, expect %s, got %s", actualParent.Parent.Id, instanceParent.Id)
		}
		if actualParent.Id != tableID {
			return nil, fmt.Errorf("TableID changed, expect %s, got %s", actualParent.Id, tableID)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &SchemaBundleIdentity{
		parent: &v1beta1.TableIdentity{
			Parent: instanceParent,
			Id:     tableID,
		},
		id: resourceID,
	}, nil
}

func ParseSchemaBundleExternal(external string) (*v1beta1.TableIdentity, string, error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "instances" || tokens[4] != "tables" || tokens[6] != "schemaBundles" {
		return nil, "", fmt.Errorf("format of BigtableSchemaBundle external=%q was not known (use projects/{{projectID}}/instances/{{instanceID}}/tables/{{tableID}}/schemaBundles/{{schemaBundleID}})", external)
	}
	p := &v1beta1.TableIdentity{
		Parent: &v1beta1.InstanceIdentity{
			Parent: &parent.ProjectParent{
				ProjectID: tokens[1],
			},
			Id: tokens[3],
		},
		Id: tokens[5],
	}
	resourceID := tokens[7]
	return p, resourceID, nil
}
