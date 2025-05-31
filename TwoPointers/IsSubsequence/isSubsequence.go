package main

func isSubsequence(s string, t string) bool {
    if len(s) > len(t) {
        return false
    }
    if len(s) == 0 || len(t) == 0{
        return true
    } 
    
    l := 0
    r := 0

    for i := 0; i < len(t); i++ {
        if t[r] == s[l] {
        if l + 1 == len(s) {
            return true
        }
            l++
            r++
        } else {
            r++
        }
    } 
    return false
}