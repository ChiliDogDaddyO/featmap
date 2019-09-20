// Code generated by go-bindata.
// sources:
// tmpl/bindata.go
// tmpl/email.tmpl
// tmpl/invite.tmpl
// tmpl/reset.tmpl
// tmpl/welcome.tmpl
// DO NOT EDIT!

package tmpl

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

var _tmplBindataGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func tmplBindataGoBytes() ([]byte, error) {
	return bindataRead(
		_tmplBindataGo,
		"tmpl/bindata.go",
	)
}

func tmplBindataGo() (*asset, error) {
	bytes, err := tmplBindataGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/bindata.go", size: 0, mode: os.FileMode(511), modTime: time.Unix(1568980251, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplEmailTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x24\xcd\x41\xaa\xc2\x30\x10\x80\xe1\x7d\x20\x77\x98\xe5\x7b\x50\xd2\x33\xb8\x50\x84\xba\x10\xc5\x03\x8c\xcd\x18\x06\xda\x24\x24\xa9\x12\x86\xb9\xbb\x14\xb7\xff\xe2\xff\xce\x3c\x58\x63\xcd\x75\x21\xac\x04\x6f\x2a\xfc\xea\xd0\xd3\x56\x20\xd2\x07\x68\x45\x5e\x00\xbd\x2f\x54\x2b\xfc\x89\xb8\xe3\x5e\x54\xff\xe1\xd9\x21\x24\x8e\x01\x5a\x02\x11\x77\xc8\xf9\xce\x8d\x1e\xb7\x8b\x2a\xce\x73\xda\x62\x1b\x7f\xb7\x51\xc4\x4d\xd4\x55\x77\x67\xe2\xe8\xa1\x50\xc0\xe2\xeb\x60\xcd\x89\xb0\xad\x98\xbf\x01\x00\x00\xff\xff\xbe\x1c\x24\xb9\x85\x00\x00\x00")

func tmplEmailTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplEmailTmpl,
		"tmpl/email.tmpl",
	)
}

func tmplEmailTmpl() (*asset, error) {
	bytes, err := tmplEmailTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/email.tmpl", size: 133, mode: os.FileMode(511), modTime: time.Unix(1565338756, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplInviteTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\xb1\x4a\xc4\x40\x10\x86\xfb\x83\xbc\xc3\x78\x95\x42\xd8\xeb\xed\x54\x14\x0f\xc5\xe2\x44\xc4\x72\x92\xfc\x97\x5b\x4d\x66\xc2\x66\x73\x21\x2c\xfb\xee\xb2\x31\x01\x0b\x0b\xcb\x99\xd9\xff\x9b\x6f\xe7\xd1\xe6\xd9\x26\xdb\x7c\xe8\x40\x27\x3e\x83\x0a\x40\xc8\xca\xd9\x7a\x54\x54\x4c\x14\x82\xd9\xff\x54\xb7\x53\x8c\x74\xf9\xbb\xbe\x6f\xd9\x36\x31\x5e\x91\x57\xfa\x54\x2b\xe4\x4f\xa0\x51\xdd\x57\xdf\x71\x09\xda\x86\x60\xde\xd7\xea\x85\x5b\xc4\xb8\x25\xf6\xf4\x00\xf6\x2d\x77\x17\xeb\xde\x92\xe5\xaf\x78\x31\x51\xad\x56\xea\x44\x0f\xc1\xdc\x74\xdd\xab\xf5\x78\x3b\x3c\xc7\xc8\x65\xa9\x83\xf8\xdd\xec\xc9\xde\xaa\xec\x42\x30\x77\x5a\x21\xc6\x44\xdd\x1f\x69\xd2\x81\xd8\x81\x04\x63\x02\x2c\x3b\xf3\xb9\x2f\x40\x95\x9a\xa5\x03\x7b\x10\x0b\x2d\x40\x43\xab\xcf\x3a\x9a\xf3\xcb\xf4\x3f\x46\xbd\xad\x65\xe8\x0c\x1d\xd0\xa2\x2d\xe0\xd2\xd3\xa1\xc7\xfc\xb5\xa3\x36\x8d\x8e\x29\x8f\x74\xb8\xeb\xc4\x58\x4e\x68\x92\xf5\x93\x95\x8a\x1c\x6a\x76\x55\x9f\x67\x9b\x45\xf9\x3b\x00\x00\xff\xff\xf9\x7c\x0a\x64\xa1\x01\x00\x00")

func tmplInviteTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplInviteTmpl,
		"tmpl/invite.tmpl",
	)
}

