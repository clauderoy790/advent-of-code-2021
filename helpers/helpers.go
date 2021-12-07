package helpers

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func GetInputStrings(dir string) []string {
	_, file, _, _ := runtime.Caller(0)
	b, err := os.ReadFile(strings.Replace(file, "/helpers/helpers.go", fmt.Sprintf("/cmd/%s/input.txt",dir), 1))
	if err != nil {
		panic(err)
	}
	strs := strings.Split(string(b), "\n")
	for i, _ := range strs {
		strs[i] = strings.TrimSpace(strs[i])
	}
	return strs
}

func ConvertStringToInts(str string, sep string) []int {
	str = strings.TrimSpace(str)
	sl := strings.Split(str, sep)
	nbs := make([]int, 0)

	for _, s := range sl {
		nb, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			continue
		}
		nbs = append(nbs, nb)
	}
	return nbs
}

func ConvertStringsToInts(strs []string) []int {
	nbs := []int{}
	for _, s := range strs {
		nb, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			continue
		}
		nbs = append(nbs, nb)
	}
	return nbs
}
