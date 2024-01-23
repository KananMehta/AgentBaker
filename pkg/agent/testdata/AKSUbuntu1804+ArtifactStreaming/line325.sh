[Unit]
Description=AKS Log Collector Slice
DefaultDependencies=no
Before=slices.target
Requires=system.slice
After=system.slice

[Slice]
CPUAccounting=true
#CPUQuota=10%
CPUShares=100
CPUWeight=1
MemoryMax=128M
