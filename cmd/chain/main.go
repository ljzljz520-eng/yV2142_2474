package main

import (
	"fmt"
	"idiomchain/internal/clock"
	"idiomchain/internal/config"
	"idiomchain/internal/report"
	"idiomchain/internal/service"
	"idiomchain/internal/store"
	"os"
)

func run(args []string) error {
	c := config.Load()
	if len(args) > 1 {
		c.Session = args[1]
	}
	s, e := store.Open(c.DBPath)
	if e != nil {
		return e
	}
	defer s.Close()
	svc := service.New(s, clock.Real{})
	session, e := svc.EnsureSession(c.Session)
	if e != nil {
		return e
	}
	if len(args) > 2 {
		_, e = svc.Record(session, args[2])
		return e
	}
	entries, e := svc.History(session, 1, 100)
	if e != nil {
		return e
	}
	rej, e := svc.Rejections(session)
	if e != nil {
		return e
	}
	fmt.Print(report.RenderHistory(session, entries, rej))
	return nil
}
func main() {
	if e := run(os.Args); e != nil {
		fmt.Fprintln(os.Stderr, e)
		os.Exit(1)
	}
}
