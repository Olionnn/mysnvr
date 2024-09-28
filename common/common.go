package common

import (
	"os"

	"github.com/Olionnn/mysnvr/helper"
)

func CreateDefaultEnv() {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		os.Create(".env")
	}

	helper.RtspUrl = os.Getenv("rtspURI")
}
