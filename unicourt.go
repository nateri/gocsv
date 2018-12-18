package main

import (
	//"encoding/csv"
	"fmt"
	//"io"
	"os"
	"path/filepath"

	"github.com/gocarina/gocsv"
)

type Entry struct {
	CaseURL                 string `csv:"Case URL"`
	CaseID                  string `csv:"Case ID"`
	CaseName                string `csv:"Case Name"`
	CaseNumber              string `csv:"Case Number"`
	FilingDate              string `csv:"Filing Date"`
	CaseTypeCategory        string `csv:"Case Type Category"`
	CaseTypeSubCategory     string `csv:"Case Type Sub Category"`
	CaseType                string `csv:"Case Type"`
	CaseStatusCategory      string `csv:"Case Status Category"`
	CaseStatus              string `csv:"Case Status"`
	Jurisdiction            string `csv:"Jurisdiction"`
	Courthouse              string `csv:"Courthouse"`
	AllPartyNames           string `csv:"All Party Names"`
	PartyName               string `csv:"Party Name"`
	PartyType               string `csv:"Party Type"`
	PartyEntityType         string `csv:"Party Entity Type"`
	PartyRepresentationType string `csv:"Party Representation Type"`
	PartyAttorney           string `csv:"Party Attorney"`
}

func main() {
	inFilename := `C:\test.csv`
	if len(os.Args) >= 2 {
		if len(os.Args[1]) > 0 {
			inFilename = os.Args[1]
		}
	}

	inFile, err := os.OpenFile(inFilename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println("Failed to open file for reading: ", inFilename)
		return
	}
	defer inFile.Close()

	inFileDir, inFilenameOnly := filepath.Split(inFilename)
	fmt.Println("inFileDir: ", inFileDir, ", inFilename: ", inFilename)
	outFilename := filepath.Join(inFileDir, "__"+inFilenameOnly)

	outFile, err := os.OpenFile(outFilename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		fmt.Println("Failed to open file for writing: ", outFilename)
		return
	}
	defer outFile.Close()

	entries := []*Entry{}
	if err := gocsv.UnmarshalFile(inFile, &entries); err != nil { // Load clients from file
		panic(err)
	}

	tmp := entries[:0]
	i := 0
	for _, e := range entries {
		if e.PartyType != "Defendant" {
			continue
		}
		if e.PartyEntityType != "Company" {
			continue
		}
		i += 1
		fmt.Println("Hello [", i, "] ", e.PartyName)
		tmp = append(tmp, e)
	}
	entries = tmp

	err = gocsv.MarshalFile(&entries, outFile) // Use this to save the CSV back to the file
	if err != nil {
		panic(err)
	}
}
