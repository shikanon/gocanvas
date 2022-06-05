![logo](https://github.com/shikanon/gocanvas/blob/main/docs/logo.png)
# gocanvas

[![GitHub license](https://img.shields.io/github/license/shikanon/gocanvas)](https://github.com/shikanon/gocanvas/blob/master/LICENSE)
[![GitHub stars](https://img.shields.io/github/stars/shikanon/gocanvas)](https://github.com/shikanon/gocanvas/stargazers)
[![GitHub forks](https://img.shields.io/github/forks/shikanon/gocanvas)](https://github.com/shikanon/gocanvas/network)
[![Language](https://img.shields.io/badge/Language-Go-blue.svg)](https://golang.org/)
[![Go Report Card](https://goreportcard.com/badge/github.com/shikanon/gocanvas)](https://goreportcard.com/report/github.com/shikanon/gocanvas)

基于golang封装的3D模型动画演示库，支撑gltf、obj、fbx等格式，支撑天空盒、灯光和脚本动画编辑。

## Example

[example](./example)

## Usage

```go
import (
	"fmt"
	"net/http"

	gocanvas "github.com/shikanon/gocanvas/pkg/canvas"
)

func render(w http.ResponseWriter, r *http.Request) {
	app := gocanvas.NewApplication(w)
	scene := app.Scene()
	// 创建一个摄像机
	camera := gocanvas.DefaultCamera()
	// 创建一个天空盒
	skybox := gocanvas.DefaultSkybox()
	// 创建一个灯光
	light := gocanvas.DefaultLight()
	// 创建一个模型
	model, err := gocanvas.LoadGLTFModel("./assets/corvette_stingray/scene.gltf")
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
	// 加入到场景中
	scene.AddCamera(camera)
	scene.AddSkyBox(skybox)
	scene.AddLight(light)
	scene.AddModel(model)
	// 渲染
	err = scene.Render()
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
}

func main() {
	http.HandleFunc("/", render)
	http.Handle("/src/", http.FileServer(http.Dir("")))
	http.Handle("/assets/", http.FileServer(http.Dir("")))
	http.ListenAndServe(":8080", nil)
}
```