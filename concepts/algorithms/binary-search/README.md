# Binary Search and Bounds

Binary search applies when a search space has a monotonic decision: after some
boundary, a predicate remains true. Searching a sorted slice is the most common
example, but the same reasoning supports binary search over an answer range.

## Invariant

`LowerBound` maintains a half-open interval `[left, right)` containing the first
index whose value could be at least the target. When the interval is empty,
`left` is the insertion point.

`UpperBound` uses the same interval but finds the first value strictly greater
than the target.

## Complexity

- Time: O(log n)
- Extra space: O(1)

## Pitfalls

- The input must be sorted in ascending order.
- Be deliberate about whether equality belongs on the left or right side.
- Prefer `left + (right-left)/2` to make the midpoint rule explicit and safe.
- A valid insertion point can equal `len(values)`.

## Related problems

No solved problems are linked yet.
