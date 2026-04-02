package floyd

func SequentialSP(adjMat [][]float64) {
	n := len(adjMat)
	floydProcessK(adjMat, 0, n, 0, n, 0, n)
}

func SequentialSPWithPath(adjMat [][]float64, prev [][]int) {
	n := len(adjMat)
	floydWithPathProcessK(adjMat, prev, 0, n, 0, n, 0, n)
}
