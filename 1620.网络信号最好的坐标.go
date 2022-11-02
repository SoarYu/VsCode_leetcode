/*
 * @lc app=leetcode.cn id=1620 lang=golang
 *
 * [1620] 网络信号最好的坐标
 */

// @lc code=start
func bestCoordinate(towers [][]int, radius int) []int {
	res, max := []int{0,0}, 0
	var getRadius func(int, int) 
	getRadius = func(i, j int) {
		quality := 0
		for _, t := range towers {
			d := (i-t[0])*(i-t[0]) + (j-t[1])*(j-t[1])
			if d <= (radius*radius) {
				quality += int(float64(t[2]) / (1 + math.Sqrt(float64(d))))
			}
		}
        if quality > max {
            max = quality
            res = []int{i, j}
        }
	}

	for x:=0; x<=50; x++ {
		for y:=0; y<=50; y++ {
            getRadius(x, y)
		}
	}
	return res
}

// @lc code=end

/*

*/
