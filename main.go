package main

import (
	"embed"

	uos "github.com/mschoebel/urban-octo-succotash"
)

//go:embed assets/static/*
var staticAssetsFS embed.FS

var (
	mLogins int
)

func main() {
	// initialize application framework components
	uos.ComponentSetup()
	defer uos.ComponentCleanup()

	// register data models (auto-migration)
	uos.RegisterDBModels(
		&person{},
	)

	// register request handler
	uos.RegisterAppRequestHandlers(
		uos.PageHandler("/", "home"),
		uos.PageHandler("examples"),
		uos.PageHandler("internal"),

		uos.PageHandler("login").AuthPage(),
		uos.PageHandler("protected").Internal(),

		uos.DialogHandler(
			hintDialog{},
			dataDialog{},
			loginDialog{},
		),
		uos.FormHandler(
			loginForm{},
			dataForm{},
		),
		uos.TableHandler(
			dataTable{},
		),

		uos.ActionHandler(),

		uos.FragmentHandler(),
		uos.MarkdownHandler(),
	)
	uos.RegisterStaticAssets("/assets/static/", staticAssetsFS)

	// register metric - count successful login attempts
	mLogins = uos.Metrics.RegisterCounter(
		"app_login_count",
		"Number of successful logins.",
	)

	// start web application
	uos.StartApp()
}
