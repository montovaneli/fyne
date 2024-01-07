//go:build !ci && (!android || !ios || !mobile) && (wasm || test_web_driver)

package app

import (
	"fyne.io/fyne/v2"
)

func (app *fyneApp) SendNotification(_ *fyne.Notification) {
	// TODO #2735
	fyne.LogError("Sending notification is not supported yet.", nil)
}

func rootConfigDir() string {
	return "/data/"
}
