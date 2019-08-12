package services_test

import (
	"testing"

	"github.com/fenr1s/back-end-take-home/domain/services"
	"github.com/stretchr/testify/assert"
)

func TestFileReader(t *testing.T) {
	t.Run("ShouldReadFile", shouldReadFile)
	t.Run("CantFindFile", cantFindFile)
}

func shouldReadFile(t *testing.T) {
	fileReader := &services.FileReader{}
	lines, err := fileReader.ReadFromFile("../../data/test/airports.csv")
	assert.Equal(t, len(lines), 6)
	assert.Nil(t, err)
}

func cantFindFile(t *testing.T) {
	fileReader := &services.FileReader{}
	_, err := fileReader.ReadFromFile("data/test/airports.csv")
	assert.NotNil(t, err)
}
