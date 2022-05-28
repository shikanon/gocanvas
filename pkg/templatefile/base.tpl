{{ define "base" }}
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
            
            // 初始化摄像头
            {{ template "light" . }}

            // 初始化模型
            {{ template "model" . }}
        </script>
    </body>
</html>
{{ end }}