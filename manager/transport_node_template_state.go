/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package manager

// transport node template application state
type TransportNodeTemplateState struct {

	// node id
	NodeId string `json:"node_id"`

	// Transport node template state on individual hosts of ComputeCollection which enabled automated transport code creation. 'FAILED_TO_CREATE' means transport node isn't created. 'IN_PROGRESS' means transport node is in progress of creation. 'FAILED_TO_REALIZE' means transport node has been created, but failed on host realization, it will repush to host by NSX later. 'SUCCESS' means transport node creation is succeeded.
	State string `json:"state,omitempty"`

	// transport node id
	TransportNodeId string `json:"transport_node_id,omitempty"`
}
