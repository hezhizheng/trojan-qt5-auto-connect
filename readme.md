## trojan-qt5-auto-connect 启动 trojan-qt5 自动连接指定节点(鸡肋 2021-12-24 更新)
> 使用 `robotgo` 模拟鼠标键盘的操作，使得 trojan-qt5 在启动时自动连接指定节点

## 使用
- 下载 [release](https://github.com/hezhizheng/trojan-qt5-auto-connect/releases)
- 参数说明
```
-cx string
    连接节点的X轴坐标 (default "79")
-cy string
    连接节点的Y轴坐标 (default "670")
-path string
    trojan-qt5.exe的绝对路径 (default "D:\\Trojan-Qt5-Windows.1.4.0\\trojan-qt5.exe")
-scroll string
    鼠标X轴方向向下滚动的值，按需设置 (default "0")
```
- 启动命令
```
./trojan-qt5-auto-connect_windows_amd64.exe -path D:\\Trojan-Qt5-Windows.1.4.0\\trojan-qt5.exe -cx 79 -cy 670
```

## 编译
```
gox -osarch="windows/amd64" -ldflags "-s -w" -gcflags="all=-trimpath=${PWD}" -asmflags="all=-trimpath=${PWD}" -output trojan-qt5-auto-connect_windows_amd64
```
## TODO
- [ ] 多屏幕坐标定位问题
- [ ] win10 下开机自启，长时间停留在输入密码界面，输入密码之后进入主页面没有自动连接节点
