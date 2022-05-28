{{ define "application" }}
    // init canvas
    var app = new pc.Application(canvas, {
        mouse: new pc.Mouse(document.body),
        keyboard: new pc.Keyboard(window),
        touch: new pc.TouchDevice(window),
    });
    app.start();

    // 在全屏模式下填满可用空间
    app.setCanvasFillMode(pc.FILLMODE_FILL_WINDOW);
    app.setCanvasResolution(pc.RESOLUTION_AUTO);

    app.scene.gammaCorrection = pc.GAMMA_SRGB;
    app.scene.toneMapping = pc.TONEMAP_ACES;

    // 确保在窗口尺寸变化同时画布也同时改变其大小
    window.addEventListener('resize', function() {
        app.resizeCanvas();
    });
{{ end }}