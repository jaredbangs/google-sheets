package sheets

import (
	"encoding/csv"
	"net/http"
)

type Reader struct {
}

func (r *Reader) ReadAll(sheetsCSVUrl string) (records [][]string, err error) {

	response, err := http.Get(sheetsCSVUrl)

	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(response.Body)

	reader.FieldsPerRecord = -1

	records, err = reader.ReadAll()

	return
}
