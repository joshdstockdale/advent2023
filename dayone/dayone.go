package dayone

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

//53924 : too high
//53866

func Parse(filePath string) int{
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	total := 0
	spelled := [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := scanner.Text()
		indexes := make([]int, len(line))
		for i, w := range spelled {
			r := regexp.MustCompile(w)
			for _,ix := range r.FindAllStringIndex(line, -1) {
				indexes[ix[0]] = i
			}
			
		}
		for i:=0; i < len(line); i++{
			val, err := strconv.Atoi(string(line[i]))
			if err == nil {			
				indexes[i] = val
			}
		}
		var clean []int

		for _, v := range(indexes) {
			if v > 0 {
				clean = append(clean, v)
			}
		}
		
		//log.Printf("after: %v", clean)
		if len(clean) == 0{
			continue
		}
		if len(clean) == 1{
			total += (clean[0]*10) + clean[0]	
		}else{
			total += (clean[0]*10) + clean[len(clean)-1]
		}
	}
	return total
}


