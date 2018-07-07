package util

import (
	"html/template"
	"os"
)

type File struct {
	Tmplate  string //*template.Template
	FilePath string //*os.File
	File     *FileDetail
}
type FileDetail struct {
	Title   string
	Content string
}

func GetTextGenerater() *File {
	return &File{}
}

// Excute テンプレートからファイルを作成する
func (f *File) Excute() error {
	tmp, _ := template.ParseFiles(f.Tmplate)
	file, err := os.Create(f.FilePath)
	defer file.Close()
	if err != nil {
		return err
	}
	err = tmp.Execute(file, f.File)
	return err
}

func (f *File) SetTmplate(dir string) {
	f.Tmplate = dir
}

func (f *File) SetFilePath(dir string) {
	f.FilePath = dir
}

func (f *File) SetContent(content string) {
	f.File.Content = content
}

func (f *File) SetTitle(title string) {
	f.File.Title = title
}
