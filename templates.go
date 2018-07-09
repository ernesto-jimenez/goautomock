// Code generated by go-bindata. DO NOT EDIT.
// sources:
// templates/simple.go.tpl
// templates/testify.go.tpl
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

var _templatesSimpleGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x54\x4d\x6b\xdb\x40\x10\x3d\x6b\x7f\xc5\x20\x4c\x91\x1a\xa3\xdc\x03\x3e\x05\x0a\x85\x26\x04\x93\xb6\x87\x90\xc3\xb2\x1a\xd9\x8b\xb5\x2b\x21\xad\xa8\xdd\xe9\xfc\xf7\xb2\x1f\x91\xea\x34\x69\x29\x6d\x2f\x46\x3b\x7e\x3b\xf3\xe6\xbd\x99\xed\xa5\x3a\xc8\x1d\x02\x51\x75\x17\x3f\x99\x85\xd0\xa6\xef\x06\x07\x85\xc8\xf2\xc6\xb8\x5c\x64\xf9\x78\xb2\x2a\x17\x44\x83\xb4\x3b\x84\x55\x2f\xdd\x7e\x0d\x2b\x2b\x0d\xc2\xd5\x06\xaa\xf7\xe1\xc2\xc8\x2c\x32\xa2\x10\x66\x86\x9c\x28\x00\x99\x73\x22\xb4\x35\xb3\x28\x85\xb8\xbc\xf4\xc5\x6e\x23\xc4\x74\xea\x20\xdc\xa9\xc7\x1f\x62\xa3\x1b\x26\xe5\x80\x44\x66\x26\x07\xbe\x70\xb5\xfd\x7c\x33\x39\x3c\x8a\x4c\xc9\xb6\x1d\xc1\xc8\xfe\x61\x74\x83\xb6\xbb\x47\x6d\x9d\x2f\x19\x69\x55\x37\xe8\xf6\x5d\x9d\x68\xa4\x7c\xef\x26\xab\xa0\x99\xac\x2a\x66\xf6\xda\xd6\x78\x5c\xc3\x2a\x14\xf6\xf4\xef\xe4\x20\xcd\xfd\xa9\xc7\x91\x99\x48\x37\x09\xc2\xbc\x86\xc4\x9c\x28\xa0\xfd\x47\x38\x97\xf0\x8b\x74\x5b\x74\xd3\x60\xff\x2c\x9f\x67\x1c\x35\xe2\xa0\xd1\x2d\x7e\x49\x2d\xc0\x37\xb8\x96\xbd\x76\xb2\xd5\x5f\xbd\x3e\x6a\x40\xe9\x70\x5c\x14\x13\xbe\xbb\xd7\x2f\x14\x25\xbc\x5d\xd4\x25\x91\x0d\x81\x1e\xbc\x99\x83\x24\xb2\xa8\xec\x15\x18\x79\xc0\xe2\x5c\xdf\x72\x2d\xb2\xc0\x8a\x68\xb5\x43\x1b\x3a\x64\x16\x2f\x88\xfe\x93\xb5\x58\x83\x09\xff\x46\x8a\x85\xf1\x4c\x7c\x92\x04\x2a\x17\xfc\x5f\xb9\xd3\x13\x3d\xc5\xe0\xbf\x3b\xe5\x25\xac\xb1\xc1\x21\x4e\x55\x38\x67\xa6\x32\x93\xab\x3e\x74\xea\x50\x94\xe1\x18\xf4\x7c\xc8\xe7\xfe\xf2\xc7\x8b\x8b\x19\xf7\xd1\xb6\x09\xc9\xfe\x47\x37\x60\xaa\xf3\x81\xdd\x6c\xc0\xea\x36\xa4\xee\xa5\xd5\xaa\xc8\x27\x8b\xc7\x1e\x95\xc3\x1a\x7c\x6e\x70\xdd\xb9\xc4\x8b\x94\xb9\xcf\xeb\xe7\x49\x37\xcf\x3a\x4c\xce\xa7\x4e\x9e\x17\xfd\x57\x16\xcc\xc2\xf9\xfa\x9f\xe4\xa0\x65\xad\x15\x73\x55\x55\xf3\xb0\xf3\xf9\x3b\x70\xdf\x39\xd9\x5e\x87\xdd\x8e\x14\x47\x70\x7b\x04\x69\xba\xc9\x3a\xe8\x1a\x88\x7b\xef\xba\x10\x7e\xa5\xed\xdf\xcf\xd8\x52\xa6\x28\x41\xdb\xf8\xc6\x04\x43\xb6\x4f\xce\x45\x63\x53\x70\xb1\x29\x09\xf7\x92\xad\xc2\xaf\x42\xdc\xdc\xef\x01\x00\x00\xff\xff\xbe\x03\x25\x3d\x4d\x05\x00\x00")

func templatesSimpleGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesSimpleGoTpl,
		"templates/simple.go.tpl",
	)
}

