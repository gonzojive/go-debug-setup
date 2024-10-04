package main

import (
	"debug/dwarf"
	"debug/elf"
	"fmt"
	"os"
	"sort"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-debug-files",                                                    // Updated command name
	Short: "Extract source file names from a Go binary with debug information", // Updated description
}

var filesCmd = &cobra.Command{
	Use:   "files",
	Short: "Print a list of source files", // Updated description
	Run: func(cmd *cobra.Command, args []string) {
		binaryPath, err := cmd.Flags().GetString("executable")
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		sourceFiles, err := getSourceFilesFromBinary(binaryPath)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		for _, file := range sourceFiles {
			fmt.Println(file)
		}
	},
}

var filesExecutableFlag = filesCmd.Flags().StringP("executable", "p", "", "Path to the Go executable file")

func init() {
	rootCmd.AddCommand(filesCmd)

	filesCmd.MarkFlagRequired("executable")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// getSourceFilesFromBinary extracts the source file names from the debug information (DWARF)
// of a Go binary. It takes the path to the binary as input and returns a sorted list of
// unique source file names, or an error if any occur during the process.
func getSourceFilesFromBinary(binaryPath string) ([]string, error) {
	f, err := elf.Open(binaryPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	dwarfData, err := f.DWARF()
	if err != nil {
		return nil, err
	}

	// Use a map to store unique source file names
	sourceFiles := make(map[string]struct{})

	// Iterate over DWARF info entries
	reader := dwarfData.Reader()
	for {
		entry, err := reader.Next()
		if err != nil {
			return nil, err
		}
		if entry == nil {
			break // reached end of DWARF info
		}

		// Check if the entry is a compile unit (CU)
		if entry.Tag == dwarf.TagCompileUnit {
			// Get the corresponding line reader for the CU
			lineReader, err := dwarfData.LineReader(entry)
			if err != nil {
				return nil, err
			}
			if lineReader == nil {
				continue
			}

			// Iterate over the line program
			for _, file := range lineReader.Files() {
				if file == nil {
					continue
				}
				// Add the source file name to the map
				sourceFiles[file.Name] = struct{}{}
			}
		}
	}

	// Convert the map keys to a slice
	var fileList []string
	for file := range sourceFiles {
		fileList = append(fileList, file)
	}
	sort.Strings(fileList) // Sort the file list

	return fileList, nil
}
