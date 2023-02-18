package logs

import "log"

type LocalLog struct{}

func (l LocalLog) Log(info string, message string) {
	log.Printf("%s: %s\n", info, message)
}
