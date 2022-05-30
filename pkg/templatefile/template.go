package templatefile

import (
	_ "embed"
	"text/template"
)

//go:embed base.tpl
var base []byte

//go:embed application.tpl
var application []byte

//go:embed camera.tpl
var camera []byte

//go:embed light.tpl
var light []byte

//go:embed model.tpl
var model []byte

//go:embed element.tpl
var element []byte

func BaseTemplate() string {
	return string(base)
}

func ApplicationTemplate() string {
	return string(application)
}

func CameraTemplate() string {
	return string(camera)
}

func LightTemplate() string {
	return string(light)
}

func ModelTemplate() string {
	return string(model)
}

func ElementTemplate() string {
	return string(element)
}

func TemplateMerge(args ...string) (tpl *template.Template, err error) {
	tpl = template.New("engine")
	for _, arg := range args {
		tpl, err = tpl.Parse(arg)
		if err != nil {
			return
		}
	}
	return tpl, err
}

func EngineTemplate() (*template.Template, error) {
	return TemplateMerge(BaseTemplate(), ApplicationTemplate(), CameraTemplate(), LightTemplate(), ModelTemplate(), ElementTemplate())
}
