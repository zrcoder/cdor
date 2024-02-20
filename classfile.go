package cdor

const GopPackage = true

type App = Cdor

type IApp interface {
	init()
	MainEntry()
}

func Gopt_App_Main(app IApp) {
	app.init()
	app.MainEntry()
}
