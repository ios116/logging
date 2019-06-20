package logging

import (
	"testing"
	"bytes"
	"time"
	"fmt"
	
)

func TestLogger(t *testing.T) {
	var b bytes.Buffer
	
	date:=time.Now().Format("2006-01-02")
	
	hwa:=HwAccepted{3456,4}
	LogOtusEvent(&hwa,&b)
	mustHave := fmt.Sprintf("%s accepted %d %d",date, hwa.ID, hwa.Grade)
	if b.String() != mustHave {
		t.Errorf("we have %s but should be %s", b.String(),mustHave)
	}
	b.Reset()

	hws:=HwSubmitted{3456,"Code","please take a look at my homework"}
	mustHave = fmt.Sprintf("%s submitted %d \"%s\"",date, hws.ID, hws.Comment)
	LogOtusEvent(&hws,&b)
	if b.String() != mustHave {
		t.Errorf("we have %s but should be %s", b.String(),mustHave)
	}
}