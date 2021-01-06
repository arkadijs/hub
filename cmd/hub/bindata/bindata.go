// Code generated by go-bindata. DO NOT EDIT.
//  memcopy: true
//  compress: true
//  decompress: once
//  metadata: true
//  asset-dir: true
//  restore: true
// sources:
//  meta/hub-well-known-parameters.yaml
//  meta/manifest.schema.json
//  cmd/hub/api/requests/aks-adapter-instance.json.template
//  cmd/hub/api/requests/aks-adapter-template.json.template
//  cmd/hub/api/requests/eks-adapter-instance.json.template
//  cmd/hub/api/requests/eks-adapter-template.json.template
//  cmd/hub/api/requests/eks-cluster-instance.json.template
//  cmd/hub/api/requests/eks-cluster-template.json.template
//  cmd/hub/api/requests/gke-adapter-instance.json.template
//  cmd/hub/api/requests/gke-adapter-template.json.template
//  cmd/hub/api/requests/hybrid-adapter-instance.json.template
//  cmd/hub/api/requests/hybrid-adapter-template.json.template
//  cmd/hub/api/requests/k8s-aws-adapter-instance.json.template
//  cmd/hub/api/requests/k8s-aws-adapter-template.json.template
//  cmd/hub/api/requests/metal-adapter-instance.json.template
//  cmd/hub/api/requests/metal-adapter-template.json.template
//  cmd/hub/api/requests/openshift-adapter-instance.json.template
//  cmd/hub/api/requests/openshift-adapter-template.json.template
//  cmd/hub/initialize/hub.yaml.template
//  cmd/hub/initialize/hub-component.yaml.template

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
	"meta/manifest.schema.json": &asset{
		name: "manifest.schema.json",
		data: "" +
			"\xec\x5a\x49\x6f\x2b\x37\x0c\xbe\xfb\x57\x18\xd3\x1e\xf3\xe2\xf6\x54\x20\xd7\x2e\xb7\x02\x05\x5a" +
			"\xf4\xf2\xe0\x83\x46\xc3\xb1\xf5\xa2\xad\x94\xe4\x3e\xa3\xc8\x7f\x2f\xbc\xc4\x8e\x6d\x69\x44\x65" +
			"\x66\x50\xa7\x76\x0e\x01\x2c\x51\x14\x45\x91\x9f\xb8\xcc\x3f\x93\xe9\x74\x3a\xad\xbe\x15\x4d\xf5" +
			"\x34\xad\x5c\xb0\x80\xcb\x50\x3f\x0a\x33\x53\x4c\x8b\x16\x9c\x7f\x74\x7c\x09\x8a\x3d\x7e\x71\x46" +
			"\x57\x0f\x7b\xf2\xdd\xd8\x66\xc9\xd2\x7b\xfb\x34\x9b\x6d\x66\x3f\xed\x29\x0d\x2e\x66\x0d\xb2\xd6" +
			"\x7f\xfa\xee\x87\xd9\x6e\xec\x9b\xd7\x95\x5e\x78\x09\x9b\x75\xbf\xee\xd9\x1f\x26\xd6\x76\x3b\x6e" +
			"\xea\x2f\xc0\x0f\xa3\xac\x69\x84\x17\x46\x33\xf9\x1b\x1a\x0b\xe8\x05\xb8\xea\x69\xda\x32\xe9\x60" +
			"\x4f\x62\xdf\x4e\xec\x8e\xb3\x1d\x5f\x01\x3a\x61\xf4\xc9\xe0\x76\x02\x74\x50\xd5\xd3\xf4\xf3\xc9" +
			"\xe8\xe6\xef\xfb\x93\x91\xf9\xe1\xd7\xcb\xc3\x91\xeb\xb3\xd0\x4d\x01\xcb\xca\x79\xc6\x9f\xab\x87" +
			"\xcb\x09\x66\xad\x14\x9c\x6d\x0e\x17\x9b\xe6\x46\x59\xa3\x41\xfb\xd8\xa4\x65\xc8\x14\x78\x40\x57" +
			"\x11\x44\x56\xe0\xd9\xa5\xc8\x51\x7d\x1f\x66\x11\xfe\x0a\x02\xa1\x89\x1f\x4a\x33\x05\x67\x3b\x9f" +
			"\xad\x4f\x5c\xca\x29\x87\xd8\xcc\x89\x6c\xce\xa3\xd0\x8b\xea\x82\xe8\x25\xa2\x93\x16\x8d\xfa\x7d" +
			"\xab\xec\x41\xd9\xbe\xda\xeb\x80\x2c\x6b\x14\xd0\x0e\xcb\xb2\x01\xc7\x51\x58\x1f\xb3\xf7\x5e\x8c" +
			"\x39\xf3\xb0\x30\xb8\x1e\x96\x6b\xca\x35\xcf\x99\x7e\x8e\xce\xee\xfd\x6a\xbb\xdd\x43\x9a\x42\x07" +
			"\x55\x03\x56\x51\x82\x39\x49\x4c\xc5\x7c\x40\xe1\x3b\x0e\x9f\xf4\xfb\x03\xc5\x82\x75\xc9\x58\x6f" +
			"\x5c\xb3\x63\x9e\x49\xbb\x64\x7d\x8e\x20\x05\x07\xed\x06\x36\x60\x67\x02\x72\x02\xcf\x3d\xb4\x5c" +
			"\xf2\x9c\xc4\x7f\xbd\x05\xad\x03\xfe\xb9\x34\x74\x31\x44\xb6\x3e\x47\x2e\xe1\x41\x25\x40\xa7\x13" +
			"\xf2\xf2\xb0\x77\x04\xae\xf8\x8d\xbd\xea\xe5\x62\x72\x1e\x83\xf0\x6e\x80\xcc\x83\x24\xe9\xf6\x12" +
			"\x37\xb8\xc7\x0c\x0b\xba\x71\xa4\x0d\xd2\x06\xbe\x33\xd3\xc8\x45\x44\xfc\x51\xca\x2a\x49\x32\xef" +
			"\xf0\x82\xf4\x95\x16\x2b\xe3\xd2\xfc\x72\x6a\xca\x18\x3b\xcd\xb0\x4a\xae\xfd\x78\x41\x02\xb3\x44" +
			"\x45\x07\xef\x38\xe6\x11\xaf\x84\x2f\xdb\x34\x7b\x66\xba\x73\x45\x56\x28\xe3\xa1\xca\x12\xcf\x09" +
			"\xbb\x17\x68\xfe\x7c\x7f\x2a\x7d\xf1\x65\x10\x2f\xe5\x8d\x3c\xed\xf5\x08\xe3\x42\xfd\x13\xd1\x42" +
			"\xfb\xc9\x33\xe9\x47\xf1\x52\x0a\x04\xef\x7a\xaf\xf6\xc6\x9d\x7e\xad\x22\x61\x74\x0a\x36\x23\x50" +
			"\x39\x2f\x7f\xe5\x62\x3a\x8e\xcb\x6e\xd1\xac\x44\xf3\x31\x65\x97\xa2\x05\xbe\xe6\x91\x18\x3d\x93" +
			"\xe1\xe4\x33\x4b\x72\x32\x53\x33\x84\x3e\x31\x22\x93\xd2\xfc\xdd\x27\xca\x5b\x01\xd6\xae\x4f\x34" +
			"\x9d\x7b\xc0\x3b\x1e\xef\x04\xf6\x12\x1e\x6d\x5a\xf4\x42\x52\x80\xc1\x06\xf0\xa6\x15\x60\x77\xb6" +
			"\x7c\xcb\x3a\x50\x4c\x37\xcc\x53\x92\xd5\xff\xb1\x12\x92\xef\x10\x3d\x3b\x8b\x72\x65\xcd\xfa\x47" +
			"\xa3\x77\x90\xf9\x9f\x41\x4d\x81\x2b\xb8\x71\xd3\x53\xf8\xea\x41\xbb\xe8\x3e\x99\x77\x27\xf7\x98" +
			"\x08\xcd\x65\x68\xe0\x96\x6d\x98\x1b\xdd\x8a\x45\xc0\x9b\x56\x42\x03\x56\x9a\x35\xd9\x8a\x13\x22" +
			"\x13\xb3\x9e\xaa\x86\xd6\x20\x90\x73\xec\x7c\xfe\x46\xaa\x09\x10\xea\x02\x84\xf4\x8e\x58\x1f\x28" +
			"\x4f\x95\x27\xef\x48\x91\x2a\xd6\x7a\xc0\xbb\x22\x07\xce\xc0\x22\x0e\x12\xf4\xdd\x45\xee\x2e\x72" +
			"\x77\x91\x82\xa8\xe5\x4d\xdf\xf0\x6a\x52\xfd\x74\x31\x9e\x9e\xa1\x97\xd4\xee\x3f\x48\x75\xfe\xd8" +
			"\xff\x1d\x6d\x8b\x15\x93\x61\x7b\x82\x74\x87\xa0\x65\x41\xfa\x2e\x12\x50\xb6\xa3\x3f\x46\xab\x7f" +
			"\xe4\x6a\x20\xf1\xd0\xbf\xeb\x64\xd1\x26\xfd\x3b\x84\x0a\x0e\x30\xd7\xd6\xf0\xc0\x97\x39\x1a\x29" +
			"\xf4\xf3\x50\x67\xeb\xee\x1e\xf7\x36\x8a\x16\x8d\xfa\x59\xaf\xc6\xdd\xe0\x17\x21\x47\x74\x1d\x18" +
			"\x53\xfc\x0e\x08\x2d\x7f\xbc\xae\xac\x71\x46\x6a\xe5\x94\xb5\x71\x12\x98\x5b\xf2\xe2\x16\x36\x6e" +
			"\xf2\xb8\xdc\xab\x23\x71\x0d\xbd\x06\x0f\xca\x4a\xe6\x61\xf0\xda\x43\x27\x6e\x12\x0a\xd9\x3c\xa0" +
			"\xec\xcc\xb9\x55\x70\x9e\xf1\x25\x74\xd1\x2c\x4c\x9f\xea\x4f\x2b\x24\xdc\x74\x29\xbc\x11\x08\xdc" +
			"\x1b\x14\xb7\xad\x06\xf8\xea\x91\x7d\x5c\x05\x0c\xfb\x25\x41\x36\x1e\x2a\x8b\x8b\xc8\xfe\x5e\xea" +
			"\xfb\x24\x1c\xc8\x87\x4e\x84\x1c\x94\x80\x15\xef\xcb\x45\x8b\xf2\x51\x62\x4e\x3a\x25\x7e\xe2\x40" +
			"\xcf\x4d\xc7\x78\xfa\x1e\xb2\x5f\xb3\x64\x71\xe9\xae\xf2\x2b\x8c\x36\x4c\xf0\x36\xf8\x7b\xc5\xe0" +
			"\xda\x2b\x06\x87\x74\x7e\x1c\xf6\xf9\x9c\xba\x6f\x56\xf8\x47\xfb\x27\xc3\xf1\xb6\x18\x39\x71\xa6" +
			"\x7c\x82\xdd\x27\xc6\x49\x79\xea\x64\xf7\xff\x65\xf2\x6f\x00\x00\x00\xff\xff",
		size: 12746,
		mode: 0644,
		time: time.Unix(1609961974, 348607212),
	},
	"cmd/hub/api/requests/aks-adapter-instance.json.template": &asset{
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
		mode: 0644,
		time: time.Unix(1556808588, 764066530),
	},
	"cmd/hub/api/requests/aks-adapter-template.json.template": &asset{
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
		mode: 0644,
		time: time.Unix(1556808625, 516396344),
	},
	"cmd/hub/api/requests/eks-adapter-instance.json.template": &asset{
		name: "eks-adapter-instance.json.template",
		data: "" +
			"\xa4\xcf\x41\x6a\xf3\x30\x10\x05\xe0\x7d\x4e\x21\xc4\xbf\xcc\xaf\x03\x04\xba\xea\x05\xba\xe9\xaa" +
			"\x64\x31\xb1\x5e\x8d\x88\x34\x12\x33\xa3\x50\x08\xbe\x7b\xa9\x9d\x42\x5b\x4c\x5b\xd3\x9d\x79\x6f" +
			"\xfc\x3d\x74\xdd\x39\xe7\x9c\x67\x2a\xf0\x07\xe7\xff\x5d\xdf\xbe\x26\xbf\x5f\x62\x43\x69\x99\x6c" +
			"\xae\xde\x33\xf0\x25\x49\xe5\x02\xb6\x8f\xb1\xd1\xa8\xfe\xe0\x9e\x3c\x45\x6a\x06\xb9\xc3\x59\xfd" +
			"\xf1\x56\x36\x12\x2a\x30\xc8\x7c\xb2\x6c\x7e\xda\x8d\xac\x21\xd6\x42\x89\x6f\xde\xdc\x5e\x28\xf7" +
			"\x65\x7b\xce\xa6\xbd\x5b\xf9\x75\xc8\xb5\xc7\x20\x18\x53\xe5\xef\xee\xce\xfd\x04\x61\x18\x34\x50" +
			"\x4b\x01\x1c\x5b\x4d\x6c\x9b\xf7\xbe\x38\x03\xdd\x43\xfe\xa4\xe0\xac\x61\xc8\x5d\x0d\xb2\xfd\xf1" +
			"\xb5\xb4\xca\x60\x0b\x89\x47\x81\x6a\xe8\x92\x1f\x04\xcf\xe9\x65\xd5\xa2\xd6\xb6\x71\xaa\xf5\xf1" +
			"\x27\x51\x7f\x47\x0e\x10\xfb\x5f\x88\x69\x84\x04\x30\x9d\x32\xe2\x1a\x69\xd2\xb1\x80\xc7\xdd\xb4" +
			"\x7b\x0d\x00\x00\xff\xff",
		size: 673,
		mode: 0644,
		time: time.Unix(1582724150, 992728231),
	},
	"cmd/hub/api/requests/eks-adapter-template.json.template": &asset{
		name: "eks-adapter-template.json.template",
		data: "" +
			"\x94\x92\xbf\x6a\x33\x31\x10\xc4\x7b\x3f\x85\x10\x5f\xe9\x3b\xf8\xba\x60\x48\x11\x82\xab\x34\x86" +
			"\x90\x2a\xb8\x58\x9f\xc6\x46\x9c\x6e\x25\x76\x75\xc1\xc1\xf8\xdd\x83\xee\x8f\x49\xc0\x24\x76\xa5" +
			"\x45\x33\xfb\x5b\xb1\xa3\xd3\xc2\x18\x63\x2c\x53\x07\xbb\x32\xf6\xdf\xa9\x54\x67\xbb\x1c\xaf\x1d" +
			"\xb4\x11\x9f\xb2\x8f\x5c\xd4\xf5\xcb\xab\x79\x72\x94\x32\x64\x76\x68\xa6\xa6\x2d\x1a\x5a\xad\x68" +
			"\xd4\x56\xff\x67\xb5\x89\x5d\x8a\x0c\xce\xba\x66\xda\x05\x38\xbb\x32\xef\xb6\x7d\xd0\x0a\xad\xda" +
			"\xa5\xb1\xd9\x87\x50\x60\xc6\xe2\x98\x21\x4c\xa1\x72\x3c\x2a\x42\xd8\xfb\xb6\x94\x0e\xc7\x72\x34" +
			"\x90\x5c\x75\xc4\x74\x80\xd8\xed\x34\xe1\x03\xb2\xd3\x81\xea\x90\x42\xfc\x2c\xc6\x9e\xa7\x7a\x36" +
			"\x65\x3a\x8c\x9e\xe9\x81\x8f\x65\xfa\x2c\x26\x12\xea\x90\x21\x83\x65\xdc\xc7\x8f\x9d\x38\xd6\xda" +
			"\xc5\x8e\x3c\xdb\x41\x3c\x2f\xcd\x15\x57\xdb\xef\x20\x8c\x0c\xad\x29\xf9\x1a\xec\x52\xf4\x9c\xef" +
			"\x68\x69\xe8\x19\x72\x6b\x03\x5a\xad\x9b\xd0\x6b\x89\xe2\x97\x8e\x4b\x00\xb5\xe7\x83\x40\xb5\xee" +
			"\x25\x6c\x04\x7b\x7f\x9c\x32\x1a\xb7\x48\xa1\x1f\xfc\x94\xd2\x7d\x38\xd5\xf8\xf6\x17\x51\x6f\x43" +
			"\x7e\xcf\xb7\xc6\xf4\x5f\xae\x21\xb3\xf4\x98\x90\x97\x84\x41\x9d\x6e\x20\x9d\x57\xf5\x91\x87\x28" +
			"\xb7\x8b\xf3\xe2\x2b\x00\x00\xff\xff",
		size: 735,
		mode: 0644,
		time: time.Unix(1593716556, 553915201),
	},
	"cmd/hub/api/requests/eks-cluster-instance.json.template": &asset{
		name: "eks-cluster-instance.json.template",
		data: "" +
			"\xac\xd2\x3d\xae\xdb\x30\x0c\x07\xf0\x3d\xa7\x30\x84\x8e\xb1\x0e\x10\xa0\x53\x2f\x90\xa5\x4b\x8b" +
			"\x0c\xb4\xc4\xb8\x82\x25\xd2\x20\x25\x37\x69\x90\xbb\x17\xb1\xd3\x8f\x07\x18\xef\x25\x81\x36\xe3" +
			"\x4f\xf3\x47\x98\xf4\x65\xd3\x34\x4d\x63\x08\x12\x9a\x5d\x63\x3e\x5d\x6e\x4f\x57\xb3\x5d\xe2\x8c" +
			"\x69\x8c\x90\xe7\xd2\x9f\x0c\x69\x0a\xc2\x94\x90\xf2\xff\x71\x86\x5e\xcd\xae\xf9\x6e\x6e\x0d\x47" +
			"\x96\xf4\x19\x07\x35\x87\x7b\x75\x04\x81\x84\x19\x65\x7e\x67\x19\xfa\x66\xb0\x27\xb5\x9e\x13\x04" +
			"\xba\x83\x73\x75\x82\x58\x96\xe1\x73\x76\xdd\x36\x2b\xad\x2e\x72\xf1\x56\xb0\x0f\xfc\x6a\x33\x4c" +
			"\x10\x22\x74\x21\x86\x7c\xfe\xc6\x84\xfa\xbc\xc3\x69\x64\x42\xca\x76\x28\x1d\x0a\x61\x46\xb5\x3f" +
			"\x59\x06\x14\xab\xe1\x17\x56\x05\x1d\x17\xca\x55\xc5\x04\xa7\x2f\xd5\x51\x1d\x39\xef\x25\xb8\xba" +
			"\x1f\x3f\x71\x2c\x09\x2b\x2e\x15\x07\xb5\x2e\x16\xcd\x28\xf5\x40\xf0\xe9\x95\x5f\xf9\x2f\x17\xa8" +
			"\x17\x54\xb5\x45\xe2\x5e\xf0\x18\x4e\xab\x16\x8c\xe3\x73\x9c\x2a\x7f\xfd\x48\xd4\xc7\x48\x70\xc9" +
			"\x22\x41\x17\xd1\xaf\x49\x47\x88\x8a\x0f\x41\x0e\x25\xb7\x09\x08\x7a\x94\xf7\xc4\x2c\xe5\x41\x70" +
			"\x39\x66\x0b\x25\xb3\x3a\x88\x95\xd8\x7f\x17\x6e\x3d\xe8\x8f\x8e\x41\x7c\x9d\x0d\xac\xca\xd2\x81" +
			"\xb3\x43\xa0\x55\xdb\x08\x82\x6f\x99\xe2\xf9\x7e\xac\xc3\xe6\xba\xf9\x1d\x00\x00\xff\xff",
		size: 1484,
		mode: 0644,
		time: time.Unix(1582734387, 457518969),
	},
	"cmd/hub/api/requests/eks-cluster-template.json.template": &asset{
		name: "eks-cluster-template.json.template",
		data: "" +
			"\x94\x91\x41\x6b\xb4\x30\x10\x86\xef\xfb\x2b\x42\xf8\x8e\x2a\x7c\xb7\x22\xf4\x54\xf6\xd4\xcb\x42" +
			"\xe9\xa9\xec\x61\xd6\x8c\x12\x4c\x26\x61\x26\x16\xcb\xe2\x7f\x2f\x31\xba\xb4\xb0\xd0\xf6\xe4\x30" +
			"\xef\xc3\x33\xfa\x7a\x3d\x28\xa5\x94\x26\xf0\xa8\x5b\xa5\xff\x5d\xf3\xb4\xe8\xaa\xac\x0d\x4a\xc7" +
			"\x36\x26\x1b\x28\xa7\xc7\xe7\x17\xf5\xe4\x26\x49\xc8\x3b\x21\x09\xba\x31\x67\x38\x4a\xfb\x7f\xdf" +
			"\x76\xc1\xc7\x40\x48\x49\x8e\x04\x17\x87\x46\xb7\xea\xad\xb0\xf5\xf8\x20\x35\x8e\xa2\x2b\xa5\x93" +
			"\x75\x2e\xab\x94\xc6\x39\x21\x13\xb8\xda\x50\x49\x18\xb0\xb7\x63\x1e\x0d\xce\xf9\xd1\x21\xa7\xda" +
			"\x03\xc1\x80\xac\xcf\xdb\x9d\x77\xe4\x8b\xac\x6e\x83\xd1\x85\x8f\x0c\x4e\xb4\xcd\x3b\x94\x60\x28" +
			"\x4c\x74\x90\xfa\xc0\xfe\x31\x9f\xdf\xd3\x08\x0c\x1e\x13\xf2\xca\x94\x3a\xbe\x55\x62\x48\x1a\x13" +
			"\x3c\x58\xd2\x6b\xb8\x54\xea\x0e\x75\xfb\xe2\xc6\xd2\xc0\x28\xd2\x4c\xec\x4e\x8c\xbd\x9d\xb7\x52" +
			"\xca\x0b\x83\x9b\x56\x1e\x62\xfc\x9b\x4e\x24\xbc\xfe\x64\x94\xdf\x29\xbf\x56\xd9\xe0\xf6\x83\xee" +
			"\x28\x13\x4f\x58\x84\xb7\x2a\x11\xbc\x9c\x90\xbd\x15\xb1\x81\xd6\xca\xce\x87\xe5\xf0\x19\x00\x00" +
			"\xff\xff",
		size: 582,
		mode: 0644,
		time: time.Unix(1593716564, 42266828),
	},
	"cmd/hub/api/requests/gke-adapter-instance.json.template": &asset{
		name: "gke-adapter-instance.json.template",
		data: "" +
			"\x9c\x8e\xc1\x4a\xc4\x30\x14\x45\xf7\xf3\x15\x21\xb8\x1c\xf3\x01\x03\xfe\x83\x1b\x37\xca\x2c\x5e" +
			"\xdb\x6b\x08\x4d\x5e\xc2\xcb\x4b\x51\x86\xfe\xbb\xd8\x56\x50\x28\x68\xdd\x85\x73\x73\xcf\x7d\xb7" +
			"\x93\x31\xc6\x58\xa6\x04\x7b\x31\xf6\xee\xf6\xf9\x9a\xed\x79\xc5\x8a\x54\x22\xe9\x12\x7d\x31\xf0" +
			"\x14\x24\x73\x02\xeb\x77\xac\xe4\xab\xbd\x98\x17\x4b\x03\x15\x85\x3c\xf8\x11\xf6\xba\x85\x85\x84" +
			"\x12\x14\xb2\x7c\x59\x37\x7f\xec\x0e\x5c\xdd\x90\x13\x05\xde\x7c\x4b\x3a\x51\x6c\xeb\xf6\xc2\xe6" +
			"\xb3\xd9\xa9\xf6\x31\xb7\xc1\x09\x7c\xc8\xff\x2d\xd3\x44\x21\x52\x17\x62\xd0\xf7\xe7\xcc\x38\xac" +
			"\x19\x5b\x07\x61\x28\xaa\xf3\x23\x5c\x1f\x5b\x55\xc8\xf1\x6b\x72\x2a\x99\xc1\xea\x02\x7b\x41\xad" +
			"\xae\x49\x7c\x14\xbc\x86\xb7\x5d\x17\x95\x72\x4c\x57\x6b\x7e\xfa\xcd\x58\xff\xa6\xec\x21\x7a\x9f" +
			"\x88\xc9\x43\x1c\x98\xba\x88\x61\x57\xa9\xd2\xb0\x29\xaf\xa7\xf9\xf4\x11\x00\x00\xff\xff",
		size: 625,
		mode: 0644,
		time: time.Unix(1591633666, 277682918),
	},
	"cmd/hub/api/requests/gke-adapter-template.json.template": &asset{
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
		mode: 0644,
		time: time.Unix(1556807810, 675331929),
	},
	"cmd/hub/api/requests/hybrid-adapter-instance.json.template": &asset{
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
		mode: 0644,
		time: time.Unix(1567628018, 941528099),
	},
	"cmd/hub/api/requests/hybrid-adapter-template.json.template": &asset{
		name: "hybrid-adapter-template.json.template",
		data: "" +
			"\x9c\x92\x4f\x8b\xdb\x30\x10\xc5\xef\xf9\x14\x42\xf4\x18\x1b\x7a\x2b\x81\x1e\xda\x50\x48\xe9\x25" +
			"\x97\x9e\x96\x1c\xc6\xd6\x24\x11\x96\x46\x62\x66\xbc\x9b\x10\xf2\xdd\x17\xff\x0b\xbb\x10\x76\xd7" +
			"\x39\xf9\xe1\x79\xbf\x27\xa1\x79\x97\x85\x31\xc6\x58\x82\x88\x76\x65\xec\xb7\x4b\xa7\xae\x76\x39" +
			"\xfc\x76\x28\x35\xfb\xac\x3e\x51\x37\xdd\x9c\x2b\xf6\xce\xfc\x72\x90\x15\xd9\xec\x13\x9b\xdf\xc0" +
			"\x58\x44\x54\x08\xe6\xc5\xeb\xd1\xac\x43\x6a\xdd\x84\x8b\x42\xdd\x74\xe0\xb1\x07\x0b\x18\xc0\xd5" +
			"\xf7\xc9\x50\xa7\x98\x13\x21\xa9\xfc\x21\xa8\x02\x3a\xbb\x32\x4f\xb6\xf9\x21\xc5\x40\xd8\xa5\xb1" +
			"\xea\x43\x40\xee\x15\x03\xee\x7d\xd3\x49\x87\xa7\xee\x53\x23\x6b\x11\x81\xe0\x80\x6c\x77\x63\xe8" +
			"\x33\x72\x25\x7d\x90\xc3\x1c\xd2\xb9\x33\xb6\x34\xea\xc9\xa4\x70\x18\x3c\xe3\x9d\x7e\x8e\x07\x4e" +
			"\xf3\x0c\x0c\x11\x15\xb9\x77\x0d\xaf\xf4\xee\xa5\x1c\x49\xe9\x52\x04\x4f\xb6\x1f\x5e\x97\xe6\x8e" +
			"\xab\x69\x2b\x64\x42\x45\x29\x21\xfb\x12\xc9\xe5\xe4\x49\x67\x20\x35\xac\x91\x67\x01\xc1\x23\xe9" +
			"\x43\xd0\x3f\x3c\x7f\xc4\xdc\xb6\x55\x7a\x3a\x30\x8a\x94\xa2\xa0\xbe\xfe\x9b\x1f\xa1\x36\x49\x74" +
			"\x1e\xd7\x72\xd8\x32\xee\xfd\x69\xac\xcf\xb0\x6d\x08\x6d\xef\x87\x3c\xf7\x1a\x92\xfe\x7f\x96\x28" +
			"\x5f\x8b\x7c\xdb\xc3\x12\xc7\x2a\xdf\x8b\x54\x6e\x71\x8c\xbc\x35\x11\x21\xca\x16\x39\x7a\x11\x9f" +
			"\xa8\xef\xdb\x6e\x71\x5d\xbc\x06\x00\x00\xff\xff",
		size: 922,
		mode: 0644,
		time: time.Unix(1571170007, 550590129),
	},
	"cmd/hub/api/requests/k8s-aws-adapter-instance.json.template": &asset{
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
	"cmd/hub/api/requests/k8s-aws-adapter-template.json.template": &asset{
		name: "k8s-aws-adapter-template.json.template",
		data: "" +
			"\x9c\x92\x31\x6b\xc3\x30\x10\x85\xf7\xfc\x0a\x21\x0a\x5d\x12\x43\xb7\x10\xe8\x10\x4a\xa7\x2c\x81" +
			"\x50\x3a\x94\x0c\x17\xeb\x62\x84\xa5\x93\xb8\x93\xdb\x84\x90\xff\x5e\x64\xcb\xa1\x85\xd0\x36\x9d" +
			"\x7c\xf8\x7d\xef\xf9\xb8\xe7\xd3\x44\x29\xa5\x34\x81\x47\xbd\x50\xfa\xee\x94\xa7\xb3\x9e\x0e\xaf" +
			"\x0d\x4a\xcd\x36\x26\x1b\x28\xab\xcb\xc6\x3a\xdc\x24\xa8\xdb\x7b\x51\xab\xf9\x46\x05\x52\xcb\xd7" +
			"\x8d\x02\x03\x31\x21\x8f\x2e\xc9\x44\xe6\xdb\xb9\xcc\x8a\xb6\x78\x18\xd5\x3a\xf8\x18\x08\x29\xc9" +
			"\x33\xc1\xce\xa1\xd1\x0b\xf5\x36\xa0\x1f\xa2\xa7\x4a\x27\xeb\x5c\x0e\x53\x1a\x0f\x09\x99\xc0\xcd" +
			"\x0c\x0d\x0a\x03\xee\x6d\x9b\x47\x83\x87\xfc\xa8\x91\xd3\xcc\x03\x41\x83\xac\xb7\xe5\x0b\xef\xc8" +
			"\x3b\xe9\x53\x0d\x46\x17\x8e\x19\xec\xa8\xcc\x23\x94\xa0\x19\x98\xb2\xe0\xe3\xb8\xc1\x08\x44\x60" +
			"\xf0\x98\x90\x7b\x6c\xb8\xd3\xb7\x5b\x19\x92\xca\x04\x0f\x96\x74\x2f\x9e\xa7\xea\x0a\xd5\x76\x3b" +
			"\x64\xc2\x84\x52\x41\xb4\x55\x0d\x4f\xc8\xe9\x16\x83\xb3\x48\xe9\x5f\xa6\x15\x1e\x7f\xf2\x5c\x8a" +
			"\xa8\x2c\x35\x8c\x22\x55\xc7\x6e\xcd\xb8\xb7\x87\xd2\xd5\x70\x4d\x70\x5d\xcf\x43\x8c\xb7\xc5\x89" +
			"\x84\x97\xdf\x12\xe5\x6f\x91\x5f\x7b\xae\xb0\xfc\x37\xd7\x22\x13\x77\x58\x22\x2f\x4d\x23\x78\x59" +
			"\x23\x7b\x2b\x62\x03\xf5\x75\x6e\x27\xe7\xc9\x67\x00\x00\x00\xff\xff",
		size: 763,
		mode: 0644,
		time: time.Unix(1593716614, 388213101),
	},
	"cmd/hub/api/requests/metal-adapter-instance.json.template": &asset{
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
		mode: 0644,
		time: time.Unix(1566034466, 888047415),
	},
	"cmd/hub/api/requests/metal-adapter-template.json.template": &asset{
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
		mode: 0644,
		time: time.Unix(1566034452, 131327161),
	},
	"cmd/hub/api/requests/openshift-adapter-instance.json.template": &asset{
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
		mode: 0644,
		time: time.Unix(1554303190, 999879579),
	},
	"cmd/hub/api/requests/openshift-adapter-template.json.template": &asset{
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
		mode: 0644,
		time: time.Unix(1554303130, 902704954),
	},
	"cmd/hub/initialize/hub.yaml.template": &asset{
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
	"cmd/hub/initialize/hub-component.yaml.template": &asset{
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
	"cmd": bintree{
		"hub": bintree{
			"api": bintree{
				"requests": bintree{
					"aks-adapter-instance.json.template":       bintree{},
					"aks-adapter-template.json.template":       bintree{},
					"eks-adapter-instance.json.template":       bintree{},
					"eks-adapter-template.json.template":       bintree{},
					"eks-cluster-instance.json.template":       bintree{},
					"eks-cluster-template.json.template":       bintree{},
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
	"meta": bintree{
		"hub-well-known-parameters.yaml": bintree{},
		"manifest.schema.json":           bintree{},
	},
}
