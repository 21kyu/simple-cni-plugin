package connector

import (
	"github.com/containernetworking/plugins/pkg/ns"
	"github.com/containernetworking/plugins/pkg/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/vishvananda/netlink"
	"testing"
)

const contIfaceName = "eth0"

func TestSetupVeth(t *testing.T) {
	br, _, _ := SetupBridge(brName)
	defer func(link netlink.Link) {
		if err := netlink.LinkDel(link); err != nil {
			t.Error(err)
		}
	}(br)

	netNs, err := testutils.NewNS()
	if err != nil {
		t.Error(err)
	}

	veth, peer, err := SetupVeth(br, netNs, "1234567890", contIfaceName)
	defer func(veth netlink.Link) {
		if err := netlink.LinkDel(veth); err != nil {
			t.Error(err)
		}
	}(veth)
	if err != nil {
		t.Error(err)
	}

	err = netNs.Do(func(_ ns.NetNS) error {
		contIface, err := netlink.LinkByName(contIfaceName)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, contIfaceName, contIface.Attrs().Name)
		assert.Equal(t, peer.Attrs().HardwareAddr.String(), contIface.Attrs().HardwareAddr.String())
		return nil
	})
	if err != nil {
		t.Error(err)
	}
}
