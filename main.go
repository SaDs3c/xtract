package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	urlRegex := regexp.MustCompile(`https?://[^\s"'<>]+`)

	for scanner.Scan() {
		inputURL := strings.TrimSpace(scanner.Text())
		if inputURL == "" {
			continue
		}

		fmt.Printf("\nFetching: %s\n", inputURL)

		resp, err := http.Get(inputURL)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error fetching %s: %v\n", inputURL, err)
			continue
		}
		defer resp.Body.Close()

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading response from %s: %v\n", inputURL, err)
			continue
		}

		body := string(bodyBytes)
		foundURLs := urlRegex.FindAllString(body, -1)

		fmt.Println("Extracted URLs:")
		for _, link := range foundURLs {
			fmt.Println(link)
		}
	}
}
