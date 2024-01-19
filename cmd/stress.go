/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/aluferraz/go-expert-stress-test/usecases/stresstest"

	"github.com/spf13/cobra"
)

var url string
var requests int
var concurrency int

// stressCmd represents the stress command
var stressCmd = &cobra.Command{
	Use:   "stress",
	Short: "Stress test an URL",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Stress Testing %s, Requests %d, Concurrency %d...\n", url, requests, concurrency)
		uc := stresstest.NewStressTest(stresstest.StressTestDTOInput{
			Url:         url,
			Requests:    requests,
			Concurrency: concurrency,
		})
		res, err := uc.Execute()
		if err != nil {
			fmt.Printf("Error executing test case %s\n", err.Error())
			return
		}
		granTotal := 0
		for statusCode, count := range res.Results {
			fmt.Printf("Total of request with Status code[%d]: %d \n", statusCode, count)
			granTotal += count
		}
		fmt.Printf("Total of request with: %d \n", granTotal)
		fmt.Printf("Execution time %fs\n", res.ExecutionTime.Seconds())

	},
}

func init() {
	rootCmd.AddCommand(stressCmd)

	// Here you will define your flags and configuration settings.
	stressCmd.Flags().StringVar(&url, "url", "", "URL to test")
	stressCmd.Flags().IntVar(&requests, "requests", 0, "Total number of requests")
	stressCmd.Flags().IntVar(&concurrency, "concurrency", 1, "Number of concurrent requests")
	stressCmd.MarkFlagRequired("url")
	stressCmd.MarkFlagRequired("requests")
	stressCmd.MarkFlagRequired("concurrency")
}
