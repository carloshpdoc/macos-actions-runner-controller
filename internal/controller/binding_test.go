/*
 * Copyright 2025 TomTom N.V.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package controller

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomtom-international/macos-actions-runner-controller/pkg/core/types"
	"github.com/tomtom-international/macos-actions-runner-controller/pkg/utils"
)

func nodeWithBindings(runnerIDs ...utils.UID) *types.Node {
	n := &types.Node{ID: "node-1"}
	for _, id := range runnerIDs {
		n.Status.Binding = append(n.Status.Binding, types.ResourceBinding{RunnerID: id})
	}
	return n
}

func bindingIDs(n *types.Node) []utils.UID {
	ids := make([]utils.UID, 0, len(n.Status.Binding))
	for _, b := range n.Status.Binding {
		ids = append(ids, b.RunnerID)
	}
	return ids
}

func TestRemoveRunnerBinding_RemovesOnlyThatRunner(t *testing.T) {
	n := nodeWithBindings("a", "b", "c")

	assert.True(t, removeRunnerBinding(n, "b"))
	assert.Equal(t, []utils.UID{"a", "c"}, bindingIDs(n))
}

func TestRemoveRunnerBinding_RemovesDuplicates(t *testing.T) {
	n := nodeWithBindings("a", "b", "a")

	assert.True(t, removeRunnerBinding(n, "a"))
	assert.Equal(t, []utils.UID{"b"}, bindingIDs(n))
}

func TestRemoveRunnerBinding_ReportsNothingToRemove(t *testing.T) {
	n := nodeWithBindings("a")

	assert.False(t, removeRunnerBinding(n, "missing"))
	assert.Equal(t, []utils.UID{"a"}, bindingIDs(n))

	empty := &types.Node{ID: "node-2"}
	assert.False(t, removeRunnerBinding(empty, "a"))
	assert.Empty(t, empty.Status.Binding)
}
