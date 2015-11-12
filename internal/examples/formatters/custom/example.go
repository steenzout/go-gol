package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mediaFORGE/gol"
)

// Custom struct for a custom formatter.
type Custom struct{}

// Format formats the log message.
func (f Custom) Format(msg *gol.LogMessage) (string, error) {
	lmsg := msg.FieldLength()
	buffer := make([]string, lmsg, lmsg)

	i := 0
	for k, v := range *msg {
		buffer[i] = fmt.Sprintf("%s:'%s'", k, v)
		i += 1
	}

	return fmt.Sprintf("%s\n", strings.Join(buffer, " ")), nil
}

var _ gol.LogFormatter = (*Custom)(nil)

var logger gol.Logger = gol.SimpleLog(nil, &Custom{}, os.Stdout)

func main() {
	// this will be written to stderr
	logger.Send(gol.NewEmergency("message", "system is down"))
	logger.Send(gol.NewAlert("message", "failed to write to disk"))
	logger.Send(gol.NewCritical("message", "high server load"))
	logger.Send(gol.NewError("message", "invalid number format"))

	// this will not be written anywhere
	logger.Send(gol.NewWarning("message", "performance close to 1s threshold"))
	logger.Send(gol.NewNotice("message", "failed to communicate with monitoring service"))
	logger.Send(gol.NewInfo("message", "requested processed in 250ms"))
	logger.Send(gol.NewDebug("debug", "var x = 10"))
}
