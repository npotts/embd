// +build ignore

// this sample uses the mcp3008 package to interface with a similar MCP3208, which is a 12 bit variant. Works without code change on bbb and rpi
package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/kidoman/embd"
	  "github.com/kidoman/embd/convertors/mcp3008"
	_ "github.com/kidoman/embd/host/all"
)

const (
	channel = 0
	speed   = 1000000
	bpw     = 8
	delay   = 0
)

func main() {
	flag.Parse()
	fmt.Println("this is a sample code for mcp3008 10bit 8 channel ADC")

	if err := embd.InitSPI(); err != nil {
		panic(err)
	}
	defer embd.CloseSPI()

	spiBus := embd.NewSPIBus(embd.SPIMode0, channel, speed, bpw, delay)
	defer spiBus.Close()

	adc := &mcp3008.MCP3008{Mode: mcp3008.SingleMode, Bus: spiBus, Bits: mcp3008.Bits12}

	// adc := mcp3008.New(mcp3008.SingleMode, spiBus)

	for i := 0; i < 20; i++ {
		time.Sleep(1 * time.Second)
		val, err := adc.AnalogValueAt(0)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		fmt.Printf("analog value is: %v\n", val)
	}
}
