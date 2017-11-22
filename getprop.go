package main

import (
	"fmt"
	"os"
)

func getprop(args []string) int {

	printUsage := func() {
		fmt.Printf("getprop: drive path cannot be empty.\n")
		fmt.Printf("Usage: getprop drive_path ...\n")
		fmt.Printf("Run \"skicka help\" for more detailed help text.\n")
	}

	var drivePaths []string
	errs := 0


	for _, arg := range args {
		switch {
		default:
			drivePaths = append(drivePaths, arg)
		}
	}

	if len(drivePaths) == 0 {
		printUsage()
		return 1
	}

	for _, path := range drivePaths {

		file, err := gd.GetFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "skicka: Error getting file found at path %s: %v\n", path, err)
			errs++
		}

		fmt.Printf("Properties for file: %v\n", file.Path)
		for _, property := range file.Properties {
			fmt.Printf("%s=%s\n", property.Key, property.Value)
		}

	}

	return errs

}

