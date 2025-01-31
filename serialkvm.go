package main

import (
	"fmt"
	"log"

	"go.bug.st/serial"
)

func (p *PluginImpl) NewSerialPort() error {
	if p.serialPort != nil {
		return nil
	}

	LoadConfig()

	port, err := serial.Open(config.Device, &config.SerialMode)
	if err != nil {
		return fmt.Errorf("failed to open serial port %s: %w", config.Device, err)
	}
	defer port.Close()

	p.serialPort = port

	log.Printf("Opened port %s at %d baud\n", config.Device, config.SerialMode.BaudRate)

	return nil
}

func (p *PluginImpl) SwitchInput(inputNumber int) error {

	if inputNumber > len(config.Inputs) {
		return fmt.Errorf("invalid input number: %d", inputNumber)
	}

	command := config.Inputs[inputNumber].ControlMessage
	_, err := p.serialPort.Write([]byte(command))
	if err != nil {
		return fmt.Errorf("failed to send command: %w", err)
	}

	log.Printf("Switched to input %d\n", inputNumber)

	return nil
}
