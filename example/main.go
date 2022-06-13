package main

import (
	"fmt"
	"log"
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
	// 创建UI，比如文字或按钮
	// element := gocanvas.DefaultElement("gocanvas car")
	// scene.AddElement(element)
	// 创建一个模型
	model, err := gocanvas.LoadGLTFModel("./assets/corvette_stingray/scene.gltf")
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
	model.Pos = engine.XYZ{
		X: 1,  // 左右
		Y: -1, // 上下
		Z: -4, // 前后
	}
	model.EulerAngle = engine.XYZ{
		X: 1,   // 沿x旋转
		Y: 330, // 沿y轴旋转
		Z: 0,   // 沿z轴旋转
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
	log.Println("the service is runing with 127.0.0.1:8080/")
	http.ListenAndServe(":8080", nil)
}
