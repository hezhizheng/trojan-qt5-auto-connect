package main

import (
	"flag"
	"github.com/go-vgo/robotgo"
	"log"
	"os/exec"
	"strconv"
)

var  (
	scroll = 0
	cx = 79
	cy = 670
	path = "D:\\Trojan-Qt5-Windows.1.4.0\\trojan-qt5.exe"
)

func init()  {
	// 初始化参数
	_path := *flag.String("path", path, "trojan-qt5.exe的绝对路径")
	_scroll := *flag.String("scroll", "0", "鼠标X轴方向向下滚动的值，按需设置")
	_cx := *flag.String("cx", "79", "连接节点的X轴坐标")
	_cy := *flag.String("cy", "670", "连接节点的Y轴坐标")
	flag.Parse()

	intScroll ,_ := strconv.Atoi(_scroll)
	intX ,_ := strconv.Atoi(_cx)
	intY ,_ := strconv.Atoi(_cy)
	scroll = intScroll
	cx = intX
	cy = intY
	path = _path

}

func main() {

	//exec.Command("cmd", "/c", "start", "www.baidu.com").()
	log.Println("启动 trojan-qt5 ")
	// -windowstyle Maximized
	runErr := exec.Command("cmd", "/c", "start", path).Run()
	if runErr != nil {
		panic("启动 trojan-qt5 失败 " + runErr.Error())
	}
	log.Println("等待锁定当前窗口")
	//等待
	robotgo.MilliSleep(3 * 1000)

	// get current Window Active
	mdata := robotgo.GetActive()
	// set Window Active
	robotgo.SetActive(mdata)

	// 全屏
	robotgo.KeyTap(`space`,`alt`)
	robotgo.KeyTap(`x`)

	// 获取鼠标初始坐标
	ix, iy := robotgo.GetMousePos()
	log.Println("获取鼠标初始坐标",ix, iy)

	if scroll != 0{
		// 鼠标移动到最小可滑动位置
		robotgo.MoveMouse(420, 320)
		robotgo.ScrollMouse(scroll, "down")
	}

	// 鼠标移动到要连接的节点坐标位置
	robotgo.MoveMouse(cx, cy)
	// 双击鼠标左键
	robotgo.MouseClick("left", true)
	log.Println("连接节点",cx, cy)

	// close current Window
	robotgo.CloseWindow()
	log.Println("关闭 trojan-qt5 主页面")

	// 鼠标回归初始位置
	robotgo.MoveMouse(ix, iy)
	log.Println("鼠标回归初始位置，程序完成")
}
