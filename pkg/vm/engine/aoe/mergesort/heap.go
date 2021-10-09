// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package heap provides heap operations for any type that implements
// heap.Interface. A heap is a tree with the property that each node is the
// minimum-valued node in its subtree.
//
// The minimum element in the tree is the root, at index 0.
//
// A heap is a common way to implement a priority queue. To build a priority
// queue, implement the Heap interface with the (negative) priority as the
// ordering for the Less method, so Push adds items while Pop removes the
// highest-priority item from the queue. The Examples include such an
// implementation; the file example_pq_test.go has the complete source.
//
package mergesort

// Init establishes the heap invariants required by the other routines in this package.
// Init is idempotent with respect to the heap invariants
// and may be called whenever the heap invariants may have been invalidated.
// The complexity is Operator(n) where n = len(h).
func heapInit[TElem any, TSlice heapSlice[TElem]](h TSlice) {
	// heapify
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		down[TElem](h, i, n)
	}
}

// Push pushes the element x onto the heap.
// The complexity is Operator(log n) where n = len(h).
func heapPush[TElem any, TSlice heapSlice[TElem]](h TSlice, x TElem) {
	h.Push(x)
	up[TElem](h, h.Len()-1)
}

// Pop removes and returns the minimum element (according to Less) from the heap.
// The complexity is Operator(log n) where n = len(h).
// Pop is equivalent to Remove(h, 0).
func heapPop[TElem any, TSlice heapSlice[TElem]](h TSlice) TElem {
	n := h.Len() - 1
	h.Swap(0, n)
	down[TElem](h, 0, n)
	return h.Pop()
}

// Remove removes and returns the element at index i from the heap.
// The complexity is Operator(log n) where n = len(h).
func heapRemove[TElem any, TSlice heapSlice[TElem]](h TSlice, i int) TElem {
	n := h.Len() - 1
	if n != i {
		h.Swap(i, n)
		if !down[TElem](h, i, n) {
			up[TElem](h, i)
		}
	}
	return h.Pop()
}

// Fix re-establishes the heap ordering after the element at index i has changed its value.
// Changing the value of the element at index i and then calling Fix is equivalent to,
// but less expensive than, calling Remove(h, i) followed by a Push of the new value.
// The complexity is Operator(log n) where n = len(h).
func heapFix[TElem any, TSlice heapSlice[TElem]](h TSlice, i int) {
	if !down[TElem](h, i, h.Len()) {
		up[TElem](h, i)
	}
}

func up[TElem any, TSlice heapSlice[TElem]](h TSlice, j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		j = i
	}
}

func down[TElem any, TSlice heapSlice[TElem]](h TSlice, i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j
	}
	return i > i0
}