// Code generated by go-bindata. DO NOT EDIT.
//  memcopy: true
//  compress: true
//  decompress: once
//  metadata: true
//  asset-dir: true
//  restore: true
// sources:
//  meta/hub-well-known-parameters.yaml
//  src/hub/api/requests/eks-adapter-instance.json.template
//  src/hub/api/requests/eks-adapter-template.json.template
//  src/hub/api/requests/k8s-aws-adapter-instance.json.template
//  src/hub/api/requests/k8s-aws-adapter-template.json.template
//  src/hub/api/requests/metal-adapter-instance.json.template
//  src/hub/api/requests/metal-adapter-template.json.template
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
		mode: 0644,
		time: time.Unix(1540400209, 965820203),
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
		mode: 0644,
		time: time.Unix(1554231556, 763394984),
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
		mode: 0644,
		time: time.Unix(1554231681, 129530988),
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
		mode: 0644,
		time: time.Unix(1554231572, 476675267),
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
		mode: 0644,
		time: time.Unix(1554231961, 439226614),
	},
	"src/hub/api/requests/metal-adapter-instance.json.template": &asset{
		name: "metal-adapter-instance.json.template",
		data: "" +
			"\xac\xce\x41\x4b\x03\x31\x10\x05\xe0\x7b\x7f\x45\x08\x1e\x6b\x7e\x40\xc1\x93\x27\xf1\xe2\xc5\x93" +
			"\xf4\x30\xdd\x3c\x97\x60\x32\x19\x26\xb3\x45\x29\xfb\xdf\xc5\xdd\x15\x14\x16\xb4\xb4\xb7\xf0\x5e" +
			"\xf2\xbd\x9c\x36\xce\x39\xe7\x99\x0a\xfc\xce\xf9\x9b\xd3\xd7\x69\xf4\xdb\x39\x36\x14\xc9\x64\x53" +
			"\xf5\x9d\x81\x8f\x49\x2b\x17\xb0\xfd\x8c\x8d\xfa\xe6\x77\xee\xc5\x53\x24\x31\xe8\x5d\x81\x51\xf6" +
			"\xfb\xa5\x16\x52\x2a\x30\xe8\x74\x69\x5e\xfd\xb5\x1c\xb9\x85\x58\x0b\x25\x5e\xc4\xa9\x3d\x52\x1e" +
			"\xe6\xf5\x29\x1b\xb7\x6e\xe5\xe9\xdb\x70\x80\x32\x0c\x2d\x90\xa4\x00\x8e\x52\x13\xdb\xa5\x4e\x47" +
			"\xf7\xd0\xcb\x95\x9c\xc0\x76\x3d\xe9\x11\x1f\x67\x43\x5d\x2d\x52\x19\x6c\x21\x71\xaf\x68\x2d\x34" +
			"\x23\x4b\xdd\x83\x5c\x81\x1a\x34\x3f\x29\x5e\xd3\xfb\xaa\x45\x22\x67\xfe\xac\xd5\xe7\xbf\xc4\xf6" +
			"\x3f\xb2\x83\xda\x6d\x21\xa6\x1e\x1a\xc0\x74\xc8\x88\xab\xa4\xe9\x80\x85\xdc\x6f\xc6\xcd\x67\x00" +
			"\x00\x00\xff\xff",
		size: 788,
		mode: 0644,
		time: time.Unix(1554231589, 634317422),
	},
	"src/hub/api/requests/metal-adapter-template.json.template": &asset{
		name: "metal-adapter-template.json.template",
		data: "" +
			"\x9c\x92\x31\x6b\xe3\x40\x10\x85\x7b\xff\x8a\x65\xb9\xd2\x16\x5c\x77\x18\xae\x48\x42\x8a\x90\xc6" +
			"\x4d\xaa\xe0\x62\xac\x1d\x9b\x45\xbb\xb3\xcb\xcc\x28\xd8\x18\xff\xf7\xb0\x5a\xc9\x24\x60\x92\x28" +
			"\x95\x06\xbd\xf7\x3e\x89\x99\x77\x5e\x18\x63\x8c\x25\x88\x68\xd7\xc6\xfe\x39\x97\xe9\x62\x97\xf5" +
			"\xb5\x43\x69\xd9\x67\xf5\x89\x8a\x7a\x0f\x8c\xab\x88\x0a\xc1\xdc\x39\xc8\x8a\x3c\x19\x45\xa1\xed" +
			"\x8a\x65\x50\x57\x50\xd5\xf5\xdf\x49\x6f\x53\xcc\x89\x90\x54\x1e\x09\x76\x01\x9d\x5d\x9b\x57\xdb" +
			"\xfd\x93\x8a\xb3\x4b\x63\xd5\x87\x50\x80\xc6\x2a\x03\xee\x7d\x57\x46\x87\xc7\xf2\x68\x91\x75\x15" +
			"\x81\xe0\x80\x6c\xb7\x23\xf3\x0d\x79\x27\x03\xc7\x61\x0e\xe9\x54\x8c\x3d\x8d\xf3\x64\x52\x38\x54" +
			"\xcf\xf8\x4b\xff\xeb\xf7\x26\x39\x03\x43\x44\x45\x1e\x4c\x75\x19\x9f\x16\xe2\x48\x1a\x97\x22\x78" +
			"\xb2\x83\x78\x59\x9a\x1b\xae\xae\xdf\x21\x13\x2a\x4a\x03\xd9\x37\x48\x2e\x27\x4f\x3a\x23\xd2\xc2" +
			"\x03\xf2\xac\x40\xf0\x48\xfa\xab\xd0\x33\x9e\xbe\xca\x5c\x6f\xd5\x78\x3a\x30\x8a\x34\xa2\xa0\xbe" +
			"\x7d\xca\xf3\x52\x3d\x87\x0d\xe3\xde\x1f\xc7\x12\xd4\xa3\x41\xe8\x07\x3f\xe4\x99\x38\x91\xf4\xf2" +
			"\x1d\x51\x7e\x86\xfc\x58\xa7\x06\xc7\x42\xde\x42\x2a\xf7\x38\x22\xaf\x85\x42\x88\xb2\x41\x8e\x5e" +
			"\xc4\x27\x1a\x7a\xb3\x5d\x5c\x16\xef\x01\x00\x00\xff\xff",
		size: 841,
		mode: 0644,
		time: time.Unix(1554231854, 163885572),
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
		mode: 0644,
		time: time.Unix(1515794706, 303261491),
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
		mode: 0644,
		time: time.Unix(1515795479, 848753595),
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
					"eks-adapter-instance.json.template":     bintree{},
					"eks-adapter-template.json.template":     bintree{},
					"k8s-aws-adapter-instance.json.template": bintree{},
					"k8s-aws-adapter-template.json.template": bintree{},
					"metal-adapter-instance.json.template":   bintree{},
					"metal-adapter-template.json.template":   bintree{},
				},
			},
			"initialize": bintree{
				"hub-component.yaml.template": bintree{},
				"hub.yaml.template":           bintree{},
			},
		},
	},
}
