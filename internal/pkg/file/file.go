package file

import "io/ioutil"

import "github.com/golang_example/internal/pkg/log"

type (
	File interface {
		GetBufFromFile() []byte
	}

	Config struct {
		Path string
	}
)

func (conf Config) GetBufFromFile() []byte {
	buf, err := ioutil.ReadFile(conf.Path)

	if err != nil {
		log.Fatal(err.Error())
	}
	return buf
}

func (conf Config) GetStringFromFile() string {
	buf, err := ioutil.ReadFile(conf.Path)

	if err != nil {
		log.Fatal(err.Error())
	}
	return string(buf)
}
