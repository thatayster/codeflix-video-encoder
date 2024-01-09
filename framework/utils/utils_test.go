package utils_test

import (
	"encoder/framework/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsJson(t *testing.T) {
	json := `{
				"id": "123",
				"file_path": "some_file.txt",
				"status": "ok"
			}`

	err := utils.IsJson(json)
	require.Nil(t, err)
	
	json = `not a json`
	err = utils.IsJson(json)
	require.Error(t, err)
}