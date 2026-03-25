package floyd

var floydTestCases = []struct {
	name         string
	input        [][]float64
	expectedDist [][]float64
	expectedPrev [][]int
}{
	{
		name: "V=5",
		input: [][]float64{
			{0, INF, INF, 10, INF},
			{8, 0, INF, 17, INF},
			{17, 9, 0, 13, INF},
			{13, INF, INF, 0, 15},
			{INF, INF, INF, 16, 0},
		},
		expectedDist: [][]float64{
			{0, INF, INF, 10, 25},
			{8, 0, INF, 17, 32},
			{17, 9, 0, 13, 28},
			{13, INF, INF, 0, 15},
			{29, INF, INF, 16, 0},
		},
	},
	{
		name: "V=9",
		input: [][]float64{
			{0, 2, 15, 2, 2, INF, INF, INF, 9},
			{15, 0, 20, INF, 8, 1, 6, 11, INF},
			{8, 4, 0, 3, 11, 14, 20, 19, 6},
			{18, 6, 20, 0, INF, 10, 15, 12, 17},
			{4, 15, INF, 3, 0, INF, 12, INF, 17},
			{17, INF, 17, INF, INF, 0, 15, 2, 18},
			{11, INF, 6, 10, 20, 4, 0, 20, 2},
			{12, 17, 20, 2, 4, 5, 16, 0, 20},
			{7, INF, 5, 19, 7, 12, 17, INF, 0},
		},
		expectedDist: [][]float64{
			{0, 2, 14, 2, 2, 3, 8, 5, 9},
			{11, 0, 12, 5, 7, 1, 6, 3, 8},
			{8, 4, 0, 3, 10, 5, 10, 7, 6},
			{17, 6, 18, 0, 13, 7, 12, 9, 14},
			{4, 6, 18, 3, 0, 7, 12, 9, 13},
			{10, 10, 17, 4, 6, 0, 15, 2, 17},
			{9, 10, 6, 8, 9, 4, 0, 6, 2},
			{8, 8, 20, 2, 4, 5, 14, 0, 16},
			{7, 9, 5, 8, 7, 10, 15, 12, 0},
		},
	},
	{
		name: "V=15",
		input: [][]float64{
			{0, 10, INF, INF, INF, INF, INF, 4, 9, INF, INF, INF, 4, INF, 9},
			{INF, 0, INF, INF, INF, INF, INF, INF, INF, INF, INF, INF, 13, 2, INF},
			{INF, INF, 0, INF, INF, INF, INF, INF, INF, 14, INF, INF, INF, INF, 8},
			{INF, INF, 8, 0, INF, 8, INF, INF, INF, INF, 14, INF, 11, INF, INF},
			{INF, INF, INF, 4, 0, INF, INF, 14, 19, INF, INF, INF, 16, INF, INF},
			{INF, INF, INF, 20, INF, 0, INF, 7, 5, 11, INF, INF, INF, INF, INF},
			{INF, INF, 10, INF, INF, INF, 0, 8, 5, INF, INF, INF, INF, INF, INF},
			{INF, INF, INF, INF, INF, 10, INF, 0, 13, INF, INF, INF, INF, 10, INF},
			{INF, 6, INF, INF, INF, INF, INF, INF, 0, INF, INF, INF, 16, INF, 4},
			{INF, INF, INF, 13, INF, INF, INF, INF, 4, 0, 1, INF, INF, 19, 11},
			{INF, 9, INF, INF, 14, INF, 18, INF, 6, INF, 0, 11, INF, 6, 12},
			{4, INF, INF, 8, 6, INF, INF, INF, INF, 14, 16, 0, 11, INF, 20},
			{INF, INF, INF, INF, INF, INF, INF, 3, 3, INF, INF, 17, 0, 10, INF},
			{INF, INF, INF, INF, INF, INF, INF, 14, 20, INF, INF, 11, 7, 0, 9},
			{13, 12, INF, 20, INF, INF, INF, INF, INF, INF, INF, INF, INF, INF, 0},
		},
		expectedDist: [][]float64{
			{0, 10, 37, 29, 27, 14, 44, 4, 7, 25, 26, 21, 4, 12, 9},
			{17, 0, 29, 21, 19, 22, 46, 12, 12, 27, 28, 13, 9, 2, 11},
			{21, 20, 0, 27, 29, 35, 33, 25, 18, 14, 15, 26, 25, 21, 8},
			{29, 19, 8, 0, 28, 8, 32, 14, 13, 19, 14, 25, 11, 20, 16},
			{33, 23, 12, 4, 0, 12, 36, 14, 17, 23, 18, 29, 15, 24, 20},
			{22, 11, 28, 20, 26, 0, 30, 7, 5, 11, 12, 23, 20, 13, 9},
			{22, 11, 10, 29, 30, 18, 0, 8, 5, 24, 25, 24, 20, 13, 9},
			{25, 19, 37, 29, 27, 10, 40, 0, 13, 21, 22, 21, 17, 10, 17},
			{17, 6, 32, 24, 25, 28, 52, 18, 0, 33, 34, 19, 15, 8, 4},
			{16, 10, 21, 13, 15, 21, 19, 17, 4, 0, 1, 12, 14, 7, 8},
			{15, 9, 26, 18, 14, 26, 18, 16, 6, 25, 0, 11, 13, 6, 10},
			{4, 14, 16, 8, 6, 16, 33, 8, 11, 14, 15, 0, 8, 16, 13},
			{20, 9, 33, 25, 23, 13, 43, 3, 3, 24, 25, 17, 0, 10, 7},
			{15, 16, 27, 19, 17, 20, 44, 10, 10, 25, 26, 11, 7, 0, 9},
			{13, 12, 28, 20, 31, 27, 52, 17, 20, 38, 34, 25, 17, 14, 0},
		},
	},
}
