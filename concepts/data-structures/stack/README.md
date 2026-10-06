# Stack

A stack stores values in last-in, first-out order. It is useful when the most
recent unfinished item must be processed first: nested delimiters, iterative
depth-first search, expression evaluation, and monotonic-stack problems.

## Invariant

Only the final element of the backing slice is the top. `Push` appends to that
end, while `Pop` removes from the same end.

## Complexity

| Operation | Time | Extra space |
| --- | --- | --- |
| Push | Amortized O(1) | O(1), excluding slice growth |
| Pop | O(1) | O(1) |
| Peek | O(1) | O(1) |

The implementation clears a removed slot before shrinking the slice so a stack
of pointers does not retain an otherwise unreachable value.

## Pitfalls

- Decide explicitly how empty `Pop` and `Peek` operations are represented.
- Do not remove from the front of a slice; that obscures the LIFO invariant and
  can retain a large backing array.
- A recursive algorithm already uses the call stack. An explicit stack is most
  useful when recursion depth is unsafe or traversal state must be controlled.

## Related problems

No solved problems are linked yet.
