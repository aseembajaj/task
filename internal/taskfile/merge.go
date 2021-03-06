package taskfile

import (
	"fmt"
	"strings"
)

// NamespaceSeparator contains the character that separates namescapes
const NamespaceSeparator = ":"

// Merge merges the second Taskfile into the first
func Merge(t1, t2 *Taskfile, namespaces ...string) error {
	if t1.Version != t2.Version {
		return fmt.Errorf(`Taskfiles versions should match. First is "%s" but second is "%s"`, t1.Version, t2.Version)
	}

	if t2.Expansions != 0 && t2.Expansions != 2 {
		t1.Expansions = t2.Expansions
	}
	if t2.Output != "" {
		t1.Output = t2.Output
	}
	for k, v := range t2.Includes {
		t1.Includes[k] = v
	}
	for k, v := range t2.Vars {
		t1.Vars[k] = v
	}
	for k, v := range t2.Tasks {
		t1.Tasks[taskNameWithNamespace(k, namespaces...)] = v
	}

	return nil
}

func taskNameWithNamespace(taskName string, namespaces ...string) string {
	return strings.Join(append(namespaces, taskName), NamespaceSeparator)
}
