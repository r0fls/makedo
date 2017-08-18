// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"errors"
	"fmt"
	"github.com/go-yaml/yaml"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

var cfgFile string

type rootOptions struct {
	doFile string
}

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "do",
	Short: "Like make but with yaml.",
	Long:  `Makefile syntax getting you down? Here you can just use yaml.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: run,
}

func loadConjureFile() ([]byte, error) {
	filenames := []string{"conjure.yaml", "conjure.yml", "Conjure.yaml", "Conjure.yaml"}
	for _, filename := range filenames {
		data, err := ioutil.ReadFile(filename)
		if err == nil {
			return data, nil
		}
	}
	return nil, errors.New("You must have a makedo file with one of the names: conjure.yaml, Conjure,yaml, conjure.yml, or Conjure.yml")
}

func run(cmd *cobra.Command, args []string) {

	command := args[0]

	if len(args) == 0 {
		fmt.Println("Must supply a command to do.")
		os.Exit(1)
	}

	var data []byte
	var err error

	if opts.doFile == "" {
		data, err = loadConjureFile()
	} else {
		data, err = ioutil.ReadFile(opts.doFile)
		if err != nil {
			panic(err)
		}
	}

	doMap := make(map[string]map[string][]string)
	err = yaml.Unmarshal([]byte(data), &doMap)

	if err != nil {
		panic(err)
	}

	if value, ok := doMap[command]; ok {
		do(doMap, value["commands"], value["depends"])
	} else {
		panic(ok)
	}
}

// Runs a list of commands
func runCommands(commands []string) {
	for _, command := range commands {
		args := strings.Fields(command)
		out, err := exec.Command(args[0], args[1:]...).Output()
		if err != nil {
			panic(err)
		}
		fmt.Print(string(out))
	}
}

// Recursively runs do() on the dependencies
func do(doMap map[string]map[string][]string, commands []string, depends []string) {
	for _, dependency := range depends {
		do(doMap, doMap[dependency]["commands"], doMap[dependency]["depends"])
	}
	runCommands(commands)
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var opts *rootOptions

func init() {
	cobra.OnInitialize(initConfig)

	opts = &rootOptions{}

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.do.yaml)")
	RootCmd.PersistentFlags().StringVarP(&opts.doFile, "file", "f", "", "do file")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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

		// Search config in home directory with name ".do" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".do")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
