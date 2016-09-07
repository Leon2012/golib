package ds

import (
	"testing"
)

func BenchmarkArraylist(b *testing.B) {

	for i := 0; i < b.N; i++ {
		arrayList := NewArrayList()
		arrayList.Add("a")
		//fmt.Println(arrayList.String())
		arrayList.Remove("a")
	}
}
