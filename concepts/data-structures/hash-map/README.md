# Hash Maps

A hash map associates unique keys with values and is useful when an algorithm
needs fast lookup by identity rather than by position. Go's built-in `map` is a
natural fit for membership tests, frequency counts, grouping, and remembering
where a value was previously seen.

`PairSumIndices` demonstrates complement lookup: for each value, it asks whether
the value needed to complete the target appeared earlier.

## Invariant

Before processing index `i`, the map contains an earlier index for each distinct
value in `values[0:i]` and no index at or after `i`. Looking up before inserting
the current value prevents one element from being used twice while still
allowing equal values at different indices to form a pair.

## Complexity

- Expected time: O(n), assuming expected O(1) map lookup and insertion
- Worst-case auxiliary space: O(n)

Hash-table operations have a theoretical pathological worst case beyond their
expected O(1) bound. For ordinary integer keys, expected bounds are the useful
model for Go algorithm problems.

## Pitfalls

- Decide whether duplicate keys should keep the first index, the latest index,
  or a collection of indices.
- Look up before insertion when the current item must not match itself.
- A missing key returns the value type's zero value, so use the two-result map
  lookup when zero is a valid stored value.
- Map iteration order is unspecified; do not use it when deterministic ordering
  is required.

## Related problems

- [Two Sum workbench](../../../workbench/main.go)
