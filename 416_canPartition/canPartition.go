package canPartition

func canPartition(nums []int) bool {
	if len(nums) == 0 {
		return true
	}

	var volumn int
	for _, num := range nums {
		volumn += num
	}
	if volumn%2 != 0 {
		return false
	}
	volumn /= 2
	// dp def
	dp := make([]bool, volumn+1)
	// dp init
	dp[0] = true
	// dp transition
	for i := 1; i <= len(nums); i++ {
		for j := volumn; j >= nums[i-1]; j-- {
			dp[j] = dp[j] || dp[j-nums[i-1]]
		}
	}
	return dp[volumn]
}
