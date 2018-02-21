// Package tomlutil provides helper functions for https://github.com/pelletier/go-toml.
package tomlutil

import (
	"github.com/pelletier/go-toml"
)

// WalkLeaves walks a TOML tree and executes the process function on each leaf.
// If tree is the root tree, keys must be an empty slice. If tree is a subtree,
// keys is its current path.
func WalkLeaves(tree *toml.Tree, keys []string, process func(leaf interface{}, keys []string)) error {

	// create a map from the tree, this is because the toml package doesn't
	// export relevant fields from the tree struct in order to walk it
	treeMap := tree.ToMap()

	for key, v := range treeMap {
		nodeKeys := append(keys, key)
		switch value := v.(type) {
		case map[string]interface{}:
			// if we need to walk deeper, convert the tree map back to a tree
			tree, err := toml.TreeFromMap(value)
			if err != nil {
				return err
			}
			WalkLeaves(tree, nodeKeys, process)
		default:
			// this is a leaf, process it
			process(value, nodeKeys)
		}
	}
	return nil
}