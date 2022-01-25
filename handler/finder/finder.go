package finder

import (
	"bufio"
	"context"
	"fmt"
	diction "diction-car/pb/diction"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type LineInfo struct {
	lineNo int
	line   string
}

type FindInfo struct {
	filename string
	lines    []LineInfo
}

type Serviceserver struct {
	diction.UnimplementedFinderServer
}

func (s *Serviceserver) FindFile(ctx context.Context, in *diction.FileRequest) (*diction.FileReply, error) {

	var result []string

	startTime := time.Now()

	kind := in.GetKind()
	Word := in.GetWord()
	Filename := in.GetFilename()

	if kind == "normal" {
		result = Normal(Word, Filename)
	} else if kind == "goroutine" {
		result = Goroutine(Word, Filename)
	} else {
		result = append(result, "normal 과 goroutine 중 골라 주세요")
	}

	return &diction.FileReply{Message: result, Time: time.Now().Sub(startTime).String()}, nil
}

func GetFileList(path string) ([]string, error) {
	return filepath.Glob(path)
}

func FindWordInAllFiles(word, path string) []FindInfo {
	findInfos := []FindInfo{}

	filelist, err := GetFileList(path)
	if err != nil {
		fmt.Println("파일을 찾을수 없습니다")
		return findInfos
	}

	for _, filename := range filelist {
		findInfos = append(findInfos, FindWordInFile(word, filename))
	}
	return findInfos
}

func FindWordInFile(word, filename string) FindInfo {
	findInfo := FindInfo{filename, []LineInfo{}}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다")
		return findInfo
	}
	defer file.Close()

	lineNo := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, word) {
			findInfo.lines = append(findInfo.lines, LineInfo{lineNo, line})
		}
		lineNo++

	}
	return findInfo
}

func Normal(word string, files []string) (result []string) {
	findInfos := []FindInfo{}
	for _, path := range files {
		findInfos = append(findInfos, FindWordInAllFiles(word, path)...)
	}
	if len(findInfos) != 0 {
		for _, findInfo := range findInfos {
			result = append(result, findInfo.filename)
			result = append(result, "-----------------------------")
			for _, lineInfo := range findInfo.lines {
				result = append(result, lineInfo.line)
			}
			result = append(result, "------------------------------")
			result = append(result, "")
		}
	}
	return
}

func RoutineFindWordInAllFiles(word, path string) []FindInfo {
	findInfos := []FindInfo{}

	filelist, err := filepath.Glob(path)
	if err != nil {
		fmt.Println("파일을 찾을수 없습니다")
		return findInfos
	}

	ch := make(chan FindInfo)
	cnt := len(filelist)
	recvCnt := 0

	for _, filename := range filelist {
		go RoutineFindWordInFile(word, filename, ch)
	}

	for findInfo := range ch {
		findInfos = append(findInfos, findInfo)
		recvCnt++
		if recvCnt == cnt {
			break
		}
	}
	return findInfos
}

func RoutineFindWordInFile(word, filename string, ch chan FindInfo) {
	findInfo := FindInfo{filename, []LineInfo{}}
	file, err := os.Open(filename)
	if err != nil {
		ch <- findInfo
		return
	}
	defer file.Close()

	lineNo := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, word) {
			findInfo.lines = append(findInfo.lines, LineInfo{lineNo, line})
		}
		lineNo++
	}
	ch <- findInfo
}

func Goroutine(word string, files []string) (result []string) {

	var findInfos []FindInfo

	for _, path := range files {
		findInfos = append(findInfos, RoutineFindWordInAllFiles(word, path)...)
	}

	if len(findInfos) != 0 {
		for _, findInfo := range findInfos {
			result = append(result, findInfo.filename)
			result = append(result, "-----------------------------")
			for _, lineInfo := range findInfo.lines {
				result = append(result, lineInfo.line)
			}
			result = append(result, "------------------------------")
			result = append(result, "")
		}
	}
	return
}
