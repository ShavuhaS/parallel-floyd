package utils

import (
	"fmt"
	"strconv"
	"strings"
	"text/tabwriter"
)

func GetMatrixString(mat [][]int, infValue int) string {
	var builder strings.Builder
	w := tabwriter.NewWriter(&builder, 5, 0, 1, ' ', tabwriter.AlignRight)
	for _, row := range mat {
		var sb strings.Builder
		for _, num := range row {
			if num == infValue {
				sb.WriteString("INF")
			} else {
				sb.WriteString(strconv.FormatInt(int64(num), 10))
			}
			sb.WriteByte('\t')
		}
		fmt.Fprintln(w, sb.String())
	}
	w.Flush()
	return builder.String()
}
