/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package manager

import (
	"github.com/vmware/go-vmware-nsxt/common"
)

type EdgeCluster struct {

	// The server will populate this field when returning the resource. Ignored on PUT and POST.
	Links []common.ResourceLink `json:"_links,omitempty"`

	Schema string `json:"_schema,omitempty"`

	Self *common.SelfResourceLink `json:"_self,omitempty"`

	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected.
	Revision int64 `json:"_revision"`

	// Timestamp of resource creation
	CreateTime int64 `json:"_create_time,omitempty"`

	// ID of the user who created this resource
	CreateUser string `json:"_create_user,omitempty"`

	// Timestamp of last modification
	LastModifiedTime int64 `json:"_last_modified_time,omitempty"`

	// ID of the user who last modified this resource
	LastModifiedUser string `json:"_last_modified_user,omitempty"`

	// Indicates system owned resource
	SystemOwned bool `json:"_system_owned,omitempty"`

	// Description of this resource
	Description string `json:"description,omitempty"`

	// Defaults to ID if not set
	DisplayName string `json:"display_name,omitempty"`

	// Unique identifier of this resource
	Id string `json:"id,omitempty"`

	// The type of this resource.
	ResourceType string `json:"resource_type,omitempty"`

	// Opaque identifiers meaningful to the API user
	Tags []common.Tag `json:"tags,omitempty"`

	// Edge cluster profile bindings
	ClusterProfileBindings []ClusterProfileTypeIdEntry `json:"cluster_profile_bindings,omitempty"`

	// This field is a readonly field which could show deployment_type of members. It would return UNKNOWN if there is no members, and return VIRTUAL_MACHINE| PHYSICAL_MACHINE if all edge members are VIRTUAL_MACHINE|PHYSICAL_MACHINE.
	DeploymentType string `json:"deployment_type,omitempty"`

	// Edge cluster is Homogeneous collection of transport nodes. Hence all transport nodes of the cluster must be of same type. This readonly field shows the type of transport nodes.
	MemberNodeType string `json:"member_node_type,omitempty"`

	// EdgeCluster only supports homogeneous members. These member should be backed by either EdgeNode or PublicCloudGatewayNode. DeploymentType and TransportNode type of these nodes should be the same.
	Members []EdgeClusterMember `json:"members,omitempty"`
}
