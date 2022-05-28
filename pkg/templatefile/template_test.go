package templatefile

import (
	_ "embed"
	"testing"
)

func TestBaseTemplate(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test base template",
			want: `{{ define "base" }}
<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<meta name='viewport' content='width=device-width, initial-scale=1, maximum-scale=1, minimum-scale=1, user-scalable=no' />
		<style>
			body {
				margin: 0;
				overflow: hidden;
			}
		</style>
		<script src='./src/playcanvas-stable.js'></script>
		<script src='./src/playcanvas-gltf.js'></script>
	</head>
	<body>
		<script>
			// create canvas dom
			var canvas = document.createElement('canvas');
			document.body.appendChild(canvas);

			// 初始化 application
			{{ template "application" . }}
			
			// 初始化摄像头
			{{ template "camera" . }}
			
			// 初始化模型
			{{ template "model" . }}
		</script>
	</body>
</html>
{{ end }}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BaseTemplate(); got != tt.want {
				t.Errorf("BaseTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}
