// Copyright © 2020 Kindly Ops, LLC <support@kindlyops.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// Debug controls whether or not debug messages should be printed.
var Debug bool

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Version: "dev",
	Use:     "terp",
	Short:   "🙈 terp is an experimental eBPF diagnostic interpreter",
	Long: `🙈 terp helps work with eBPF IR code.

Brought to you by

_  ___           _ _        ___
| |/ (_)_ __   __| | |_   _ / _ \ _ __  ___
| ' /| | '_ \ / _| | | | | | | | | '_ \/ __|
| . \| | | | | (_| | | |_| | |_| | |_) __ \
|_|\_\_|_| |_|\__,_|_|\__, |\___/| .__/|___/
                      |___/      |_|
use at your own risk.
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

func Execute(v string) {
	rootCmd.SetVersionTemplate(v)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.terp.yaml)")

	rootCmd.PersistentFlags().BoolVarP(&Debug, "debug", "d", false, "Print debug messages while working")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".terp" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".terp")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
