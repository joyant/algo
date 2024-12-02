package other

/**
x := overlap(201401, 201406, 201403, 201409)
fmt.Println(x)
*/

func overlap(start1, end1, start2, end2 int) int {
    if end2 < start1 {
        return 0
    } else if start2 <= start1 {
        return min(end1, end2) - start1 + 1
    } else if start2 > start1 && start2 <= end1 {
        return end1 - start2 + 1
    } else {
        return 0
    }
}
