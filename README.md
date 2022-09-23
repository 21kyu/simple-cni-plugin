# simple-cni-plugin

The following features will be implemented
- IP management and assignment via CRD without seperate IPAM module
- CNI spec implementation by iptables (also eBPF if possible)
- Provide network policy APIs
- Configure BGP, IPIP, VXLAN for routing information distribution
- Use [leader election](https://pkg.go.dev/k8s.io/client-go/tools/leaderelection) if necessary

#### 1. Create a bridge

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

#### 2. Create a veth pair

We need to create a pair of network interfaces.

```shell
ip link add veth0 type veth peer name peer0
```

The Interfaces are create as an interconnected pair.
The `veth0` interface is attached to the host network namespace,
and the `peer0` interface will be attached to the container network namespace.

#### 3. Add the veth pair to the bridge

The `veth0` interface remains in the host network namespace
and should be added to the `cni0` network bridge interface.

```shell
ip link set veth0 up
ip link set veth0 master cni0
```

#### 4. Move the peer interface to the container network namespace

Now, we move the `peer0` interface to the container network namespace.
Then, we rename the interface to `eth0` and set the interface to up.

```shell
ip link set peer0 netns $netns
ip netns exec $netns ip link set peer0 name eth0
ip netns exec $netns ip link set eth0 up
```

#### 5. Assign an IP address to the container