package router

import (
	"sort"
	"strings"
)

const (
	maxLevel = 255
)

type parser struct {
	fields map[uint8]records
	static map[string]Handle
}

type record struct {
	key    uint16
	handle Handle
	parts  []string
}

type records []*record

func (n records) Len() int {
	return len(records)
}

func (n records) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n records) Less(i, j int) bool {
	return n[i].key < n[j].key
}

func newParser() *parser {
	return &parser{
		fields: make(map[uint8]records),
		static: make(map[string]Handle),
	}
}

func (p *parser) register(path string, handle Handle) bool {

}
