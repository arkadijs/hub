// Copyright (c) 2022 EPAM Systems, Inc.
// 
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package initialize

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/agilestacks/hub/cmd/hub/bindata"
	"github.com/agilestacks/hub/cmd/hub/config"
	"github.com/agilestacks/hub/cmd/hub/util"
)

const initializeTemplates = "cmd/hub/initialize"

func InitStack(manifestDir string) {
	initManifest(manifestDir, "hub.yaml")
}

func InitComponent(manifestDir string) {
	initManifest(manifestDir, "hub-component.yaml")
}

func initManifest(dir string, template string) {
	file, manifest, project := initFile(dir, template)
	bytes := bindata.MustAsset(fmt.Sprintf("%s/%s.template", initializeTemplates, template))
	yaml := string(bytes)
	yaml = strings.Replace(yaml, "${project}", project, -1)

	wrote, err := strings.NewReader(yaml).WriteTo(file)
	if err != nil || wrote != int64(len(yaml)) {
		file.Close()
		log.Fatalf("Unable to save `%s`; wrote %d bytes: %v", manifest, wrote, err)
	}
	file.Close()
	log.Printf("Wrote `%s` for project %s", manifest, project)
}

func initFile(dir string, template string) (*os.File, string, string) {
	const defaultModeDir = 0775
	const defaultModeFile = 0664

	info, err := os.Stat(dir)
	if err != nil {
		if !util.NoSuchFile(err) {
			log.Fatalf("Unable to init in `%s` directory: %v", dir, err)
		}
		err = os.Mkdir(dir, defaultModeDir)
		if err != nil {
			log.Fatalf("Unable to init in `%s` directory: mkdir failed: %v", dir, err)
		}
	} else {
		if !info.IsDir() {
			log.Fatalf("Unable to init in `%s`: not a directory", dir)
		}
	}
	abs, err := filepath.Abs(dir)
	if err != nil {
		log.Fatalf("Unable to determine absolute path to `%s`: %v", dir, err)
	}
	project := filepath.Base(abs)

	manifest := filepath.Join(dir, template)
	info, err = os.Stat(manifest)
	if err != nil {
		if !util.NoSuchFile(err) {
			log.Fatalf("Unable to init `%s`: %v", manifest, err)
		}
	} else {
		if info.IsDir() {
			log.Fatalf("Unable to init `%s`: is a directory", manifest)
		}
		if !config.Force {
			log.Fatalf("`%s` exist, add --force to override", manifest)
		}
	}
	file, err := os.OpenFile(manifest, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, defaultModeFile)
	if err != nil {
		log.Fatalf("Unable to init `%s`: %v", manifest, err)
	}

	return file, manifest, project
}
