package main

import (
	"fmt"
	"github.com/21kyu/simple-cni-plugin/plugin/cni/types"
	"runtime"
	"syscall"

	"github.com/containernetworking/cni/pkg/skel"
	cniTypes "github.com/containernetworking/cni/pkg/types"
	current "github.com/containernetworking/cni/pkg/types/100"
	"github.com/containernetworking/cni/pkg/version"
	"github.com/containernetworking/plugins/pkg/utils/buildversion"
	"github.com/containernetworking/plugins/pkg/utils/sysctl"
	"github.com/vishvananda/netlink"
)

func init() {
	// this ensures that main runs only on main thread (thread group leader).
	// since namespace ops (unshare, setns) are done for a single thread, we
	// must ensure that the goroutine does not jump from OS thread to thread
	runtime.LockOSThread()
}

func bridgeByName(name string) (*netlink.Bridge, error) {
	link, err := netlink.LinkByName(name)
	if err != nil {
		return nil, fmt.Errorf("could not lookup %q: %v", name, err)
	}
	bridge, ok := link.(*netlink.Bridge)
	if !ok {
		return nil, fmt.Errorf("%q already exists but is not a bridge", name)
	}
	return bridge, nil
}

func ensureBridge(brName string) (*netlink.Bridge, error) {
	br := &netlink.Bridge{
		LinkAttrs: netlink.LinkAttrs{
			Name: brName,
			// Let kernel use default txqueuelen; leaving it unset
			// means 0, and a zero-length TX queue messes up FIFO
			// traffic shapers which use TX queue length as the
			// default packet limit
			TxQLen: -1,
		},
	}

	// ip link add name cni0 type bridge
	err := netlink.LinkAdd(br)
	if err != nil && err != syscall.EEXIST {
		return nil, fmt.Errorf("could not add %q: %v", brName, err)
	}

	// Re-fetch link to read all attributes and if it already existed,
	// ensure it's really a bridge with similar configuration
	br, err = bridgeByName(brName)
	if err != nil {
		return nil, err
	}

	// we want to own the routes for this interface
	// for security, even if our IPv6 support is disabled, try to disable RAs on the interface.
	_, _ = sysctl.Sysctl(fmt.Sprintf("net/ipv6/conf/%s/accept_ra", brName), "0")

	// ip link set dev cni0 up
	if err := netlink.LinkSetUp(br); err != nil {
		return nil, err
	}

	return br, nil
}

func setupBridge(conf *types.NetConf) (*netlink.Bridge, *current.Interface, error) {
	br, err := ensureBridge(conf.BrName)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create bridge %q: %v", conf.BrName, err)
	}

	return br, &current.Interface{
		Name: br.Attrs().Name,
		Mac:  br.Attrs().HardwareAddr.String(),
	}, nil
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

	_, _, err = setupBridge(conf)
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
