// Package cmd implements the oasis-core-ledger tool.
package cmd

import (
	// "fmt"

	"github.com/spf13/cobra"
	// flag "github.com/spf13/pflag"
	// "github.com/spf13/viper"
	// "github.com/oasisprotocol/oasis-core/go/common/logging"
	// "github.com/oasisprotocol/oasis-core/go/oasis-node/cmd/common"
)

// const cfgLogLevel = "log.level"

var rootCmd = &cobra.Command{
	Use:     "oasis-core-ledger",
	Short:   "Oasis Ledger Tool",
	Version: "0.0.1",
}

// rootFlags = flag.NewFlagSet("", flag.ContinueOnError)

// RootCommand returns the root (top level) cobra.Command.
func RootCommand() *cobra.Command {
	return rootCmd
}

// Execute spawns the main entry point after handling the command line arguments.
func Execute() {
	// Blah, The fact that oasis-core imports the ledger library causes
	// issues, re-enable logging once the ledger support has been purged.
	/*
		var logLevel logging.Level
		if err := logLevel.Set(viper.GetString(cfgLogLevel)); err != nil {
			common.EarlyLogAndExit(fmt.Errorf("root: failed to set log level: %w", err))
		}

		if err := rootCmd.Execute(); err != nil {
			common.EarlyLogAndExit(err)
		}
	*/
	_ = rootCmd.Execute()
}

func init() { // nolint: gochecknoinits
	/*
		logLevel := logging.LevelInfo
		rootFlags.Var(&logLevel, cfgLogLevel, "log level")
		_ = viper.BindPFlags(rootFlags)
		rootCmd.PersistentFlags().AddFlagSet(rootFlags)
	*/

	// Register all of the sub-commands.
	rootCmd.AddCommand(listCmd)
}