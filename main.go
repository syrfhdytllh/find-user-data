package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	names := []string{"Tono", "Reza", "Tibo", "Shog"}
	address := []string{"Jakpus", "Jakbar", "Jaktim", "Jaksel"}
	reason := []string{"A", "B", "C", "D"}
	job := []string{"Programmer", "Pembalap", "Petinju", "Pengarang"}

	payload := os.Args[1]
	dataIndex := FindIndex(names, payload)

	addressData := address[dataIndex]

	reasonData := reason[dataIndex]

	jobData := job[dataIndex]

	fmt.Println("name : ", payload)
	fmt.Println("Address : ", addressData)
	fmt.Println("Job : ", jobData)
	fmt.Println("Reason : ", reasonData)
}

func FindIndex(names []string, element string) int {
	for ind, val := range names {
		if val == element {
			return ind
		}
	}
	// if not found, return -1
	return -1
}
