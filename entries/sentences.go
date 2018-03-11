package fukushu

import (
	"encoding/csv"
	"fmt"
	"github.com/gocarina/gocsv"
	"io"
	"os"
)

type EngSentence struct {
	ID       int    `csv:"id"`
	LangID   string `csv:"lang_id"`
	Sentence string `csv:"sentence"`
}

type JaSentence struct {
	ID       int    `csv:"id"`
	EngID    int    `csv:"eng_id"`
	Sentence string `csv:"sentence"`
}

type Pair struct {
	Eng *EngSentence
	Ja  *JaSentence
}

func (p Pair) ToString() string {
	return fmt.Sprintf(`%s
		%s`, p.Ja.Sentence, p.Eng.Sentence)
}

var (
	//local marshalled
	eng = []*EngSentence{}

	//mapped with ids
	Eng = map[int]*EngSentence{}
	Ja  = []*JaSentence{}
)

func init() {
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.LazyQuotes = true
		r.Comma = '\011'
		return r // Allows use tabs as delimiter and use quotes in CSV
	})

	engFile, err := os.OpenFile("en.csv", os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer engFile.Close()

	if err := gocsv.UnmarshalFile(engFile, &eng); err != nil {
		panic(err)
	}

	for _, en := range eng {
		Eng[en.ID] = en
	}

	jaFIle, err := os.OpenFile("ja.csv", os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer jaFIle.Close()

	if err := gocsv.UnmarshalFile(jaFIle, &Ja); err != nil {
		panic(err)
	}
}