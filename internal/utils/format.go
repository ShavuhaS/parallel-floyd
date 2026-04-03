package utils

import (
	"fmt"
	"strings"
	"text/tabwriter"
)

func GetMatrixString[T comparable](mat [][]T, infValue T) string {
	var builder strings.Builder
	w := tabwriter.NewWriter(&builder, 5, 0, 1, ' ', tabwriter.AlignRight)
	for _, row := range mat {
		var sb strings.Builder
		for _, num := range row {
			if num == infValue {
				sb.WriteString("INF")
			} else {
				sb.WriteString(fmt.Sprintf("%v", num))
			}
			sb.WriteByte('\t')
		}
		fmt.Fprintln(w, sb.String())
	}
	w.Flush()
	return builder.String()
}

func GetPathString[T comparable](adjMat [][]T, path []int, totalLength T) string {
	l := len(path)

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Path from %v to %v (Total: %v): [#%v]", path[0], path[l-1], totalLength, path[0]))
	for i := 1; i < l; i++ {
		u, v := path[i-1], path[i]
		edgeWeightStr := NumberToSubscript(fmt.Sprint(adjMat[u][v]))
		sb.WriteString(fmt.Sprintf(" --%v--> [#%v]", edgeWeightStr, v))
	}

	return sb.String()
}

func NumberToSubscript(numStr string) string {
	replacer := strings.NewReplacer(
		"0", "₀", "1", "₁", "2", "₂", "3", "₃", "4", "₄",
		"5", "₅", "6", "₆", "7", "₇", "8", "₈", "9", "₉",
		"-", "₋",
	)

	return replacer.Replace(numStr)
}
