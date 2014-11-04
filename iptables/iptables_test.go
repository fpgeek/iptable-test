package iptables

import (
	"testing"
	"net"
	"strings"
)

func TestForward(t *testing.T) {
	if err := RemoveExistingChain("mesos"); err != nil {
		t.Errorf("RemoveExistingChain : %s", err.Error())
	}
//	return
	chain, err := NewChain("mesos", "eth1")
	if err != nil {
		t.Errorf("NewChain Failed : %s", err.Error())
	}

	inter, err := net.InterfaceByName("eth1")
	if err != nil {
		t.Errorf("InterfaceByName : %s", err.Error())
	}
	addrs, _ := inter.Addrs()
	if len(addrs) == 0 {
		t.Error("Addrs()")
	}

	if index := strings.Index(addrs[0].String(), "/"); index == -1 {
		t.Error("Wrong Address")
	}

	destAddr := strings.Split(addrs[0].String(), "/")[0]

	if err := chain.Forward(Add, net.ParseIP("0.0.0.0"), 9080, "tcp", destAddr, 8080); err != nil {
		t.Error("chain.Forward Add failed")
	}
}
