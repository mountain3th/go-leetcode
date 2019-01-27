package main

import "fmt"

func max(p... int) int {
    var x int = 0
    for _, n := range p {
        if n >= x {
            x = n
        }
    }
    return x
}

func rob(nums []int) int {
    dp := make([]int, len(nums), len(nums))
    for i, n := range nums {
        if i >= 3 {
            dp[i] = max(dp[i-1], dp[i-2]+n, dp[i-3]+n)
        } else if i >= 2 {
            dp[i] = max(dp[i-1], dp[i-2]+n)
        } else if i >= 1{
            dp[i] = max(dp[i-1], n)
        } else {
            dp[i] = n
        }
    }
    if len(dp) > 0 {
        return dp[len(nums)-1]    
    }
    return 0
}