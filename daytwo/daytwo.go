package daytwo

import (
	"bufio"
	"log"
	"os"
	"regexp"

	"slices"
	"strconv"
	"strings"
)

type Game map[string]int


func Parse(filePath string) int{
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	colors := []string{"red", "green", "blue"}
	scanner := bufio.NewScanner(file)
	
	total := 0 
	index := 1
	for scanner.Scan(){
		line := scanner.Text()
		r := regexp.MustCompile(":")
		colon := r.FindStringIndex(line)
		line = line[colon[0] + 2:]
		line = strings.ReplaceAll(line, ",", "")
		curNum := 0
		rounds := strings.Split(line, ";")
		game := Game{"red": 0, "green": 0, "blue":0}
		for _, r := range rounds {
			r = strings.Trim(r, " ")
			words := strings.Split(r, " ")
			for _,w := range words {

				if curNum > 0 {
					found := slices.Index(colors, w)
					if found > -1 {
						if curNum > game[w] {
							game[w] = curNum
						}
						curNum = 0
					}
				}
				num, err := strconv.Atoi(w)
				if err == nil{
					curNum = num	
				}
			}
		}
		total = total + getTotal(game)
		
		index++
	}
	
	return total
}

func getTotal(game Game) int {
	total:=0
	for _,k := range game {
		if total == 0{
			total = k
		}else{
			total = total * k
		}
	}
	return total
}

