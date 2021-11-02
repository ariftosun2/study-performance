package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	//"strings"
	//"encoding/hex"
)

type User struct {
	firstName string
	lastName  string
	workhours int
}

func sortingEmploye(sorting []User) []User {
	var sortingUser []User
	for i, v := range sorting {
		for _, b := range sorting[i:] {
			if v.workhours > b.workhours {

				sortingUser = append(sortingUser, v)
				break
			}
		}
	}
	return sortingUser
}

func RecordFile(recordsl [][]string) []User {
	var readdosya []User
	for _, record := range recordsl {

		hours, _ := strconv.Atoi(record[2])
		//fmt.Printf("%d\n", hours)
		user := User{
			firstName: record[0],
			lastName:  record[1],
			workhours: hours,
		}
		readdosya = append(readdosya, user)
	}
	//fmt.Printf("%d\n", readdosya)
	return readdosya

}

func workingHoursSum(studyuser []User) []User {
	var working []User
	for i, v := range studyuser {
		for _, b := range studyuser[i:] {
			if v.firstName == b.firstName && v.lastName == b.lastName {
				sumworkhours := v.workhours + b.workhours
				v.workhours = sumworkhours

			}
		}
		working = append(working, v)

	}
	return working
}

func calculate(readfile string, createfilename string) {
	records, err := readData(readfile)

	if err != nil {
		log.Fatal(err)
	}
	recordUser := RecordFile(records)
	maprecodscalculate := workingHoursSum(recordUser)
	sortingUsur := sortingEmploye(maprecodscalculate)
	writeUserFile(createfilename, sortingUsur)

}

func main() {

	calculate("books.csv", "textcvs")

}

func readData(fileName string) ([][]string, error) {

	f, err := os.Open(fileName)

	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	//fmt.Println(r)

	records, err := r.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}

func writeUserFile(filename string, userwork []User) error {
	file, err := os.OpenFile(filename, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	dataWriter := bufio.NewWriter(file)
	for _, v := range userwork {
		s := fmt.Sprintf("%s,%s,%d\n", v.lastName, v.firstName, v.workhours)
		_, err = dataWriter.WriteString(s)
		if err != nil {
			return err
		}
	}

	dataWriter.Flush()
	return nil

}
