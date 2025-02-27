// Code generated by go-bindata. (@generated) DO NOT EDIT.

// Package jsonscheme generated by go-bindata.
// sources:
// ../../../../../../../../resources/component-descriptor-ocm-v3-schema.yaml
package jsonscheme

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _ResourcesComponentDescriptorOcmV3SchemaYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x1a\xef\x6f\xdb\xb8\xf5\xbb\xfe\x8a\x87\x4b\x01\x39\x4d\x64\x37\x29\x3a\xe0\xfc\x25\xc8\x7a\x18\x50\x6c\x77\x39\xb4\xdd\x3e\x2c\xcd\x0a\x5a\x7a\xb6\xd9\xa3\x48\x8f\xa4\x9c\xb8\xbd\xfe\xef\x03\x49\x51\xa2\x64\xc9\x3f\x93\x6e\xc3\x35\x5f\x62\x52\xef\x37\xdf\x2f\x3e\xe9\x19\xcd\xc6\x10\xcf\xb5\x5e\xa8\xf1\x68\x34\x23\x32\x43\x8e\x72\x98\x32\x51\x64\x23\x95\xce\x31\x27\x6a\x94\x8a\x7c\x21\x38\x72\x9d\x64\xa8\x52\x49\x17\x5a\xc8\x44\xa4\x79\xb2\x7c\x49\xd8\x62\x4e\x2e\xe2\xe8\x99\x83\x0d\x68\x7d\x52\x82\x27\x6e\x77\x28\xe4\x6c\x94\x49\x32\xd5\xa3\xcb\x17\x97\x2f\x92\x8b\xcb\x92\x74\x1c\x79\x82\x54\xf0\x31\xc4\x37\xaf\x7f\x86\xd7\x9e\x19\xfc\x54\x31\x83\xe5\x4b\xa8\x31\xa6\x94\x53\x83\xa0\xc6\x11\x40\x8e\x9a\x98\xff\x00\x7a\xb5\xc0\x31\xc4\x62\xf2\x09\x53\x1d\xdb\xad\x26\xf5\x4a\x0d\x58\xa2\x54\x54\x70\x8b\x9c\x11\x4d\x1c\xb4\xc4\x7f\x17\x54\x62\xe6\xc8\x01\x24\x10\x73\x92\x63\x5c\x2f\x4b\x3c\xb7\x43\xb2\xcc\x8a\x41\xd8\xaf\x52\x2c\x50\x6a\x8a\x6a\x0c\x53\xc2\x14\xda\xe7\x8b\x7a\xb7\xa4\x60\xa8\xf9\xdf\x00\xcf\x24\x4e\xc7\x10\x9f\x8c\x02\x8d\x6a\x53\xff\x12\x70\x2e\xd9\x6e\x41\x95\xc8\xc8\x03\x66\xef\x30\x5f\xa2\xf4\xa8\x8c\x4c\x90\xa9\x2d\x98\x0e\xc8\xa3\x2c\xa4\x58\xd2\x0c\xe5\x16\x24\x0f\xe6\xd1\x52\x89\xc4\x3c\x79\x4f\x43\x25\xdd\xa1\x28\x2d\x29\x9f\x55\x9b\x53\x21\x73\xa2\xc7\x90\x11\x8d\x89\xa6\x39\x46\x51\x53\xd2\xf2\x28\x89\x94\x64\xe5\xe8\x53\x8d\x79\xa5\x46\xbf\x12\xb1\x27\xd4\xeb\x12\x3b\x1c\x32\x61\x45\xb9\xde\x76\x84\x1d\xda\x59\xec\x31\x7c\xf9\xda\x77\x76\x0b\xa2\x35\x4a\xe3\x8f\xff\x5a\xde\xbe\x48\x7e\xbc\x3b\x7b\xe6\x99\x2b\x3a\xe3\x94\xcf\xda\xf4\xe3\x89\x10\x0c\xc9\x0e\x6e\x17\x01\x34\x1c\xa8\x61\x05\x27\xa6\x23\x92\x93\x87\xbf\x21\x9f\xe9\xf9\x18\x2e\x5f\xbd\x8a\x5a\x72\xdd\x92\xe4\xf3\xdd\x6d\x42\x92\xcf\x46\xbe\xe7\x83\xdb\xe1\x5d\x6b\xeb\xf4\xb9\xdf\xfb\x72\x79\xfe\x75\x30\x6a\x3c\xfe\xd8\x81\xf2\xd1\xe0\x9c\x1a\x55\x23\x00\x9a\x21\xd7\x54\xaf\xae\xb5\x96\x74\x52\x68\xfc\x2b\xae\x9c\xa8\x39\xe5\x95\x5c\x5d\x52\x19\xe6\x83\xdb\xe4\xe3\x99\x17\xc4\x6f\x9e\x5e\x39\xd2\x8d\x20\x70\x34\x4f\x40\x93\xdf\x90\xc3\x54\x8a\x1c\x94\x7d\x60\x12\x12\x10\x9e\x01\xc9\x3e\x15\x4a\x63\x06\x5a\x00\x61\x4c\xdc\x03\xe1\x20\x16\xce\xbe\xc0\x90\x64\x94\xcf\x20\x5e\xc6\xe7\x90\x93\x4f\x26\xeb\x71\xb6\x3a\xb7\xa8\x76\x3d\xcc\x29\x2f\x77\x3d\xaf\x39\x55\x90\x23\xe1\x0a\xf4\x1c\x61\x2a\x0c\x55\x43\xc4\x99\x5f\x01\x91\x68\x58\x19\x47\xa1\x59\x53\x5e\xe5\x05\xbe\x18\x5e\x0e\x5f\x86\xbf\x93\xa9\x10\x67\x13\x22\xcb\xbd\x65\x08\xb0\xec\x82\xb8\x18\x5e\xfa\x5f\x15\x58\x00\x5f\xfd\x6c\xa0\x85\xc6\x5e\xde\x5d\x0d\x5e\xfc\x7e\x7b\x91\xfc\x78\xf7\x21\x7b\x7e\x3a\xb8\x1a\x7f\x18\x86\x1b\xa7\x57\xdd\x5b\xc9\x60\x70\x35\xae\x37\x7f\xff\x90\xd9\x33\xba\x4e\xfe\x99\xdc\x19\x77\xf7\xbf\x3d\xc9\x1d\x81\x4f\x3d\xc7\xb3\x41\xf8\xe0\xcc\x12\x69\xec\x58\xc8\x32\xa4\x5a\x9e\xdf\xe5\x7a\xbd\x89\xa2\x8c\xfd\x95\x89\x23\x35\x86\x2f\xdd\x59\xa7\xcb\x95\x63\xf8\xea\x5c\x71\x21\x14\xd5\x42\xae\x5e\x0b\xae\xf1\x41\xef\x93\x93\x0c\x54\x5f\x0e\xb2\x14\xda\x39\x22\xd0\x51\xa4\xf4\x6d\x37\x6f\xc2\xd8\xcd\xb4\xe6\xd2\x53\x46\x5a\xa8\x75\x6a\x6c\xcb\x59\xca\x3a\x21\x0a\xff\x2e\x59\x5c\xa7\xb8\x35\x91\xcd\x5f\x09\x16\x6e\x75\xe6\xa6\xb2\x9c\x84\x79\xec\x67\xb2\x58\x34\x12\xe3\x46\x54\x00\xe4\x45\x3e\x86\xdb\xb8\x90\xec\x57\xa2\xe7\xf1\x39\xc4\x6a\x4e\x2e\x5f\xfd\x29\xc9\xe8\x0c\x95\x8e\xef\xa2\x16\x9d\x7d\x29\x5b\x1b\xcf\xa8\xd2\x72\x65\xa8\xdf\xbc\x7e\x53\x2d\xef\xcc\x19\x90\x34\x45\xa5\x76\x6c\x4c\x8c\x65\x2c\x94\xa9\x8c\x25\x2a\x2a\x18\x98\x15\x3e\x68\xe4\xa6\x82\xa8\xd3\x2d\xce\x12\x01\xcc\xa8\x9e\x17\x93\xeb\xcd\xbc\x37\x7a\x9b\x5d\x1a\x17\x08\x0e\xd4\xee\x4c\x0f\xf2\xc6\xb6\xd9\x9c\x80\x95\xf9\x4b\x46\x5b\xd0\x8d\x97\x6e\x86\x48\x45\x9e\x53\xbd\x29\x26\xb8\xe0\x78\x8c\x5d\x8e\xd4\xfb\x17\xc1\xd1\x39\x86\x12\x85\x4c\xf1\xa7\x2a\xe0\xf6\x10\xc7\xf4\x1e\xd5\xa2\xec\x2b\xaa\xb5\xa1\x50\x2d\x9c\x0b\xed\xd1\xc2\xac\x09\xbe\x7b\xb2\x2b\x51\xf0\x41\x4b\xf2\xa6\x04\xd8\xd2\x3a\xae\xd1\x79\x84\x46\x77\x87\xe3\x38\xa0\x17\x0e\xc3\xd8\xae\xf9\xea\x66\xda\x4c\x7f\x9d\x54\x1c\x5e\xbc\x1d\x30\x8c\xd8\x1d\xc0\xcd\xe5\xca\x03\x47\x00\x2e\x9b\xbd\x5b\x60\xba\x87\x1b\xcd\x89\x9a\x5f\xb3\x99\x90\x54\xcf\xf3\xda\xb9\x4c\x4f\xce\xa8\xb2\x3d\xfc\xfa\x63\xdb\xd6\x1e\x78\xed\x69\x30\xdc\xd8\x3c\x77\x0b\xb1\x43\xbf\xdd\x0d\x11\xb9\x96\x9a\xe8\x42\xe2\x9e\x46\x22\x1b\x2c\x60\x56\x39\x66\x94\xbc\xf7\x31\xb7\xbf\x4d\xc8\xd1\xca\xb9\xad\x4a\x8e\x1a\xaa\x59\x5b\xde\xcf\xd1\x01\xb9\x02\x23\xa6\xb6\x2d\xad\xcc\x02\xc1\x6d\x67\xa3\xfd\x0e\xcd\x53\xce\x45\xab\x65\x45\xef\x40\xbb\x6d\xbd\x7f\x39\x7e\x5b\x82\xbc\x8e\x9b\xf0\xea\x15\xe8\xd9\x8b\xd9\xf0\x27\x1b\x83\x4a\xa6\x6f\x7d\x81\xda\x5a\xe9\x89\x29\x66\x28\x91\xa7\x68\xaf\x1c\x30\xa8\x67\x2b\x4c\xa4\x84\x9d\x96\x05\xa2\xaf\xea\xf8\xd4\xf9\x0e\x19\xa6\x5a\x6c\xbb\xa4\xf7\x66\xda\xbd\x72\xa1\x6d\x66\x4b\xb1\x0f\x55\xb4\xd2\x73\xd7\x6b\x78\xe7\x24\xe4\xf8\x19\x4c\xc7\xfd\xb8\x57\xff\x4e\x11\x36\x95\x4f\x38\x01\x92\xea\x82\x30\xb6\x1a\xd7\x9c\x12\x1b\x79\xf7\x23\x50\x0b\x4c\x29\x61\x20\xd1\xc0\xa7\x96\xc9\xff\x6f\xc5\x3d\xa0\x9c\xb6\x83\x53\x70\x6c\x97\xd3\xd2\xa0\xbc\x60\x6c\x87\x7a\x18\x06\xb2\xf5\x52\x17\x3d\x75\x42\xdc\xb3\xf7\xf6\x04\xd4\xbe\x13\x41\x38\xb1\xf8\x36\x86\x6b\x2a\xe7\xe5\x38\xa0\x50\x1a\x72\xa2\xd3\x79\x10\x06\x6a\xad\x85\x5b\x6f\xc3\x99\x2d\x84\xc1\x56\xd8\x57\x7c\xef\xec\x2a\xad\x5c\x0e\x56\x6b\x50\xc1\x00\x11\xda\x43\xc4\x5e\x21\x1c\xb1\xfa\xf2\xe1\x0e\x61\xe7\x56\xdf\xba\x80\xb9\x13\x9a\x9b\x9b\xe4\x84\x55\xb7\x9d\xff\xc5\xfe\x53\xa4\xf4\xcf\x4c\xec\xde\x80\x5a\xed\xfe\x42\x19\xaa\x95\xd2\x98\xef\x8f\x7b\xd3\xc5\xf0\xa9\xf3\x82\x48\xe9\x9b\x9c\xcc\x8e\xba\x01\xda\x25\x35\x54\xde\xfa\xca\xf6\x28\x57\xc3\x70\x92\xe0\x3d\xa5\xc9\x66\xcb\xac\xa7\x36\xe7\x11\x8a\x31\xb2\xf2\x11\x77\x9c\x3e\x10\x97\x22\xc5\x50\xdf\xf2\xa7\x7d\xdd\xe9\xb5\x51\xa0\xd9\x2a\x98\xf6\x34\x27\x9c\x4e\x51\xe9\x76\x5f\xda\x62\x7a\x60\xf3\xeb\x2c\xe3\x52\xb3\x0b\x14\x27\x81\x02\x2d\xb6\x70\x6c\x3b\xea\x3a\x3b\x07\xe1\x59\x69\x22\x67\xa8\x31\x83\x54\x70\x5d\x35\x3f\xbd\xe4\x15\xfd\xbc\x51\x17\xf3\x1c\x28\x87\xc9\x4a\xa3\xf2\x3c\x26\xc6\xd8\x6d\xba\xbc\xc8\x27\xe6\x40\x23\x80\xde\x90\x3d\xc2\x5d\xa6\x94\x61\x5d\x09\x8f\xf5\x98\x0e\x09\x6b\xef\xf1\xac\xfa\xec\xe2\x9f\x87\xe6\x00\x3d\x27\x1a\xa8\xb2\xba\x1b\xf3\x53\x6e\x9f\xfd\x60\x1e\xaa\x1f\x20\xa3\xd2\x76\xcf\xab\xde\xf3\xf0\x76\xbb\x79\xa4\xf8\x7a\x02\x83\xdd\xb4\xe3\x6c\xb3\x73\x36\x1d\xd3\xc6\x3b\xdc\x53\x3d\x2f\x4d\x93\x16\x52\x22\xd7\x75\x83\x02\xf5\xbb\xde\x4d\x56\xf2\xa9\xf5\x6d\xd9\xf3\x1c\xf3\xe2\x2d\xec\xec\xbb\x8c\xf8\xbd\xfb\xd9\x5e\x4b\xec\x61\x3c\x66\xcb\xd1\xd7\x36\x04\x05\xf5\xdb\x94\xf1\x08\xa0\x1e\x7f\x1d\x11\x8a\x85\x9f\x6c\x1f\x59\xb8\x8d\x30\x95\xa1\x8b\x0d\x53\xec\x08\x60\x86\x1c\x25\x4d\xff\x8b\x13\xe8\x52\x02\x37\x84\x2e\x17\xdf\x3a\x66\x1f\x67\xdc\xf3\x07\x8b\xe9\xfa\xe0\xdc\xfe\x53\x85\x74\xc3\x45\xbf\x55\x63\xde\xfc\xd6\x64\x5f\x0f\x7c\x12\x7f\xda\x77\x32\xa6\x36\x0d\x96\x9b\x25\xd8\xce\x7f\xa6\x34\xb5\x17\x4a\x5f\x89\xcb\xce\xd0\x2c\x83\x29\x99\x77\x2f\x7d\xa8\xa6\xe5\x04\xe2\x91\xae\xc4\xad\x97\x56\xc1\x9b\x39\xd7\xb8\x3f\x12\x1f\xd9\xbc\x59\xd5\x03\x9d\xfd\xe9\xaf\xdd\x94\x37\xbc\xf0\xae\x87\x46\xf1\x2e\x08\xed\x96\x67\x27\xa4\x56\xca\x8d\xa3\xa8\xe5\x2e\xa1\xa7\x9b\xbc\xb9\xa0\xff\xa8\x73\x6b\x02\xf1\x6f\x94\x67\xe5\xcf\xf0\xb3\xb5\xc4\xb9\x55\x1c\x35\x5d\xa0\x46\x6f\xf8\x66\xe8\xea\xc1\x85\x2d\x1f\xb6\xbe\xfc\xab\x3e\xec\x3b\x77\x8f\x95\x98\xea\x7b\x22\xb1\x7e\x60\xbb\x4e\x23\x53\x2f\xfd\x54\x70\xa5\xc7\x10\x57\x1f\xf4\x05\xfa\x78\x0d\x1c\x72\xa7\xc1\x0c\x48\xdc\xf5\x19\xc5\x6e\xdf\x88\xb5\xce\xbf\xff\x28\xd7\x3e\x95\x88\xe1\xc4\x77\xc3\x6c\x75\x0e\xf7\x08\x82\xb3\x55\xf9\x79\x90\xbd\x34\x0a\x8e\x8d\xc0\xef\x8e\x99\xf2\xed\x42\xf5\xc6\xe0\x88\x6f\xdb\x2a\x1a\xf1\x7f\x02\x00\x00\xff\xff\x7f\xb7\x13\x95\xb0\x29\x00\x00")

