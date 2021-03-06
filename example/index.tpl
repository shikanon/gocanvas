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
            var canvas = document.createElement('canvas');
            document.body.appendChild(canvas);

            var app = new pc.Application(canvas, {});
            app.start();

            // fill the available space at full resolution
            app.setCanvasFillMode(pc.FILLMODE_FILL_WINDOW);
            app.setCanvasResolution(pc.RESOLUTION_AUTO);

            // ensure canvas is resized when window changes size
            window.addEventListener('resize', function() {
                app.resizeCanvas();
            });

            // create camera entity
            var camera = new pc.Entity('camera');
            camera.addComponent('camera', {
                clearColor: [ 1, 1, 1 ]
            });
            app.root.addChild(camera);
            camera.setLocalPosition(0, 1, 20);

            // create directional light entity
            var light = new pc.Entity('light');
            light.addComponent('light',);
            app.root.addChild(light);
            light.setEulerAngles(45, 0, 45);

            // rotator script
            var Rotate = pc.createScript('rotate');
            Rotate.prototype.update = function (deltaTime) {
                this.entity.rotate(0, deltaTime * 20, 0);
            };

            // glTF scene root that rotates
            var gltfRoot = new pc.Entity();
            gltfRoot.addComponent('script');
            gltfRoot.script.create('rotate');
            app.root.addChild(gltfRoot);

            app.assets.loadFromUrl('assets/studded_tyre/scene.gltf', 'json', function (err, asset) {
                var json = asset.resource;
                var gltf = JSON.parse(json);
                loadGltf(gltf, app.graphicsDevice, function (err, res) {
                    // add the loaded scene to the hierarchy
                    gltfRoot.addComponent('model');
                    gltfRoot.model.model = res.model;
                }, {
                    basePath: 'assets/studded_tyre/'
                });
            });
        </script>
    </body>
</html>