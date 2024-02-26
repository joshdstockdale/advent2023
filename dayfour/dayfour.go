package dayfour

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"slices"
	"strings"
	//"regexp"
)

func Parse(filePath string) int{
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	total := 0 
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line:= scanner.Text();
		_, round, _ :=strings.Cut(line, ":")
		////log.Printf("round: %v:", round)
		dif := strings.Index(round, "|")
		win := round[:dif]
		play := round[dif+1:]
		
		r := regexp.MustCompile("[0-9]+")
		wins := r.FindAllString(win, -1)
		plays := r.FindAllString(play, -1)
		//log.Printf("plays: %v", plays)
		curTot := 0
		for _,p := range plays {
			if slices.Index(wins, p) > -1{
				if curTot == 0{
					curTot = 1
				}else{
					curTot = curTot *2
				}
			}
		}
		//log.Printf("p: %v", curTot)
		total = curTot + total
		//round := r.Split(line, -1)
		//r = regexp.MustCompile(`|`)
		//plays := r.Split(round[1], -1)
		//r = regexp.MustCompile(``)
		//wins := r.Split(plays[0], -1)
		//cards := r.Split(plays[1],-1)

		}
	return total

}
