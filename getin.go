package main

import (
	"strconv"
)

type node interface {
	String() string
	Children() []node
	Lookup(string) (node, bool)
}

type mapnode struct {
	m map[string]node
}

func (m mapnode) String() string {
	result := "{"
	for k, v := range m.m {
		result += k + "," + v.String()
	}
	result += "}"
	return result
}

func (m mapnode) Children() []node {
	keys := make([]node, len(m.m))

	i := 0
	for _, v := range m.m {
		keys[i] = v
	}
	return keys
}

func (m mapnode) Lookup(key string) (node, bool) {
	v := m.m[key]
	if v != nil {
		return v, true
	}
	return nil, false
}

type setnode struct {
	m []node
}

func (m setnode) String() string {
	result := "{"
	for _, k := range m.m {
		result += k.String() + ","
	}
	result += "}"
	return result
}

func (m setnode) Children() []node {
	return m.m
}

//does this have a lookup
func (m setnode) Lookup(key string) (node, bool) {
	return nil, false
}

type stringnode struct {
	s string
}

func (m stringnode) String() {
	return
}

func (m stringnode) Children() []node {
	return nil

}

func (m stringnode) Lookup(key string) (node, bool) {
	return nil, false
}

type intnode struct {
	v int64
}

func (m intnode) String() string {
	return strconv.FormatInt(m.v, 10)
}

func (m intnode) Children() []node {
	return nil
}

func (m intnode) Lookup(key string) (node, bool) {
	return nil, false
}

// status?
func insert(n node, path []string, value node) {
	h := n.(*mapnode)
	for _, i := range path[:len(path)-1] {
		z, ok := n.Lookup(i)
		if !ok {
			m := &mapnode{make(map[string]node)}
			h.m[i] = m
			h = m
		} else {
			h = z.(*mapnode)
		}
	}
	h.m[path[len(path)-1]] = value
}

func lookup(n node, path []string) (node, bool) {
	var ok bool
	for _, i := range path {
		n, ok = n.Lookup(i)
		if !ok {
			return nil, false
		}
	}
	return n, true
}