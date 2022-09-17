# simple-cni-plugin

The following features will be implemented
- IP management and assignment via CRD without seperate IPAM module
- CNI spec implementation by iptables (also eBPF if possible)
- Provide network policy APIs
- Configure BGP, IPIP, VXLAN for routing information distribution
- Use [leader election](https://pkg.go.dev/k8s.io/client-go/tools/leaderelection) if necessary
