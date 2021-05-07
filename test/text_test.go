package test

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"strconv"
	"testing"
)


func TestText(t *testing.T)  {
	e := simplifiedchinese.GBK
	encoder := e.NewEncoder()
	s := "赵信"
	fmt.Println("before encoding")
	fmt.Println(s)
	s1, ok := encoder.String("赵信")
	if ok != nil {
		fmt.Println("error")
	}
	fmt.Println("after encoding")
	fmt.Println(s1)
}

func TestSearchMatrix(t *testing.T) {
	m := [][]int{{1,3,5}}
	r := searchMatrix(m, 4)
	fmt.Println(r)
}

func searchMatrix(matrix [][]int, target int) bool {
	row1 := 0
	row2 := len(matrix) - 1
	clu1 := 0
	clu2 := len(matrix[0]) - 1
	fmt.Println(len(matrix))
	fmt.Println(len(matrix[0]))
	for row2 - row1 > 0 || clu2 - clu1 > 0 {
		r := row1
		c := clu1
		for r <row2 {
			if matrix[r][c] == target {
				return true
			} else if matrix[r][c] < target {
				r ++
				continue
			}else {
				row2 = r
				break
			}
		}

		r = row1
		c = clu1

		for c < clu2 {
			if matrix[r][c] == target {
				return true
			} else if matrix[r][c] < target {
				c ++
				continue
			}else {
				clu2 = c
				break
			}
		}

		if row2 - row1 > 0 {
			row1 ++
		}
		if clu2 - clu1 > 0 {
			clu1 ++
		}
	}

	if matrix[row1][clu1] == target {
		return true
	}
	if matrix[row2][clu2] == target {
		return true
	}

	return false
}

func numJewelsInStones(jewels string, stones string) int {
	num := 0
	m := make(map[byte]int)
	stBs := []byte(stones)
	for _, b := range stBs {
		if _, ok := m[b];ok{
			m[b] ++
		}else {
			m[b] = 1
		}
	}
	jBs := []byte(jewels)
	for _, jb := range jBs {
		if n, ok := m[jb]; ok {
			num += n
		}
	}
	return num
}

func TestNumDecoding(t *testing.T)  {
	tests := []struct {
		Input string
		Want int
	}{
		{"12", 2},
		{"226", 3},
		{"0", 0},
		{"06", 0},
		{"2101", 1},
	}

	for _, ti := range tests {
		r := numDecodings(ti.Input)
		if r != ti.Want {
			t.Errorf("input %s, want %d , but %d\n", ti.Input, ti.Want, r)
		}
	}
}

func numDecodings(s string) int {
	num := len(s)
	max2 := num / 2
	total := 0
	ss := make([]string, len(s))
	for i := 0; i < len(s); i++ {
		ss[i] = string(s[i])
	}
	if ok := check(ss); ok {
		total ++
	}
	for i := 1; i <= max2; i++ {
		pullIndex := make([]int, i+1)
		pullIndex[0] = -2
		pullBlock(ss, 1, pullIndex, &total)
	}
	return total
}

//currentPullIndex从1开始
//分割符从0开始计算： 0，1，2，3
//pullIndex[0]的value是-2，这是个哨兵
func pullBlock(ss []string, currentPullIndex int, pullIndex []int, total *int)  {
	if currentPullIndex == len(pullIndex) - 1{	//到达最后一个了
		startIndex := pullIndex[currentPullIndex - 1] + 2
		for i := startIndex; i < len(ss) - 1; i++ {
			pullIndex[currentPullIndex] = i
			newSS := newStrS(ss, pullIndex)
			ok := check(newSS)
			if ok {
				*total ++
			}
		}
		return
	}

	startIndex := pullIndex[currentPullIndex - 1] + 2
	surplus := len(pullIndex) - 1 - currentPullIndex
	for i := startIndex; i < len(ss) - (surplus * 2); i++{
		pullIndex[currentPullIndex] = i
		pullBlock(ss, currentPullIndex+1, pullIndex, total)
	}
}

func newStrS(ss []string, indexs []int) []string {
	currentIndex := 0
	newSS := make([]string, 0)
	for i := 1; i < len(indexs); i++ {
		index := indexs[i]
		for currentIndex <= index {
			if currentIndex == index {
				ns := ss[currentIndex] + ss[currentIndex + 1]
				newSS = append(newSS, ns)
				currentIndex += 2
				break
			}
			newSS = append(newSS, ss[currentIndex])
			currentIndex ++
		}
	}
	for i := currentIndex; i < len(ss); i++ {
		newSS = append(newSS, ss[i])
	}
	return newSS
}

func check(ss []string) bool {
	for _, s := range ss {
		if s[0] == '0' {
			return false
		}
		v, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		if v > 26 {
			return false
		}
	}
	return true
}

