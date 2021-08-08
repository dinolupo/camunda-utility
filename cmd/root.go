package cmd

import (
	"fmt"
	"os"

	"github.com/dinolupo/camunda-utility/pkg/camunda/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var host string
var port int
var Camunda *client.Client

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "camunda-utility",
	Short: "Camunda command line utility",
	Long: `Camunda Utility is a command line tool that permits to execute
	administrative tasks like deleting all definitions and instances.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig, initClient)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.camunda-utility.yaml)")
	rootCmd.PersistentFlags().StringVar(&host, "host", "localhost", "Camunda Host")
	rootCmd.PersistentFlags().IntVar(&port, "port", 8080, "Camunda Port")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func initClient() {
	Camunda = client.NewClient(client.ClientOptions{
		Host: host,
		Port: port,
	})
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".camunda-utility" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".camunda-utility")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
