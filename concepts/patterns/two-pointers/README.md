# Two Pointers

The two-pointers pattern tracks two meaningful positions instead of repeatedly
scanning a range. Common forms move inward from both ends, move at different
speeds, or maintain the boundary of a valid window.

`PairSumSorted` demonstrates inward-moving pointers on ascending data.

## Invariant

For indices outside `[left, right]`, no pair can meet the target. When the sum
is too small, advancing `left` is the only move that can increase it. When the
sum is too large, decreasing `right` is the only move that can reduce it.

## Complexity

- Time: O(n)
- Extra space: O(1)

## Pitfalls

- The inward-moving pair-sum version requires sorted input.
- Define whether indices or values must be returned.
- Negative values are valid when the data remains sorted. Do not reuse this
  movement rule when the problem's predicate breaks monotonicity.
- Ensure adding two input values cannot overflow `int` for the problem limits.

## Related problems

No solved problems are linked yet.
