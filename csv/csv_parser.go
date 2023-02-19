package csv

import (
	"encoding/csv"
	"github.com/jedib0t/go-pretty/v6/table"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

type People struct {
	Id         int
	FirstName  string
	LastName   string
	Email      string
	Profession string
}

const FileName = "./generated-file.csv"

func ProcessFile() {
	absPath, _ := filepath.Abs(FileName)
	f, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}

	peopleList := scanFile(f)

	printTheOutPut(peopleList)

}

func scanFile(file *os.File) []People {
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	people := make([]People, len(data)-1)
	for i, line := range data {
		if i > 0 {
			id, err := strconv.Atoi(line[0])
			if err != nil {
				log.Fatal(err)
			}
			people[i-1] = People{
				Id:         id,
				FirstName:  line[1],
				LastName:   line[2],
				Email:      line[3],
				Profession: line[3],
			}
		}
	}

	return people
}

func printTheOutPut(people []People) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "First Name", "Last Name", "Email", "Profession"})
	for _, val := range people {
		t.AppendRows([]table.Row{{val.Id, val.FirstName, val.LastName, val.Email, val.Profession}})
		t.AppendSeparator()
	}

	t.Render()

}
