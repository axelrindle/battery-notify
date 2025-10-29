package main

import (
	"fmt"
	"os"
	"path"
)

var layout []string

func init() {
	layout = []string{
		"charge_now",
		"charge_full",
	}
}

func validateDevice(dir string) error {
	for _, file := range layout {
		p := path.Join(dir, file)
		stat, err := os.Stat(p)

		if err != nil {
			return err
		}
		if stat.Mode().IsRegular() {
			return fmt.Errorf("required property %s not found in %s", file, dir)
		}
	}

	return nil
}
