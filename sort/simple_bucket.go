package sort

func simpleBucket(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }
    maxVal := arr[0]
    for _, val := range arr {
        if val > maxVal {
            maxVal = val
        }
    }
    slice := make([]int, maxVal+1)
    for _, val := range arr {
        slice[val]++
    }
    newArr := make([]int, 0, len(arr))
    for index, val := range slice {
        if val > 0 {
            newArr = append(newArr, index)
        }
    }
    return newArr
}
