package main

import (
	"embed"

	uos "github.com/mschoebel/urban-octo-succotash"
)

//go:embed assets/static/*
var staticAssetsFS embed.FS

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

	// start web application
	uos.StartApp()
}