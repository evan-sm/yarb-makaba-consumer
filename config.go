package main

import "os"

var (
	ginPortMakabaConsumer = os.Getenv("GIN_PORT_MAKABA_CONSUMER")
	Passcode              = os.Getenv("PASSCODE")
	yarbBasicAuthUser     = os.Getenv("YARB_BASIC_AUTH_USER")
	yarbBasicAuthPass     = os.Getenv("YARB_BASIC_AUTH_PASS")
)
