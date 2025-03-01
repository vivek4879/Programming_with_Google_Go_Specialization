/*Write a program which reads information from a file and represents it in a slice of structs.
Assume that there is a text file which contains a series of names.
Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.
Your program will define a name struct which has two fields, fname for the first name, and lname for the last name.
Each field will be a string of size 20 (characters).
Your program should prompt the user for the name of the text file.
Your program will successively read each line of the text file and create a struct which contains the first and last names
found in the file. Each struct created will be added to a slice, and after all lines have been read from the file,
your program will have a slice containing one struct for each line in the file. After reading all lines from the file,
your program should iterate through your slice of structs and print the first and last names found in each struct.*/

package main

import (
	"fmt"
	"log"
	"os"
)

type Name struct {
	fname string
	lname string
}

func main() {
	var file_name string
	fmt.Println("Enter file name")
	fmt.Scan(&file_name)
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fileinfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	filesize := fileinfo.Size()
	// fmt.Print(filesize)
	barr := make([]byte, filesize)
	nb, err := file.Read(barr)
	if err != nil {
		log.Fatal(err)
	}
	start := 0
	end := 20
	res := make([]Name, 0)
	// fmt.Print(nb)
	for nb > 0 {
		var current Name
		current.fname = string(barr)[start:end]
		start = end + 2
		end = end + 20
		current.lname = string(barr)[start:end]
		start = end + 2
		end = end + 20
		nb = nb - 42
		// fmt.Print(nb)
		res = append(res, current)
	}
	for _, name := range res {
		fmt.Printf("First Name: %s, Last Name:%s\n", name.lname, name.lname)
	}

}
