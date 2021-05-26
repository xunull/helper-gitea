package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/xunull/goc/reflectx"
	"github.com/xunull/helper-gitea/pkg/helper_api"
	"reflect"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "api",
	Long:  `api`,
	PreRun: func(cmd *cobra.Command, args []string) {
		helper_api.InitDefaultApi()
	},
	Run: func(cmd *cobra.Command, args []string) {
		v := reflect.ValueOf(helper_api.Api)
		method := v.MethodByName(args[0])
		if len(args) > 1 {
			vs := make([]reflect.Value, 0)
			mt := method.Type()
			for i, item := range args[1:] {
				vs = append(vs, reflectx.StringToCommon(item, mt.In(i)))
			}
			method.Call(vs)
		} else {
			method.Call(nil)
		}
	},
}

var apiListCmd = &cobra.Command{
	Use:   "list",
	Short: "list wiki api",
	Long:  `list wiki api`,
	Run: func(cmd *cobra.Command, args []string) {
		vt := reflect.TypeOf(helper_api.Api)
		for i := 0; i < vt.NumMethod(); i++ {
			method := vt.Method(i)
			t := method.Type.NumIn()

			argStr := ""
			for j := 1; j < t; j++ {
				argStr += fmt.Sprintf("%s  ", method.Type.In(j).String())
			}
			if len(argStr) > 2 {
				argStr = argStr[:len(argStr)-2]
			}
			color.New(color.FgBlue).Printf("%s(", method.Name)
			color.New(color.FgYellow).Printf("%s", argStr)
			color.New(color.FgBlue).Printf(")\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
	apiCmd.AddCommand(apiListCmd)
}
