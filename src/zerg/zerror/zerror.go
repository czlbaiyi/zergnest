package zerror

import "fmt"
import "zerg/log"

func New(title string, info string) error {
	errorinfo := fmt.Sprintf("title : %s , info : %s ", title, info)
	log.Error(errorinfo)
	errorinfo = title + info
	return &Zerror{errorinfo}
}

// errorString is a trivial implementation of error.
type Zerror struct {
	s string
}

func (e *Zerror) Error() string {
	return e.s
}
