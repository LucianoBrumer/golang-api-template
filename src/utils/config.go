package utils

import "os"

func Config() {
	os.Setenv("PORT", "3000")
}
