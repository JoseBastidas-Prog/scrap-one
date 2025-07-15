package main

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	// Get the bs to usd exchange value from the Central Bank of Venezuela
	resp, err := http.Get("https://www.bcv.org.ve")

	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}

	defer resp.Body.Close()
	fmt.Println("Response Status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		//string.Contains hace que el scope sea la línea completa
		location := strings.Split(scanner.Text(), `<div id="dolar" class="col-sm-12 col-xs-12 ">`)
		fmt.Println("Location:", location)
		if strings.Contains(scanner.Text(), `id="dolar"`) {
			fmt.Println("Valor del Dolar:", strings.Split(scanner.Text(), ">")[1])
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
}
