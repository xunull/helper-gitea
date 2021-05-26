package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xunull/helper-gitea/pkg/global"
	"log"
	"os"
	"runtime/debug"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "helper-gitea",
	Short: "helper-gitea",
	Long:  `helper-gitea`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.helper-gitea.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		cobra.CheckErr(err)
		viper.AddConfigPath(home)
		viper.SetConfigName(".helper-gitea")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		var config global.AppConfig
		err := viper.Unmarshal(&config)
		if err != nil {
			debug.PrintStack()
			log.Fatal(err)
		}
		global.InitConfig(&config)
	} else {
		log.Fatal(err)
	}
}
