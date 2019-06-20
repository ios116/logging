package logging

import (
	"io"
	"fmt"
	"time"
)

// HwAccepted this event homework is accepted
type HwAccepted struct {
	ID    int
	Grade int
}

// Log 2019-01-01 accepted 3456 4
func (hwa *HwAccepted) Log() string {
   return fmt.Sprintf("%s accepted %d %d",time.Now().Format("2006-01-02"), hwa.ID, hwa.Grade)
}

// HwSubmitted this event homework is submitted
type HwSubmitted struct {
	ID      int
	Code    string
	Comment string
}

// Log 2019-01-01 submitted 3456 "please take a look at my homework"
func (hws *HwSubmitted) Log() string {
	return fmt.Sprintf("%s submitted %d \"%s\"",time.Now().Format("2006-01-02"), hws.ID, hws.Comment)
 }
 
// OtusEvent interface for logging
type OtusEvent interface{
	Log() string
}

// LogOtusEvent function logs event to file
func LogOtusEvent(e OtusEvent, w io.Writer) {
     w.Write([]byte(e.Log()))
}