package formats_test

import (
	"testing"

	"github.com/mentai-mayo/gotools/formats"
	"github.com/stretchr/testify/assert"
)

func TestCountFormatArgs(t *testing.T) {
	var result int
	result = formats.CountFormatArgs("%%%")
	if result != 2 {
		t.Errorf("expect 2, got %d\n", result)
	}
	result = formats.CountFormatArgs("%s%d%%%x")
	if result != 4 {
		t.Errorf("expect 4, got %d\n", result)
	}
}

func TestFormats(t *testing.T) {
	var result []string
	var err error

	// #1
	result, err = formats.Formats("this %s correct", "is")
	assert.NoError(t, err)
	assert.Equal(t, []string{"this is correct"}, result)

	// #2
	result, err = formats.Formats("this %s correct", "is", "yattaze")
	assert.NoError(t, err)
	assert.Equal(t, []string{"this is correct", "yattaze"}, result)

	// #3
	result, err = formats.Formats("this %s correct", "is", "totemo %s %s", "ureshi-", "na")
	assert.NoError(t, err)
	assert.Equal(t, []string{"this is correct", "totemo ureshi- na"}, result)

	// #4
	result, err = formats.Formats("this %s %s correct", "is" /*, "not" */)
	assert.Error(t, err)

	// #5
	result, err = formats.Formats("not string -> %d, %X", 12, 15)
	assert.NoError(t, err)
	assert.Equal(t, []string{"not string -> 12, F"}, result)
}
