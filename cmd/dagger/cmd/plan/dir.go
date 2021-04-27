package plan

import (
	"dagger.io/go/cmd/dagger/logger"
	"dagger.io/go/dagger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var dirCmd = &cobra.Command{
	Use:   "dir PATH",
	Short: "Load plan from a local directory",
	Args:  cobra.ExactArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		// Fix Viper bug for duplicate flags:
		// https://github.com/spf13/viper/issues/233
		if err := viper.BindPFlags(cmd.Flags()); err != nil {
			panic(err)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		lg := logger.New()
		ctx := lg.WithContext(cmd.Context())

		updateEnvironmentPlan(ctx, dagger.DirInput(args[0], []string{"*.cue", "cue.mod"}))
	},
}

func init() {
	if err := viper.BindPFlags(dirCmd.Flags()); err != nil {
		panic(err)
	}
}
