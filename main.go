package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/caarlos0/env"
	"github.com/sourcegraph/jsonrpc2"
	"go.bug.st/serial"
	"serialkvm/plugin"
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
		return nil, fmt.Errorf("failed to connect to socket: %w", err)
	}

	impl := &PluginImpl{}
	client := jsonrpc2.NewConn(ctx, jsonrpc2.NewPlainObjectStream(conn), plugin.HandleRPC(impl))
	impl.client = client

	return impl, nil
}

func main() {
	log.Println("Starting SerialKVM plugin")

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
		log.Fatalln(err)
	}
	defer impl.client.Close()

	if len(config.Inputs) < 1 {
		log.Fatalln("No input configured")
	}

	err = impl.NewSerialPort()
	if err != nil {
		log.Fatalln(err)
	}
	defer impl.serialPort.Close()

	//...

	log.Println("plugin started")

	<-ctx.Done()
}
