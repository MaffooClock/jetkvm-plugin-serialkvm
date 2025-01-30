package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/MaffooClock/jetkvm-plugin-serialkvm/plugin"
	"github.com/caarlos0/env"
	"github.com/sourcegraph/jsonrpc2"
	"go.bug.st/serial"
)

var version = "0.0.1"

type PluginImpl struct {
	client     *jsonrpc2.Conn
	serialPort serial.Port
}

var PluginConfig struct {
	PluginSocket string `env:"JETKVM_PLUGIN_SOCK" envDefault:"./tmp/plugin.sock"`
}

func connect(ctx context.Context) (*PluginImpl, error) {
	conn, err := net.Dial("unix", PluginConfig.PluginSocket)
	if err != nil {
		return nil, err
	}

	impl := &PluginImpl{}
	client := jsonrpc2.NewConn(ctx, jsonrpc2.NewPlainObjectStream(conn), plugin.HandleRPC(impl))
	impl.client = client

	return impl, nil
}

func main() {
	env.Parse(&PluginConfig)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, os.Interrupt)
		<-signalChan
		fmt.Println("Received an interrupt, stopping services...")
		cancel()
	}()

	log.Default().SetPrefix("[jetkvm-plugin-serialkvm v" + version + "] ")

	impl, err := connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer impl.client.Close()

	log.Println("client started")

	err = impl.NewSerialPort()
	if err != nil {
		log.Fatal(err)
	}
	defer impl.serialPort.Close()

	//...

	<-ctx.Done()
}
