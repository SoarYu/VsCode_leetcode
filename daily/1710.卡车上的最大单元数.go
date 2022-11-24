package daily

import "sort"

// @lc code=start
func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	res, curBox := 0, 0
	for i := range boxTypes {
		num, size := boxTypes[i][0], boxTypes[i][1]
		if curBox+num <= truckSize {
			res += num * size
			curBox += num
		} else {
			res += (truckSize - curBox) * size
			break
		}
	}
	return res
}

// @lc code=end

func maximumUnits1(boxTypes [][]int, truckSize int) int {
	sizeMap := make(map[int]int)
	max := 1
	for _, row := range boxTypes {
		sizeMap[row[1]] += row[0]
		if row[1] > max {
			max = row[1]
		}
	}
	res, curBox := 0, 0
	for i := max; i >= 1; i-- {
		if curBox+sizeMap[i] <= truckSize {
			res += sizeMap[i] * i
			curBox += sizeMap[i]
		} else {
			res += (truckSize - curBox) * i
		}
	}
	return res
}
