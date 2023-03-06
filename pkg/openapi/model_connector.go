// Copyright 2022 Cisco Systems, Inc. and its affiliates
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
//
// SPDX-License-Identifier: Apache-2.0

/*
 * Flame REST API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Connector struct {
	Name string `json:"name"`

	Description string `json:"description,omitempty"`

	Connection map[string]interface{} `json:"connection"`
}

// AssertConnectorRequired checks if the required fields are not zero-ed
func AssertConnectorRequired(obj Connector) error {
	elements := map[string]interface{}{
		"name":       obj.Name,
		"connection": obj.Connection,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseConnectorRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Connector (e.g. [][]Connector), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseConnectorRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aConnector, ok := obj.(Connector)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertConnectorRequired(aConnector)
	})
}
