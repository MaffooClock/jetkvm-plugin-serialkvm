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
	PluginSocket     string `env:"JETKVM_PLUGIN_SOCK" envDefault:"./tmp/plugin.sock"`
	PluginWorkingDir string `env:"JETKVM_PLUGIN_WORKING_DIR" envDefault:"./tmp"`
}

func connect(ctx context.Context) (*PluginImpl, error) {
	conn, err := net.Dial("unix", PluginConfig.PluginSocket)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to socket: %w", err)
	}

	p := &PluginImpl{}
	client := jsonrpc2.NewConn(ctx, jsonrpc2.NewPlainObjectStream(conn), plugin.HandleRPC(p))
	p.client = client

	return p, nil
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

	p, err := connect(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer p.client.Close()

	log.Println("rpc client started")

	// TODO: maybe we don't need to open the port right now,
	//       and do it later when `SwitchInput` is called?
	//err = p.OpenSerialPort()
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//defer p.CloseSerialPort()
	//
	//log.Println("plugin started")

	<-ctx.Done()

	log.Println("we go bye-bye?")
}
