package main

import (
	"fmt"
	"os"
	"sort"
)

func saveToFile(filePath string, values []string) {
	f, err := os.Create(filePath)
	if err != nil {
		return
	}
	defer f.Close()
	for _, value := range values {
		fmt.Fprint(f, value+"\n")
	}
}

func main() {
	var n int
	var list []string
	var temp string
	var i int = 0
	fmt.Println("Ingrese la cantidad de cadenas")
	fmt.Scanln(&n)
	fmt.Println("Ingresa ", n, "cadenas")
	for i < n {
		fmt.Scanln(&temp)
		list = append(list, temp)
		i = i + 1
	}
	sort.Sort(sort.Reverse(sort.StringSlice(list)))
	saveToFile("descendente.txt", list)
	sort.Strings(list)
	saveToFile("asecendente.txt", list)

}
