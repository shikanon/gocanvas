{{ define "element" }}
    // 将字体添加到assets
    var fontAsset = new pc.Asset('courier.json', "font", { url: "assets/fonts/courier.json" });
    app.assets.add(fontAsset);
    // app.assets.load(fontAsset);
    {{- range .ElementResource }}
    // 创建一个屏幕，然后将文字或图片加上去
    var screen = new pc.Entity();
    screen.setLocalScale(0.01, 0.01, 0.01);
    screen.addComponent("screen", {
             referenceResolution: new pc.Vec2(1280, 720),
             screenSpace: true
            });
    app.root.addChild(screen);
    var element{{ .Name }} = new pc.Entity('element{{ .Name }}');
    element{{ .Name }}.addComponent("element", {
                    pivot: new pc.Vec2(0.5, 0.5),
                    anchor: new pc.Vec4(0, 0.4, 1, 1),
                    margin: new pc.Vec4(0, 0, 0, 0),
                    fontAsset: fontAsset,
                    fontSize: 20,
                    text: `{{ .Text }}`,
                    type: pc.ELEMENTTYPE_TEXT
                });
    screen.addChild(element{{ .Name }})
    {{- end }}
{{ end }}