func templatesSimpleGoTpl() (*asset, error) {
	bytes, err := templatesSimpleGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/simple.go.tpl", size: 1357, mode: os.FileMode(420), modTime: time.Unix(1531135330, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesTestifyGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x53\x41\x8b\xdb\x3c\x10\x3d\x4b\xbf\x62\x30\xf9\x3e\xec\x12\x94\x7b\xa0\xa7\x3d\x94\x1e\x76\x59\xda\xd2\xbb\x2a\x8f\x13\xb1\xb6\x62\xa4\xf1\xee\xa6\x53\xfd\xf7\x32\x8a\xe3\x6c\x49\xbb\x14\x5a\xe8\x4d\xcc\xbc\xf7\x66\xde\x93\x34\x5a\xf7\x60\x77\x08\xcc\xe6\xfe\x74\xcc\x59\x6b\x3f\x8c\x87\x48\x50\x6b\x55\x75\x03\x55\x5a\x0d\x07\xf7\x00\xd5\xce\xd3\x7e\xfa\x62\xdc\x61\xd8\x24\x8a\x48\x6e\x1f\x37\x84\x89\x7c\x77\xdc\x08\xa2\xd2\xcc\xd1\x86\x1d\xc2\x6a\xb4\xb4\x5f\xc3\x2a\xd8\x01\x61\xfb\x16\xcc\xfb\xa2\x98\x72\xd6\x8a\xb9\x94\x73\x86\x8a\xb9\x00\x73\xae\x98\x31\xb4\x39\xeb\x46\xeb\xcd\x46\xb6\xb9\x3b\x41\x44\x56\xd3\x71\xc4\x17\xb5\x44\x71\x72\x04\x7c\x5a\xcb\xdc\x0a\x24\x17\xde\x1d\x3e\xcd\x30\xf8\x06\x37\x76\xf4\x64\x7b\xff\x55\x38\x2e\xa2\x25\x4c\x17\x15\xdd\x4d\xc1\xfd\x9a\x50\x37\xf0\xe6\x32\x91\xb5\x8a\x48\x53\x0c\xf0\xff\x52\xe4\x2c\x43\x99\x57\x3b\x0c\xc5\x62\xce\x8b\x7d\x73\x8b\xb4\x3f\xb4\x62\xf7\xca\x0d\xb6\x30\x94\xee\x69\x83\x7a\x90\x41\x22\x32\x83\x9a\x0b\xbe\x5e\xe2\xf4\xa1\xc5\xe7\x35\xac\x4a\x12\x32\xec\xde\x46\x3b\x7c\x3a\x8e\x98\x72\x66\xf6\xdd\x0c\xc9\x79\x0d\x73\x94\x23\xf3\xb9\x06\xcc\x85\x29\xd0\xd2\x6b\xe0\x15\xe9\x0f\xc5\xe9\xab\xda\xd7\x7a\xac\x0b\xf2\x47\x72\x49\x4d\x34\x07\x73\x63\xfb\x1e\xdb\xbf\x65\x68\x59\x43\x46\x7e\xb6\xd1\xdb\xd6\xbb\x9c\x8d\x31\xe7\x85\xe4\x99\xfd\x96\x41\xad\x1e\x6d\x84\xf8\xb3\xb0\xb4\x4a\x4f\x9e\xdc\x1e\x22\x26\x21\x46\x24\xf3\x0e\xa9\xbe\x60\x1b\x53\x0b\x54\xec\x2b\x67\x13\x42\xf0\xfd\x76\x3e\x2e\x32\x5b\xad\xd4\x4b\x7d\x11\x4a\x5a\xb5\xd8\xd9\xa9\x27\xe9\x8e\x36\x78\x57\x77\x03\x99\x8f\x63\xf4\x81\xba\xba\x9a\x02\x3e\x8f\xe8\x08\x5b\x10\x95\x2d\xfc\xf7\x58\xad\x85\xd8\x34\x5a\x95\x4f\x74\xfa\x2f\xe7\x77\xf9\x47\xd7\x19\xaf\x92\xd5\xcc\xd8\xa7\x92\xc1\x3f\xb8\xbb\xf3\x12\x79\x39\x7d\x0f\x00\x00\xff\xff\xcf\x68\x72\x7c\xa7\x04\x00\x00")

func templatesTestifyGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesTestifyGoTpl,
		"templates/testify.go.tpl",
	)
}

func templatesTestifyGoTpl() (*asset, error) {
	bytes, err := templatesTestifyGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/testify.go.tpl", size: 1191, mode: os.FileMode(420), modTime: time.Unix(1531135338, 0)}
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
	"templates/simple.go.tpl": templatesSimpleGoTpl,
	"templates/testify.go.tpl": templatesTestifyGoTpl,
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
	"templates": &bintree{nil, map[string]*bintree{
		"simple.go.tpl": &bintree{templatesSimpleGoTpl, map[string]*bintree{}},
		"testify.go.tpl": &bintree{templatesTestifyGoTpl, map[string]*bintree{}},
	}},
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

