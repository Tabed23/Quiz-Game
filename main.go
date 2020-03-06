package main

import (
	"encoding/csv"
	"flag"
	"os"
	"fmt"
	"strings"
)
func main(){
	
	csvFilename :=  flag.String("csv", "problems.csv", "Csv File of Q/A")
	flag.Parse()
	file, err := os.Open(*csvFilename)

	if err != nil {
		exit("fail to open the Csv flile")
		
	}
	reader :=  csv.NewReader(file)
	lines, err :=  reader.ReadAll()
	if err != nil{
		exit("fail to read the file")
	}
	problems := cvsParser(lines)
	fmt.Printf("you have scored %d out of %d \n", checkAnswers(problems), len(problems))
}
func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
}

type  problems struct{
	question string
	answers string
}

func cvsParser(lines[][]string)[]problems{

	retur :=  make([] problems, len(lines))

	for key, line := range lines {
		retur[key] = problems{
			question: line[0],
			answers: strings.TrimSpace(line[1]),
		}
	}
	return retur
}
func checkAnswers(prob [] problems)int{
	var correct int
	var ans string
	for i, p := range prob{
		fmt.Printf("Problems #%d: %s=\n", i+1, p.question)
		fmt.Scanf("%s\n", &ans)
		if ans == p.answers{
			correct++
		}
	}
	return correct
}