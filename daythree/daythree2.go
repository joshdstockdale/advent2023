package daythree

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

/*
	Check for *
		> check grid for 2 numbers
			> check left and right for total num
			> multiply numbers
			> add to total
*/
func ParseV2(filePath string) int{
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	total := 0
	//var gears []string
	var table [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){

		line := scanner.Text()
		r := regexp.MustCompile(``)
		table = append(table, r.Split(line, -1))
	}
	
	for row, r:= range table {
		for col, char := range r {

			if char == "*" {
				gs := checkGridV2(row, col, table)
				if len(gs) == 2{
					ratio := -1
					for _,g := range gs{
						//log.Printf("Gear found.. %v, ratio: %v",g, ratio)
						
						num, err := strconv.Atoi(g)
						if err == nil {
							if ratio > -1{
								total = total + (num * ratio)
								ratio = -1
							}else{
								ratio = num
							}
						}
					}
				}
			}			
		}
	}
	return total
}

func checkGridV2(row int, col int, table [][]string) []string {
	dirs := [][]int{{-1,-1}, {-1,0}, {-1,1}, {0,-1}, {0,1}, {1,-1}, {1,0}, {1,1} }
	
	var numStrs []string
	for _, d := range dirs {
		if row+d[0] >= 0 && row+d[0] < len(table) && col+d[1] >= 0 && col+d[1] < len(table[row]) {
			r:=row+d[0]
			c:=col+d[1]
			if isNumber(table[r][c]){
				//log.Printf("isNumber: %v", table[r][c])
				//Right
				curNum := []string{table[r][c]} 
				rightNum := slices.Clip(findNumber(r, c, table, 1, []string{}))
				//log.Printf("Right fullNum: %v", rightNum)
				//Left
				leftNum := slices.Clip(findNumber(r, c, table, -1, []string{}))
				slices.Reverse(leftNum)

				final := leftNum
				final = append(final, curNum[:]...)
				final = append(final, rightNum[:]...)

				finalNumStr := strings.Join(final,"")
				//}
				if slices.Index(numStrs, finalNumStr) == -1 {
					numStrs = append(numStrs, finalNumStr)
				}
			}
		}
	}
	//log.Printf("--------- numStrs: %v", numStrs)
	if len(numStrs) == 2{
		return numStrs
	}
	return []string{}
}

func isNumber(char string) bool{
	_, err := strconv.Atoi(char)
	if err != nil {
		return false
	}
	return true
}

func findNumber(row int, col int, table [][]string, dir int, numStr []string) []string {
		
	c:=col+dir

	if c >= 0 && c < len(table[row]) {
		if isNumber(table[row][c]) {
			numStr = append(numStr, table[row][c])
			return findNumber(row, c, table, dir, numStr)
		}else{
			return numStr
		}
	}else{
		return numStr
	}
}
