package container

import (
	"github.com/opencontainers/runtime-spec/specs-go"
	"golang.org/x/sys/unix"
)

var namespaceCloneFlags = map[specs.LinuxNamespaceType]int{
	specs.CgroupNamespace: unix.CLONE_NEWCGROUP,
	specs.IPCNamespace: unix.CLONE_NEWIPC,
	specs.MountNamespace: unix.CLONE_NEWNS,
	specs.NetworkNamespace: unix.CLONE_NEWNET,
	specs.PIDNamespace: unix.CLONE_NEWPID,
	specs.UserNamespace: unix.CLONE_NEWUSER,
	specs.UTSNamespace: unix.CLONE_NEWUTS,
	specs.TimeNamespace: unix.CLONE_NEWTIME,
}

func getNamespaceCloneFlags(ns []specs.LinuxNamespace) uintptr {
	var flag int
	for _, n := range ns {
		if n.Path != "" {
			continue
		}
		flag |= namespaceCloneFlags[n.Type]
	}
	return uintptr(flag)
}