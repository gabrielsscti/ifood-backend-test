package clients

import (
	"github.com/joho/godotenv"
	"path"
	"path/filepath"
	"runtime"
)

func getBasePath() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Join(filepath.Dir(d), "..")
}

func TryLoadEnvironmentFile() {
	godotenv.Load(getBasePath() + "/.env")
}
