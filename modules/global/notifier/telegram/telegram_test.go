package telegram

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetName(t *testing.T) {
	n := &GlobalNotifierTelegram{}
	require.Equal(t, MODULE_NAME, n.GetName())
}

func TestGetConfig(t *testing.T) {
	n := &GlobalNotifierTelegram{}
	require.Equal(t, &Config{}, n.GetConfig())
}
func TestClose(t *testing.T) {
	n := &GlobalNotifierTelegram{}
	assert.Equal(t, nil, n.Close())
}

func TestInitPipe(t *testing.T) {
	n := &GlobalNotifierTelegram{}
	bufInput := bytes.NewBuffer([]byte("hello world"))
	bufOutput := &bytes.Buffer{}
	require.NoError(t, n.InitPipe(bufOutput, bufInput))
}
