package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	// inputFile, outputFile string
	inputFile string
	rootCmd   = &cobra.Command{
		Use:   "namemash",
		Short: "Creating a user name list for brute force attacks",
		Run: func(cmd *cobra.Command, args []string) {
			file, err := os.Open(inputFile)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := scanner.Text()
				line = FilterNonAlphaCharacters(line)

				tokens := strings.Fields(line)
				if len(tokens) < 1 {
					continue
				}

				fname := tokens[0]
				var lname string
				if len(tokens) == 2 {
					lname = tokens[1]
				} else {
					lname = strings.Join(tokens[1:], "")
				}
				fmt.Println(fname, lname)

				candidates := BuildCandidates(fname, lname)
				for _, candidate := range candidates {
					fmt.Println(candidate)
				}
			}

			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
		},
	}
)

func FilterNonAlphaCharacters(s string) string {
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		b := s[i]
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			('0' <= b && b <= '9') ||
			b == ' ' {
			result.WriteByte(b)
		}
	}
	return result.String()
}

func BuildCandidates(fname string, lname string) []string {
	var candidates []string
	candidates = append(candidates, fname+lname)
	candidates = append(candidates, lname+fname)
	candidates = append(candidates, fname+"."+lname)
	candidates = append(candidates, lname+"."+fname)
	candidates = append(candidates, lname+string(fname[0]))
	candidates = append(candidates, string(fname[0])+lname)
	candidates = append(candidates, string(lname[0])+fname)
	candidates = append(candidates, string(fname[0])+"."+lname)
	candidates = append(candidates, string(lname[0])+"."+fname)
	candidates = append(candidates, fname)
	candidates = append(candidates, lname)
	return candidates
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&inputFile, "input", "", "File to read initial names from")
	// rootCmd.PersistentFlags().StringVar(&outputFile, "output", "", "File to write output to (optional)")
	rootCmd.MarkPersistentFlagRequired("input")
}
