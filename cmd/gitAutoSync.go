
package cmd

import (
	"flag"
	"fmt"
	types "git-auto-sync/common"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"os"
	"time"
)
func NewGitSync() *cobra.Command {

	opts := newGitOptions()

	cmd := &cobra.Command{
		Use:   "gitAutoSync",
		Short: "gitAutoSync is a tool about sync by git",
		Run: func(cmd *cobra.Command, args []string) {
			run(opts)
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {

			flag.CommandLine.Parse([]string{})
		},
	}
	flag.Set("logtostderr", "true")
	cmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)

	cmd.AddCommand()
	AddFlag(cmd, opts)
	return cmd
}

func AddFlag(cmd *cobra.Command,opts * types.GitOptions){
	cmd.Flags().DurationVar(&opts.CommitInterval, "commit-interval", time.Minute,"Specify the commit interval")
	cmd.Flags().DurationVar(&opts.PushInterval, "push-interval", time.Hour,"Specify the gpush interval")
	cmd.Flags().StringVarP(&opts.Directory, "path","p", ".","Specify the git path")
	cmd.Flags().StringVarP(&opts.CommitName, "name","n", "GIT_AUTO_SYNC","Specify commit name")
	cmd.Flags().StringVarP(&opts.CommitEmail, "email","e", "","Specify commit email")
}

func newGitOptions() * types.GitOptions{
	ops := &types.GitOptions{}
	return ops
}


func Execute() {
	rootCmd := NewGitSync()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(opt * types.GitOptions){

	stop := make(chan int)
	tc := time.NewTicker(opt.CommitInterval)
	tp := time.NewTicker(opt.PushInterval)

	go func (){
		AutoCommit(opt)
		for {
			select {
			case <-tc.C:
				glog.V(4).Info("try to commit")
				AutoCommit(opt)
			}
		}
	}()

	go func (){
		AutoPush(opt)
		time.Sleep(time.Second * 10)
		for {
			select {
			case <-tp.C:
				glog.V(4).Info("try to push")
				AutoPush(opt)
			}
		}
	}()

	<-stop
}

func AutoCommit(opt * types.GitOptions){
	r, err := git.PlainOpen(opt.Directory)
	if err != nil{
		glog.Errorf(err.Error())
		return
	}

	w, err := r.Worktree()
	if err != nil{
		glog.Errorf(err.Error())
		return
	}
	status, err := w.Status()
	if len(status) ==0 {
		return
	}

	for k,_ := range status{
		_, err = w.Add(k)
	}

	commit, err := w.Commit("update", &git.CommitOptions{
		Author: &object.Signature{
			Name:  opt.CommitName,
			Email: opt.CommitEmail,
			When:  time.Now(),
		},
	})

	if err != nil{
		glog.Errorf(err.Error())
		return
	}

	_, err = r.CommitObject(commit)
	if err != nil{
		glog.Errorf(err.Error())
		return
	}
	glog.Info("commit success")
}

func AutoPush(opt * types.GitOptions){
	r, err := git.PlainOpen(opt.Directory)
	if err != nil{
		glog.Errorf(err.Error())
		return
	}
	err = r.Push(&git.PushOptions{})
	if err != nil{
		glog.Errorf(err.Error())
		return
	}
	glog.Info("push success")
}