/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package upgrade

import (
	"github.com/vmware/go-vmware-nsxt/common"
)

type AwsAccountsListResult struct {

	// The server will populate this field when returning the resource. Ignored on PUT and POST.
	Links []common.ResourceLink `json:"_links,omitempty"`

	Schema string `json:"_schema,omitempty"`

	Self *common.SelfResourceLink `json:"_self,omitempty"`

	// Opaque cursor to be used for getting next page of records (supplied by current result page)
	Cursor string `json:"cursor,omitempty"`

	// Count of results found (across all pages), set only on first page
	ResultCount int64 `json:"result_count,omitempty"`

	SortAscending bool `json:"sort_ascending,omitempty"`

	// Field by which records are sorted
	SortBy string `json:"sort_by,omitempty"`

	// Instance statistics across all accounts managed by CSM
	AllAccountsInstanceStats *InstanceStats `json:"all_accounts_instance_stats,omitempty"`

	// Vpc statistics across all accounts managed by CSM
	AllAccountsVpcStats *VpcStats `json:"all_accounts_vpc_stats,omitempty"`

	// AWS accounts list result
	Results []AwsAccount `json:"results"`
}
