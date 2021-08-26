package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/mutsuki333/goi/client"
	"github.com/mutsuki333/goi/modules/log"
	"github.com/mutsuki333/goi/modules/spa"
	"github.com/mutsuki333/goi/modules/value"
	"github.com/zserge/lorca"
)

var ui lorca.UI
var err error
var done chan struct{}
var quit chan os.Signal

func init() {
	quit = make(chan os.Signal, 1)
	done = make(chan struct{}, 1)
}

func main() {
	/*
	* setups
	* setup system vars
	 */

	value.SetDefault("log.level", 15)

	value.Read()
	log.Level(value.GetInt("log.level"))
	log.Info("Staring...")

	f := &spa.SpaFS{
		FS:   client.WebFS,
		Root: "dist",
	}

	// ui instanse
	ui, err = lorca.New("http://localhost:10088", "", 800, 600)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	bindJS()

	// start http server in goroutine
	go func() {
		http.ListenAndServe("localhost:10088", f)
	}()

	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-quit:
	case <-ui.Done():
	}

	// post app actions
	go func() {
		log.Debug("Before terminate.")
		value.Save()
		done <- struct{}{}
	}()

	<-done
	log.Info("Program exited.")
}
