/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package manager

type DneRuleList struct {

	// List of DNE rules in the section. Only Homogeneous rules are supported.
	Rules []DneRule `json:"rules"`
}
