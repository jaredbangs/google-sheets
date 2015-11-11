package sheets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadAllDoesNotError(t *testing.T) {

	sheetsUrl := "https://docs.google.com/spreadsheets/u/0/d/16PQqxZN1TZIlqVeWNHV9JI2V3295F_Vwq1xsv8o9UGE/export?format=csv&id=16PQqxZN1TZIlqVeWNHV9JI2V3295F_Vwq1xsv8o9UGE&gid=0"

	reader := Reader{}

	records, err := reader.ReadAll(sheetsUrl)

	assert.Nil(t, err, "error should be nil")
	assert.NotNil(t, records, "records should not be nil")
}

func TestReadAllRecordsAreCorrect(t *testing.T) {

	sheetsUrl := "https://docs.google.com/spreadsheets/u/0/d/16PQqxZN1TZIlqVeWNHV9JI2V3295F_Vwq1xsv8o9UGE/export?format=csv&id=16PQqxZN1TZIlqVeWNHV9JI2V3295F_Vwq1xsv8o9UGE&gid=0"

	reader := Reader{}

	records, _ := reader.ReadAll(sheetsUrl)

	assert.Equal(t, "ABC", records[1][1], "B2 should be ABC")
	assert.Equal(t, "123", records[2][2], "C3 should be 123")
}
