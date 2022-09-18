package cni

import (
	"encoding/json"
	"fmt"
	"net"
	"runtime"

	"github.com/containernetworking/cni/pkg/skel"
	"github.com/containernetworking/cni/pkg/types"
	current "github.com/containernetworking/cni/pkg/types/100"
	"github.com/containernetworking/cni/pkg/version"
	"github.com/containernetworking/plugins/pkg/utils/buildversion"
)

type NetConf struct {
	types.NetConf

	// Add plugin-specifc flags here
	//
}

func init() {
	// this ensures that main runs only on main thread (thread group leader).
	// since namespace ops (unshare, setns) are done for a single thread, we
	// must ensure that the goroutine does not jump from OS thread to thread
	runtime.LockOSThread()
}

// parseConfig parses the supplied configuration (and prevResult) from stdin
func parseConfig(stdin []byte) (*NetConf, error) {
	conf := NetConf{}

	if err := json.Unmarshal(stdin, &conf); err != nil {
		return nil, fmt.Errorf("failed to parse network configuration: %v", err)
	}

	if err := version.ParsePrevResult(&conf.NetConf); err != nil {
		return nil, fmt.Errorf("could not parse prevResult: %v", err)
	}

	// Do any validation here
	//

	return &conf, nil
}

// cmdAdd is called for ADD requests
func cmdAdd(args *skel.CmdArgs) error {
	conf, err := parseConfig(args.StdinData)
	if err != nil {
		return err
	}

	if conf.PrevResult != nil {
		return fmt.Errorf("must be called as the first plugin")
	}

	// Generate some fake container IPs and add to the result
	result := &current.Result{
		CNIVersion: current.ImplementedSpecVersion,
	}
	result.Interfaces = []*current.Interface{
		{
			Name:    "if0",
			Sandbox: args.Netns,
			Mac:     "00:11:22:33:44:55",
		},
	}
	ip, ipNet, err := net.ParseCIDR("10.244.0.2/24")
	gatewayIp := "10.244.0.1"
	result.IPs = []*current.IPConfig{
		{
			Address: net.IPNet{ip, ipNet.Mask},
			Gateway: net.ParseIP(gatewayIp),
			// Interface is an index into the Interfaces array of the Interface element this IP allocates to
			Interface: current.Int(0),
		},
	}

	// Implement your plugin here

	// Pass through the result for the next plugin
	return types.PrintResult(result, conf.CNIVersion)
}

func cmdDel(args *skel.CmdArgs) error {
	conf, err := parseConfig(args.StdinData)
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
