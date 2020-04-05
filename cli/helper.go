package cli

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"regexp"
	"text/template"
)

func createFileWithTemplate(
	directoryPath string,
	fileName string,
	tplStr string,
	data interface{},
) (err error) {
	file, err := os.Create(fmt.Sprintf("%s/%s", directoryPath, fileName))
	if err != nil {
		return
	}
	defer file.Close()

	tpl := template.Must(template.New(fileName).Parse(tplStr))
	err = tpl.Execute(file, data)
	if err != nil {
		return
	}

	return
}

func createDir(
	dirPath string,
) (err error) {
	_, err = os.Stat(dirPath)
	if err == nil {
		return
	}
	if !os.IsNotExist(err) {
		return
	}

	err = os.MkdirAll(dirPath, 0755)
	if err != nil {
		return
	}

	return
}

func appendToFile(filePath string, content string) (err error) {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	defer file.Close()

	fmt.Fprintln(file, content)
	return
}

func appendToFileWithTemplate(
	filePath string,
	tplStr string,
	data interface{},
) (err error) {
	content, err := execTemplateToString(tplStr, data)
	if err != nil {
		return
	}

	err = appendToFile(filePath, content)

	return
}

func execTemplateToString(
	tplStr string,
	data interface{},
) (result string, err error) {
	var buf bytes.Buffer

	tpl := template.Must(template.New(tplStr).Parse(tplStr))
	err = tpl.Execute(&buf, data)
	if err != nil {
		return
	}

	result = buf.String()
	return
}

func getPakcageName(dir string) (name string, err error) {
	// Get from go.mod file
	modFilePath := fmt.Sprintf("%s/go.mod", dir)

	if _, err = os.Stat(modFilePath); err == nil {
		var modFile *os.File
		modFile, err = os.Open(modFilePath)
		if err != nil {
			return
		}

		reader := bufio.NewReaderSize(modFile, 4096)
		r := regexp.MustCompile(`module (.*)`)

		for {
			var line []byte
			line, _, err = reader.ReadLine()
			if err == io.EOF {
				break
			}
			if err != nil {
				return
			}

			result := r.FindStringSubmatch(string(line))

			if len(result) > 0 {
				name = path.Base(result[1])
				break
			}
		}
	}

	return
}
