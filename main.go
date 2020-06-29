/**
 * Auth :   liubo
 * Date :   2020/6/29 11:21
 * Comment:
 */

package main

import (
	"flag"
	"fmt"
	"github.com/keybase/go-ps"
	"os"
	"strings"
)

var ls = flag.Bool("ls", false, "-ls")
var kill = flag.Int("kill", 0, "--kill=pid")
var search = flag.String("s", "", "--s=appName")

func main() {
	flag.Parse()

	if *ls {
		listProcess()
		return
	}

	if *kill > 0 {
		killProcess(*kill)
	}

	if len(*search) > 0 {
		searchProcess(*search)
	}
}

func listProcess() {
	var psList, _ = ps.Processes()
	dumpProcessList(psList)
}

func dumpProcessList(psList []ps.Process) {
	fmt.Printf("i\tPid\tName\tPath\n")
	for i, v := range psList {
		var pathName, _ = v.Path()
		fmt.Printf("%d\t%d\t%s\t%s\n", i, v.Pid(), v.Executable(), pathName)
	}
}

func killProcess(pid int) {

	var p, _ = os.FindProcess(pid)
	if p == nil {
		fmt.Println("没有此进程: ", pid)
		return
	}
	var e = p.Kill()
	if e != nil {
		fmt.Println("杀掉进程失败：", e.Error())
		return
	}
	fmt.Println("杀掉进程成功:", pid)
}

func searchProcess(psName string) {
	var psList, err = ps.Processes()
	var result []ps.Process
	if err == nil {
		for _, v := range psList {
			var exeName = strings.ToLower(v.Executable())
			var pathName, _ = v.Path()
			pathName = strings.ToLower(pathName)
			if strings.Contains(exeName, strings.ToLower( psName)) ||
				strings.Contains(pathName, strings.ToLower( psName)) {
				result = append(result, v)
			}
		}
	}

	dumpProcessList(result)
}