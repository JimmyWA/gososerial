package main

import (
	"encoding/hex"
	"flag"
	"fmt"

	gososerial "github.com/JimmyWA/Gososerial"
	"github.com/JimmyWA/Gososerial/log"
	"github.com/JimmyWA/Gososerial/tool"
	"github.com/JimmyWA/Gososerial/ysoserial/gadget"
)

const (
	version = "0.1"
	author  = "4ra1n"
)

func main() {
	log.PrintLogo(version, author)
	var payload string
	var command string
	var list bool
	flag.StringVar(&payload, "payload", "", "use which payload")
	flag.StringVar(&command, "cmd", "", "command")
	flag.BoolVar(&list, "list", false, "show payload list")
	flag.Parse()
	if list {
		log.Info("payload list: ")
		all := gososerial.GetAllNames()
		for _, v := range all {
			fmt.Printf("\t%s\n", v)
		}
		return
	}
	if command == "" || payload == "" {
		log.Error("error input")
		return
	}
	fmt.Println(command)
	switch payload {
	case gadget.CC1:
		log.Info("get payload: %s", gadget.CC1)
		bytePayload := gososerial.GetCC1(command)
		output := tool.FormatPayload(hex.EncodeToString(bytePayload))
		log.Info("payload: %s", output)
	case gadget.CC2:
		log.Info("get payload: %s", gadget.CC1)
		bytePayload := gososerial.GetCC1(command)
		output := tool.FormatPayload(hex.EncodeToString(bytePayload))
		log.Info("payload: %s", output)
	case gadget.CC3:
		log.Info("get payload: %s", gadget.CC1)
		bytePayload := gososerial.GetCC1(command)
		output := tool.FormatPayload(hex.EncodeToString(bytePayload))
		log.Info("payload: %s", output)
	case gadget.CC4:
		log.Info("get payload: %s", gadget.CC1)
		bytePayload := gososerial.GetCC1(command)
		output := tool.FormatPayload(hex.EncodeToString(bytePayload))
		log.Info("payload: %s", output)
	case gadget.CC5:
		log.Info("get payload: %s", gadget.CC1)
		bytePayload := gososerial.GetCC1(command)
		output := tool.FormatPayload(hex.EncodeToString(bytePayload))
		log.Info("payload: %s", output)
	case gadget.CC6:
		log.Info("get payload: %s", gadget.CC1)
		bytePayload := gososerial.GetCC1(command)
		output := tool.FormatPayload(hex.EncodeToString(bytePayload))
		log.Info("payload: %s", output)
	case gadget.CC7:
		log.Info("get payload: %s", gadget.CC1)
		bytePayload := gososerial.GetCC1(command)
		output := tool.FormatPayload(hex.EncodeToString(bytePayload))
		log.Info("payload: %s", output)
	default:
		log.Error("error payload")
		return
	}
}