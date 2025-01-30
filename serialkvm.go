package main

import (
	"fmt"
	"log"
	"time"

	"go.bug.st/serial"
)

type SerialConfig struct {
	BaudRate int
	DataBits int
	StopBits int
	Parity   int
	Timeout  time.Duration
}

func DefaultConfig() *serial.Mode {
	return &serial.Mode{
		BaudRate: 115200,
	}
}

func (p *PluginImpl) NewSerialPort() error {
	port, err := serial.Open("/dev/ttyS3", DefaultConfig())
	p.serialPort = port
	if err != nil {
		return err
	}
	defer port.Close()

	return nil
}

func (p *PluginImpl) SwitchInput(inputNumber string) error {

	command := fmt.Sprintf("X%s,1$", inputNumber)
	_, err := p.serialPort.Write([]byte(command))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Switched to input ", inputNumber)

	return nil
}
