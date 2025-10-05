func sortColors(nums []int) {
    count := [3]int{}

    for _, num := range nums {
        count[num]++
    }

    index := 0
    for i := 0; i < 3; i++ {
        for j := 0; j < count[i]; j++ {
            nums[index] = i
            index++
        }
    }
}