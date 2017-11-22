package main

import (
	"fmt"
	"os"
)

func print_usage() {
	fmt.Printf("setprop: drive path cannot be empty.\n")
	fmt.Printf("Usage: setprop key value drive_path ...\n")
	fmt.Printf("Run \"skicka help\" for more detailed help text.\n")
}

func setprop(args []string) int {

	var drivePaths []string
	errs := 0

	// Need at least key + value + drive_path
	if len(args) < 3 {
		print_usage()
		return 1
	}

	key := args[0]
	val := args[1]

	// Shift slice to the right by two slots
	args = args[2:]

	for _, arg := range args {
		switch {
		default:
			drivePaths = append(drivePaths, arg)
		}
	}

	if len(drivePaths) == 0 {
		print_usage()
		return 1
	}

	for _, path := range drivePaths {

		file, err := gd.GetFile(path)

		fmt.Printf("file: %+v", file)

		if err != nil {
			fmt.Fprintf(os.Stderr, "skicka: Error getting file found at path %s: %v\n", path, err)
			errs++
		}

		if err := gd.UpdateProperty(file, key, val); err != nil {
			fmt.Fprintf(os.Stderr, "skicka: Error setting property for file %s: %v\n", path, err)
			errs++
		}

	}

	return errs

}
