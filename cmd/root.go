/*
Copyright © 2020 NAME HERE <kodrat.meden@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/satesate-dev/go-restful-boilerplate/util"

	"github.com/satesate-dev/go-restful-boilerplate/helper/database"
	"github.com/spf13/cobra"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	DBPool  *sql.DB
	logger  *util.Logger
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-restful-boilerplate",
	Short: "A brief description of your application",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig, initLoad)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/..config.toml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetConfigType("toml")
	// Search config in root directory with name "..config.toml" (without extension).
	viper.AddConfigPath(".")
	viper.SetConfigName(".config")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		log.Fatal("Config file not found")
	}
}

func initLoad() {
	splash()
	initLogger()
	initDatabase()
}

func initDatabase() {

	//Setup DB Connection
	dbConfig := database.NewDatabase(
		viper.GetString("database.db"),
		viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetString("database.port"),
		viper.GetString("database.name"),
		viper.GetString("database.timezone"),
		viper.GetString("database.ssl_mode"),
		viper.GetString("database.ssl_cert"),
		viper.GetString("database.ssl_key"),
		viper.GetString("database.ssl_root_cert"),
	)

	// Connect to DB
	var err error
	DBPool, err = dbConfig.Connect()
	if err != nil {
		logger.Err.Fatalf("invalid db config : %v", err)
	}
	// Checking database connection
	if err := DBPool.Ping(); err != nil {
		logger.Err.Fatalf("failed connect to db : %v", err)
	}
}

func initLogger() {
	logger = util.NewLogger()
}

func splash() {
	fmt.Println(`
	██╗  ██╗ █████╗  ██████╗██╗  ██╗████████╗ ██████╗ ██████╗ ███████╗██████╗ ███████╗███████╗███████╗████████╗
	██║  ██║██╔══██╗██╔════╝██║ ██╔╝╚══██╔══╝██╔═══██╗██╔══██╗██╔════╝██╔══██╗██╔════╝██╔════╝██╔════╝╚══██╔══╝
	███████║███████║██║     █████╔╝    ██║   ██║   ██║██████╔╝█████╗  ██████╔╝█████╗  █████╗  ███████╗   ██║
	██╔══██║██╔══██║██║     ██╔═██╗    ██║   ██║   ██║██╔══██╗██╔══╝  ██╔══██╗██╔══╝  ██╔══╝  ╚════██║   ██║
	██║  ██║██║  ██║╚██████╗██║  ██╗   ██║   ╚██████╔╝██████╔╝███████╗██║  ██║██║     ███████╗███████║   ██║
	╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝   ╚═╝    ╚═════╝ ╚═════╝ ╚══════╝╚═╝  ╚═╝╚═╝     ╚══════╝╚══════╝   ╚═╝
	`)
}
