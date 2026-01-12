# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This repository contains Go solutions for NeetCode problems (LeetCode-style algorithm and data structure problems).

## Build and Test Commands

```bash
# Run all tests
go test ./...

# Run tests for a specific problem
go test ./problems/arrays_and_hashing/01_contains_duplicate

# Run tests with verbose output
go test -v ./...

# Run a specific test function
go test -v -run TestContainsDuplicate ./problems/arrays_and_hashing/01_contains_duplicate
```

## Code Organization

Solutions are organized by NeetCode category:
- `problems/arrays_and_hashing/` - Arrays & Hashing problems
- `problems/two_pointers/` - Two Pointers problems
- `problems/sliding_window/` - Sliding Window problems
- `problems/stack/` - Stack problems
- `problems/binary_search/` - Binary Search problems
- `problems/linked_list/` - Linked List problems
- `problems/trees/` - Trees problems
- `problems/tries/` - Tries problems
- `problems/heap/` - Heap / Priority Queue problems
- `problems/backtracking/` - Backtracking problems
- `problems/graphs/` - Graphs problems
- `problems/dynamic_programming/` - Dynamic Programming problems
- `problems/greedy/` - Greedy problems
- `problems/intervals/` - Intervals problems
- `problems/math_and_geometry/` - Math & Geometry problems
- `problems/bit_manipulation/` - Bit Manipulation problems

Each problem has its own directory with a numbered prefix (e.g., `01_contains_duplicate`).

## File Naming Convention

- Solution file: `solution.go`
- Test file: `solution_test.go`
- Package name: `main`
