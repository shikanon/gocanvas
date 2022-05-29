{{ define "camera" }}
    {{- range .CameraResource }}
    {{- if ne .Name ""}}
    var camera{{ .Name }} = new pc.Entity('camera{{ .Name }}');
    camera{{ .Name }}.addComponent('camera', {
        clearColor: [ 1, 1, 1 ]
    });
    {{- if .Pos }}
    camera{{ .Name }}.setLocalPosition({{ .Pos.X }}, {{ .Pos.Y }}, {{ .Pos.Z }});
    {{- end }}

    // Make the camera interactive
    {{- if .Interactive }}
    camera{{ .Name }}.addComponent('script');
    app.assets.loadFromUrl('./src/orbit-camera.js', 'script', function (err, asset) {
        camera{{ .Name }}.script.create('orbitCamera');
        camera{{ .Name }}.script.create('keyboardInput');
        camera{{ .Name }}.script.create('mouseInput');
        camera{{ .Name }}.script.create('touchInput');

        if (this.cameraPosition) {
            camera{{ .Name }}.script.orbitCamera.distance = this.cameraPosition.length();
        } else if (this.gltf) {
            camera{{ .Name }}.script.orbitCamera.focusEntity = this.gltf;
        }
    }.bind(this));
    {{- end }}

    app.root.addChild(camera{{ .Name }});
    {{- end }}
    {{- end }}
{{ end }}