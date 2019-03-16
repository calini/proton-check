package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const url = "https://mail.protonmail.com/api/users/available"

func main() {

	client := &http.Client{}

	scanner := bufio.NewScanner(os.Stdin)

	var available []string
	for scanner.Scan() {
		str := scanner.Text()

		if isAvailable(client, str) {
			fmt.Printf("[%s] available!\n", str)
			available = append(available, str)
		} else {
			fmt.Printf("[%s] unavailable :(\n", str)
		}

		time.Sleep(100 * time.Millisecond)
	}

	// Print only available at the end
	if len(available) > 0 {
		fmt.Printf("Available ones:\n")

		for _, x := range available {
			fmt.Printf("%s\n", x)
		}
	}

}
func isAvailable(client *http.Client, x string) bool {
	req, _ := http.NewRequest("GET", url+"?Name="+x, nil)
	req.Header.Set("x-pm-apiversion", "3")
	req.Header.Set("x-pm-appversion", "Web_3.15.20")

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("[%s] idk what happened :(\n", x)
		log.Fatal(err)
		return false
	}

	return res.StatusCode == 200
}
