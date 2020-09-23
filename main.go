package main

import (
	"git-auto-sync/cmd"
)

func main() {
	//pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	//flag.Parse()
	//defer glog.Flush()

	cmd.Execute()
}