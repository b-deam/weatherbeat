// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package haproxy

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "haproxy", asset.ModuleFieldsPri, AssetHaproxy); err != nil {
		panic(err)
	}
}

// AssetHaproxy returns asset data.
// This is the base64 encoded gzipped contents of ../metricbeat/module/haproxy.
func AssetHaproxy() string {
	return "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"
}
