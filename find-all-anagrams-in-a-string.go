func findAnagrams(s string, p string) []int {
		m := make(map[[26]byte]int)
        temp := [26]byte{}
        var result []int
        left, right := 0, len(p)
        for _, val := range p {
            temp[val-'a']++
        }
        m[temp]++
        for right <= len(s) {
            temp = [26]byte{}
            for _, ch := range s[left:right] {
                temp[ch-'a']++
            }
            if _, found := m[temp]; found {
                result = append(result, left)
            }
            left++
            right++
        }
        return result
}