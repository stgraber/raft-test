// Copyright 2017 Canonical Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rafttest

import "github.com/hashicorp/raft"

// Config sets a hook for tweaking the raft configuration of individual nodes.
func Config(f func(int, *raft.Config)) Knob {
	return &configKnob{
		f: f,
	}
}

// configKnob gives access to the FileSnapshotStore objects used by the
// various nodes.
type configKnob struct {
	f func(int, *raft.Config)
}

func (k *configKnob) init(cluster *cluster) {
	for i, node := range cluster.nodes {
		k.f(i, node.Config)
	}
}

func (k *configKnob) cleanup(cluster *cluster) {
}