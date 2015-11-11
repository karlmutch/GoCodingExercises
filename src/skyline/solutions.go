// Package skyline contains the implementation of the skyline problem
//
// The Skyline problem has a good description found at
// https://leetcode.com/problems/the-skyline-problem/
//
// This problem poses a problem where given a data structure the contains the
// highest points of a number of buildings along with their left and right sided
// positions on a line the objective is to return a set of vectors that can be used
// to trace the outlines of the buildings when they are viewed as a skyline.
//
// To trace the outline the output is a set of vectors that represent the left end
// of each line segment for the heights
//
// More details can be found on the wikipedia page
//
package skyline

// Naieve contains the dead simple brute force solution to the problem.
// In this case we will create an array that covers all the positions
// on the line  then then as each rectangle is seen we simply fill in the
// height positions on the line left to right if they are larger than the heights
// ready on the line.
//
// This is expensive on space but is the naieve solution so that is expected
//
// The input array contains a number of buildings, each is an array of length three
// containing a left side position, right side position and a height
//
func Naieve(buildings [][]int) (output [][]int) {

	rightEdge := 0

	for _, building := range buildings {
		if len(building) != 3 {
			return output
		}
		if building[1] <= building[0] {
			return output
		}

		// Get right most building edge we can find
		if building[1] > rightEdge {
			rightEdge = building[1]
		}
	}

	// Make an array of 0's for the heights as we start
	heights := make([]int, rightEdge+1)

	for _, building := range buildings {
		for i := building[0]; i < building[1]; i++ {
			if heights[i] < building[2] {
				heights[i] = building[2]
			}
		}
	}

	// Now we have the maximum heights start by taking the important transitions in heights
	// and turning them into simple vectors
	output = [][]int{}

	lastHeight := 0
	for x, y := range heights {
		if lastHeight != y {
			output = append(output, []int{x, y})
			lastHeight = y
		}
	}

	return output
}

// TransitionPoints will take the approach that is more space efficent
//
func TransitionPoints(input [][]int) (output [][]int) {
	return [][]int{{}}
}
