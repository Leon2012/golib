package ds

import (
	"fmt"
	"testing"
)

func TestArraylist(t *testing.T) {
	arrayList := NewArrayList()
	arrayList.Add("a")
	fmt.Println(arrayList.String())
}
