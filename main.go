package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/korylprince/ipnetgen"
)
func main() {
	var cidrInput string

	// Check if CIDR is provided as a command line argument
	if len(os.Args) > 1 {
		cidrInput = os.Args[1]
	} else {
		// Read CIDR from stdin if not provided as an argument
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			cidrInput = scanner.Text()
		}
	}

	cidrInput = strings.TrimSpace(cidrInput)

	if cidrInput == "" {
		fmt.Println("No Input")
		return
	}

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
