package rotation

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/flamefatex/log"
	"github.com/flamefatex/log/impl/w_zap"
)

func TestRotation(t *testing.T) {
	dir, _ := ioutil.TempDir("", "test")
	defer os.RemoveAll(dir) // clean up

	log.InitLogger(w_zap.NewZapLogger, &log.Config{
		ServiceName:    "test",
		Level:          log.LevelDebug,
		EnableRotation: true,
		EnableConsole:  true,
		RotationConfig: &RotationConfig{
			Filename: filepath.Join(dir, "error.log"),
			MaxSize:  1,
		},
	})
	s := strings.Repeat("1", 600*1024)
	log.Debug(s)
	// trigger log rotation
	log.Debug(s)

	files, _ := ioutil.ReadDir(dir)
	assert.Equal(t, 2, len(files))
}
