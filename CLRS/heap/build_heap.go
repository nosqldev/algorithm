/* © Copyright 2012 jingmi. All Rights Reserved.
 *
 * +----------------------------------------------------------------------+
 * | Build heap                                                           |
 * +----------------------------------------------------------------------+
 * | Author: nosqldev@gmail.com                                           |
 * +----------------------------------------------------------------------+
 * | Created: 2012-11-24 15:25                                            |
 * +----------------------------------------------------------------------+
 */

package main

import (
    "fmt"
)

func printSepa() {
    fmt.Println("---------------")
}

func tryPrintHeap() {
    heap := []int{9, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    PrintHeap(heap)
    printSepa()
}

func tryMaxHeapify() {
    heap := []int{6, 5, 6, 3, 3, 2, 1}
    PrintHeap(heap)
    MaxHeapify(heap, 1)
    PrintHeap(heap)
    printSepa()
}

func main() {
    tryPrintHeap()
    tryMaxHeapify()
}
