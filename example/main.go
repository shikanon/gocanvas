package main

import (
	"fmt"
	"net/http"

	gocanvas "github.com/shikanon/gocanvas/pkg/canvas"
	"github.com/shikanon/gocanvas/pkg/engine"
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
	// 创建UI
	element := gocanvas.DefaultElement("hello world")
	// 创建一个模型
	model, err := gocanvas.LoadGLTFModel("./assets/corvette_stingray/scene.gltf")
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
	model.Pos = engine.Position{
		X: 0,  // 左右
		Y: -2, // 上下
		Z: -5, // 前后
	}

	// 加入到场景中
	scene.AddCamera(camera)
	scene.AddSkyBox(skybox)
	scene.AddLight(light)
	scene.AddModel(model)
	scene.AddElement(element)
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
