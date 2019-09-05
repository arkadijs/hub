// Code generated by go-bindata. DO NOT EDIT.
//  memcopy: true
//  compress: true
//  decompress: once
//  metadata: true
//  asset-dir: true
//  restore: true
// sources:
//  meta/hub-well-known-parameters.yaml
//  src/hub/api/requests/aks-adapter-instance.json.template
//  src/hub/api/requests/aks-adapter-template.json.template
//  src/hub/api/requests/eks-adapter-instance.json.template
//  src/hub/api/requests/eks-adapter-template.json.template
//  src/hub/api/requests/gke-adapter-instance.json.template
//  src/hub/api/requests/gke-adapter-template.json.template
//  src/hub/api/requests/hybrid-adapter-instance.json.template
//  src/hub/api/requests/hybrid-adapter-template.json.template
//  src/hub/api/requests/k8s-aws-adapter-instance.json.template
//  src/hub/api/requests/k8s-aws-adapter-template.json.template
//  src/hub/api/requests/metal-adapter-instance.json.template
//  src/hub/api/requests/metal-adapter-template.json.template
//  src/hub/api/requests/openshift-adapter-instance.json.template
//  src/hub/api/requests/openshift-adapter-template.json.template
//  src/hub/initialize/hub.yaml.template
//  src/hub/initialize/hub-component.yaml.template

package bindata

import (
	"bytes"
	"compress/flate"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/tmthrgd/go-bindata/restore"
)

type asset struct {
	name string
	data string
	size int64
	mode os.FileMode
	time time.Time

	once  sync.Once
	bytes []byte
	err   error
}

func (a *asset) Name() string {
	return a.name
}

func (a *asset) Size() int64 {
	return a.size
}

func (a *asset) Mode() os.FileMode {
	return a.mode
}

func (a *asset) ModTime() time.Time {
	return a.time
}

func (*asset) IsDir() bool {
	return false
}

