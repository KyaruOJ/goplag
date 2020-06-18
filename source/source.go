package source

import (
	"goplag/token"
	"io/ioutil"
	"path"
)

// Source 源文件
type Source struct {
	Path string
	Ext  string
	Code string
	Tokens []token.Token
	Fingerprints map[int]string
}

// Open 打开源文件
func Open(p string) (*Source, error) {
	content, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, err
	}

	return &Source{
		Path: p,
		Ext:  path.Ext(p),
		Code: string(content),
	}, nil
}

// Files 获取源文件列表
func Files(files []string) ([]*Source, error) {
	fileSrcs := []*Source{}

	for _, file := range files {
		src, err := Open(file)
		if err != nil {
			return nil, err
		}
		fileSrcs = append(fileSrcs, src)
	}

	return fileSrcs, nil
}
