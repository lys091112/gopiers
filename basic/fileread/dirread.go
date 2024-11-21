package fileread

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

type targetFiles struct {
	FilePaths  []string
	regexList  []string
	ignoreDirs []string
}

func New(rootPath string, regexList, ignoreDirs []string) *targetFiles {
	return &targetFiles{
		FilePaths:  []string{},
		regexList:  regexList,
		ignoreDirs: ignoreDirs,
	}
}

func (tf *targetFiles) FindFile(path string) error {
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	separator := string(os.PathSeparator)
	for _, fileInfo := range fileInfos {
		// 当前文件名
		currentFileName := fileInfo.Name()
		currentFilePath := path + separator + currentFileName

		if fileInfo.IsDir() {
			// 如果是包含在忽略文件夹内
			if tf.isIgnoreDir(currentFileName) {
				continue
			}
			// 如果是文件夹，递归去处理
			err = tf.FindFile(currentFilePath)
			if err != nil {
				return err
			}
		} else {
			if tf.isMatchRegex(currentFileName) {
				tf.FilePaths = append(tf.FilePaths, currentFilePath)
			}
		}
	}

	return nil
}

func (tf *targetFiles) isIgnoreDir(dirName string) bool {
	for _, str := range tf.ignoreDirs {
		if str == dirName {
			return true
		}
	}
	return false
}

func (tf *targetFiles) isMatchRegex(str string) bool {
	for _, regex := range tf.regexList {
		isIn, err := regexp.MatchString(regex, str)
		if err != nil {
			log.Fatalf("egexp.MatchString err : %v", err)
		}
		if isIn {
			return true
		}
	}
	return false
}

// 获取当前可执行文件所在目录
func GetCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	//i := strings.LastIndex(path, "/")
	//if i < 0 {
	//  i = strings.LastIndex(path, "\\")
	//}
	i := strings.LastIndex(path, string(os.PathSeparator))
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\"`)
	}
	return path[0 : i+1], nil
}
