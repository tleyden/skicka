package main

import (
	"fmt"
	"os"
)

func setprop(args []string) int {

	var drivePaths []string
	errs := 0

	for _, arg := range args {
		switch {
		default:
			drivePaths = append(drivePaths, arg)
		}
	}

	if len(drivePaths) == 0 {
		fmt.Printf("setprop: drive path cannot be empty.\n")
		fmt.Printf("Usage: setprop drive_path ...\n")
		fmt.Printf("Run \"skicka help\" for more detailed help text.\n")
		return 1
	}

	for _, path := range drivePaths {

		file, err := gd.GetFile(path)
		fmt.Printf("file: %+v", file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "skicka: Error getting file found at path %s: %v\n", path, err)
			errs++
		}

		//currentProperties := file.Properties
		//fmt.Printf("file properties: %+v", currentProperties)
		//
		//property := gdrive.Property{
		//	Key: "testprop",
		//	Value: "testval",
		//}
		//
		//file.Properties = append(currentProperties, property)

		if err := gd.UpdateProperty(file, "testprop", "testval3"); err != nil {
			fmt.Fprintf(os.Stderr, "skicka: Error setting property for file %s: %v\n", path, err)
			errs++
		}

	}

	return errs

}
