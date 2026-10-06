# Concept Curriculum

This index is the learning roadmap and coverage tracker. Linked entries have a
dedicated explanation and tested Go example. Unlinked entries are planned and
should be added when studying the topic or when a solved problem needs them.

Every concept page should explain:

- what the concept models and when to recognize it;
- its core invariant or decision rule;
- time and space complexity;
- an idiomatic Go template or implementation when applicable;
- common mistakes and boundary cases;
- links to representative solved problems.

## Foundations

- Complexity analysis: time, auxiliary space, amortized cost
- Correctness: invariants, induction, and loop reasoning
- Recursion and call-stack behavior
- Integer bounds, overflow, and modular arithmetic
- Bit operations and binary representation

## Data structures

- Arrays, slices, and strings
- Linked lists
- [Stack](data-structures/stack/)
- Queues and deques
- Hash maps and sets
- Heaps and priority queues
- Trees and binary search trees
- Tries
- Disjoint-set union
- Graph representations
- Fenwick trees and segment trees

## Algorithms

- Sorting and selection
- [Binary search and bounds](algorithms/binary-search/)
- Breadth-first and depth-first search
- Topological sorting
- Shortest paths
- Minimum spanning trees
- Divide and conquer
- Greedy algorithms
- Backtracking
- Dynamic programming
- String matching
- Number theory and combinatorics

## Problem-solving patterns

- [Two pointers](patterns/two-pointers/)
- Sliding window
- Prefix sums and difference arrays
- Fast and slow pointers
- Monotonic stacks and queues
- Interval merging
- Cyclic placement
- Sweep line
- State-space search

## Go techniques

- Judge-friendly input and output
- Table-driven tests
- Sorting and custom comparators
- `container/heap`
- Generics for reusable structures
- Recursion limits and iterative alternatives
- Allocation, slice reuse, and memory awareness
