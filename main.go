package main

import (
	"time"

	"github.com/ZedElysium/GoPoker/p2p"
)

func main() {
	cfg := p2p.ServerConfig {
		Version: "GOPOKER V0.1",
		ListenAddr: ":3000",
	}
	server := p2p.NewServer(cfg)
	go server.Start()

	time.Sleep(1 * time.Second)

	remoteCfg := p2p.ServerConfig {
		Version: "GOPOKER V0.1",
		ListenAddr: ":4000",
	}
	remoteServer := p2p.NewServer(remoteCfg)
	go remoteServer.Start()
	if err := remoteServer.Connect(":3000"); err != nil {
		panic(err)
	}

	select{}
}
