/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package upgrade

import (
	"github.com/vmware/go-vmware-nsxt/common"
)

type UpgradeUnitAggregateInfo struct {

	// The server will populate this field when returning the resource. Ignored on PUT and POST.
	Links []common.ResourceLink `json:"_links,omitempty"`

	Schema string `json:"_schema,omitempty"`

	Self *common.SelfResourceLink `json:"_self,omitempty"`

	// This is component version e.g. if upgrade unit is of type edge, then this is edge version.
	CurrentVersion string `json:"current_version,omitempty"`

	// Name of the upgrade unit
	DisplayName string `json:"display_name,omitempty"`

	// List of errors occurred during upgrade of this upgrade unit
	Errors []string `json:"errors,omitempty"`

	// Info of the group to which this upgrade unit belongs
	Group *UpgradeUnitGroupInfo `json:"group,omitempty"`

	// Identifier of the upgrade unit
	Id string `json:"id,omitempty"`

	// Metadata about upgrade unit
	Metadata []common.KeyValuePair `json:"metadata,omitempty"`

	// Indicator of upgrade progress in percentage
	PercentComplete float32 `json:"percent_complete,omitempty"`

	// Status of upgrade unit
	Status string `json:"status,omitempty"`

	// Upgrade unit type
	Type_ string `json:"type,omitempty"`

	// List of warnings indicating issues with the upgrade unit that may result in upgrade failure
	Warnings []string `json:"warnings,omitempty"`
}
