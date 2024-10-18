package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

var baseDir string

func main() {
	// 创建项目
	initProject("test")
	// 创建基础目录
	// 创建内部子目录
	// 基于配置进行创建基础服务
	// Web框架
	// ORM框架
	// 配置文件
	// 构建工具
	// 模板目录
	// 日志目录
	// 其他目录
}

func getPwd() string {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前目录失败:", err)
		return ""
	}
	fmt.Println("当前目录是:", currentDir)
	return currentDir
}

// 初始化项目
func initProject(projectName string) {
	baseDir = getPwd()

	// 切换回原始工作目录
	err := os.Chdir(baseDir)
	if err != nil {
		log.Fatalf("切换回原始工作目录失败: %v", err)
	}

	// 打印 Hello World!
	cmd := exec.Command("echo", "Hello World!")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("执行 echo 命令失败: %v", err)
	}
	log.Printf("欢迎使用: %s", output)

	// 打印当前工作目录
	cmd = exec.Command("pwd")
	output, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("执行 pwd 命令失败: %v", err)
	}
	log.Printf("当前工作目录: %s", output)

	// 初始化 Go 模块
	cmd = exec.Command("go", "mod", "init", projectName)
	output, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("初始化 Go 模块失败: %v", err)
	}
	log.Printf("Go 模块初始化成功: %s", output)
}

// 在父路径下进行创建内部目录
func createDirList(parentDir string, dirList []int) {
	// 创建内部子目录
	for _, subDir := range dirList {
		dirPath := fmt.Sprintf("%s/internal/%s", parentDir, subDir)
		err := os.Mkdir(dirPath, os.ModePerm)
		if err != nil {
			fmt.Printf("创建目录 %s 失败: %v\n", dirPath, err)
		} else {
			fmt.Printf("成功创建目录: %s\n", dirPath)
		}
	}
}
