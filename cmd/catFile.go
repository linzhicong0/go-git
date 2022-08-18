/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"got/builtin"
)

// catFileCmd represents the catFile command
var catFileCmd = &cobra.Command{
	Use:   "cat-file",
	Short: "Describe the cache object of the given sha1 value",
	Long:  ``,
	Run:   builtin.CatFileRun,
	Args: cobra.ExactValidArgs(1),
}

func init() {
	rootCmd.AddCommand(catFileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// catFileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// catFileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	catFileCmd.Flags().BoolP("pretty", "p", false, "pretty-print object's content")
	catFileCmd.Flags().BoolP("type", "t", false, "show object type")
	catFileCmd.Flags().BoolP("size", "s", false, "show object size")

//catFileCmd.SetUsageFunc()

//catFileCmd.SetUsageTemplate()
}
