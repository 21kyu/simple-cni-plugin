package main

import (
	"fmt"
	"runtime"

	"github.com/21kyu/simple-cni-plugin/pkg/connector"
	"github.com/21kyu/simple-cni-plugin/plugin/cni/types"

	"github.com/containernetworking/cni/pkg/skel"
	cniTypes "github.com/containernetworking/cni/pkg/types"
	current "github.com/containernetworking/cni/pkg/types/100"
	"github.com/containernetworking/cni/pkg/version"
	"github.com/containernetworking/plugins/pkg/utils/buildversion"
)

func init() {
	// this ensures that main runs only on main thread (thread group leader).
	// since namespace ops (unshare, setns) are done for a single thread, we
	// must ensure that the goroutine does not jump from OS thread to thread
	runtime.LockOSThread()
}

// cmdAdd is called for ADD requests
func cmdAdd(args *skel.CmdArgs) error {
	conf, err := types.LoadNetConf(args.StdinData)
	if err != nil {
		return err
	}

	if conf.PrevResult != nil {
		return fmt.Errorf("must be called as the first plugin")
	}

	_, _, err = connector.SetupBridge(conf.BrName)
	if err != nil {
		return err
	}

	result := &current.Result{
		CNIVersion: current.ImplementedSpecVersion,
	}

	// Pass through the result for the next plugin
	return cniTypes.PrintResult(result, conf.CNIVersion)
}

func cmdDel(args *skel.CmdArgs) error {
	conf, err := types.LoadNetConf(args.StdinData)
	if err != nil {
		return err
	}
	_ = conf

	// Do your delete here

	return nil
}

func cmdCheck(args *skel.CmdArgs) error {
	return fmt.Errorf("not implemented")
}

func main() {
	skel.PluginMain(cmdAdd, cmdCheck, cmdDel, version.All, buildversion.BuildString("Simple CNI plugin"))
}
