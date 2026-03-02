package main

import (
	"os"
	"path/filepath"
	"encoding/binary"
	"io"
)

// Goal: Build a page manager that reads and writes fixed-size pages
// to a single binary file

const PageSize = 4096 // organize data into fixed-size pages 4KB as OS pages

type Page [PageSize]byte

type PageManager struct {
	file *os.File // the binary file to be created
	numPages uint32
}

// Opens or creates the db file
func PageManagerCreator(dbPath string) (*PageManager, error) {
	// Ensures that the containing dir is existing
	err := os.MkdirAll(filepath.Dir(dbPath), 0755)
	if err != nil {
		return nil, err
	}

	file, err := os.OpenFile(dbPath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}

	pm := &PageManager{file:file}

    var header [4]byte
    if _, err := file.ReadAt(header[:], 0); err != nil && err != io.EOF {
        return nil, err
    }
    pm.numPages = binary.BigEndian.Uint32(header[:])

    return pm, nil
}

