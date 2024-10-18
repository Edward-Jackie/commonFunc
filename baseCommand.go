package main

import (
	"log"
	"os/exec"
)

func welcome() {
	cmd := exec.Command("echo", "Hello World!")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("执行 echo 命令失败: %v", err)
	}
	log.Printf("欢迎使用: %s", output)
}

// 转化路径
func cd() {

}

// 执行命令
func execute() {

}
