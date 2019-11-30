package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("need to pass in the name of the csv file")
	}
	filename := os.Args[1]
	csvStuff(filename)
}

func csvStuff(filename string) {
	inLayout := "02 Jan 06"
	outLayout := "01/02/2006"
	f, err := os.Open(filename)
	check(err)
	fileReader := bufio.NewReader(f)
	r := csv.NewReader(fileReader)
	//read header line first and ignore it
	r.Read()
	fmt.Println("Date,Payee,FI Payee,Amount,Memo?,Category")
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		check(err)
		date, err := time.Parse(inLayout, record[0])
		check(err)
		formattedDate := date.Format(outLayout)
		creditDebit := "debit"
		moneyOut, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			creditDebit = "credit"
		}
		moneyIn, err := strconv.ParseFloat(record[2], 64)
		fmt.Printf("\"")
		fmt.Printf(formattedDate)
		fmt.Printf("\",\"")
		fmt.Printf(record[1])
		fmt.Printf("\",\"")
		fmt.Printf(record[1])
		fmt.Printf("\",\"")
		fmt.Printf("%f", moneyIn-moneyOut)
		fmt.Printf("\",\"")
		fmt.Printf(creditDebit)
		fmt.Printf("\",\"")
		fmt.Printf("\"")
		fmt.Printf("\n")
	}

}
