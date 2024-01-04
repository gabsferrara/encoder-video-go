package utils_test

import (
	"testing"

	"github.com/encoder-video-go/framework/utils"
	"github.com/stretchr/testify/require"
)

func TestIsJson(t *testing.T) {
	json := `{
		"id": "test_id",
		"file_path": "tet_file_path",
		"status": "test_status"
	}`

	err := utils.IsJson(json)
	require.Nil(t, err)

	json = `test`
	err = utils.IsJson(json)
	require.Error(t, err)
}
