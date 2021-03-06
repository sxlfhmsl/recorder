# 简单介绍
该工具是一个记录关键步骤截图并生成GIF的简易工具，名为recorder。recorder支持图上标记（可多点标记），并可在录制的过程中做暂停以及恢复操作。操作及构建方法如下述：

# 使用方法
首先启动，后开始录制，过程中可暂停及恢复，完成后可生成GIF。

## 启动与停止
* 同时按住```    command(win键)+alt+s    ```启动录制。
* 启动录制后，同时按住```    command(win键)+alt+s    ```停止录制，后在程序运行目录下生成一个GIF文件（录制结果）。

## 暂停与恢复
* 同时按住```    command(win键)+alt+z    ```暂停录制，暂停后不会再截屏即记录步骤。
* 暂停后，同时按住```    command(win键)+alt+z    ```恢复录制。

## 录制（截屏）
* 鼠标左键点击：会进行截屏，并在鼠标位置绘制特定标记<?xml version="1.0" standalone="no"?><!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd"><svg t="1638721810522" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="2633" width="32" height="32" xmlns:xlink="http://www.w3.org/1999/xlink"><defs><style type="text/css"></style></defs><path d="M455.68 404.032l91.264 529.152c0 0 67.456-69.44 123.136-117.888l122.432 163.904c4.928 6.656 15.68 7.104 23.872 1.088l52.288-38.208c8.256-6.016 10.944-16.32 5.952-22.976l-119.104-159.424c62.208-25.088 164.672-53.632 164.672-53.632L455.68 404.032zM308.352 648.384l-135.872 99.328c-20.544 15.04-24.256 43.968-8 65.408 16.256 21.376 46.272 27.008 66.752 12.032l135.872-99.328c20.992-15.36 24.512-45.504 8.256-66.88C359.168 637.504 329.344 633.024 308.352 648.384zM949.696 238.976c-16.256-21.376-45.632-26.176-67.072-10.496l-134.912 98.688c-21.44 15.68-25.152 44.672-8.896 66.048 16.256 21.376 46.272 27.008 67.712 11.328l134.912-98.688C962.88 290.176 965.952 260.352 949.696 238.976zM319.296 136.832c-15.936-20.928-45.248-25.728-66.752-10.048-20.096 14.72-24.256 43.968-8.32 64.896l105.536 138.816c15.936 20.992 45.696 25.408 65.792 10.688 21.44-15.68 25.216-44.608 9.28-65.6L319.296 136.832zM585.792 301.76c26.176 4.224 50.24-13.376 53.632-39.232l21.184-167.808c3.392-25.792-14.976-49.984-41.536-54.656-26.176-4.224-50.24 13.376-53.632 39.168l-21.248 167.872C540.928 272.96 559.296 297.088 585.792 301.76zM329.728 489.024c2.56-25.92-15.808-50.048-41.536-54.656l-170.048-27.968c-27.072-3.584-50.688 13.696-53.632 39.232-3.904 26.944 14.464 51.072 41.536 54.656l170.048 27.968C301.824 532.736 325.504 515.456 329.728 489.024z" p-id="2634" fill="#d81e06"></path></svg>。
* 鼠标中键点击：连续的点击为一组，将使用第一次点击时的截图作为底图，然后将若干次连续的点击绘制在底图上，采用标记<?xml version="1.0" standalone="no"?><!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd"><svg t="1638722127524" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="3687" width="32" height="32" xmlns:xlink="http://www.w3.org/1999/xlink"><defs><style type="text/css"></style></defs><path d="M301.4 643.5L98 440.2v-0.1l54.1-54.1c16.5-16.5 41.2-21.4 62.7-12.5l136.9 56.7h0.1L616 226.5v-0.1l-70.7-70.7v-0.1l40.6-40.6c22.5-22.5 58.9-22.5 81.4 0l40.6 40.6 0.2-0.2 162.6 162.7-0.2 0.2 40.6 40.6c22.5 22.5 22.5 58.9 0 81.4l-40.6 40.6h-0.1l-70.8-70.7h-0.1L595.8 674.3v0.1l56.7 137c8.9 21.5 4 46.2-12.5 62.7l-54.1 54.1h-0.1L382.7 724.9h-0.1l-203 203H98.2v-81.3l203.1-203.1h0.1z" p-id="3688" fill="#d81e06"></path></svg>。
* 同时按住```    command(win键)+alt+x    ```：将会进行一次截图且不绘制任何标记。

## 最后
* 将所有截图按先后顺序组合起来，生成gif。

# 构建及运行说明
```shell
go get -d github.com/rakyll/statik
go install github.com/rakyll/statik
cd ./cmd/recorder
# 仅在windows下执行该语句
rsrc -manifest ../../config/nac.mainfest -ico ../../assets/favicon.ico -o nac.syso
go generate
go build
./recorder
```
