package toml

import (
	"errors"
	"fmt"
	"io"
	"os"

	gotoml "github.com/BurntSushi/toml"
)

type Reader struct {
	tomlFile io.Reader
}

func MustFromFile(filePath string) *Reader {
	reader, err := FromFile(filePath)
	if err != nil {
		panic(err)
	}

	return reader
}

func FromFile(filePath string) (*Reader, error) {
	var err error

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading TOML file %s: %w", filePath, errors.New("File not found"))
	}

	return FromIOReader(file), nil
}

func FromIOReader(reader io.Reader) *Reader {
	return &Reader{
		tomlFile: reader,
	}
}

func (reader *Reader) MustReadStrict(value interface{}) {
	if err := reader.ReadStrict(value); err != nil {
		panic(err)
	}
}

func (reader *Reader) ReadStrict(value interface{}) error {
	meta, err := gotoml.DecodeReader(reader.tomlFile, value)
	if err != nil {
		return fmt.Errorf("error decoding TOML: %v: %w", err, errors.New("file read error"))
	}
	if len(meta.Undecoded()) > 0 {
		return fmt.Errorf("unrecognized configuration in TOML: %q", meta.Undecoded())
	}

	return nil
}
func (reader *Reader) MustRead(value interface{}) {
	if err := reader.Read(value); err != nil {
		panic(err)
	}
}

func (reader *Reader) Read(value interface{}) error {
	if _, err := gotoml.DecodeReader(reader.tomlFile, value); err != nil {
		return fmt.Errorf("error decoding TOML: %v: %w", err, errors.New("file read error"))
	}

	return nil
}
