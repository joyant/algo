package sort

func quickSort(arr []int, low, high int) {
    if low < high {
        pi := partition(arr, low, high)

        quickSort(arr, low, pi-1)
        quickSort(arr, pi+1, high)
    }
}

func partition(arr []int, low, high int) int {
    i, j := low, high
    for i < j {
        for i < j && arr[j] >= arr[low] {
            j--
        }
        for i < j && arr[i] <= arr[low] {
            i++
        }
        arr[i], arr[j] = arr[j], arr[i]
    }
    arr[i], arr[low] = arr[low], arr[i]
    return i
}

func partition2(arr []int, low, high int) int {
    pivot := arr[high]
    i := low - 1
    for j := low; j <= high-1; j++ {
        if arr[j] < pivot {
            i++
            arr[i], arr[j] = arr[j], arr[i]
        }
    }
    arr[i+1], arr[high] = arr[high], arr[i+1]
    return i + 1
}

func quickSort2(arr []int, low, high int) {
    if low < high {
        pi := partition2(arr, low, high)
        quickSort2(arr, low, pi-1)
        quickSort2(arr, pi+1, high)
    }
}
