package source

import (
	"io/ioutil"
	"path"
)

// Source 源文件
type Source struct {
	Ext  string
	Code string
}

// Open 打开源文件
func Open(p string) (*Source, error) {
	content, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, err
	}

	return &Source{
		Ext:  path.Ext(p),
		Code: string(content),
	}, nil
}
