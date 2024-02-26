package dayfour

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"slices"
	"strings"
)

func ParseV2(filePath string) int{
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//total := 0 
	cardArr := []int{0}
	scanner := bufio.NewScanner(file)
	card := 1
	for scanner.Scan(){
		if len(cardArr)-1 < card {
			cardArr = append(cardArr, 1)
		}else{
			cardArr[card]++
		}

		line:= scanner.Text();
		_, round, _ :=strings.Cut(line, ":")
		dif := strings.Index(round, "|")
		win := round[:dif]
		play := round[dif+1:]

		r := regexp.MustCompile("[0-9]+")
		wins := r.FindAllString(win, -1)
		plays := r.FindAllString(play, -1)
		curTot := 0
		for _,p := range plays {
			if slices.Index(wins, p) > -1{
				curTot++
			}
		}

		if curTot > 0{
			for i:=0; i < cardArr[card];i++ {
				for _, c := range getNextNums(card, curTot) {
					if len(cardArr) - 1 < c{
						cardArr = append(cardArr, 1)
					}else{
						cardArr[c]++
					}
				}
			}
		}

		card++
	}
	total:=0
	for _, v := range cardArr{
		total = total +v
	}
	return total
}

func getNextNums(n int, more int) []int{
	var total []int
	for i := 1; i <= more; i++ {
		total = append(total, n+i)	
	}
	return total
}