func (*asset) Sys() interface{} {
	return nil
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]*asset{
	"meta/hub-well-known-parameters.yaml": &asset{
		name: "hub-well-known-parameters.yaml",
		data: "" +
			"\x8c\x93\x4d\x6e\xeb\x30\x0c\x84\xf7\x3e\x05\x2f\xe0\x1c\xc0\xbb\xf7\x9a\x45\x81\x00\x46\x51\x77" +
			"\xd5\x1d\x25\x31\xad\x1a\xfd\x04\x94\x94\xc2\xb7\x2f\x1c\x4b\xb1\xe2\x24\x40\x36\x06\x2c\xce\x37" +
			"\x34\xe9\x51\xdb\xb6\xcd\x11\x19\x2d\x45\xe2\xd0\x35\x2d\x38\xb4\xd4\x81\x34\x3e\xa9\x06\x40\xb0" +
			"\xa6\x7d\x07\x2f\xf9\xf5\xa0\x9d\xea\x20\x05\xe2\x06\xa0\xe6\x00\x0a\x39\x49\x1a\x80\x6b\x14\x76" +
			"\xe5\x54\xd1\x1e\x93\x89\x1d\xe0\x6f\xa8\xa8\x23\xfb\x93\x56\x67\xdb\x15\xf9\x56\x57\x1e\xd0\x28" +
			"\xa5\x4f\x2e\xde\xc2\xff\xe6\x02\xf4\xc9\x0a\x62\xf0\x3c\xd9\xfd\x90\x8c\x15\xcc\xf4\xa5\xbd\xbb" +
			"\x65\xdf\xcb\xf9\xa5\xcb\x09\xb5\x41\xa1\x8d\x8e\xe3\xa7\x77\x74\xa7\x5d\xa5\x80\x2c\x29\x74\x08" +
			"\xdf\x3b\x1a\x6b\x66\x18\x5e\x61\x47\x23\xf4\x68\xe9\xb2\x78\xe5\xc2\xb2\xf6\x6d\x3f\x3c\xb7\xf4" +
			"\xe9\x59\x5b\x6f\xfb\x61\xb6\xbd\x5a\x5a\xa4\x50\x0f\x2e\x30\xd0\xd6\x5b\xd4\x6e\x8d\xfe\xc7\x40" +
			"\x50\x95\x2e\x0e\x21\xa2\x3c\x84\x8d\x22\xa3\x4f\xc4\x63\x65\xa6\xee\x1a\x65\x8f\x22\x8a\xc4\x8c" +
			"\x7b\xcf\x76\x23\x92\x3c\x50\x5c\x26\xfd\x28\x95\xa9\x45\x24\x30\x5e\x62\x9c\xd7\xff\xdc\xc4\x6b" +
			"\x83\xb9\x41\x91\x3d\xfe\xd7\x0f\xb8\x2c\x6c\x44\x72\xca\x50\x75\x31\xe6\xe0\xad\xbf\xe9\x7c\x5d" +
			"\x36\x4b\x0c\xcb\x09\x2f\x19\xca\x92\xdb\x0c\x95\x52\x0e\xc8\x5f\x00\x00\x00\xff\xff",
		size: 914,
		mode: 0664,
		time: time.Unix(1564685966, 41485166),
	},
	"src/hub/api/requests/aks-adapter-instance.json.template": &asset{
		name: "aks-adapter-instance.json.template",
		data: "" +
			"\x9c\x8e\xc1\x4a\xf4\x30\x14\x85\xf7\x7d\x8a\x10\xfe\xe5\xfc\x79\x80\x01\xd7\xee\x44\x04\x57\x32" +
			"\x8b\x3b\xed\xb1\x94\x26\x37\xe1\xde\x9b\x41\x1c\xfa\xee\x32\x6d\x05\x85\x82\xd6\x5d\x38\x5f\xce" +
			"\x77\xee\xb5\x71\xce\x39\xcf\x94\xe0\x8f\xce\xff\xbb\xde\x5e\x93\x3f\x2c\xb1\x21\x95\x48\x36\xa3" +
			"\xcf\x0c\x7c\x19\x24\x73\x02\xdb\xd7\xd8\xa8\x57\x7f\x74\x2f\x9e\x3a\x2a\x06\xb9\xa3\x51\xfd\x69" +
			"\x85\x85\x84\x12\x0c\x32\x7f\x59\x36\xbf\xed\x76\xac\xa1\xcb\x89\x06\x5e\x7d\x33\xbd\x50\xac\xcb" +
			"\xf6\x9c\x4d\x07\xb7\x51\x6d\x63\xae\x5d\x10\xf4\x43\xfe\x6b\x99\xde\xab\xe0\x09\x9a\xab\xb4\xb8" +
			"\x97\x5c\xcb\xc3\x8d\xef\x95\x8d\xf5\x0c\x61\x18\x34\xd0\xa8\xa1\x8d\x55\x0d\xb2\xff\xa6\x9c\x4a" +
			"\x66\xb0\x85\x81\x7b\x81\x6a\xa8\x12\x1f\x05\xaf\xc3\xdb\xa6\x8b\x4a\xd9\xa7\x53\xcd\xcf\x3f\x19" +
			"\xf5\x77\xca\x16\x62\xff\x13\x31\xf5\x90\x00\xa6\x73\x44\xb7\xa9\x34\xa9\x58\x95\xa7\x66\x6a\x3e" +
			"\x02\x00\x00\xff\xff",
		size: 631,
		mode: 0664,
		time: time.Unix(1564685966, 42485155),
	},
	"src/hub/api/requests/aks-adapter-template.json.template": &asset{
		name: "aks-adapter-template.json.template",
		data: "" +
			"\x94\x92\x4d\x4b\x33\x31\x14\x85\xf7\xfd\x15\x21\xbc\xcb\x36\xf0\xee\xa4\xe0\xa2\x0b\x71\x21\x48" +
			"\x51\x5c\x49\x17\xb7\xc9\xed\x10\xf2\xc9\xbd\x89\x54\xcb\xfc\x77\xc9\x7c\x88\x42\x69\x75\x35\x97" +
			"\x39\xcf\x9c\x3b\x39\x27\xa7\x85\x10\x42\xc8\x08\x01\xe5\x5a\xc8\x7f\xa7\x36\xf5\x72\x39\xbe\x36" +
			"\xc8\x9a\x6c\x2e\x36\xc5\xa6\x6e\x1e\x9e\xc5\xc6\x40\x2e\x48\x33\xc1\x05\xb4\x6b\x1a\x38\x5e\xc1" +
			"\xa8\xad\xff\xcf\xaa\x4e\x21\xa7\x88\xb1\xf0\x5d\x84\xbd\x47\x23\xd7\xe2\x55\xba\x1b\x5e\x81\x63" +
			"\xb9\x14\xb2\x58\xef\x9b\x99\x90\x85\x00\x0f\xd6\xb5\xd1\xe0\xb1\x3d\x34\x52\x59\x05\x88\xd0\x21" +
			"\xc9\xdd\xe4\xf8\x86\xb4\xe7\xc1\xc5\x60\xf6\xe9\xbd\x81\x35\x4e\xf3\x0c\x15\xe8\x46\x66\xfa\xa1" +
			"\xdb\xb6\x6d\x16\x33\x10\x04\x2c\x48\x03\x32\x9e\xff\x47\x06\x26\xb2\x32\x29\x80\x8d\x72\x10\xfb" +
			"\xa5\x38\x43\x69\x9f\xaa\x51\x84\x5d\xcb\xe6\x2a\x07\x1f\x95\xf0\x09\x39\x55\xd2\x78\x4f\xa9\xe6" +
			"\xc7\xa6\x5f\xf8\xce\xd5\x3d\x52\xc4\x82\xac\xc0\xb1\xd2\xbe\x72\x8b\xfd\xd2\xa6\x39\x6c\x65\x63" +
			"\x47\xc8\xac\x2a\xf9\x2d\xe1\xc1\x1e\xa7\x3e\xc6\x04\xc1\xd7\x81\x87\x9c\xff\x66\xc7\x9c\x5e\xae" +
			"\x39\xf2\xef\x2c\xbf\x77\xab\x70\xba\x1b\xe7\x2c\x0b\xd5\x39\xa6\xaf\x76\x11\x02\x6f\x91\x82\x65" +
			"\xb6\x29\x0e\x35\xee\x16\xfd\xe2\x33\x00\x00\xff\xff",
		size: 715,
		mode: 0664,
		time: time.Unix(1564685966, 42485155),
	},
	"src/hub/api/requests/eks-adapter-instance.json.template": &asset{
		name: "eks-adapter-instance.json.template",
		data: "" +
			"\xa4\xcd\x41\x6a\xc3\x30\x10\x05\xd0\xbd\x4f\x21\x44\x97\xa9\x0e\x10\xe8\xaa\x17\xe8\xa6\xab\x92" +
			"\xc5\xc4\xfe\x35\xc2\xd2\x48\xcc\x8c\x42\x21\xf8\xee\xa5\xb6\x0b\x2d\x18\xda\x90\x9d\xf8\x5f\xf3" +
			"\xfe\xb5\x73\xce\x39\xcf\x94\xe1\x8f\xce\x3f\x5c\xbf\x5e\xb3\x3f\xac\xb1\x21\xd7\x44\xb6\x54\xdf" +
			"\x19\xf8\x12\xa5\x70\x06\xdb\xcf\xd8\x68\x54\x7f\x74\x6f\x9e\x06\xaa\x06\x79\xc2\xa4\xfe\xb4\x95" +
			"\x95\x84\x32\x0c\xb2\x7c\x59\x37\x7f\xed\x0e\xac\x61\x28\x99\x22\x6f\xde\xd2\x5e\x28\xb5\x75\x7b" +
			"\xc9\xe6\x83\xdb\x39\x9d\xda\x19\xc2\x30\x68\xa0\x1a\x03\x78\xa8\x25\xb2\xdd\xeb\xf4\xf4\x0c\xb9" +
			"\x4b\xc1\xa4\xa1\x4f\x4d\x0d\x72\x33\xd3\x97\x5c\x0b\x83\x2d\x44\x1e\x05\xaa\xa1\x49\x7a\x11\xbc" +
			"\xc7\x8f\x5d\x8b\x6a\xbd\x8d\x53\x2d\xaf\x7f\x89\xfa\x3f\xb2\x87\xd8\x63\x26\xa6\x11\x12\xc0\x74" +
			"\x4e\x18\x76\x49\x93\x86\x8d\x3c\x75\x73\xf7\x19\x00\x00\xff\xff",
		size: 635,
		mode: 0664,
		time: time.Unix(1564685966, 42485155),
	},
	"src/hub/api/requests/eks-adapter-template.json.template": &asset{
		name: "eks-adapter-template.json.template",
		data: "" +
			"\x94\x92\x3f\x6b\x23\x31\x10\xc5\x7b\x7f\x0a\x21\xae\xb4\x17\xae\x3b\x0c\x57\x1c\x87\xab\x34\x86" +
			"\x90\x2a\xb8\x18\xaf\x9e\x8d\x58\x69\x24\x66\xb4\xc1\xc1\xf8\xbb\x07\xed\x1f\x93\x80\x49\xec\x6a" +
			"\x87\x7d\x6f\x7e\xbb\x7a\x4f\xe7\x85\x31\xc6\x58\xa6\x08\xbb\x36\xf6\xd7\xb9\x4e\x17\xbb\x1c\x5f" +
			"\x3b\x68\x2b\x3e\x17\x9f\xb8\xaa\x9b\xa7\x67\xf3\xcf\x51\x2e\x90\xd9\xa1\x85\xda\xae\x6a\xe8\x74" +
			"\x45\xa3\xb6\xfe\x3d\xab\x6d\x8a\x39\x31\xb8\xe8\x86\x69\x1f\xe0\xec\xda\xbc\xda\xee\x8f\xae\xd0" +
			"\xa9\x5d\x1a\x5b\x7c\x08\x15\x66\x6c\x11\xc2\xc1\x77\x75\x74\x38\xd5\x47\x0b\x29\xab\x48\x4c\x47" +
			"\x88\xdd\x4d\xc4\x37\xc8\x5e\x07\x8a\x43\x0e\xe9\xbd\x1a\x7b\x9e\xe6\xd9\x54\xe8\x38\x7a\xa6\x1f" +
			"\xfa\x5b\xbf\x36\x8b\x99\x84\x22\x0a\x64\xb0\x8c\xe7\xff\x92\x81\x63\x6d\x5c\x8a\xe4\xd9\x0e\xe2" +
			"\x65\x69\x6e\xb8\xba\x7e\x0f\x61\x14\x68\x43\xd9\x37\x60\x97\x93\xe7\xf2\xc0\x4a\x4b\xff\x21\xf7" +
			"\x2e\xa0\xd3\xa6\x0d\xbd\xd6\xe8\xbf\xd9\xb8\x06\xde\x78\x3e\x0a\x54\x9b\x5e\xc2\x56\x70\xf0\xa7" +
			"\xa9\x93\x31\x45\x0a\xfd\xe0\xa7\x9c\x1f\xc3\xa9\xa6\x97\x9f\x88\x7a\x1f\xf2\x73\xbf\x0d\xa6\xfb" +
			"\x71\x0b\x59\xa4\xc7\x84\xbc\x36\x0c\x8a\xba\x85\x44\xaf\xea\x13\x0f\x55\xee\x16\x97\xc5\x47\x00" +
			"\x00\x00\xff\xff",
		size: 719,
		mode: 0664,
		time: time.Unix(1564685966, 42485155),
	},
	"src/hub/api/requests/gke-adapter-instance.json.template": &asset{
		name: "gke-adapter-instance.json.template",
		data: "" +
			"\x94\x8e\x41\x6a\xc3\x30\x10\x45\xf7\x3e\x85\x10\x5d\xa6\x3a\x40\xa0\x77\xe8\xa6\xab\x92\xc5\xc4" +
			"\xfa\x15\xc2\xd2\x48\x8c\x46\xa1\x10\x7c\xf7\x52\xdb\x85\x16\x0c\x8d\x77\xe2\x7d\xfd\xf7\xe7\x3e" +
			"\x18\x63\x8c\x65\xca\xb0\x67\x63\x9f\xee\xdf\xaf\xd9\x9e\x56\xac\xc8\x35\x91\x2e\xd1\x0f\x03\xdf" +
			"\xa2\x14\xce\x60\xfd\x8d\x95\x42\xb3\x67\xf3\x6e\xc9\x53\x55\xc8\x4b\x98\x60\x2f\x5b\x58\x49\x28" +
			"\x43\x21\xcb\x97\x75\xf3\xcf\xae\xe7\xe6\x7c\xc9\x14\x79\xf3\x2d\xe9\x8d\x52\x5f\xb7\x17\x36\x9f" +
			"\xcc\x4e\x75\x4c\xa5\x7b\x27\x08\xb1\x1c\x2f\x4f\xfd\x0a\x61\x28\x9a\x0b\x13\xdc\x98\x7a\x53\xc8" +
			"\xf1\x1b\x4a\xae\x85\xc1\xea\x22\x07\x41\x6b\xae\x4b\x7a\x15\x7c\xc4\xcf\x5d\x17\xd5\x7a\x4c\xd7" +
			"\x5a\x79\xfb\xcf\xd8\x1e\x53\x8e\x10\x7d\xce\xc4\x14\x20\x0e\x4c\xd7\x04\xbf\xab\x54\xe9\xd8\x94" +
			"\x97\x61\x1e\xbe\x02\x00\x00\xff\xff",
		size: 554,
		mode: 0664,
		time: time.Unix(1564685966, 42485155),
	},
	"src/hub/api/requests/gke-adapter-template.json.template": &asset{
		name: "gke-adapter-template.json.template",
		data: "" +
			"\x94\x91\x4d\x4b\x03\x31\x10\x86\xef\xfd\x15\x21\x78\x6c\x03\xde\xa4\xe0\xc1\x43\xf1\xe0\xa5\x17" +
			"\x4f\xd2\xc3\x74\x33\x0d\x21\x9f\xcc\x24\x52\x29\xfb\xdf\x25\xbb\x59\x51\x28\x55\x4f\x3b\xec\xfb" +
			"\xf0\x24\x79\xe7\xb2\x12\x42\x08\x19\x21\xa0\xdc\x0a\x79\x77\x69\xd3\x28\xd7\xf3\x6f\x8d\x3c\x90" +
			"\xcd\xc5\xa6\xd8\xd2\xe7\x97\x9d\x78\xd2\x90\x0b\xd2\x42\x70\x81\xc1\xb5\xcc\x38\xdc\xc0\x9c\x6d" +
			"\xef\x97\x74\x48\x21\xa7\x88\xb1\xf0\x2e\xc2\xd1\xa3\x96\x5b\xf1\x26\xdd\x03\x6f\x8c\x43\xb9\x16" +
			"\xb2\x58\xef\x9b\x4c\xc8\x42\x80\x27\xeb\xda\xa8\xf1\xdc\x3e\x03\x52\xd9\x04\x88\x60\x90\xe4\xa1" +
			"\x1b\xdf\x91\x8e\x3c\x59\x34\x66\x9f\x3e\x1a\x58\x63\x9f\x17\xa8\x80\x99\x99\x7e\xa1\xc7\x76\xda" +
			"\x12\x66\x20\x08\x58\x90\x26\x64\x7e\xff\x8f\x0e\x74\x64\xa5\x53\x00\x1b\xe5\x14\x8e\x6b\x71\x85" +
			"\x1a\x7c\xaa\x5a\x11\x9a\xd6\xcd\x0d\xce\xd5\x23\x52\xc4\x82\xac\x8c\x43\x35\xf8\xca\xad\xbe\x5b" +
			"\xe6\xa5\x34\x65\xa3\x21\x64\x56\x95\xfc\x9e\xf0\x64\xcf\xbd\xd7\xb9\x09\xf0\x75\xe2\x21\xe7\xff" +
			"\xe9\x98\xd3\xeb\x6f\x46\xfe\x9b\xf2\xfb\x8e\x14\xf6\x1d\x5f\x53\x16\xaa\xd8\x95\x5f\x5b\x42\x08" +
			"\xbc\x47\x0a\x96\xd9\xa6\x38\xad\xe3\xb0\x1a\x57\x9f\x01\x00\x00\xff\xff",
		size: 659,
		mode: 0664,
		time: time.Unix(1564685966, 42485155),
	},
	"src/hub/api/requests/hybrid-adapter-instance.json.template": &asset{
		name: "hybrid-adapter-instance.json.template",
		data: "" +
			"\xac\xce\xc1\x4a\xf4\x30\x14\x05\xe0\x7d\x9f\x22\x84\x7f\x39\x7f\x1e\x60\xc0\x95\x1b\xc5\x8d\x1b" +
			"\x57\x32\x8b\xdb\xe6\x58\x83\xc9\x4d\xb8\xb9\x1d\x1c\x86\xbe\xbb\xd8\x56\x50\x28\xe8\x30\xdd\x85" +
			"\x73\x92\xef\xe4\xdc\x18\x63\x8c\x65\x4a\xb0\x7b\x63\xff\x9d\x3f\x4f\xa3\xdd\xcd\xb1\x22\x95\x48" +
			"\x3a\x55\x5f\x19\xf8\x18\x24\x73\x02\xeb\xf7\x58\xa9\xaf\x76\x6f\x9e\x2d\x79\x2a\x0a\xb9\x79\x3d" +
			"\xb5\x12\xbc\x3d\x2c\x7d\x21\xa1\x04\x85\x4c\xb7\xe6\xd9\x1f\xd3\x9e\xab\xf3\x39\x51\xe0\x85\x9c" +
			"\xda\x23\xc5\x61\x9e\x9f\xb2\x71\x67\x56\x9e\xbe\x0d\x2d\x84\xa1\xa8\x8e\x4a\x70\x60\x5f\x72\x60" +
			"\xbd\xd6\xe9\xe8\x16\x72\xbd\x12\x03\x58\xb7\x93\x1e\x70\xba\x18\xea\x72\x2a\x99\xc1\xea\x02\xf7" +
			"\x82\x5a\x5d\x55\xd2\xd0\xdd\x97\xcd\xa8\xbb\x5c\x75\x03\x6c\x90\xf8\x28\x78\x09\xef\xab\x16\x95" +
			"\x72\xe1\xdf\x6a\x7e\xfa\x4d\xac\x7f\x23\x3b\x88\xfe\x4f\xc4\xd4\x43\x1c\x98\xda\x08\xbf\x4a\xaa" +
			"\x0c\x58\xc8\x43\x33\x36\x1f\x01\x00\x00\xff\xff",
		size: 866,
		mode: 0664,
		time: time.Unix(1567627758, 465469116),
	},
	"src/hub/api/requests/hybrid-adapter-template.json.template": &asset{
		name: "hybrid-adapter-template.json.template",
		data: "" +
			"\x9c\x92\xcf\x6a\xf3\x30\x10\xc4\xef\x79\x0a\x21\xbe\x63\x6c\xf8\x6e\x25\xd0\x43\x5b\x0a\x2d\xbd" +
			"\xe4\xd2\x53\xc9\x61\x6d\x6d\x52\x61\x69\x25\x76\xd7\x25\x21\xe4\xdd\x8b\xff\x85\x16\x42\x5b\xe7" +
			"\xe4\xc5\xf3\x9b\x91\xd0\xce\x71\x61\x8c\x31\x96\x20\xa2\x5d\x19\xfb\xef\xd8\x4d\x27\xbb\x1c\x7e" +
			"\x3b\x94\x9a\x7d\x56\x9f\xa8\x53\xef\x81\xb1\x88\xa8\x10\xcc\x9d\x83\xac\xc8\x13\x28\x0a\x75\xd3" +
			"\x21\xef\x87\x8a\xbd\x2b\x60\x90\x57\xff\x27\xa0\x4e\x31\x27\x42\x52\x79\x24\xa8\x02\x3a\xbb\x32" +
			"\x6f\xb6\xb9\x91\x62\x70\xd8\xa5\xb1\xea\x43\xe8\x22\x8d\x55\x06\xdc\xfa\xa6\x1b\x1d\xee\xbb\x4f" +
			"\x8d\xac\x45\x04\x82\x1d\xb2\xdd\x8c\xa1\x1f\xc8\x95\xf4\x41\x0e\x73\x48\x87\x0e\x6c\x69\x9c\x27" +
			"\x48\x61\x37\x30\xe3\x9d\x6e\xc7\x03\x27\x3d\x03\x43\x44\x45\xee\xa9\xe1\x3d\xbe\xbd\x89\x23\x29" +
			"\x5d\x8a\xe0\xc9\xf6\xe2\x69\x69\x2e\x50\x4d\x5b\x21\x13\x2a\x4a\x09\xd9\x97\x48\x2e\x27\x4f\x3a" +
			"\xc3\x52\xc3\x03\xf2\x2c\x43\xf0\x48\x7a\x95\xe9\x05\x0f\x3f\x79\xce\xdb\x2a\x3d\xed\x18\x45\x4a" +
			"\x51\x50\x5f\x3f\xe7\x6b\x5c\x4f\x49\x74\x9e\xaf\xe5\xb0\x66\xdc\xfa\xfd\x58\x9f\x61\xdb\x10\xda" +
			"\x9e\x87\x3c\xf7\x1a\x92\x5e\x7f\x4b\x94\xbf\x45\x7e\xed\x61\x89\x63\x95\x2f\x45\x2a\xb7\x38\x46" +
			"\x9e\x9b\x88\x10\x65\x8d\x1c\xbd\x88\x4f\xd4\xf7\x6d\xb3\x38\x2d\x3e\x03\x00\x00\xff\xff",
		size: 900,
		mode: 0664,
		time: time.Unix(1567693714, 430704897),
	},
	"src/hub/api/requests/k8s-aws-adapter-instance.json.template": &asset{
		name: "k8s-aws-adapter-instance.json.template",
		data: "" +
			"\xac\xcd\x31\x4b\x03\x41\x10\x05\xe0\xfe\x7e\xc5\xb2\x58\x26\x5b\x4b\xc0\xca\xd2\xc6\xc6\x4a\x52" +
			"\x4c\xee\x9e\xc7\x92\xdd\xd9\x65\x66\x2e\x2a\xe1\xfe\xbb\x78\x77\x82\xc2\x81\x86\xd8\x2d\x6f\xde" +
			"\x7e\xef\xdc\x38\xe7\x9c\x67\xca\xf0\x3b\xe7\x6f\xce\x9f\xaf\xd1\x6f\xe6\xd8\x90\x6b\x22\x9b\x4e" +
			"\x5f\x19\xf8\x14\xa5\x70\x06\xdb\xf7\xd8\xa8\x57\xbf\x73\xcf\x9e\x3a\xaa\x06\xb9\x3b\xde\xea\x96" +
			"\x5e\xd5\xef\x97\x42\x25\xa1\x0c\x83\x4c\xb5\x79\xf7\xc7\x76\xc7\x1a\xba\x92\x29\xf2\x62\x4e\xd7" +
			"\x13\xa5\x61\xde\x9f\xb2\x71\xe3\x56\xbe\x1e\x87\x03\x84\x61\xd0\x40\x35\x06\x70\x57\x4b\x64\xbb" +
			"\xd6\x69\xe9\x1e\x72\xbd\x92\x22\xd8\xfe\x4f\x7a\xc0\xfb\xc5\x50\x5b\x72\x2d\x0c\xb6\x10\xb9\x17" +
			"\xa8\x86\x41\xd2\xa3\xe0\x25\xbe\xad\x5a\x54\xeb\x65\x9c\x6a\x79\xfa\x4d\xd4\xbf\x91\x2d\xc4\xb6" +
			"\x99\x98\x7a\x48\x00\xd3\x21\xa1\x5b\x25\x4d\x06\x2c\xe4\xbe\x19\x9b\x8f\x00\x00\x00\xff\xff",
		size: 715,
		mode: 0664,
		time: time.Unix(1564685966, 42485155),
	},
	"src/hub/api/requests/k8s-aws-adapter-template.json.template": &asset{
		name: "k8s-aws-adapter-template.json.template",
		data: "" +
			"\x9c\x92\xc1\x6a\x02\x31\x10\x86\xef\x3e\x45\x08\x85\x5e\x74\xa1\x37\x11\x7a\x90\xd2\x93\x17\x41" +
			"\x4a\x0f\xc5\xc3\xb8\x19\x97\xb0\xc9\x24\xcc\x64\x5b\x45\x7c\xf7\x92\xdd\xac\xb4\x20\x6d\xed\x69" +
			"\x87\x9d\x6f\xbe\x84\xf9\x73\x9a\x28\xa5\x94\x26\xf0\xa8\x17\x4a\xdf\x9d\x72\x75\xd6\xd3\xe1\xb7" +
			"\x41\xa9\xd9\xc6\x64\x03\xe5\xee\xb2\xb1\x0e\x37\x09\xea\xf6\x5e\xd4\x6a\xbe\x51\x81\xd4\xf2\x75" +
			"\xa3\xc0\x40\x4c\xc8\xe3\x94\x64\x22\xf3\xed\x5c\x66\xa5\xb7\x78\x18\xbb\x75\xf0\x31\x10\x52\x92" +
			"\x67\x82\x9d\x43\xa3\x17\xea\x6d\x40\x3f\x44\x4f\x95\x4e\xd6\xb9\x2c\x53\x3a\x31\xe0\xde\xb6\xb9" +
			"\x34\x78\xc8\x9f\x1a\x39\xcd\x3c\x10\x34\xc8\x7a\x5b\x8c\xef\xc8\x3b\xe9\x2d\x06\xa3\x0b\xc7\x0c" +
			"\x76\x54\xea\x11\x4a\xd0\x0c\x4c\xb9\xd0\xe3\x78\xe2\x08\x44\x60\xf0\x98\x90\x7b\x6c\xd8\xcb\xb7" +
			"\xdd\x18\x92\xca\x04\x0f\x96\x74\xdf\x3c\x4f\xd5\x15\xaa\xed\x76\xc8\x84\x09\xa5\x82\x68\xab\x1a" +
			"\x9e\x90\xd3\x2d\x03\xce\x22\xa5\x7f\x0d\xad\xf0\xf8\xd3\xcc\x65\xf1\x95\xa5\x86\x51\xa4\xea\xd8" +
			"\xad\x19\xf7\xf6\x50\xb2\x19\xb6\x09\xae\xeb\x79\x88\xf1\x36\x9d\x48\x78\xf9\xcd\x28\x7f\x53\x7e" +
			"\xcd\xb9\xc2\xf2\x4e\xae\x29\x13\x77\x58\x94\x97\xa4\x11\xbc\xac\x91\xbd\x15\xb1\x81\xfa\x38\xb7" +
			"\x93\xf3\xe4\x33\x00\x00\xff\xff",
		size: 747,
		mode: 0664,
		time: time.Unix(1564685966, 42485155),
	},
	"src/hub/api/requests/metal-adapter-instance.json.template": &asset{
		name: "metal-adapter-instance.json.template",
		data: "" +
			"\xac\xce\x41\x4b\xc4\x30\x10\x05\xe0\x7b\x7f\x45\x08\x1e\xd7\xfc\x80\x05\x4f\x5e\x14\x2f\x5e\x3c" +
			"\xc9\x1e\x66\x9b\x67\x09\x26\x93\x30\x99\x2e\xca\xd2\xff\x2e\xb6\x15\x14\x0a\xee\xd2\xde\xc2\x7b" +
			"\xc9\xf7\x72\x6e\x8c\x31\xc6\x32\x25\xd8\xbd\xb1\x37\xe7\xef\xd3\x60\x77\x53\xac\x48\x25\x92\x8e" +
			"\xd5\x4f\x06\x3e\x05\xc9\x9c\xc0\xfa\x3b\x56\xea\xaa\xdd\x9b\x57\x4b\x9e\x8a\x42\xee\x12\x94\xa2" +
			"\x3d\xcc\x75\x21\xa1\x04\x85\x8c\x97\xa6\xd5\x3f\xcb\x9e\xab\xf3\x39\x51\xe0\x59\x1c\xdb\x13\xc5" +
			"\x7e\x5a\x1f\xb3\x61\x67\x16\x9e\xbe\xf7\x47\x08\x43\x51\x1d\x95\xe0\xc0\xbe\xe4\xc0\xba\xd6\x69" +
			"\xe9\x1e\xb2\x5e\x89\x01\xac\xdb\x49\x4f\xf8\xbc\x1a\x6a\x73\x2a\x99\xc1\xea\x02\x77\x82\x5a\x5d" +
			"\x55\xd2\xd0\x3e\x96\xcd\xa8\x87\x5c\x75\x03\xac\x97\xf8\x2c\x78\x0b\x1f\x8b\x16\x95\x72\xe5\xdf" +
			"\x6a\x7e\xf9\x4f\xac\x97\x91\x2d\x44\x6f\x13\x31\x75\x10\x07\xa6\x63\x84\x5f\x24\x55\x7a\xcc\xe4" +
			"\xa1\x19\x9a\xaf\x00\x00\x00\xff\xff",
		size: 865,
		mode: 0664,
		time: time.Unix(1566264221, 744967922),
	},
	"src/hub/api/requests/metal-adapter-template.json.template": &asset{
		name: "metal-adapter-template.json.template",
		data: "" +
			"\x9c\x92\x41\x4b\x03\x31\x10\x85\xef\xfd\x15\x21\x78\x6c\x17\xbc\x49\xc1\x83\x8a\xa0\x78\xe9\xc5" +
			"\x93\xf4\x30\xdd\x4c\x4b\xd8\x64\x12\x66\x66\xa5\xa5\xf4\xbf\xcb\x6e\x76\x8b\x42\x51\xb7\xa7\x1d" +
			"\xf6\xbd\xf7\x25\x64\xde\x71\x66\x8c\x31\x96\x20\xa2\x5d\x1a\x7b\x73\xec\xa6\x93\x9d\x97\xdf\x0e" +
			"\xa5\x66\x9f\xd5\x27\xea\xd4\x47\x60\x5c\x44\x54\x08\xe6\xc1\x41\x56\xe4\xd1\x28\x0a\x75\xd3\x59" +
			"\x7a\x75\x01\x45\x5d\xde\x8e\x7a\x9d\x62\x4e\x84\xa4\xf2\x4c\xb0\x09\xe8\xec\xd2\x7c\xd8\xe6\x4e" +
			"\x0a\xce\xce\x8d\x55\x1f\x42\x07\x34\x56\x19\x70\xeb\x9b\x6e\x74\xb8\xef\x3e\x35\xb2\x2e\x22\x10" +
			"\xec\x90\xed\x7a\x60\x7e\x22\x6f\xa4\xe7\x38\xcc\x21\x1d\x3a\x63\x4b\xc3\x3c\x9a\x14\x76\xc5\x33" +
			"\x5c\xe9\xbe\x9c\x37\xca\x19\x18\x22\x2a\x72\x6f\x2a\x8f\xf1\xe3\x41\x1c\x49\xe5\x52\x04\x4f\xb6" +
			"\x17\x4f\x73\x73\xc1\xd5\xb4\x1b\x64\x42\x45\xa9\x20\xfb\x0a\xc9\xe5\xe4\x49\x27\x44\x6a\x78\x42" +
			"\x9e\x14\x08\x1e\x49\xaf\x0a\xbd\xe1\xe1\xb7\xcc\x79\x57\x95\xa7\x1d\xa3\x48\x25\x0a\xea\xeb\xd7" +
			"\x7c\x4d\xea\x25\x89\x4e\xcb\xb5\x1c\x56\x8c\x5b\xbf\x1f\xca\x53\x96\x0d\xa1\xed\xfd\x90\xa7\x5e" +
			"\x43\xd2\xfb\x5f\x44\xf9\x1f\xf2\x7b\x0d\x2b\x1c\x8a\x7c\x09\xa9\xdc\xe2\x80\x3c\x17\x11\x21\xca" +
			"\x0a\x39\x7a\x11\x9f\xa8\xef\xdb\x7a\x76\x9a\x7d\x05\x00\x00\xff\xff",
		size: 897,
		mode: 0664,
		time: time.Unix(1566264221, 744967922),
	},
	"src/hub/api/requests/openshift-adapter-instance.json.template": &asset{
		name: "openshift-adapter-instance.json.template",
		data: "" +
			"\xa4\xcf\xc1\x6a\xc3\x30\x0c\x06\xe0\x7b\x9e\xc2\x98\x1d\x3b\x3f\x40\x61\xa7\xbd\xc0\x2e\x3b\x8d" +
			"\x1e\xd4\xe4\x6f\x66\x1a\xcb\x42\x52\xca\xa0\xe4\xdd\xc7\x92\x0c\x36\x08\x6c\xa5\x37\xf3\x4b\xfe" +
			"\x7e\xfb\xda\x84\x10\x42\x64\x2a\x88\xfb\x10\x1f\xae\x5f\xa7\x29\xee\x96\xd8\x51\x64\x20\x9f\x47" +
			"\xdf\x19\xf8\x92\xb5\x72\x01\xfb\xcf\xd8\xa9\xb7\xb8\x0f\x6f\x91\x3a\x12\x87\x3e\x55\x01\xdb\x7b" +
			"\x3e\x79\x3c\xac\x2b\x42\x4a\x05\x0e\x9d\x17\x97\xe6\x5f\xed\x1d\x5b\xea\x6a\xa1\xcc\xab\x3a\x4f" +
			"\x2f\x34\x8c\xcb\x0b\xe6\x6c\xda\x85\x8d\xab\xe7\xf1\x08\x65\x38\x2c\x91\xe4\x04\xee\xa4\x66\xf6" +
			"\x7b\x9d\x96\x9e\xa1\x77\x2b\x5e\xcf\xb8\xfd\x4b\x6d\x2d\x52\x19\xec\x29\x73\xaf\x30\x4b\xa3\x0e" +
			"\x2f\x8a\x53\xfe\xd8\xb4\x48\xe4\x36\xce\xac\xbe\xfe\x25\xda\xff\xc8\x16\xea\x8f\x85\x98\x7a\x68" +
			"\x02\xd3\x71\x40\xb7\x49\xba\x8e\x58\xc9\x43\x33\x35\x9f\x01\x00\x00\xff\xff",
		size: 639,
		mode: 0664,
		time: time.Unix(1564685966, 42485155),
	},
	"src/hub/api/requests/openshift-adapter-template.json.template": &asset{
		name: "openshift-adapter-template.json.template",
		data: "" +
			"\x94\x92\x4f\x6b\xe3\x30\x10\xc5\xef\xf9\x14\x42\xec\x31\x31\xec\x6d\x09\xec\x61\x59\x7a\x6e\xa0" +
			"\xf4\x54\x72\x98\x58\x2f\xa9\xb0\x3d\x12\x33\xe3\x92\x12\xf2\xdd\x8b\xfc\x27\x6d\x21\xb4\xcd\xc9" +
			"\x83\xf5\x9b\x9f\xc5\x7b\x3e\x2d\x9c\x73\xce\x33\x75\xf0\x6b\xe7\x7f\x9d\xca\x74\xf6\xcb\xf1\x75" +
			"\x80\xd6\x12\xb3\xc5\xc4\xe5\xf4\x3e\x83\x1f\x9e\xe3\xde\xdc\xbf\x40\xd9\x20\x33\xa7\x46\x75\x53" +
			"\x88\x94\xc1\x5a\x88\x15\x8d\xc4\xfa\xf7\xcc\xd4\xa9\xcb\x89\xc1\xa6\x77\x4c\xbb\x16\xc1\xaf\xdd" +
			"\x93\x6f\xfe\xe8\xea\xb2\xe4\x97\xce\x5b\x6c\xdb\x22\x76\xde\x84\xb0\x8f\x4d\x19\x03\x8e\xe5\x51" +
			"\x43\x6c\xd5\x11\xd3\x01\xe2\xb7\x93\xf7\x05\xb2\xd3\xc1\x15\x90\xdb\xf4\x5a\xc0\x9e\xa7\x79\x86" +
			"\x8c\x0e\x23\x33\x5d\xeb\xef\xfb\x37\x67\x24\x93\x50\x07\x83\x0c\xe0\x98\xcb\xa7\x6c\x02\x6b\x15" +
			"\x52\x47\x91\xfd\x70\x78\x5e\xba\x2b\x54\xd3\xef\x20\x0c\x83\x56\x94\x63\x05\x0e\x39\x45\xb6\x1b" +
			"\x56\x6a\xfa\x0f\xb9\x65\xc1\x52\x83\x2f\xef\x74\x89\xbe\x8a\x7c\x10\xa8\x56\xbd\xb4\x1b\xc1\x3e" +
			"\x1e\xa7\x76\xc6\x24\xa9\xed\x07\x9e\x72\xbe\x4d\xa7\x9a\x1e\xbf\x33\xea\xcf\x94\x1f\x3b\xae\x30" +
			"\xfd\x29\xd7\x94\x26\x3d\x26\xe5\xa5\x65\x50\xa7\x1b\x48\x17\x55\x63\xe2\xa1\xc8\xed\xe2\xbc\x78" +
			"\x0b\x00\x00\xff\xff",
		size: 741,
		mode: 0664,
		time: time.Unix(1564685966, 42485155),
	},
	"src/hub/initialize/hub.yaml.template": &asset{
		name: "hub.yaml.template",
		data: "" +
			"\x7c\x90\xc1\x6e\x33\x21\x0c\x84\xef\x7e\x0a\x1f\xfe\x63\x60\x95\x2b\xaf\x12\xe5\x40\xc0\x2b\xf1" +
			"\x87\xc5\xc8\x98\xad\xa2\xaa\xef\x5e\xb1\x9b\xa6\x4d\x54\x75\x2e\x88\x19\xbe\x01\x6c\x8c\x81\x95" +
			"\xa4\x25\x2e\x0e\x8f\x70\x4d\x25\x3a\x6c\xea\xc3\x15\x16\x52\xef\x00\xb1\xf8\x85\x1c\xfe\x7b\xaf" +
			"\xc2\xff\x29\xe8\x87\x3b\x02\xe2\x45\x12\xcd\x3f\x5d\x40\x8c\xd4\x82\xa4\xaa\x5b\x97\xb5\x16\x10" +
			"\x1b\x77\x09\x34\x5a\x86\x62\x12\x87\x76\x02\x08\xbc\x54\x2e\x54\xb4\x8d\xc8\xdc\xaf\x98\x93\x34" +
			"\x35\x8f\x6c\x83\x7e\x2b\xb0\xd3\x37\x3f\xbd\x42\x90\xd3\x4c\xe1\x16\xf2\x06\xad\x24\x97\xe6\xf0" +
			"\x14\xa9\x66\xbe\x1d\x70\x5f\x8d\x52\xd3\x03\xf6\xb2\x6f\xcf\x80\xc8\x12\x49\x1c\x9e\x5e\xea\xce" +
			"\x00\xdc\xb5\xf6\x3f\x1f\xea\x62\x69\x36\xf2\xe2\x53\x01\xa8\x5e\xfc\x42\x4a\xf2\x44\x84\xcc\x3d" +
			"\x6e\x9f\x78\xce\x87\xbe\xce\x54\xe1\x35\x45\x92\xbb\x3d\xe6\x39\xfb\x9e\xd5\xa1\x7f\x6b\x0f\x73" +
			"\xf5\xb9\xd3\x6e\x7d\x06\x00\x00\xff\xff",
		size: 444,
		mode: 0664,
		time: time.Unix(1564685966, 45485123),
	},
	"src/hub/initialize/hub-component.yaml.template": &asset{
		name: "hub-component.yaml.template",
		data: "" +
			"\x74\x90\xbf\x6e\x02\x31\x0c\xc6\x77\x3f\x85\x85\x98\x2a\x5d\x50\xd7\x6c\x48\xed\xc0\x00\xed\x1b" +
			"\x54\x81\x98\x36\xe5\xe2\x04\xc7\xb9\x05\xf1\xee\xd5\xfd\xa1\x9c\x50\x9b\xc9\xf1\xf7\x7d\xd6\xcf" +
			"\x6e\x9a\x06\x3a\x92\x12\x12\x5b\x7c\x86\x53\x60\x6f\xf1\x90\x62\x4e\x4c\xac\x10\x49\x9d\x05\x44" +
			"\x76\x91\x2c\x2e\x2f\x59\xd2\x37\x1d\xf4\x0a\x88\x7b\x09\x74\x7c\xe8\x95\x54\xe5\x40\x7d\xa0\x7f" +
			"\x3e\x88\x45\xb3\x02\x10\x3a\xd7\x20\x54\x7a\xa1\xc1\x53\xdd\x93\x30\x29\x95\xe1\x1b\xf8\x53\xa8" +
			"\x14\x80\x2c\xa9\x0b\xfe\xe6\x9a\xcd\x85\xec\xc4\x45\x52\x92\x49\x1b\x69\x7e\x29\xcd\x34\xc2\x1c" +
			"\xcf\x9e\x67\x06\xcf\xc5\xf8\x14\x5d\xe0\x01\x88\xb8\xb3\xf8\xf2\xb6\x5d\x6f\x76\x1f\xbb\xf5\xf6" +
			"\xf5\xc1\xd9\x57\x77\xdf\x60\x00\xa5\x98\x5b\xa7\x23\xd3\x31\xb4\x63\xd1\xe7\x16\x4f\xe6\x26\x2e" +
			"\x00\x52\xd5\x5c\xf5\x1f\xba\xfb\x26\xc6\xe5\x30\xe4\xa7\xdb\xad\xdf\x37\x48\xec\x73\x0a\xac\x43" +
			"\xbf\x73\x6d\x25\x8b\x5f\xaa\xd9\xae\x56\xb3\xe0\xf2\xf2\xf7\xb6\x57\xf8\x09\x00\x00\xff\xff",
		size: 449,
		mode: 0664,
		time: time.Unix(1564685966, 45485123),
	},
}

