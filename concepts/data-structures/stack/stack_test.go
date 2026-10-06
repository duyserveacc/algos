package stack

import "testing"

func TestStackUsesLIFOOrder(t *testing.T) {
	var values Stack[int]
	values.Push(10)
	values.Push(20)

	if got, ok := values.Peek(); !ok || got != 20 {
		t.Fatalf("Peek() = (%d, %t), want (20, true)", got, ok)
	}
	if got := values.Len(); got != 2 {
		t.Fatalf("Len() = %d, want 2", got)
	}

	for _, want := range []int{20, 10} {
		got, ok := values.Pop()
		if !ok || got != want {
			t.Fatalf("Pop() = (%d, %t), want (%d, true)", got, ok, want)
		}
	}

	if !values.IsEmpty() {
		t.Fatal("IsEmpty() = false after popping every value")
	}
}

func TestEmptyStackOperations(t *testing.T) {
	var values Stack[string]

	if got, ok := values.Peek(); ok || got != "" {
		t.Fatalf("Peek() = (%q, %t), want (\"\", false)", got, ok)
	}
	if got, ok := values.Pop(); ok || got != "" {
		t.Fatalf("Pop() = (%q, %t), want (\"\", false)", got, ok)
	}
}
