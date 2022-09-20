package connector

import (
	"testing"

	"github.com/vishvananda/netlink"

	"github.com/stretchr/testify/assert"
)

const brName = "bridge0"

func TestSetupBridge(t *testing.T) {
	br, brInterface, err := SetupBridge(brName)
	defer func(link netlink.Link) {
		err := netlink.LinkDel(link)
		if err != nil {
			t.Error(err)
		}
	}(br)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, brInterface.Name, brName)
	assert.NotNil(t, brInterface.Mac)
}
