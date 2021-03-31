package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"homework/internal/web"
	"os"
	"path/filepath"
	"strconv"
)

var gitRevision string	//git的版本信息

// version 执行version命令
func version() {
	gitRevision = "version:1.0"
	log.Info().Str("gitRevision",gitRevision).Send()
}

// newVersionCmd 创建version操作命令，返回cobra.Command类型指针
func newVersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "to show version",
		Run: func(cmd *cobra.Command, args []string) {
			version()
		},
	}

	return cmd
}

// newWebCme 创建web操作命令，返回cobra.Command类型指针
func newWebCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "web",
		Short: "to show a website",
		Run: func(cmd *cobra.Command, args []string) {
			web.Start()
		},
	}

	return cmd
}

// initZeroLog设置日志信息， level作为日志等级，返回error信息
func initZeroLog(level string) error {
	zerolog.TimestampFieldName = "ts"
	zerolog.LevelFieldName = "l"
	zerolog.MessageFieldName = "m"
	zerolog.CallerFieldName = "c"
	zerolog.TimeFieldFormat = "0102-15:04:05Z07"
	zerolog.ErrorFieldName = "e"

	zerolog.CallerMarshalFunc = func(file string, line int) string {
		return filepath.Base(file) + ":" + strconv.Itoa(line)
	}

	lvl, _ := zerolog.ParseLevel(level)
	if lvl == zerolog.NoLevel {
		lvl = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(lvl)
	log.Logger = log.With().Caller().Logger()

	return nil
}

func main() {
	rootCmd := &cobra.Command{
		Use: "home",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {

			err := initZeroLog("trace")
			if err != nil {
				log.Error().Err(err).Msg("initZerolog fail")
				return err
			}

			return nil
		},
	}

	rootCmd.AddCommand(newVersionCmd())
	rootCmd.AddCommand(newWebCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
}
