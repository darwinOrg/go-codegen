package dgen

import (
	"os"
	"path/filepath"
	"strings"
)

func GetCurrentPackageNameSimply() string {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dirs := strings.Split(currentDir, "/")
	return dirs[len(dirs)-1]
}

func GetPackagePrefixSimply() string {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	targetFile := "go.mod"
	startDir := currentDir

	for {
		// 检查当前目录是否包含目标文件
		targetPath := filepath.Join(currentDir, targetFile)
		_, err = os.Stat(targetPath)
		if err == nil {
			// 找到目标文件，返回从目标文件父目录到当前目录的路径
			relPath, err := filepath.Rel(filepath.Dir(currentDir), startDir)
			if err != nil {
				panic(err)
			}
			return relPath
		}

		// 如果当前目录是根目录，则停止查找
		if currentDir == "/" {
			panic("target folder not found")
		}

		// 移动到父目录
		currentDir = filepath.Dir(currentDir)
	}
}
