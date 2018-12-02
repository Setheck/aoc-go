package days

import (
	"strings"
)

func ListChecksum(list []string) int {
	duo,trio := 0,0
	for _,s := range list {
		two, three := false, false
		for _, ch := range strings.Split(s, "") {
			count := strings.Count(s, ch)
			if count == 2 && !two {
				duo++
				two = true
			}
			if count == 3 && !three {
				trio++
				three = true
			}
		}
	}
	return duo * trio
}

func CommonInList(input []string) string {
	common := make([]rune,0)
	for idx,str := range input {
		for j := idx + 1; j < len(input); j++ {
			rstr1 := []rune(str)
			rstr2 := []rune(input[j])
			if len(rstr1) != len(rstr2) {
				continue
			}
			diff := 0
			subcommon := make([]rune, 0, len(rstr1))
			for i,r1 := range rstr1 {
				if r1 != rstr2[i] {
					diff++
				} else {
					subcommon = append(subcommon, r1)
				}
			}
			if diff == 1 {
				common = append(common, subcommon...)
			}
		}
	}
	return string(common)
}