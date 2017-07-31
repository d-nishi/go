package main

import "fmt"
import "strings"

func reversestring(domain []string) []string {

	for i, j := 0, len(domain)-1; i < j; i, j = i+1, j-1 {
		domain[i], domain[j] = domain[j], domain[i]
	}
	return domain
}

func main() {

	knowndomain := []string{
		"uk",
		"co.uk",
		"bbc.co.uk",
		"gov.uk",
		"com",
		"ca.gov",
	}

	inputdomain := []string{
		"www.google.com",
		"www.bbc.co.uk",
	}

	root := buildtree(knowndomain)
	fmt.Println(root)

	for _, inputdomain := range inputdomain {
		fmt.Println(longestMatch(root, inputdomain))
	}

}

func buildtree(knowndomain []string) map[string]interface{} {
	root := make(map[string]interface{})

	for i := range knowndomain {
		domain := strings.Split(knowndomain[i], ".")
		domain = reversestring(domain)
		current := root
		for _, domain := range domain {
			_, ok := current[domain]
			if ok == false {
				current[domain] = make(map[string]interface{})
			}
			current = current[domain].(map[string]interface{})
		}
	}
	return root
}

func longestMatch(root map[string]interface{}, inputdomain string) string {
	output := make([]string, 0)
	domain := strings.Split(inputdomain, ".")
	domain = reversestring(domain)
	current := root
	for _, domain := range domain {
		_, ok := current[domain]
		if ok == false {
			output = append(output, domain)
			break
		}
		output = append(output, domain)
		current = current[domain].(map[string]interface{})
	}

	reversestring(output)
	output_new := strings.Join(output, ".")
	return output_new
}
