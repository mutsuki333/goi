package main

import "github.com/mutsuki333/goi/modules/log"

/**
 * setup client js methods
 *
**/

func bindJS() {
	ui.Bind("add", func(a, b int) int {
		log.Error("test")
		return a + b
	})
}
