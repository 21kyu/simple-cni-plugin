# simple-cni-plugin

The following features will be implemented
- IP management and assignment via CRD without seperate IPAM module
- CNI spec implementation by iptables (also eBPF if possible)
- Provide network policy APIs
- Configure BGP, IPIP, VXLAN for routing information distribution
- Use [leader election](https://pkg.go.dev/k8s.io/client-go/tools/leaderelection) if necessary

### Create a Bridge

First, create a network bridge interface named `cni0` on the host
using that has the same effect as the following command:

```shell
ip link add cni0 type bridge
```

By default, `cni0` network bridge interface will be administratively down,
which means traffic will not flow through the network bridge interface.
To enable traffic flow, we need to set the network bridge interface to:

```shell
ip link set cni0 up
```