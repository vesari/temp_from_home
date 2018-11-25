package main

import "fmt"

func main() {
	a := []string{"A", "B"}
	sep := ","
	result := myJoin(a, sep)
  fmt.Printf("%s\n",result)
}

func myJoin(a []string, sep string) string {

	output := ""
	for i, v := range a {
		output += v
		if i != len(a)-1 {
			output += sep

		}

	}
	return output
}
