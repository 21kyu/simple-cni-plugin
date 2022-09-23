package connector

import (
	"fmt"

	"github.com/vishvananda/netlink"

	"github.com/containernetworking/plugins/pkg/ns"
)

func createIfName(prefix, containerID string) string {
	if len(containerID) < 5 {
		return prefix + containerID
	}
	return prefix + containerID[:5]
}

func setupVethRemoteNs(netNs ns.NetNS, srcIfName, dstIfName string) error {
	return netNs.Do(func(_ ns.NetNS) error {
		link, err := netlink.LinkByName(srcIfName)
		if err != nil {
			return err
		}
		err = netlink.LinkSetName(link, dstIfName)
		if err != nil {
			return err
		}
		err = netlink.LinkSetUp(link)
		if err != nil {
			return err
		}
		return nil
	})
}

// SetupVeth creates a veth pair and moves one end into the container's namespace
func SetupVeth(br *netlink.Bridge, netNs ns.NetNS, contID, ifName string) (*netlink.Veth, netlink.Link, error) {
	var err error

	hostIfName := createIfName("veth", contID)
	tempIfName := createIfName("temp", contID)

	veth := &netlink.Veth{
		LinkAttrs: netlink.LinkAttrs{
			Name:   hostIfName,
			TxQLen: 1000,
		},
		PeerName: tempIfName,
	}

	// ip link add $veth type veth peer name $peer
	if err := netlink.LinkAdd(veth); err != nil {
		return nil, nil, err
	}
	defer func() {
		if err != nil {
			if err = netlink.LinkDel(veth); err != nil {
				fmt.Printf("failed to delete veth: %v", err)
			}
		}
	}()

	// ip link set $veth up
	if err = netlink.LinkSetUp(veth); err != nil {
		return nil, nil, err
	}

	// ip link set $veth master $br
	if err = netlink.LinkSetMaster(veth, br); err != nil {
		return nil, nil, err
	}

	peer, err := netlink.LinkByName(tempIfName)
	if err != nil {
		return nil, nil, err
	}

	// ip link set $peer netns $netns
	if err = netlink.LinkSetNsFd(peer, int(netNs.Fd())); err != nil {
		return nil, nil, err
	}

	// ip netns exec $netns ip link set $peer name $ifname
	// ip netns exec $netns ip link set $peer up
	if err = setupVethRemoteNs(netNs, tempIfName, ifName); err != nil {
		return nil, nil, err
	}

	return veth, peer, nil
}