func tmplInviteTmpl() (*asset, error) {
	bytes, err := tmplInviteTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/invite.tmpl", size: 417, mode: os.FileMode(511), modTime: time.Unix(1565255956, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplResetTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8e\x41\x6a\xc3\x40\x0c\x45\xf7\x06\xdf\xe1\x2f\x5b\x30\xce\x19\xba\x29\x2d\xe9\x2a\xa1\x07\x50\x62\xd9\x16\xd8\xa3\xe9\x48\x6e\x19\xcc\xdc\xbd\x8c\x09\x5d\x76\x2d\xfd\xf7\xde\x9b\x74\x6d\xd3\x36\x57\x5d\xf9\xa6\x43\xc6\xd3\xac\x91\xc7\x6d\x59\x32\xb2\x6e\xcf\x98\xc9\x90\xf8\x6b\x63\x73\x1e\x40\x88\x64\xf6\xa3\x69\x40\x62\x63\xc7\xa8\xa9\xfe\x25\xd0\xfd\xae\x5b\x70\x48\xc0\x2b\x93\xaf\x14\x7b\xbc\x8f\xf5\x86\x99\xbe\x19\x41\xfd\x1f\x4e\x87\xb8\x30\x19\x43\xa6\xa0\x89\xe1\xb3\x18\x78\x25\x59\xfa\x5a\x77\x39\x5c\x87\xe7\x6f\x77\xcb\x98\x54\xc2\x04\x57\xec\x7b\xff\x12\xe3\x55\x9c\x3f\x2f\x1f\xa5\x3c\x5a\x4e\x07\xfa\xb4\xef\xfd\x99\x73\x29\x15\x74\x96\x50\x8d\x13\xa5\xc1\xba\xb6\x79\x94\xfe\x06\x00\x00\xff\xff\xa0\xcd\x5c\x41\x04\x01\x00\x00")

func tmplResetTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplResetTmpl,
		"tmpl/reset.tmpl",
	)
}

func tmplResetTmpl() (*asset, error) {
	bytes, err := tmplResetTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/reset.tmpl", size: 260, mode: os.FileMode(511), modTime: time.Unix(1565255962, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplWelcomeTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\xce\x31\x4b\xc4\x40\x10\xc5\xf1\xda\x85\xfd\x0e\xaf\x53\xe1\xc8\xf5\x76\x16\xca\xc1\x59\x29\x2a\x96\x73\xd9\xb9\x30\x98\xec\x86\xd9\xc9\x1d\x61\xd8\xef\x2e\x11\x6c\x6c\x1f\x8f\x1f\xff\x83\xec\x62\x38\xb0\xf2\x6d\x45\x2d\x13\x43\xf2\xb9\xc0\x0a\x06\x36\xac\x65\x41\x35\x52\xe3\xf4\x10\x43\x0c\x5f\x65\x51\x64\xbe\xe2\x5a\xf4\xbb\xce\xd4\x33\xa4\x82\x2e\x24\x23\x9d\x46\x06\x19\xdc\xbb\xc7\x79\x7e\x13\xe3\xf7\xd7\x97\xd6\xdc\xbb\xcf\xbf\x6f\x6b\x31\xdc\xc4\xf0\xc1\x2a\xe7\x75\xb3\x15\x3c\x91\x8c\xa0\x94\x94\x6b\xc5\x9d\x7b\xf7\xb4\x2d\xad\xdd\xe3\xb4\x62\x28\x92\x87\xad\xe5\x1f\x4a\x7d\x5f\x96\x6c\xfb\xcb\xaf\xb4\x77\xef\x8e\xbc\x6e\x7a\x0c\x47\xc9\x09\xca\x03\x69\xaa\xbb\x18\x9e\x99\x6c\xa2\xf9\x27\x00\x00\xff\xff\xa0\xd2\x25\x86\xe5\x00\x00\x00")

func tmplWelcomeTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplWelcomeTmpl,
		"tmpl/welcome.tmpl",
	)
}

func tmplWelcomeTmpl() (*asset, error) {
	bytes, err := tmplWelcomeTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/welcome.tmpl", size: 229, mode: os.FileMode(511), modTime: time.Unix(1565255966, 0)}
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
	"tmpl/bindata.go":   tmplBindataGo,
	"tmpl/email.tmpl":   tmplEmailTmpl,
	"tmpl/invite.tmpl":  tmplInviteTmpl,
	"tmpl/reset.tmpl":   tmplResetTmpl,
	"tmpl/welcome.tmpl": tmplWelcomeTmpl,
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
	"tmpl": &bintree{nil, map[string]*bintree{
		"bindata.go":   &bintree{tmplBindataGo, map[string]*bintree{}},
		"email.tmpl":   &bintree{tmplEmailTmpl, map[string]*bintree{}},
		"invite.tmpl":  &bintree{tmplInviteTmpl, map[string]*bintree{}},
		"reset.tmpl":   &bintree{tmplResetTmpl, map[string]*bintree{}},
		"welcome.tmpl": &bintree{tmplWelcomeTmpl, map[string]*bintree{}},
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
