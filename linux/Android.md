<!--查看当前软件包名-->

 adb -s  127.0.0.1:5555 shell dumpsys window windows |findstr name=

<!--发送文件-->

adb push 电脑路径 手机文件夹路径

<!--拉取文件-->

##### adb pull 手机路径 电脑路径

<!--获取app启动时间-->

adb shell am start -w 包名
