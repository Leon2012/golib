package hash

import (
	"fmt"
	"testing"

	"github.com/zhenjl/cityhash"
)

func Test_Get(t *testing.T) {
	ring := New(10, cityhash_32)
	var ids []string
	for i := 0; i < 10; i++ {
		ids = append(ids, fmt.Sprintf("id=%d", i))
	}
	ring.Add(ids...)
	res := ring.Get("c")
	t.Log(res)
}

func BenchmarkGet8(b *testing.B)   { benchmarkGet(b, 8) }
func BenchmarkGet32(b *testing.B)  { benchmarkGet(b, 32) }
func BenchmarkGet128(b *testing.B) { benchmarkGet(b, 128) }
func BenchmarkGet512(b *testing.B) { benchmarkGet(b, 512) }

func benchmarkGet(b *testing.B, buckets int) {

	ring := New(53, cityhash_32)

	var ids []string
	for i := 0; i < buckets; i++ {
		ids = append(ids, fmt.Sprintf("id=%d", i))
	}

	ring.Add(ids...)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ring.Get(ids[i&(buckets-1)])
	}
}

func cityhash_32(data []byte) uint32 {
	return uint32(cityhash.CityHash32(data, uint32(len(data))))
}
