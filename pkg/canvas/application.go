package canvas

import (
	"io"

	"github.com/shikanon/gocanvas/pkg/engine"
)

func NewApplication(w io.Writer) *Application {
	scene := &ApplicationScene{
		engine.HtmlEngine{
			Writer: w,
		},
	}
	return &Application{
		scene: scene,
	}
}

type Application struct {
	scene *ApplicationScene
}

func (p *Application) Scene() *ApplicationScene {
	return p.scene
}

type ApplicationScene struct {
	engine.HtmlEngine
}

func (s *ApplicationScene) AddCamera(c engine.Camera) {
	s.CameraResource = append(s.CameraResource, c)
}

func (s *ApplicationScene) AddModel(m engine.Model) {
	s.ModelResource = append(s.ModelResource, m)
}

func (s *ApplicationScene) AddLight(l engine.Light) {
	s.LightResource = append(s.LightResource, l)
}

func (s *ApplicationScene) AddSkyBox(box engine.SkyBox) {
	s.SkyBoxResource = box
}
