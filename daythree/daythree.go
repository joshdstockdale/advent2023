package daythree

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)
// too low: 523948
func Parse(filePath string) int{
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	total := 0

	var table [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){

		line := scanner.Text()
		r := regexp.MustCompile(``)
		table = append(table, r.Split(line, -1))
	}
	curNum := ""
	valid := false
	for row, r:= range table {
		for col, char := range r {
			_, err := strconv.Atoi(char)
			if err == nil {
				curNum = curNum + char
				if !valid {
					if checkGrid(row, col, table){
						valid = true
//						log.Printf("isSpecial.. %v",char)
					}
				}
			}else{
				if curNum != "" && valid {
					cnum, err := strconv.Atoi(curNum)
					if err == nil {
//						log.Printf("Add to total.. %v", cnum)
						total = total + cnum
					}
				}
				curNum = ""
				valid = false
			}			
		}
	}
	return total
}

func checkGrid(row int, col int, table [][]string) bool {
	dirs := [][]int{{-1,-1}, {-1,0}, {-1,1}, {0,-1}, {0,1}, {1,-1}, {1,0}, {1,1} }

	for _, d := range dirs {
		if row+d[0] >= 0 && row+d[0] < len(table) && col+d[1] >= 0 && col+d[1] < len(table[row]) {
			if isSpecial(table[row+d[0]][col+d[1]]){
				return true
			}
		}
	}
	return false
}

func isSpecial(char string) bool{
	if char == "." {
		return false
	}
	_, err := strconv.Atoi(char)
	if err == nil {
		return false
	}
	return true
}
