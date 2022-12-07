package Functions

import (
	"Hangman-web/HangmanModule"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

func CsvReader() [][]string {
	f, err := os.Open("web/assets/data/data.csv")

	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	var r [][]string
	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// do something with read line
		r = append(r, rec)
	}
	f.Close()

	return r
}

func CsvWritter(data HangmanModule.Session) {
	oldCsv := CsvReader()
	file, err := os.Create("web/assets/data/data.csv")
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	w := csv.NewWriter(file)
	for _, line := range oldCsv {
		if err := w.Write(line); err != nil {
			log.Fatalln("error writing record to file", err)
		}
		w.Flush()
	}
	defer w.Flush()
	// Using Write
	if data.NLoose == 0 {
		row := []string{data.Email, data.Mdp, data.Pseudo, strconv.Itoa(data.NWin), strconv.Itoa(data.NLoose), strconv.Itoa(1), strconv.Itoa(data.NWin + data.NLoose), strconv.Itoa(data.Points)}
		if err := w.Write(row); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	} else {
		row := []string{data.Email, data.Mdp, data.Pseudo, strconv.Itoa(data.NWin), strconv.Itoa(data.NLoose), strconv.Itoa(data.NWin / data.NLoose), strconv.Itoa(data.NWin + data.NLoose), strconv.Itoa(data.Points)}
		if err := w.Write(row); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}

}

func CsvEditor(data HangmanModule.Session) {
	oldCsv := CsvReader()
	file, err := os.Create("web/assets/data/data.csv")
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	w := csv.NewWriter(file)
	for _, line := range oldCsv {
		if line[0] == data.Email {
			log.Println(data.Points)
			line[4] = strconv.Itoa(data.NLoose)
			line[3] = strconv.Itoa(data.NWin)
			line[5] = strconv.Itoa(data.Ratio)
			line[6] = strconv.Itoa(data.NLoose + data.NWin)
			line[7] = strconv.Itoa(data.Points)
		}
		if err := w.Write(line); err != nil {
			log.Fatalln("error writing record to file", err)
		}
		w.Flush()
	}
}
