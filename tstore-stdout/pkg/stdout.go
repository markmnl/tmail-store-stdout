package tstdout

import (
	"fmt"
	"time"
	"github.com/markmnl/tmail-store/tstore/pkg"
)

// Store pretty prints msg supplied to stdout
func Store(msg *tstore.Msg) error {
	_, err := fmt.Printf(`
id:     %s
pid     %s
from:   %s
to:     %s
time:   %s
topic:  %s
type:	%s

%s
`, msg.ID, msg.PID, msg.From, msg.To, time.Unix(msg.Time, 0), msg.Topic, msg.Type, msg.Content)
	return err
}
