// +build gofuzz

package zookeeper

import (
	"github.com/coredns/coredns/plugin/pkg/fuzz"
)

// Fuzz fuzzes cache.
func Fuzz(data []byte) int {
	w := Zookeeper{}
	return fuzz.Do(w, data)
}
