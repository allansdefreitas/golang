package main

import "fmt"

func main() {

	siteURLs := []string{"https://www.alura.com.br"}
	siteURLs = append(siteURLs, "https://httpbin.org/status/200")
	siteURLs = append(siteURLs, "https://httpbin.org/status/400")
	siteURLs = append(siteURLs, "https://ieadpe.org")

	for i := 0; i < len(siteURLs); i++ {

		fmt.Println("i: ", i, siteURLs[i])
	}

	fmt.Print("\n\n\n")

	for i, url := range siteURLs {

		fmt.Println(i, url) // Position (i) and Value itself (url)
		// fmt.Println(url)
		// fmt.Println(siteURLs[i])
	}

}
