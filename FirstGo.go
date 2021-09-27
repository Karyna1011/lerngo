package main

import (
	"fmt"

	"io/ioutil"
	"log"
	"net/http"
	//"os"
	"strings"
	"math")

type items struct {
	items2 map[string]string{
	"id": "max_contracts_count",
	"type": "key-value-entries"
}

}
type mydata struct {
	data map[string] items {
		"data":
}


type DataArray []mydata

var perPage =3

var numbers = [10]string{
	"First", "Second", "Third", "Fourth", "Fifth", "Sixth", "Seventh", "Eighth", "Ninth", "Tenth"}

func paginate(currentPage int, total int) []string {
	firstEntry := (currentPage - 1) * perPage
	lastEntry := firstEntry + perPage

	if lastEntry > total {
		lastEntry = total
	}
    slice:=numbers[firstEntry:lastEntry]

	return slice
}
func pageCount (total float64) int {
	return int(math.Ceil(float64(10) / float64(perPage)))
}
func main() {
	max := pageCount(10)
	for i := 1; i <= max ; i++ {
		fmt.Println(paginate(i, 10))
	}
 var array[1000] string
	response, err := http.Get("http://localhost:8000/_/api/v3/key_values")

	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	dataInBytes, err := ioutil.ReadAll(response.Body)
	pageContent := string(dataInBytes)
    //fmt.Println(pageContent)

	words := strings.Fields(pageContent)
	for i, word := range words {
		array[i]=word;
		fmt.Printf("Word %d is: %s\n", i, word)}

}