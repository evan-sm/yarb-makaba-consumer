package main

import (
	"fmt"
	"os"
)

var (
	ginPortMakabaConsumer = os.Getenv("GIN_PORT_MAKABA_CONSUMER")
	Passcode              = os.Getenv("PASSCODE")
	yarbBasicAuthUser     = os.Getenv("YARB_BASIC_AUTH_USER")
	yarbBasicAuthPass     = os.Getenv("YARB_BASIC_AUTH_PASS")
	YarbDBIp              = os.Getenv("YARB_DB_IP")
	YarbDBPort            = os.Getenv("YARB_DB_PORT")
	YarbDBApiURL          = fmt.Sprintf("%v:%v", YarbDBIp, YarbDBPort)
)
