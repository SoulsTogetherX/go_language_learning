package main

import (
	"fmt"
	"os"
	"strconv"
)

func processArgs() []int {
    args := os.Args
    var ret []int
    if (len(args) > 1) {
        for i := 1; i < len(args); i++ {
            num, err := strconv.Atoi(args[i]);
            if err != nil {
                ret = append(ret, 0);
            } else {
                ret = append(ret, num);
            }
        }
    }
    return ret;
}

func partition(arr []int , low int, high int) int {
    var pivot int = arr[high];
    var i int = low - 1;

    for j := low; j <= high - 1 ; j++ {
        if(arr[j] < pivot) {
            i++;
            arr[i], arr[j] = arr[j], arr[i]
        }
    }
    arr[i+1], arr[high] = arr[high], arr[i+1]
    return i + 1;
}
           
func quickSort(nums []int, low int, high int) {
    if (low < high) {
        var pivot int = partition(nums, low, high);
        quickSort(nums, low, pivot - 1);
        quickSort(nums, pivot + 1, high);
    }
}

func quicksortStart(nums []int) {
    quickSort(nums, 0, len(nums) - 1)
}

func main() {
    var nums []int = processArgs()
    fmt.Println("Unsorted: ", nums);

    quicksortStart(nums);
    fmt.Println("Sorted: ", nums);
}