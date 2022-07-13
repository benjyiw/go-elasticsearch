// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741

package types

// NodeInfoMemory type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/nodes/info/types.ts#L315-L318
type NodeInfoMemory struct {
	Total        string `json:"total"`
	TotalInBytes int64  `json:"total_in_bytes"`
}

// NodeInfoMemoryBuilder holds NodeInfoMemory struct and provides a builder API.
type NodeInfoMemoryBuilder struct {
	v *NodeInfoMemory
}

// NewNodeInfoMemory provides a builder for the NodeInfoMemory struct.
func NewNodeInfoMemoryBuilder() *NodeInfoMemoryBuilder {
	r := NodeInfoMemoryBuilder{
		&NodeInfoMemory{},
	}

	return &r
}

// Build finalize the chain and returns the NodeInfoMemory struct
func (rb *NodeInfoMemoryBuilder) Build() NodeInfoMemory {
	return *rb.v
}

func (rb *NodeInfoMemoryBuilder) Total(total string) *NodeInfoMemoryBuilder {
	rb.v.Total = total
	return rb
}

func (rb *NodeInfoMemoryBuilder) TotalInBytes(totalinbytes int64) *NodeInfoMemoryBuilder {
	rb.v.TotalInBytes = totalinbytes
	return rb
}