package canvas

import (
	"errors"
	"path"
	"strings"

	"github.com/shikanon/gocanvas/pkg/engine"
)

// 默认相机参数
func DefaultCamera() engine.Camera {
	return engine.Camera{
		Name: "camera",
		Pos: engine.Position{
			X: 0,
			Y: 0,
			Z: 10,
		},
	}
}

// 默认天空盒
func DefaultSkybox() engine.SkyBox {
	return engine.SkyBox{
		Name: "helipad",
		URL:  "./assets/cubemap/6079289/Helipad.dds",
		Textures: [6]string{
			"./assets/cubemap/6079292/Helipad_posx.png",
			"./assets/cubemap/6079290/Helipad_negx.png",
			"./assets/cubemap/6079293/Helipad_posy.png",
			"./assets/cubemap/6079298/Helipad_negy.png",
			"./assets/cubemap/6079294/Helipad_posz.png",
			"./assets/cubemap/6079300/Helipad_negz.png",
		},
		MagFilter:   1,
		MinFilter:   5,
		Anisotropy:  1,
		RGBM:        true,
		Prefiltered: "Helipad.dds",
	}
}

// 默认灯光
func DefaultLight() engine.Light {
	return engine.Light{
		Name: "light",
	}
}

// 加载gltf模型
func LoadGLTFModel(modelpath string) (m engine.Model, err error) {
	if strings.Contains(modelpath, "http://") {
		// TODO download to local path
		modelpath = func(string) string {
			return ""
		}(modelpath)
	}
	basepath, modelfile := path.Split(modelpath)
	if !strings.Contains(modelfile, "gltf") {
		err = errors.New("format error, must be gltf")
		return
	}
	m = engine.Model{
		Name:      "gltf",
		Type:      engine.ModelType_GLTF,
		ModelPath: modelpath,
		BasePath:  basepath,
	}
	return
}