// AssetAndInfo loads and returns the asset and asset info for the
// given name. It returns an error if the asset could not be found
// or could not be loaded.
func AssetAndInfo(name string) ([]byte, os.FileInfo, error) {
	a, ok := _bindata[filepath.ToSlash(name)]
	if !ok {
		return nil, nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
	}

	a.once.Do(func() {
		fr := flate.NewReader(strings.NewReader(a.data))

		var buf bytes.Buffer
		if _, a.err = io.Copy(&buf, fr); a.err != nil {
			return
		}

		if a.err = fr.Close(); a.err == nil {
			a.bytes = buf.Bytes()
		}
	})
	if a.err != nil {
		return nil, nil, &os.PathError{Op: "read", Path: name, Err: a.err}
	}

	return a.bytes, a, nil
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	a, ok := _bindata[filepath.ToSlash(name)]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
	}

	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	data, _, err := AssetAndInfo(name)
	return data, err
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

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}

	return names
}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	return restore.Asset(dir, name, AssetAndInfo)
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	return restore.Assets(dir, name, AssetDir, AssetAndInfo)
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

	if name != "" {
		var ok bool
		for _, p := range strings.Split(filepath.ToSlash(name), "/") {
			if node, ok = node[p]; !ok {
				return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
			}
		}
	}

	if len(node) == 0 {
		return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
	}

	rv := make([]string, 0, len(node))
	for name := range node {
		rv = append(rv, name)
	}

	return rv, nil
}

type bintree map[string]bintree

var _bintree = bintree{
	"meta": bintree{
		"hub-well-known-parameters.yaml": bintree{},
	},
	"src": bintree{
		"hub": bintree{
			"api": bintree{
				"requests": bintree{
					"aks-adapter-instance.json.template":       bintree{},
					"aks-adapter-template.json.template":       bintree{},
					"eks-adapter-instance.json.template":       bintree{},
					"eks-adapter-template.json.template":       bintree{},
					"gke-adapter-instance.json.template":       bintree{},
					"gke-adapter-template.json.template":       bintree{},
					"hybrid-adapter-instance.json.template":    bintree{},
					"hybrid-adapter-template.json.template":    bintree{},
					"k8s-aws-adapter-instance.json.template":   bintree{},
					"k8s-aws-adapter-template.json.template":   bintree{},
					"metal-adapter-instance.json.template":     bintree{},
					"metal-adapter-template.json.template":     bintree{},
					"openshift-adapter-instance.json.template": bintree{},
					"openshift-adapter-template.json.template": bintree{},
				},
			},
			"initialize": bintree{
				"hub-component.yaml.template": bintree{},
				"hub.yaml.template":           bintree{},
			},
		},
	},
}
