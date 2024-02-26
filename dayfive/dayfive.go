package dayfive

import (
	"bufio"
	"cmp"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"runtime"
	"slices"
	"strconv"
	"strings"
)
type Category struct {
	source string
	dest string	
}
type SrcDestLen struct {
	src int
	dest int
	len int
	category Category
}
type Node struct {
	id int
	next int
	category Category
} 

type NodeList struct {
	nodes []Node
}

func Parse(filePath string) int{
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines, err := getLines(file)

	seeds, err := getSeeds(lines[0])
	log.Println("got seeds")
	sdl, err := getSrcDestLen(lines[1:])	
	log.Println("sdl")
	var nl NodeList
	nl.getNodes(seeds, sdl, 0)
	log.Println("got NODES")
// 			// Print the memory statistics after allocating memory
// 			printMemStats()

// 			// Force a garbage collection cycle
//  			runtime.GC()

// 			// Print the memory statistics after garbage collection
//  			printMemStats()
	//nl.filterNodesByCategory(6)
	slices.SortFunc(nl.nodes, func(a,b Node) int {return cmp.Compare(a.next, b.next) })
	return nl.nodes[0].next
}
func (n *Node) findNextSeedStep(step int, source int, dest int, length int) {
	if n.id == n.next || n.next == 0 {
	
		if step < source || step > source + length{
		// out of bounds, return step
		////log.Println("OUTOFBOUNDS")
			n.next = step
		}else{
		////log.Println("RETURN: %v", (dest - source + step))
			n.next = dest - source + step 
		}
	}
}
//func walk(seed int, nodes []Node) int {
//	step := seed
//	for i:=0; i < len(nodes); i++ {
//		if nodes[i].next != 0 {
//			step = nodes[i].next
//		}
//	}
//	return step
//}

func getLines(file io.Reader)([]string, error){
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

func getSeeds(line string)([]int, error){
	//found := strings.Index(line, "seeds: ") 
	//if found > -1{
		r := regexp.MustCompile("[0-9]+")
		checkSeeds := convertAllToint(r.FindAllString(line, -1))
		var rangeSeeds []int
		for i:=0; i <= len(checkSeeds)/2; i=i+2  {	
			r := makeRange(checkSeeds[i],checkSeeds[i]+checkSeeds[i+1] )
			rangeSeeds = append(rangeSeeds, r...)
			break;
		}
		return rangeSeeds, nil
	//}
	//return nil, fmt.Errorf("")
}

func getSrcDestLen(lines []string)([]SrcDestLen, error){
	var sdl []SrcDestLen
	var category Category
	for _, l := range lines{
		cat, err := getSourceDest(l)
		if(err == nil){
			category = cat
		}else{
			r := regexp.MustCompile("[0-9]+")
			strs := r.FindAllString(l, -1)
			if len(strs) > 0{
				ints := convertAllToint(strs)
				sdl = append(sdl, SrcDestLen{src: ints[1], dest: ints[0], len: ints[2], category: category})
			}			
		}
	}
	return sdl, nil
}

func getSourceDest(line string)(Category, error){
	found := strings.Index(line, "-to-")
	if found > -1 {	
		dest := strings.TrimRight(line[found+4:], " map:")
		return Category{source: line[0:found], dest: dest}, nil		
	}else{
		return Category{}, fmt.Errorf("No Number")
	}
}
	var (
		categories = [...]string{
			"seed",
			"soil",
			"fertilizer",
			"water",
			"light",
			"temperature",
			"humidity",
			"location",
		}
	)	
func (nl *NodeList) getNodes(steps []int, sdl []SrcDestLen, level int){
	//var nodes []Node
	category := Category{source: categories[level], dest: categories[level+1]}
	var levels []SrcDestLen
	for _, s := range sdl {
		if s.category == category{
			levels = append(levels, s)
		}
	}
	log.Printf("category: %v\n", category)
	// log.Printf("Level: %+v\n", levels)
	for _, l := range levels {
		for i:= 0; i < l.len; i++{
			// log.Printf("src: %v\n", l.src+i)
			if slices.Index(steps, l.src+i) > -1{
					nl.nodes= append(nl.nodes, Node{id: l.src+i, next: l.dest+i, category: l.category})
			}
		}
	}
	//check steps for not specified
	for _, s := range steps{
		if slices.IndexFunc(nl.nodes, func(n Node)bool{return n.id == s && n.category == category}) == -1{
			// log.Printf("FOUND>>>: %v/n", s)
			nl.nodes = append(nl.nodes, Node{id:s, next: s, category: category})
		}
	}
	nl.filterNodesByCategory(level)
	//Update steps
	steps = nil
	for _, n := range nl.nodes {
		steps = append(steps, n.next)
	}
	// log.Printf("steps: %+v\n", steps)
	// log.Printf("Nodes: %+v\n", nl.nodes)
	if level == 6 {
		//n.filterNodesByCategory(level)
		return
	}
	nl.getNodes(steps, sdl, level+1)
}

func (nl *NodeList) filterNodesByCategory(level int) {
	var ns []Node
	for _,n := range nl.nodes {
		if n.category.source == categories[level]{
			ns = append(ns, n)
		}
	} 
	nl.nodes = ns
}
// func makeNodes(src int, dest int, category Category)[]Node{

// }

func convertAllToint(strs []string) []int{
	var ints []int
	for _,n := range strs {
		num, err := strconv.Atoi(n)
		if err == nil{
			ints = append(ints, num)
		}
	}
	return ints
}
func makeRange(min, max int) []int {
	a := make([]int, max-min)
	for i := range a {
		a[i] = min + int(i)
	}
	return a
}
func printMemStats() {
 var mem runtime.MemStats
 runtime.ReadMemStats(&mem)
 log.Printf("Alloc = %v MiB, TotalAlloc = %v MiB, Sys = %v MiB, NumGC = %v\n",
 bToMb(mem.Alloc), bToMb(mem.TotalAlloc), bToMb(mem.Sys), mem.NumGC)
}

func bToMb(b uint64) uint64 {
 return b / 1024 / 1024
}