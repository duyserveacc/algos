package linkedlist

import (
	"reflect"
	"testing"
)

func TestFromSliceAndToSlicePreserveOrder(t *testing.T) {
	want := []int{2, 4, 3}
	if got := ToSlice(FromSlice(want)); !reflect.DeepEqual(got, want) {
		t.Fatalf("round trip = %v, want %v", got, want)
	}
}

func TestReverse(t *testing.T) {
	want := []int{3, 2, 1}
	got := ToSlice(Reverse(FromSlice([]int{1, 2, 3})))
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("reversed list = %v, want %v", got, want)
	}
}

func TestEmptyList(t *testing.T) {
	if got := FromSlice[int](nil); got != nil {
		t.Fatalf("FromSlice(nil) = %#v, want nil", got)
	}
	if got := Reverse[int](nil); got != nil {
		t.Fatalf("Reverse(nil) = %#v, want nil", got)
	}
	if got := ToSlice[int](nil); len(got) != 0 {
		t.Fatalf("ToSlice(nil) = %v, want empty", got)
	}
}
