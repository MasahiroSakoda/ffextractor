package ffmpeg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsInstalled(t *testing.T) {
	assert.True(t, IsInstalled())
}
