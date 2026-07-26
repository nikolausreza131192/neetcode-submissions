func carFleet(target int, position []int, speed []int) int {
	pair := make([][2]int, len(position))
	for i := 0; i < len(position); i++ {
		pair[i] = [2]int{position[i], speed[i]}
	}

	sort.Slice(pair, func(i, j int) bool {
		return pair[i][0] > pair[j][0]
	})

	fleets := 1
	prevTime := float64(target - pair[0][0]) / float64(pair[0][1])
	for i := 0; i < len(position); i++ {
		currTime := float64(target - pair[i][0]) / float64(pair[i][1])
		if currTime > prevTime {
			fleets++
			prevTime = currTime
		}
	}
	return fleets
}