func ResourcesComponentDescriptorOcmV3SchemaYamlBytes() ([]byte, error) {
	return bindataRead(
		_ResourcesComponentDescriptorOcmV3SchemaYaml,
		"../../../../../../../../resources/component-descriptor-ocm-v3-schema.yaml",
	)
}

func ResourcesComponentDescriptorOcmV3SchemaYaml() (*asset, error) {
	bytes, err := ResourcesComponentDescriptorOcmV3SchemaYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../../../../../../../../resources/component-descriptor-ocm-v3-schema.yaml", size: 10672, mode: os.FileMode(420), modTime: time.Unix(1675519154, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
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
	"../../../../../../../../resources/component-descriptor-ocm-v3-schema.yaml": ResourcesComponentDescriptorOcmV3SchemaYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("nonexistent") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"..": {nil, map[string]*bintree{
		"..": {nil, map[string]*bintree{
			"..": {nil, map[string]*bintree{
				"..": {nil, map[string]*bintree{
					"..": {nil, map[string]*bintree{
						"..": {nil, map[string]*bintree{
							"..": {nil, map[string]*bintree{
								"..": {nil, map[string]*bintree{
									"resources": {nil, map[string]*bintree{
										"component-descriptor-ocm-v3-schema.yaml": {ResourcesComponentDescriptorOcmV3SchemaYaml, map[string]*bintree{}},
									}},
								}},
							}},
						}},
					}},
				}},
			}},
		}},
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
