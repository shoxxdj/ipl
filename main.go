package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/korylprince/ipnetgen"
)

func displayIps(cidrInput string){
	cidrInput = strings.TrimSpace(cidrInput)
	if(strings.Contains(cidrInput,"/")){
		gen, err := ipnetgen.New(cidrInput)
		if err != nil {
			//do something with err
		}
		for ip := gen.Next(); ip != nil; ip = gen.Next() {
			fmt.Println(ip)
		}
	}else{
		fmt.Println(cidrInput)
	}
}

func main() {
	var cidrInput string
	var input string 

	//Check here if one or more, and  convert to strings
	if len(os.Args) > 1 {
		cidrInput = os.Args[1]
		displayIps(cidrInput)
	}else{
		for {
			_, err := fmt.Scanln(&input)
			displayIps(input)
			if err != nil {
				break
			}
		}
	}
}
