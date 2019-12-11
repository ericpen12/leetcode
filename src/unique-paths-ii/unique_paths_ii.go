package unique_paths_ii

func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	length := len(obstacleGrid[0])
	mn := make([]int, length)
	mn[0] = 1
	for _, v := range obstacleGrid {
		for j := 0; j < length; j++ {
			if v[j] == 1 {
				mn[j] = 0
			}else if j > 0 {
				mn[j] += mn[j-1]
			}
		}
	}
	return mn[length-1]
}

// 分析
// 1. 动态规划问题
// 2. 一般到某点方式 = 某点上面的方式 + 某点右边的方式
// 3. dp 方程：mn[m][n] = mn[m-1][n] + mn[m][n-1]
// 3. 障碍点：到障碍点的方式为 0， 到障碍点挨着的点的途径只可能是上面或右边的一种
// 4. 碰到障碍点时，可以将到障碍点的方式设为 0

// 方案
// 一维数组
