{{ define "camera" }}
    {{- range .CameraResource }}
    var camera{{ .Name }} = new pc.Entity('camera{{ .Name }}');
    camera{{ .Name }}.addComponent('camera', {
        clearColor: [ 1, 1, 1 ]
    });
    {{- if .Pos }}
    camera{{ .Name }}.setLocalPosition({{ .Pos.X }}, {{ .Pos.Y }}, {{ .Pos.Z }});
    {{- end }}
    app.root.addChild(camera{{ .Name }});
    {{- end }}
{{ end }}