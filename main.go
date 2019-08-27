func twoSum(nums []int, target int) []int {
	var preMap = make(map[int]int)
	for indexA, a := range(nums) {
			if indexB, ok := preMap[target-a]; ok {
                return []int{indexB, indexA}
			} else {
				preMap[a] = indexA
			}
	}
    return []int{}
}
