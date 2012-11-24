/* © Copyright 2012 jingmi. All Rights Reserved.
 *
 * +----------------------------------------------------------------------+
 * | Lib of heap                                                          |
 * +----------------------------------------------------------------------+
 * | Author: nosqldev@gmail.com                                           |
 * +----------------------------------------------------------------------+
 * | Created: 2012-11-24 15:28                                            |
 * +----------------------------------------------------------------------+
 */

// NOTE: First element in heap array/slice is the size of the heap

package main

import (
    "fmt"
)

func left_child(pos int) int {
    return pos * 2
}

func right_child(pos int) int {
    return pos * 2 + 1
}

func parent(pos int) int {
    return pos / 2
}

func size(heap []int) int {
    return heap[0]
}

func PrintHeap(heap []int) {
    var pos int = 1
    var length int = 1

    for {
        if pos + length < len(heap) {
            fmt.Println(heap[pos : pos + length])
            pos += length
            length *= 2
        } else {
            fmt.Println(heap[pos : len(heap)])
            break
        }
    }
}

func MaxHeapify(heap []int, pos int) {
    left := left_child(pos)
    right := right_child(pos)
    var largest int = pos

    if (left <= size(heap)) && (heap[left] > heap[pos]) {
        largest = left
    }
    if (right <= size(heap)) && (heap[right] > heap[largest]) {
        largest = right
    }

    if largest != pos {
        heap[largest], heap[pos] = heap[pos], heap[largest]
        MaxHeapify(heap, largest)
    }
}

func BuildMaxHeap(heap []int) {
    for i := size(heap) / 2; i > 0; i-- {
        MaxHeapify(heap, i)
    }
}

func HeapSort(heap []int) {
    BuildMaxHeap(heap)

    for i := size(heap); i > 0; i-- {
        heap[1], heap[size(heap)] = heap[size(heap)], heap[1]
        heap[0] --
        MaxHeapify(heap, 1)
    }

    heap[0] = len(heap) - 1
}
