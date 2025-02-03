package main

import (
	"fmt"
	"log"

	"go.bug.st/serial"
)

func (p *PluginImpl) CloseSerialPort() error {
	if p.serialPort == nil {
		return nil
	}

	err := p.serialPort.Close()
	p.serialPort = nil

	log.Printf("closed port %s", config.Device)

	return err
}

func (p *PluginImpl) OpenSerialPort() error {

	if p.serialPort != nil {
		return nil
	}

	LoadConfig()

	if len(config.Inputs) < 1 {
		return fmt.Errorf("no inputs configured; serial port will not be opened")
	}

	port, err := serial.Open(config.Device, &config.SerialMode)
	if err != nil {
		return fmt.Errorf("failed to open serial port %s: %w", config.Device, err)
	}

	p.serialPort = port

	log.Printf("Opened port %s at %d baud", config.Device, config.SerialMode.BaudRate)
	return nil
}

func (p *PluginImpl) SwitchInput(inputNumber int) error {

	LoadConfig()

	if inputNumber < 1 || inputNumber > len(config.Inputs) {
		return fmt.Errorf("invalid input number: %d", inputNumber)
	}

	if err := p.OpenSerialPort(); err != nil {
		return err
	}

	command := config.Inputs[inputNumber-1].ControlMessage
	log.Printf("Sending command: %s", command)

	_, err := p.serialPort.Write([]byte(command))
	if err != nil {
		return fmt.Errorf("failed to send command: %w", err)
	}

	log.Printf("Switched to input %d", inputNumber)

	return nil
}
