package config

import "os"

func Config() {
	os.Setenv("PORT", "3000")
}
