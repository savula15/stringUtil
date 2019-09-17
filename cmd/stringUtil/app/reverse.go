package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var reverseCmd = &cobra.Command{
	Use:   "reverse",
	Short: "reverse string and chars within word in-place",
	Long: `reverse string and chars within word in-place. 
	For example:
	'today is going to be awesome' will result in
	'emosewa eb ot gniog si yadot`,

	Run: func(cmd *cobra.Command, args []string) {
		testStr, _ := rootCmd.Flags().GetString("testString")
		if testStr != "" {
			fmt.Println("Calling reverse fucntion")
			fmt.Println(reverse(testStr))
		}
	},
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func init() {
	rootCmd.AddCommand(reverseCmd)
}
