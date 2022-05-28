package engine

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math/rand"

	tfile "github.com/shikanon/gocanvas/pkg/templatefile"
)

var (
	ErrNil                   = errors.New("value is nil")
	ErrInvalid               = errors.New("value is invalid value")
	ModelType_GLTF ModelType = "gltf"
	ModelType_GLB  ModelType = "glb"
	ModelType_OBJ  ModelType = "obj"
)

type Engine interface {
	Render() ([]byte, error)
	Check() error
}

type TemplateValue interface {
	Check() error
	FillDefault()
}

type EngineMeta struct{}
type Camera struct {
	Name string
	Pos  Position
}

func (c *Camera) Check() (err error) {
	if c.Name == "" {
		err = fmt.Errorf("camera name %w", ErrNil)
		return
	}
	if c.Pos == (Position{}) {
		err = fmt.Errorf("camera position %w", ErrNil)
		return
	}
	return
}

func (c *Camera) FillDefault() {
	if c.Name == "" {
		c.Name = fmt.Sprintf("%d", rand.Int())
	}
	if c.Pos == (Position{}) {
		c.Pos = Position{
			X: 0,
			Y: 0,
			Z: 0,
		}
	}
}

type Light struct {
	Name string
}

func (l *Light) Check() (err error) {
	if l.Name == "" {
		err = fmt.Errorf("light name %w", ErrNil)
		return
	}
	return
}

func (l *Light) FillDefault() {
	if l.Name == "" {
		l.Name = fmt.Sprintf("%d", rand.Int())
	}
}

type ModelType string

type Model struct {
	Name      string
	Type      ModelType
	BasePath  string
	ModelPath string
}

func (m *Model) Check() (err error) {
	if (m.Name == "") || (m.Type == "") || (m.BasePath == "") || (m.ModelPath == "") {
		err = fmt.Errorf("model name %w", ErrNil)
		return
	}
	if m.Type == "" {
		err = fmt.Errorf("model type %w", ErrNil)
		return
	}
	return
}

type Script struct{}
type SkyBox struct {
	Name        string
	URL         string
	Textures    [6]string
	MagFilter   int
	MinFilter   int
	Anisotropy  int
	RGBM        bool
	Prefiltered string
}

type HtmlEngine struct {
	Writer         io.Writer
	Meta           []EngineMeta // 基本属性，比如是否iframe、占据全屏
	CameraResource []Camera     // 相机资源
	LightResource  []Light      // 灯光资源
	ModelResource  []Model      // 模型资源
	SkyBoxResource SkyBox       // 天空盒
}

func (e *HtmlEngine) Check() (err error) {
	for i, c := range e.CameraResource {
		err = c.Check()
		if errors.Is(err, ErrNil) {
			c.FillDefault()
			err = nil
		}
		e.CameraResource[i] = c
	}

	return
}

func (e *HtmlEngine) render(w io.Writer) (err error) {
	tpl, err := tfile.EngineTemplate()
	if err != nil {
		return
	}

	return tpl.ExecuteTemplate(w, "base", e)
}

// 渲染
func (e *HtmlEngine) Render() (err error) {
	err = e.render(e.Writer)
	return
}

// 渲染成字节输出
func (e *HtmlEngine) RenderBytes() (result []byte, err error) {
	buff := &bytes.Buffer{}
	err = e.render(buff)
	return buff.Bytes(), err
}
