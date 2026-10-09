# Linked Lists

A singly linked list stores each value in a node that points to the next node.
It is useful when an algorithm must advance through a sequence without random
access, or when it must splice nodes without shifting later values.

## Invariant

Starting at the head and repeatedly following `Next` visits nodes in list
order. An acyclic list eventually reaches `nil`. Code that builds a result with
a tail pointer maintains a second invariant: `tail` is always the final node,
so appending a node takes constant time.

## Common techniques

- Use a dummy head when the first result node should not need special handling.
- Advance separate pointers independently when input lists can have different
  lengths.
- Use fast and slow pointers when detecting a cycle or locating a midpoint.
- Reverse a list by retaining the next node before redirecting `Next`.
- For reverse-order digit arithmetic, process corresponding nodes with a carry:
  output `sum % 10` and carry `sum / 10` to the next position.

## Complexity

| Operation | Time | Extra space |
| --- | --- | --- |
| Traverse or search | O(n) | O(1) |
| Insert after a known node | O(1) | O(1) |
| Reverse iteratively | O(n) | O(1) |
| Copy to or from a slice | O(n) | O(n) |

These space bounds exclude nodes created for the returned list. A recursive
traversal additionally consumes O(n) call-stack space.

## Pitfalls

- Save `node.Next` before overwriting it during a reversal.
- Handle an empty head and a one-node list explicitly through loop conditions.
- Do not lose the head while advancing a working pointer.
- A malformed cycle prevents ordinary `nil`-terminated traversal from ending.
- Count recursion depth as auxiliary space; prefer iteration when lists may be
  long.

## Related problems

- [Add Two Numbers (Easy)](../../../problems/easy_add_two_numbers.go)
