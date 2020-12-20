package graph

import "strings"

func extractID(in string) (string, string) {
	idx := strings.Index(in, ":")
	if idx < 0 {
		return "", ""
	}
	return in[:idx], in[idx+1:]
}
