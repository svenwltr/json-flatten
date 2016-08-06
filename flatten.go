package main

import (
	"fmt"
	"sort"
)

type ByPath Assignments

func (a ByPath) Len() int {
	return len(a)
}

func (a ByPath) Less(i, j int) bool {
	return fmt.Sprint(a[i].Path) < fmt.Sprint(a[j].Path)
}

func (a ByPath) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func Flatten(data interface{}) Assignments {
	a := flatten([]Name{}, data)
	sort.Sort(ByPath(a))
	return a
}

func flatten(path []Name, value interface{}) Assignments {
	switch value := value.(type) {
	case map[string]interface{}:
		return flattenMap(path, value)
	case []interface{}:
		return flattenArray(path, value)
	default:
		return Assignments{&Assignment{path, value}}
		//panic(fmt.Sprintf("Unknown type: %#v", value))
	}
}

func flattenMap(path []Name, value map[string]interface{}) Assignments {
	out := make(Assignments, 0)
	for k, v := range value {
		p := append(path, ObjectName(k))
		out = append(out, flatten(p, v)...)
	}

	return out
}

func flattenArray(path []Name, value []interface{}) Assignments {
	out := make(Assignments, 0)
	for i, v := range value {
		p := append(path, ArrayIndex(i))
		out = append(out, flatten(p, v)...)
	}

	return out

}
