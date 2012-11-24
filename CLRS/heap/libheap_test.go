/* © Copyright 2012 jingmi. All Rights Reserved.
 *
 * +----------------------------------------------------------------------+
 * | Tests of lib heap                                                    |
 * +----------------------------------------------------------------------+
 * | Author: nosqldev@gmail.com                                           |
 * +----------------------------------------------------------------------+
 * | Created: 2012-11-24 22:44                                            |
 * +----------------------------------------------------------------------+
 */

package main

import (
    "testing"
)

func TestParent(t *testing.T) {
    if parent(2) != 1 {
        t.Errorf("parent() error")
    }
}

func TestMaxHeapify(t *testing.T) {
    heap := []int{5, 3, 9, 5, 6, 8}
    max_heap := []int{5, 9, 8, 5, 6, 3}

    MaxHeapify(heap, 1)

    for i, v := range(heap) {
        if v != max_heap[i] {
            t.Errorf("MaxHeapify() error")
        }
    }
}

func TestBuildMaxHeap(t *testing.T) {
    heap := []int{5, 3, 6, 5, 9, 8}
    max_heap := []int{5, 9, 8, 5, 6, 3}

    BuildMaxHeap(heap)

    for i, v := range(heap) {
        if v != max_heap[i] {
            t.Errorf("BuildMaxHeap() error: %v, %v", v, max_heap[i])
        }
    }
}

func TestHeapSort(t *testing.T) {
    toBeSortedHeap := []int{8, 8, 2, 3, 1, 9, 10, 6, 7}
    sortedHeap := []int{8, 1, 2, 3, 6, 7, 8, 9, 10}

    HeapSort(toBeSortedHeap)

    for i, v := range(toBeSortedHeap) {
        if v != sortedHeap[i] {
            t.Errorf("HeapSort() error: [%v], %v, %v", i, v, sortedHeap[i])
        }
    }
}
