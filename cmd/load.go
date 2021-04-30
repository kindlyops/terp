// Copyright © 2018 Kindly Ops, LLC <support@kindlyops.com>
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
	"log"
	"os"

	"github.com/spf13/cobra"
)

var loadCmd = &cobra.Command{
	Use:   "load <program.foo>",
	Short: "Load and examine a program.",
	Long:  `Examine the instruction stream from an eBPF IR code stream.`,
	Run:   load,
	Args:  cobra.ExactArgs(1),
}

func load(cmd *cobra.Command, args []string) {

	target := args[0]
	_, _ = os.Stat(target)

	log.Fatal("TODO")

}

func init() {
	rootCmd.AddCommand(loadCmd)
}
