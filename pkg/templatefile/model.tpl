{{ define "model" }}
    // add model
    {{- range .ModelResource }}
    var model{{ .Name }} = new pc.Entity("model{{ .Name }}");
    app.root.addChild(model{{ .Name }});
    
    // load model for glTF scene root that rotates
    {{- if eq .Type "gltf" }}
    app.assets.loadFromUrl('{{ .ModelPath }}', 'json', function (err, asset) {
        var json = asset.resource;
        var gltf = JSON.parse(json);
        loadGltf(gltf, app.graphicsDevice, function (err, res) {
            // add the loaded scene to the hierarchy
            model{{ .Name }}.addComponent('model');
            model{{ .Name }}.model.model = res.model;
        }, {
            basePath: '{{ .BasePath }}'
        });
    });
    {{- end }}
    // set position
    {{- if .Pos }}
    model{{ .Name }}.setLocalPosition({{ .Pos.X }}, {{ .Pos.Y }}, {{ .Pos.Z }});
    {{- end }}
    {{- end }}
{{ end }}