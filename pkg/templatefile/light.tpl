{{ define "light" }}
    // create directional light entity
    {{- range .LightResource }}
    var light{{ .Name }} = new pc.Entity('light{{ .Name }}');
    light{{ .Name }}.addComponent('light',);
    {{- if .EulerAngle }}
    light{{ .Name }}.setEulerAngles({{ .EulerAngle.X }}, {{ .EulerAngle.Y }}, {{ .EulerAngle.Z }});
    {{- end }}
    app.root.addChild(light{{ .Name }});
    {{- end }}

    // Set a prefiltered cubemap as the skybox
    {{- with .SkyBoxResource}}
    {{- if ne .Name ""}}
    var cubemapAsset{{ .Name }} = new pc.Asset('{{ .Name }}', 'cubemap', {
        url: "{{ .URL }}"
    }, {
        "textures": [
            {{- range .Textures}}
            "{{.}}",
            {{- end }}
        ],
        "magFilter": {{ .MagFilter }},
        "minFilter": {{ .MinFilter }},
        "anisotropy": {{ .Anisotropy }},
        "name": "{{ .Name }}",
        "rgbm": true,
        "prefiltered": "{{ .Prefiltered }}"
    });
    app.assets.add(cubemapAsset{{ .Name }});
    app.assets.load(cubemapAsset{{ .Name }});

    cubemapAsset{{ .Name }}.ready(function () {
        app.scene.skyboxMip = 2;
        app.scene.setSkybox(cubemapAsset{{ .Name }}.resources);
    });
    {{- end }}
    {{- end }}
{{ end }}