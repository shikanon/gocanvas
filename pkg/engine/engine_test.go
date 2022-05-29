package engine

import (
	"fmt"
	"os"
	"testing"
)

func TestHtmlEngine_Render(t *testing.T) {
	tests := []struct {
		name    string
		e       *HtmlEngine
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test html engine",
			e: &HtmlEngine{
				Writer: os.Stdout,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.e.Render(); (err != nil) != tt.wantErr {
				t.Errorf("HtmlEngine.Render() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHtmlEngine_RenderBytes(t *testing.T) {
	tests := []struct {
		name       string
		e          *HtmlEngine
		wantResult []byte
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			name: "test html engine for camera",
			e: &HtmlEngine{
				Writer:         os.Stdout,
				CameraResource: []Camera{{Name: "01", Pos: Position{X: 1, Y: 1, Z: 30}, Interactive: true}},
				LightResource:  []Light{},
			},
			// wantResult: []byte(""),
			wantErr: false,
		},
		{
			name: "test html engine for light",
			e: &HtmlEngine{
				Writer:         os.Stdout,
				CameraResource: []Camera{},
				LightResource:  []Light{{Name: "01"}},
			},
			// wantResult: []byte(""),
			wantErr: false,
		},
		{
			name: "test html engine for model",
			e: &HtmlEngine{
				Writer:         os.Stdout,
				CameraResource: []Camera{},
				ModelResource: []Model{
					{
						Name:      "01",
						Type:      ModelType_GLTF,
						BasePath:  "assets/",
						ModelPath: "assets/scene.glt",
					},
				},
			},
			// wantResult: []byte(""),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := tt.e.RenderBytes()
			if (err != nil) != tt.wantErr {
				t.Errorf("HtmlEngine.RenderBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(string(gotResult))
			// if !reflect.DeepEqual(gotResult, tt.wantResult) {
			// 	t.Errorf("HtmlEngine.RenderBytes() = %v, want %v", gotResult, tt.wantResult)
			// }
		})
	}
}

func TestHtmlEngine_Check(t *testing.T) {
	tests := []struct {
		name    string
		e       *HtmlEngine
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test check camera",
			e: &HtmlEngine{
				CameraResource: []Camera{
					{Name: "1"},
					{Pos: Position{}},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.e.Check(); (err != nil) != tt.wantErr {
				t.Errorf("HtmlEngine.Check() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
