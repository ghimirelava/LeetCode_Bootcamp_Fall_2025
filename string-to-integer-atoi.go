func myAtoi(s string) int {
    for len(s) > 0 && s[0] == ' ' {
        s = s[1:]
    }

    sign := int32(1)
    switch {
        case len(s) == 0: 
            return 0
        
        case s[0] == '+': 
            s = s[1:]
        
        case s[0] == '-': 
            sign = -1
            s = s[1:]
    }

    for len(s) > 0 && s[0] == '0' {
        s = s[1:]
    }

    result := int32(0)
    for len(s) > 0 {
        digit := int32(s[0]) - int32('0');
        
        if digit < 0 || digit > 9 {
            break
        }

        digit *= sign

        if sign > 0 && result > (math.MaxInt32 - digit) / 10 {
            result = math.MaxInt32
            break;
        } else if sign < 0 && result < (math.MinInt32 - digit) / 10 {
            result = math.MinInt32
            break;
        }

        result = result * 10 + digit
        s = s[1:]
    }

    return int(result)
}