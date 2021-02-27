package console

import (
	"golang.org/x/sys/windows"
	"log"
)

func enableVirtualTerminalProcessing(console windows.Handle) error {
	var mode uint32
	err := windows.GetConsoleMode(console, &mode)
	if err != nil {
		return err
	}
	return windows.SetConsoleMode(console, mode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
}

func Init() {
	err := enableVirtualTerminalProcessing(windows.Stdout)
	if err != nil {
		log.Println(err)
	}
}
