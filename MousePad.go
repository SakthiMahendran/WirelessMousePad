package main

import (
	"fmt"
	"net"
	"time"

	"github.com/SakthiMahendran/WirelessMousePad/core"
	ct "github.com/daviddengcn/go-colortext"
	"github.com/go-vgo/robotgo"
)

func main() {
	go preventScreenSleep()

	printWelcomeScreen()

	mouseEvents := make(chan core.MouseEvent)

	webServer := core.WebServer{}
	go webServer.Start(mouseEvents)

	for e := range mouseEvents {
		robotgo.Scroll(0, int(e.ScrollY))
	}

}

func preventScreenSleep() {
	ticker := time.NewTicker(time.Second * 10)

	for range ticker.C {
		robotgo.MoveRelative(-1, -1)
		robotgo.MoveRelative(1, 1)
	}
}

func getIPAddress() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Error retrieving IP address:", err)
		return ""
	}

	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			return ipNet.IP.String()
		}
	}

	return ""
}

func printWelcomeScreen() {
	printData("WARNING", "          This beta release may contain bugs", ct.Red)
	printData("APPLICATION", "      Wireless Mouse Pad", ct.Green)
	printData("SUPPORT DEVELOPER", "https://github.com/SakthiMahendran", ct.Green)
	printData("SERVER URL", "       http://"+getIPAddress(), ct.Green)
}

func printData(label, info string, color ct.Color) {
	ct.Background(color, true)
	ct.Foreground(ct.White, true)
	fmt.Print(label, " :")

	ct.ResetColor()

	ct.Foreground(color, true)
	fmt.Println(" " + info)

	ct.ResetColor()
}
