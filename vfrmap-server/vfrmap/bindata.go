// Code generated by go-bindata.
// sources:
// html/index.html
// DO NOT EDIT!

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5c\xdd\x93\xdb\xb8\x91\x7f\x4f\x55\xfe\x07\x2c\x52\xce\x68\x62\x12\xe2\x87\xa8\x91\x34\xd2\x24\x5e\xef\x47\xf6\xce\x5e\xbb\x6c\x6f\xf6\xee\x6c\x97\x0a\x22\x21\x8a\x31\x48\xf0\x08\x68\x34\x63\xd7\xfc\xef\x57\x00\x49\x09\xe0\x87\xa4\xd9\xdd\x4b\x5e\xd6\x0f\x1e\x09\xe8\x6e\x34\x7e\xdd\x68\x34\x1a\xa4\xe6\x5f\x7d\xf3\xea\xf9\xbb\xff\x7e\xfd\x2d\xd8\x88\x94\xde\xfc\xf1\x0f\xf3\xea\x2f\x00\xf3\x0d\xc1\x91\xfa\x04\xc0\x5c\x24\x82\x92\x9b\x94\xaf\xb9\xe7\x78\x8e\x1d\xb3\xe1\xed\xba\x48\x71\x3e\x1f\x96\x3d\x15\x19\x4d\xb2\x4f\xa0\x20\x74\x01\x93\x90\x65\x10\x88\xfb\x9c\x2c\x60\x92\xe2\x98\x0c\xf3\x2c\x86\x60\x53\x90\xf5\x02\x46\x58\xe0\xd9\xbe\xf5\x7a\x85\x39\x19\x8f\xac\xe4\x1f\x5f\xbf\x7a\xb3\x73\xfe\xf3\xfb\x98\x2d\x60\x5b\x22\x17\xf7\x94\xf0\x0d\x21\xa2\x16\x43\x09\x5e\x53\x22\xfe\xc9\x87\xd5\x27\x14\x72\x0e\xc1\xb0\xe6\x4d\x89\xc0\x20\xc3\x29\x59\x40\x9c\xe7\x94\xd8\x29\x5b\x25\x94\xd8\x3b\xb2\xb2\x71\x9e\xdb\x21\xce\xf1\x8a\x12\x08\x42\x96\x09\x92\x89\x05\xbc\x27\x1c\x9e\xcb\xcd\x05\x16\x5b\x6e\xaf\x70\x61\x2b\xd5\x34\x31\x2b\x8a\xc3\x4f\x67\x0b\x52\x08\x6a\xdc\xff\xf8\xee\xcd\x4b\x9c\xef\xd9\x79\x58\x24\xb9\x00\xbc\x08\xbb\x66\xfc\x4f\x0e\x6f\xe6\xc3\x92\xe6\x3c\x8e\x82\x09\x2c\x48\xf4\x12\x17\x9f\x48\xd1\xcd\x2f\xa7\x53\xd9\x4e\x90\x3b\x31\x94\xb0\x56\x9d\x40\x39\x8a\x05\x56\x2c\xba\x07\x5f\xea\x36\x00\x36\x24\x89\x37\x62\x06\x5c\xc7\x79\x72\x7d\x68\x4e\x71\x11\x27\xd9\x0c\x38\x5a\x5b\x8e\xa3\x28\xc9\x62\xb3\x71\xcd\x32\xc9\xed\xe5\x77\x43\x17\x05\x00\xfe\x9d\xd0\x5b\x22\x92\x10\x83\x1f\xc9\x96\x40\x0b\x3c\x2b\x12\x4c\x2d\xb0\x6f\xb7\x00\xc7\x19\xb7\x39\x29\x92\x75\x2d\xe6\xe1\x8f\x7f\xa8\x3f\x36\xf5\x8b\x12\x9e\x53\x7c\x3f\x03\x6b\x4a\xee\xf4\x71\x29\xb9\xb3\xa3\xa4\x20\xa1\x48\x58\x36\x03\x21\xa3\xdb\x34\xbb\xee\x90\xf8\xa7\x14\xe7\xba\xc4\x15\x0e\x3f\xc5\x05\xdb\x66\x91\x1d\x32\xca\x8a\x19\x88\x0b\x7c\xdf\x14\x1d\x17\x6c\x37\x03\x6e\xa7\x40\x41\x28\xc9\x59\x21\xec\x9c\xe5\xdb\x1c\x18\xd2\x6b\x90\x1c\xe4\x91\xf4\x04\xa0\xd2\x46\x36\xa6\x49\x2c\xf5\x27\x99\x20\xc5\x79\xc3\x51\xbc\x22\xb4\x63\x50\xbb\x28\x6d\x69\x8e\xad\x89\x42\x95\x2f\xd9\xd2\x67\x0b\x46\x6d\x8a\xef\x49\xc1\x75\x51\xc3\xbf\x88\x02\x67\x7c\xcd\x8a\x74\x06\x78\x88\x29\x19\x78\x97\x40\xb5\x51\x2c\xc8\xc0\xf6\x82\x27\x16\xf0\x82\x27\x97\xd7\x7f\x19\x6a\x53\x69\x32\xf9\x06\x93\xef\x3c\xb1\x80\xef\x3c\xb9\xd4\xb5\x2a\x3f\xf4\xea\xf6\x99\xb1\xf4\x31\x9a\xfd\x12\xc5\x3a\xf5\xda\x03\xbf\xd9\x46\xb5\x36\x06\x46\x39\xe3\x49\xe9\x75\x78\xc5\x19\xdd\x0a\xa2\x19\x75\xc5\x84\x60\xe9\x0c\xb8\x86\xf9\x2b\xcb\x98\x8d\x9f\xed\x24\x8b\xc8\x9d\x5a\x7c\xba\x5f\xb4\x3d\x74\xb7\x49\x1a\x83\x14\x11\x29\x66\xc0\xcd\xef\x00\x67\x34\x89\x80\x0a\x5d\x9d\x8b\x15\x05\xc6\xa0\xe7\x19\xd8\x7e\xb4\x85\x7d\xc9\x75\xdc\xc4\x49\x96\x6f\xc5\x7b\x15\x9f\xc2\x0d\x09\x3f\xad\xd8\xdd\x47\xab\xa3\xb3\xc0\x51\xc2\x3e\x76\x46\x81\x8c\x65\xa4\x3d\xc0\xf1\x21\xc0\x53\xb5\x60\xba\x46\xd2\x68\x78\x8e\xb3\x23\xca\xf4\xcb\xa8\x09\xa4\x80\x6e\x2f\x29\x08\xc5\x22\xb9\x25\xad\x70\x60\x77\xba\x45\xbd\x9a\x29\x59\x8b\x19\xf0\x90\x67\x1a\x70\x1f\xb1\x8d\x56\x9a\x64\xc4\xee\xee\xda\x72\x52\xd8\x9c\x50\x12\x8a\x36\x80\xad\xe5\xa7\x82\x8c\x5d\x6d\x6a\xfa\x7c\xd2\x24\xb3\x77\x49\x24\x36\x33\xe0\x3b\x4e\xae\x07\xe3\xaa\x79\x9b\x71\x22\xc0\x57\x49\x2a\xa3\x15\xce\xc4\x23\xc6\x31\xd1\xed\x21\x5a\x6d\x85\x60\xd9\x29\x2a\xcd\x34\x6a\x17\x84\x86\x27\xc9\xfd\xca\xe6\xc9\x67\xa2\x66\xf1\xa4\x4b\xc5\x23\x3e\x34\x9b\xad\xc8\x9a\x15\xe4\x84\x2f\x55\xb4\x78\x2d\x48\x71\x86\xdb\x9d\x29\xb5\x24\xed\x15\x6a\x38\xea\x31\x91\x0d\xc2\x93\x02\x4f\xa9\x68\xd2\x29\x71\x3a\xe2\x95\x5d\x66\xe0\xe2\xe2\xfa\xcc\x18\xaa\xe2\x4a\xdd\x4d\x29\x70\x90\xcb\x01\xc1\xfc\x71\x21\xb0\xa4\xb0\xa5\x7a\x5b\xde\x5c\x14\xa7\x62\xec\x6f\xea\x0f\xe7\x22\x78\xb6\xe9\x74\x81\x9d\xd9\x9c\x31\xd7\x6a\x79\x9a\x09\x49\x19\x5d\x8c\x74\x84\xe5\x7a\xc3\xb9\x08\x9c\xef\xe5\xe7\xba\xee\x99\x1e\xd9\xf2\xb4\xb6\x4d\x9b\x2e\x51\xe3\xe3\xa0\x71\x17\x42\xcd\xe6\x0a\xa3\x46\x2a\x57\xe2\xd4\x97\x63\x29\x75\x67\x6a\xee\x24\x7a\x8c\xd7\xf4\xf2\xf4\x7b\x8f\xc9\x72\x86\x17\x75\x33\xf4\x79\x53\xb5\x80\xf6\x58\x6e\xc9\xd1\x05\x64\x10\xf4\xed\xcd\x27\x66\xd9\x61\xfa\x53\x58\x9e\x64\xe9\xf0\x9a\x47\xc0\xd8\xf2\x32\x2d\x11\xda\xe7\x3f\xff\x35\x70\x49\x7a\xd9\x8e\x3d\x7d\xe9\xdb\xe9\xe0\x23\xff\xcc\x87\xea\x44\x77\x53\x3b\xd7\xdc\x38\xee\x01\x00\xb7\x9c\x00\x2e\x8a\x24\x14\xf0\xfa\xe0\x81\x94\x08\x90\xe2\xfc\xda\xfc\x2e\x4f\x8d\x1d\x4d\xef\xaa\xc3\x45\x47\xd7\x0f\x21\xcb\x8c\x66\xb5\xcf\x1a\x2d\x3b\x6e\x12\x50\x9c\x91\x65\x9b\xac\x3e\xc1\x74\x74\xad\x19\xa5\x6c\xb7\x54\x9c\x60\x01\xd6\x98\x6a\x51\x5e\x12\x50\xcc\xc5\xb2\x50\xec\x60\x01\xbe\x3c\xec\x3b\x43\x96\x71\xb9\xdb\x27\x22\xc1\x74\x99\x33\x0e\x16\xe0\x05\xa2\x58\xbc\xc8\xe2\x41\xe0\xa1\xd1\x95\x77\x65\xd9\x2e\xba\x0a\x3c\xff\xd2\x10\xc9\x6f\xe3\xd7\x72\x3c\x39\xc1\xb7\xa2\x48\xb2\x18\x2c\xc0\xc5\xfc\xaf\x77\x29\x05\xb7\xa4\xe0\x09\xcb\x16\xd0\x45\x0e\x04\x24\x0b\x99\xcc\xc5\x16\xf0\xa7\x77\xdf\xd9\x13\x08\xb8\xc0\x59\x84\x29\xcb\xc8\x02\x66\x0c\xfe\xf5\x66\xce\x6f\x63\x70\x97\xd2\x8c\x2f\xe0\x46\x88\x7c\x36\x1c\xee\x76\x3b\xb4\xf3\x11\x2b\xe2\xa1\xe7\x38\xce\x90\xdf\xc6\xb0\x8a\x3c\x0b\xe8\x8d\xa6\x68\x32\x82\x65\xc4\x91\x5f\x27\xc8\x0b\xa0\x39\xec\x8d\x2a\x3d\x44\x58\x60\x90\x44\x0b\x58\x7f\x99\xc2\xe1\xcd\x3c\xc7\x62\xa3\x5a\xe5\x87\xe0\xca\x1b\x41\x10\x2d\xe0\x4b\xe0\x8d\xae\x50\xe0\x8e\x9c\x91\xe5\xca\xb9\x3b\xde\x78\x0c\x5c\x7f\x8a\x9c\xe0\x6a\xe2\x5a\x57\x2e\x9a\x38\xce\x74\x34\x06\x21\x70\xd0\xc4\xf1\xc6\x13\xcb\x76\x3d\x34\x0a\xdc\xc9\x28\x00\x2e\xf2\xbd\xd1\x95\x6f\xd9\x23\x07\x79\xc1\x58\xf2\x3a\x68\x12\x8c\x24\xd5\x28\x40\x23\xf7\x2a\x98\x4e\x81\xed\xa3\xe9\xc8\xf1\x47\x96\x3d\xf2\x91\x37\x1e\x8f\xc6\x1e\xb0\x7d\x17\x79\xbe\xe3\x4e\x2c\xdb\x1b\xa1\xb1\xef\xb8\xee\xd4\x57\xad\xa3\x89\xef\x07\x96\x1d\x20\xdf\x73\xfc\xf1\x15\xb0\x1d\xe4\x8c\xa7\xbe\x15\x20\x6f\xe2\xfa\x63\x17\xd8\x2e\x72\xdc\xc0\xf1\x2c\xdf\x43\xc1\x74\xe2\x4f\x26\xb2\xc9\x75\x46\x57\xae\x15\x38\x68\xe2\x8f\xc7\x9e\x07\x5e\xc8\x48\x3b\x19\x79\x57\xee\x95\xe5\x06\x23\xe4\x5f\x05\x63\x0f\x38\x96\x3b\x71\x90\x3b\x0d\xae\x02\x40\x81\xeb\x3a\x28\x70\x9c\x60\x62\xd9\x81\x83\x46\x13\xcf\x9f\x02\x1f\x4d\xa7\xbe\xef\x59\x13\x07\x79\x53\x77\x2c\x75\xf2\x90\x33\xf2\x82\xf1\x95\xe5\x79\x68\xea\x4f\xdc\xb1\xd4\xc9\x73\xfc\xc9\x28\xb0\xdc\x31\x9a\x4c\xc7\x53\x1f\x8c\x3c\x24\xc7\xba\xf2\x2c\xdb\x75\x51\x30\x0d\x14\x16\x8e\x33\xb1\x1c\xe4\xfa\xd3\x40\x12\x5c\xb9\xbe\xeb\x5a\xae\x83\xa6\xee\xe4\x6a\x2a\xa5\x04\xce\xd4\x9b\x5a\xb6\x94\x32\xf1\xdc\x72\xb0\x51\xe0\x5f\x8d\x2c\xdb\xf3\x90\x3f\x9d\x3a\x3e\xf0\xd0\xd8\x75\x7d\xcf\xb2\x27\x0e\xf2\x03\xcf\x09\x80\xeb\xba\xc8\x0f\xa6\xd3\xc0\x1a\x4d\x50\xe0\x8c\x5d\x57\xca\xba\xf2\x47\x53\xc9\x17\xa0\xab\x2b\x6f\x1a\x80\xcf\x10\xac\x13\x4a\xed\x62\x4b\xc9\x02\x92\x5b\x92\xb1\x28\x2a\xdb\x16\x70\xb9\x7c\xfe\xea\xc5\xab\x37\xcb\xa5\xf4\x0d\xe9\x69\x37\x17\x66\x20\xc8\x6b\x47\xff\x5a\x6e\x85\x6a\x89\x24\x21\xcb\x06\x5a\x40\x93\xdf\x7f\x2a\xe8\xac\xf4\x76\xf2\xd3\x9b\x1f\x06\x7a\xa5\x8f\xdf\xc6\x4f\xef\x52\x6a\x41\xf0\xb4\x63\xe5\x5c\xa2\x82\xe4\x14\x87\x64\x00\xff\x04\x2d\xf8\xc4\xf3\xa1\xd6\x74\x50\xcf\x02\x55\x75\xed\xd2\x32\x47\x7e\xab\x72\xf1\xf7\xe3\x91\x05\xc6\xa3\xc3\x71\xf0\xe1\xb2\x1d\x59\xe4\xa8\x3f\xcb\x40\xf9\x6f\x9d\x86\x0a\xd5\xbf\x76\x1a\xdf\x17\x84\x64\xff\xd6\x69\xc4\x52\x83\x5f\x30\x8d\x32\xec\xa6\x38\x5f\x16\x44\x25\xee\x09\xcb\x64\xe8\xd5\xb3\xd0\x24\xde\xcc\xf4\x06\x00\x44\x42\xc9\xb2\x3c\x76\x79\xc1\xd8\xd2\xbb\x3e\x33\x96\x2e\xd9\x7a\xcd\x89\xcc\xb5\x0e\x3d\x0f\x1a\x55\x4a\xa2\x64\x9b\xf6\xcb\x0c\x5c\xaf\x5f\xa6\xed\x76\x0b\xa5\x6c\xd7\x2f\xd1\x75\xbc\xd1\x11\x91\x9e\x26\x72\x0f\x52\x63\x0f\xd5\x11\x02\x8b\x26\x64\xa8\x9c\x93\xb6\x6f\xaf\xb7\x99\x2a\x66\x02\x96\x93\x6c\x99\x64\xcb\x98\xb1\x98\x92\x65\x8a\x73\x3e\xb8\xd4\x35\xbd\xc5\x05\xd8\x16\x14\x2c\x80\xda\x73\x78\xb5\xe9\x94\xf4\x28\x64\xe9\x50\xf2\x0c\xff\x26\x9d\x44\xdb\x3f\xe5\xe6\x98\x88\x6d\x44\xc0\x53\x00\xad\x56\x27\xcb\x62\xb3\x37\xc5\x39\x8a\x89\xf8\x1f\xc6\xd2\xc1\xa5\x6c\xfd\x0c\xf5\x94\x39\x8b\xd8\x0e\x49\x5d\x07\xdb\x82\x5a\x17\xcb\x15\xc5\xd9\xa7\x8b\xce\x52\xda\x7e\x6a\xdb\x3c\xc2\x82\xbc\xc4\xf9\x20\xe5\xf1\xa5\x89\xbe\x9c\x55\x63\x1b\x4f\x79\xbc\x57\xda\x02\xea\x5b\xad\xa5\x9e\x6d\x81\x2a\x63\x41\x9c\x88\x17\x25\x67\xce\x78\x1f\xc5\x1b\x26\xb0\x54\xe6\x59\x16\x53\xa2\x86\xd8\x10\x2c\x77\xf9\x4b\xcd\x18\xf2\x9f\x96\xce\xa0\x9c\x71\x94\x64\x99\x4c\x98\xee\x64\x22\xf2\xe3\x36\x5d\x91\x42\x0e\x23\x15\xbc\x44\x82\x7d\x97\xdc\x91\x68\x30\xbe\xdc\xc3\xa7\x93\xc8\xa5\x79\x20\x69\x8c\x93\xac\xc1\x40\xcf\x83\x1a\xb8\x00\x65\x88\x1c\x67\xef\x58\x7b\x5a\x0f\x1d\x68\xef\x24\x86\x19\xd9\x81\x9f\xc9\xea\x2d\x0b\x3f\x11\x31\x80\x3b\xe9\x24\x52\xab\xca\x6e\x94\x85\x0a\x04\xb4\x61\x5c\x64\x38\x55\x56\x9f\x75\x11\xa8\xcc\xeb\x29\x80\xc3\x1d\x87\x87\xb1\x77\x1c\xb1\x4c\x1a\x5f\x26\x6d\x95\x75\x4d\x27\x1d\x0e\x65\xa8\x60\x94\x20\xca\x62\xa9\x80\xf2\x6b\x4d\xc4\x83\x29\x2c\xa4\x8c\x93\x47\x48\x53\xf4\xfd\xe2\x52\xc2\x39\x8e\x0d\x81\xa4\xb5\x88\x52\x2e\x53\xbf\xff\x78\xfb\xea\x47\x94\xe3\x82\x93\x01\x41\x32\xd6\xea\x10\x9b\xf9\x67\xca\x63\xc3\x78\xd2\x74\x29\xce\xc1\x57\x8b\x05\xd8\x66\x11\x59\x27\x19\x89\x1a\xf6\x33\x7d\xfe\xba\x27\x72\xb4\x96\x8a\xcc\x6c\x25\x53\x4b\xe9\x72\x8d\x68\x89\xaf\xa1\x91\xa4\x60\x3c\xad\x3c\xe0\x05\x7a\x97\x50\xf2\x02\xdf\x93\x62\x50\xe7\xa7\x5f\xf8\x03\x92\x91\x4e\xad\x5d\x2e\x0a\x42\x84\x74\x30\x99\xaf\x7e\xf9\xfc\x30\xfc\x72\xf7\x30\xfc\x72\xff\x80\xf2\x2c\x86\x96\x39\x93\x14\xdf\xc9\x60\x30\x03\xee\xc4\x88\x8d\x69\x92\x95\xed\x66\x14\x96\x83\x94\x1b\x8a\x19\xfb\xd0\x3e\xce\xb6\x22\xec\xab\x2a\xc0\x36\x18\xb4\xe0\x6b\xb0\xc8\x63\x18\x16\x33\xa0\x5d\x55\x1a\xfd\x7c\xbb\x8a\x58\x8a\x93\x8c\xcf\xc0\x7b\x88\x55\x16\x22\xff\x0b\xe1\x47\xcd\x0c\xba\x4d\x14\x7c\x39\xc9\x70\x92\x2f\x43\x1c\x6e\x48\xb4\x5c\x61\x4e\xa4\x8d\xcf\x40\x54\x86\x5d\x54\xb1\xa3\x8c\x88\x61\x4c\xd8\x8e\xac\x94\xa0\x21\x27\xc5\x6d\x12\x92\xa1\x48\xf9\xd0\x45\x0e\x72\x86\xf5\x40\xd5\x08\x7f\xfb\xf6\xf5\xdb\xef\x9f\xf8\xcf\xa6\x8e\x33\x75\xfd\xbf\xe5\xd9\x23\x0c\x32\xea\x36\xc8\xe8\x5f\x6c\x10\x91\x72\x79\x2c\xde\x9a\x82\x86\xc3\x88\x08\x12\x8a\x37\x44\x24\x19\xee\x20\xd0\xed\x04\x5d\x0f\x3e\xca\xc8\xea\x14\x9e\xe3\x42\x15\x19\xa5\xe8\x23\xa6\xe5\x02\xa7\x24\x5b\xaa\x54\x74\xb9\xab\x72\xc9\x93\x66\x2d\xb9\xd4\xc6\x2a\x58\x46\x8a\xdf\xd7\x89\x0e\xa6\x20\x45\x81\x93\xec\xb1\x40\x96\x5c\xbf\x43\xa9\x43\xb9\xc3\x82\x14\x8f\x04\x52\xf1\xa8\x32\xd2\xef\x58\x2a\x2c\x43\x5c\x08\xb6\x8c\x70\xf1\xa9\x0f\x49\x99\x0b\x29\xaa\x68\x65\x57\xb1\x97\xdb\x12\xdb\x98\xb2\x15\xa6\x88\x73\x8a\xd6\x98\x0b\x7a\xaf\x82\xb8\x14\xb5\xc4\x94\xfe\x8e\x2f\xd0\x8e\x64\xda\x8e\x28\xd3\x14\x98\xe2\xbc\x09\x48\xf9\x78\xc1\x0c\xbc\x97\x19\x89\xd5\xb7\xaf\x7e\x34\x74\x2a\x1f\x84\x98\xc9\x2c\xa7\x35\x7d\x79\x2c\x33\x1a\xb1\x10\x45\xb2\x52\x20\x3c\x2f\x2f\xea\x67\x65\x09\xb1\xc5\x6a\x76\x37\x66\x67\xfa\x4f\x29\x54\x1d\x43\xaa\xdb\x7f\xa4\x8d\x33\xf8\x72\xb8\xba\x82\xe5\x65\x3f\x25\x6b\x01\x0d\x98\x4a\x7a\x84\xa3\xe8\x99\xc6\x09\xe7\xb8\x7c\x02\xea\x83\x71\x6c\x6b\xa7\x61\x21\xcb\xef\xd5\xd5\xf0\x07\x08\x04\x2e\x62\x22\x16\x1f\x60\x79\xbe\xfa\x00\x81\x2a\x0a\x2f\x3e\xc0\x0f\xf0\xe6\x55\x4e\xe4\xc9\x9f\x10\x99\x29\xce\x87\xf8\x06\xfe\x2a\x2d\xaa\xac\xe5\xd4\xa8\x92\xf4\xd9\x0f\xaf\x7f\xd9\x78\xb3\xa1\x3a\x9c\x6a\x11\xec\xd4\x70\x6f\x15\xe5\x2f\x9f\x9d\x5a\xe9\x2a\x54\x9e\x1a\xe9\xb9\xa4\xfc\xe5\x03\xc5\x89\xd8\x6c\x57\x6a\x24\x9a\xe0\x6c\xa8\x3d\x76\x77\x6a\x64\x49\xff\x9b\x0c\xbc\xba\x4f\xf1\x5d\x92\x6e\xf9\x61\x74\xc1\x18\x5d\xe1\xc2\x2e\x0f\x76\xb6\x20\x69\x4e\xb1\x20\xa7\x54\xda\x4b\xfa\x4d\xf4\x7a\xbe\x29\x12\x2e\x12\x9c\xb9\xd3\xc9\xe8\x31\xc8\x18\x8c\x2d\x55\xba\x94\x7a\xc7\xe4\x91\xac\xbd\xae\x65\xbc\x79\x89\xf3\x46\xb1\x4a\xfe\x83\xc6\x3a\x82\x33\x15\xae\x1a\x24\xa5\x17\x82\x77\x65\xda\x02\x67\x8d\xec\xa7\x8f\x5c\xa6\x8b\x07\x62\x2d\xef\xec\x61\xf8\x59\xee\xe6\x07\x06\xb5\xb9\x37\x49\x95\x9b\x82\x6f\xe4\x06\x37\xf8\x51\x86\x09\xf0\x92\x45\xe4\x12\xce\xb4\xad\x4f\xe3\xd9\x1f\x8c\x1b\xa7\x9c\x5b\x52\x50\x7c\xdf\x83\xc8\x8f\xf8\x36\x89\x55\x01\x00\x53\xf0\x0d\x16\x58\xa2\xd2\x19\xbf\xbb\x46\x02\x40\x0b\x9e\xe5\x2e\x30\xa8\xe1\xb7\xf4\x91\x2f\xfb\x0c\x56\x96\x6b\x54\x10\x2e\x3f\x0e\xe4\x86\x60\xaa\x99\x84\x32\x08\x9b\x15\x6e\x03\xab\x42\xaf\xf4\xcc\x80\xd3\xd9\xf9\xaa\x48\xd4\x93\x79\xb0\xdc\x79\xf4\x1d\xd2\x88\xe9\x55\x01\xc9\xd0\xb7\xd1\xb7\x4a\xb2\xe8\x35\xcb\xb7\xf9\xe0\x05\x52\x65\xa3\xc1\x17\xbc\x15\xec\x35\xce\xaa\x7d\xe7\xe1\xb2\x51\xa2\x92\x5f\x9f\x97\x0f\x48\x0c\xf4\x72\x93\xdc\x94\x2f\xdb\x2e\x5c\x8e\xf3\xda\x2c\x92\x39\x96\xd3\xd6\xa5\xbe\xd7\xd3\x11\xdc\x73\x5b\xe0\x4b\xc7\xd4\x6a\x96\x63\x53\xdc\xd3\x9c\x39\xd5\x7a\x6e\xe6\x0d\xe0\x61\x7a\xb5\x74\x4e\xc4\x72\x4f\xd3\x54\xb7\xc4\xc1\x48\x3d\x10\xcb\x06\x17\x51\x81\x63\x2e\x70\x21\x2e\xac\xbe\x22\x4f\x29\xb9\x2c\xae\x0d\x94\x5e\x46\x01\xa6\xe9\x73\x07\xb9\x24\x8b\x8e\x49\x2d\x6b\x3a\x6f\x05\x2b\x48\xf4\x5c\xf9\xcd\xe0\x0c\xc1\x32\x11\xf9\x75\x82\x01\xa0\x2c\xc4\x54\xf6\xe3\x98\x48\x84\x7f\x10\x24\x1d\xc0\x6c\x29\x65\x43\xcb\xac\xd9\x9e\xa1\x52\x48\x93\xf0\xd3\x29\xfc\x9a\x96\x21\xd2\xf5\x68\x16\x9f\x21\x5f\xae\x7a\x15\x00\xc2\x0d\xce\x62\x72\x6c\xa4\x64\x0d\x06\x04\xa9\x5a\xe4\x62\xd1\x1f\xe5\x3a\xaa\xa3\x75\x65\x57\x46\x81\x81\x79\x55\xd4\xa8\x96\x02\x42\x39\x69\x8e\x64\xc4\xea\xb3\xc5\xab\x2b\x9c\x4e\xf1\x67\x0a\x50\xf1\xaa\xa7\x9a\x6b\xc4\x9e\xae\x42\x7a\x97\x59\x2a\xa3\x18\x0f\xed\x99\xab\xf6\x10\x7c\xda\x06\x6c\x2c\xd1\x38\xe7\xe8\x16\xd3\x2d\x01\x0b\x50\x12\x4b\xa3\xef\x8b\xd8\x93\x4b\xed\x12\xa1\xec\xcd\x62\xad\xf7\xda\xac\x91\xea\x17\x0d\x98\x96\x05\xfd\x06\xd2\x8d\xf1\x6b\x2a\x4d\x89\xb6\x88\xce\x62\x6a\x07\x5a\x7b\xd9\x1b\x52\x90\x76\x45\x55\x2b\x03\x73\x75\x8f\x96\xac\xef\x07\xba\x6e\xcd\xdd\x51\xdc\xe7\x04\xce\x00\xac\xe5\xc2\xe6\x16\x4d\xb1\x80\x33\xa0\x6a\xca\xdf\x51\x86\x5b\x01\x70\x8f\x2e\xe2\x39\x4d\xc4\x00\x5a\xf0\xf2\xbd\xf3\xf1\xb2\x25\x27\x8b\x1f\x2f\xc7\x6d\xcb\xa9\x01\x3b\x2a\xcc\x84\x5c\x1a\xd8\x41\x81\xd5\xe3\x9f\xba\x85\x77\x1c\x71\x92\x45\x66\x7d\xbb\xcf\x6b\xeb\x60\xac\xfe\x18\xa6\x38\xdc\x83\x34\x5c\x43\xdf\x19\x4b\x02\xe3\x2e\x06\x46\x2c\xbb\xa8\x1f\x25\x29\x89\xa1\xee\x19\x1d\xab\xf2\x94\xc4\x5e\x59\x7a\x8c\x6b\x3e\xbb\xa2\xbe\xea\x57\x07\x5d\xa1\xfa\x62\x55\x21\x20\x43\x61\x39\xd9\xa3\x90\x09\x16\xc7\x94\xd4\xa8\x19\xc0\x68\x68\x7e\x65\xdc\x1f\x9d\xb4\x01\x4e\x8a\x12\x81\x32\x3e\x2c\x6f\x13\x9e\xac\x12\x9a\x88\xfb\x81\xfa\x48\x49\xcb\x30\x5d\xed\x7a\x6c\x7b\x95\xe3\x50\xf2\xbb\x9d\x37\x1c\xa0\xc3\x06\x6d\x5e\xc7\x8c\x86\x47\xb7\x6e\xdd\x10\xdd\x7b\xe2\x6a\xc9\x37\x6c\xb7\x9f\x2b\xb4\x40\x3d\x87\x73\x00\xdf\x43\xa4\x61\x43\x5a\xf0\x1f\x01\x92\xa0\xea\xb1\x32\x6d\x38\xd0\x31\x5e\x99\x76\xee\x25\x99\x26\xa6\x44\x80\x70\x05\x16\x20\x62\xe1\x36\x25\x99\x40\xff\xbb\x25\xc5\xfd\x5b\xf5\x34\x37\x2b\x06\xd0\x78\x49\xc1\x96\x13\xb6\xf7\x13\x6e\x06\xe1\x70\xd5\xb0\x5e\xb8\xaa\x75\x04\x0b\x55\x39\xef\x03\xf8\xc4\x54\x25\x6b\x33\x9b\xab\x0c\x57\x77\x99\x53\x3a\x7e\x85\xd5\xd8\x30\xc0\x9f\xff\xdc\x7d\x8b\xdd\x6c\xdf\x5f\x0d\x37\xd6\xba\x99\x2e\x77\x89\xb2\x7a\x04\xf5\xe1\xd1\x73\x21\xdb\x7f\xe3\x6d\x26\x74\x0d\x2b\xa8\x27\x2a\x4a\x37\x28\x9f\x13\x90\x19\xdc\xa3\x92\x3f\xa5\xbd\xdc\x73\xac\x4a\x8e\xba\x94\x7e\x04\x2b\x53\xd5\xcc\x9a\x37\x8b\x8f\xbc\xc5\xb1\x9f\x19\x65\x38\x2a\xe7\xf5\x56\x60\xd1\x70\xdc\x72\x52\xc6\x02\x94\x7b\xb8\xae\x44\x7c\x88\x88\x06\xdd\x45\xd3\x6f\x4d\x29\xc6\x55\xab\xf4\x81\x76\x77\xb6\xa5\xb4\x23\x8d\x3d\xe2\xc1\x0d\x45\x17\x00\x4a\xc7\xad\x56\x50\x87\xb1\x7e\x8b\x25\xd9\xb3\x28\x1b\xcb\xb2\x47\xb3\xde\xb4\x51\x77\xd3\x52\xd9\x6a\x2f\xeb\x47\xbf\xda\x8f\x9a\xb0\x57\x7c\x2d\xbc\xb5\xf6\x0e\xa0\x41\x7b\x93\x3f\x68\xdd\x36\x57\x03\xea\xd6\x7c\xda\xdb\x46\x47\x70\x39\x32\xf9\xbd\x7b\xf7\xce\x5f\x5b\x02\x06\x04\x3a\x3f\x16\x27\xd9\xb1\x68\x01\x68\x6c\xca\x86\x17\xd5\x91\x4b\x29\xd6\x02\xd8\xec\x92\x18\x77\x33\x63\xd1\xc7\x5b\xf5\x74\x9a\xc7\x78\x8e\x42\x05\xb1\x19\xb4\xf6\x7c\xd6\x61\xf4\xe6\xf3\x33\x75\xc8\x33\x23\x69\x83\xe7\x84\x39\xd4\x4b\x7d\xfd\x50\xca\xee\x16\x8c\x8a\xa7\x35\xd1\x7d\x6b\xc7\x24\xa5\xae\xbc\x3a\x04\x4b\xba\x13\x4a\x65\xf8\x76\xa9\x9e\xb6\xed\x5f\x22\x35\x49\x4b\xb9\x3d\x6f\x4b\x41\xa3\xa7\x43\x49\x73\xec\xe5\xd1\x90\xd2\xf3\xe2\x66\xf5\x56\x17\x2b\xca\x27\xd7\x2f\xea\x87\xe3\x2f\x3e\x76\x04\x1a\x6d\xa4\x96\x4f\x68\x7d\x48\x55\x04\x06\xfd\xe7\xd2\x36\x7e\x38\x14\xc9\xad\x7a\x4e\xed\x88\x69\x0f\x44\x1d\xab\x6c\x3f\x7e\xb1\xe2\xbd\x30\x3c\xa3\xf4\x6c\x24\xd4\x43\xfc\x0d\x18\x24\x08\x9a\xaa\x2d\x83\x35\xfa\xd4\xba\xd3\xec\x28\x75\x43\x94\x64\xb1\xd8\x80\x1b\xe0\x34\x38\xe6\x5d\x74\x0d\x9c\x75\x8a\xf7\x07\xde\x8f\x1d\x90\x1f\x3d\xcd\x16\x24\x4e\xb8\x20\xc5\xdf\x71\x16\x51\x52\xf0\xae\x9d\xf7\x5f\xe1\x57\xc7\xbd\x4a\xf7\x29\x1c\x45\xdf\xde\x92\x4c\xbc\x90\x7a\x67\xa4\x18\xc0\xb2\x1c\x04\x2d\x20\x93\xea\xc5\x4d\xd3\x21\xfb\x32\xfa\x5a\x28\xb4\x4c\x9f\xad\x32\x6d\xc3\x67\xcf\x5c\xf5\xff\xaf\x3e\xb7\x66\x05\x18\xc8\x8c\x37\x01\x0b\xe0\x5c\x83\xa4\xdb\x51\xae\x41\xf2\xf4\xe9\x31\x6f\x49\x3e\xfe\x56\x18\xea\x2b\x11\x5a\x20\x39\x8a\x59\x0b\xbc\x3d\x4a\x6d\x6d\xbe\x79\xf5\xb2\xaa\xf6\xbe\x60\x38\x22\x11\xd4\x4b\x7d\x92\xd4\x98\x9f\x76\x00\x6f\xde\x3e\xa4\x38\xc9\x66\x87\x91\x62\x22\xbe\xa5\x44\x7e\xfc\xfa\xfe\x87\x68\x00\x15\x67\xf9\xce\x27\x34\xeb\x1c\x39\xe3\xe7\xf1\xd9\x39\xe3\x0d\xde\x38\xc5\xf9\x99\xcc\x92\xb4\xc1\x5d\x6e\xf4\x67\xf2\x97\xc4\x86\x84\x87\xde\x3a\xdc\x23\xc1\x31\x7f\x2b\xa0\xa1\x25\xdf\xae\xd2\x44\x9c\xcd\x6d\x97\xf4\x4d\xa0\xf2\x63\x20\x37\x24\xc4\x79\x13\xe7\xba\xc6\x74\xbe\x8c\x7d\xe9\xaa\x09\x98\x16\x88\xea\x47\x25\x8d\xe2\x4b\xe3\x74\xa2\xf5\xb5\x03\xe8\xe1\xb8\x53\x7f\x34\x7e\x5c\x63\x3e\xdc\xff\xa8\xca\x7c\xc5\xa2\xfb\xfa\x17\x37\xa2\xe4\xb6\x7c\x45\x07\xe7\xf0\x66\x3e\x8c\x92\xdb\xc3\xcb\x5a\xb2\xaf\xbc\xc7\x84\xf5\xeb\xeb\xea\xe5\xeb\xc3\x2f\x72\xec\xd9\x75\x97\xbe\x39\xa8\x39\xcf\x6f\xe6\x1b\xbf\x49\xa1\x9c\xf7\x66\x3e\xdc\xf8\x37\xf3\x61\xde\x20\x2f\x5f\x91\x6e\xb1\x28\x97\x05\x2c\x53\x1b\xcd\x02\x76\x3e\x60\x7e\x5d\x5e\xec\x83\x24\x03\x65\xbb\x4c\xa3\xf8\x7c\x58\x8a\x3c\x7f\xac\xca\xbd\x0f\xa3\x35\x2a\x58\xd7\x52\xf9\x0e\xa1\x26\x7c\x1a\x3a\x0d\x9f\x6e\x68\x51\xfe\x18\xc6\x9a\x15\x4d\x42\xe5\x7c\x37\xdf\xbf\x7e\x3b\x9b\x0f\x15\xd1\xcd\x5c\xbd\xed\xa7\xfd\x40\x0a\xec\x90\x5f\xb2\xb5\x67\xdb\x3f\x0e\xa6\x02\xde\x3c\xab\xfd\xfa\x71\x83\xed\x9d\xbb\x1f\xdf\x52\x42\xf9\xa5\x53\x46\xb5\x4c\x35\xc0\xcd\x6a\xf7\x35\xbc\xa9\x5b\x8e\x02\xdf\x36\xc2\xde\x04\xfa\x89\x56\xf3\x5f\x9e\xe3\xcc\x9c\x68\x9d\x29\xc0\x16\x5b\xe3\x20\x2c\xb5\x55\xdb\xd7\xde\x3f\xba\x0a\x6e\x62\x93\xf0\xcb\x6b\x08\xea\xe3\xf0\xd0\xb0\xc3\x11\xf1\x37\x6f\x37\x6c\x07\x9e\x55\x5f\xf7\x36\x19\x2a\x8d\x1b\xfa\xf7\xe2\x6c\xc8\x2f\xdd\xd7\x9c\x40\x09\x77\xab\x80\x77\x0d\x6f\xca\xea\x8d\x54\x20\x2c\xf0\x5a\xc7\x5d\xd3\xe0\x00\xfc\x7c\x58\x45\x95\xf9\xb0\xfa\x1d\xa7\xff\x0b\x00\x00\xff\xff\x0d\xa8\x9e\x3d\xe1\x49\x00\x00")

func indexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_indexHtml,
		"index.html",
	)
}

func indexHtml() (*asset, error) {
	bytes, err := indexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.html", size: 18913, mode: os.FileMode(438), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"index.html": indexHtml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"index.html": &bintree{indexHtml, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

