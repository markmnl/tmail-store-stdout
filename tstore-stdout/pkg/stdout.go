package tstdout

import (
	"fmt"
	"time"
	"github.com/markmnl/tmail-store/tstore/pkg"
)

// ParentExists always returns true since this store doesn't know
func ParentExists(msg *tstore.Msg) (bool, error) {
	return true, nil
}

// Store pretty prints msg supplied to stdout
func Store(msg *tstore.Msg) error {
	_, err := fmt.Printf(`
id:     %s
pid:    %s
from:   %s
to:     %s
time:   %s
topic:  %s
type:	%s

%s
`, msg.ID, msg.PID, msg.From, msg.To, time.Unix(msg.Time, 0), msg.Topic, msg.Type, msg.Content)
	return err
}
