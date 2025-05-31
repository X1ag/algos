package main

import "strings"

func isPalindrome(s string) bool {
    l, r := 0, len(s)-1
    for l < r {
        for l < r && !isAlphaNumeric(s[l]) {
            l++
        }
        for l < r && !isAlphaNumeric(s[r]) {
            r--
        }
        slLow := strings.ToLower(string(s[l]))
        srLow := strings.ToLower(string(s[r]))
        if slLow != srLow {
            return false
        }
        l++
        r--
    }
    return true
}

func isAlphaNumeric(c byte) bool {
    return (c >= 'a' && c <= 'z') || 
           (c >= 'A' && c <= 'Z') || 
           (c >= '0' && c <= '9')
}