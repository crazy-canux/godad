package sort

import (
	"fmt"
	"testing"
)

func TestInsertion(t *testing.T) {
	seq := []int64{23, 78, -25, 82, 0, 589}
	if res := InsertionSort(seq); res[0] != -25 || res[5] != 589 {
		fmt.Println(res)
		t.Error("Failed: sort failed.")
	} else {
		fmt.Println(res)
		t.Log("sort succeed.")
	}
}
