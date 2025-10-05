func productExceptSelf(nums []int) []int {
    result := []int{}
    size := len(nums)

    for i := 0; i < size; i++ {
        product := 1
        for j := 0; j < size; j++ {
            if j != i {
                product *= nums[j]
            }
        }
        result = append(result, product)
    }

    return result
}