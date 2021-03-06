// Copyright (c) 2021 The Srpmproc Authors
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package modulemd

import (
	"github.com/go-git/go-billy/v5"
	"gopkg.in/yaml.v2"
)

type ModuleMd struct {
	Document string `yaml:"document,omitempty"`
	Version  int    `yaml:"version,omitempty"`
	Data     struct {
		Name          string `yaml:"name,omitempty"`
		Stream        string `yaml:"stream,omitempty"`
		Version       int64  `yaml:"version,omitempty"`
		StaticContext bool   `yaml:"static_context,omitempty"`
		Context       string `yaml:"context,omitempty"`
		Arch          string `yaml:"arch,omitempty"`
		Summary       string `yaml:"summary,omitempty"`
		Description   string `yaml:"description,omitempty"`
		Servicelevels struct {
			Rawhide struct {
				Eol struct {
				} `yaml:"eol,omitempty"`
			} `yaml:"rawhide,omitempty"`
			StableAPI struct {
				Eol struct {
				} `yaml:"eol,omitempty"`
			} `yaml:"stable_api,omitempty"`
			BugFixes struct {
				Eol struct {
				} `yaml:"eol,omitempty"`
			} `yaml:"bug_fixes,omitempty"`
			SecurityFixes struct {
				Eol struct {
				} `yaml:"eol,omitempty"`
			} `yaml:"security_fixes,omitempty"`
		} `yaml:"servicelevels,omitempty"`
		License struct {
			Module  []string `yaml:"module,omitempty"`
			Content []string `yaml:"content,omitempty"`
		} `yaml:"license,omitempty"`
		Xmd          map[string]interface{} `yaml:"xmd,omitempty"`
		Dependencies []struct {
			Buildrequires map[string][]string `yaml:"buildrequires,omitempty,omitempty"`
			Requires      map[string][]string `yaml:"requires,omitempty,omitempty"`
		} `yaml:"dependencies,omitempty"`
		References struct {
			Community     string `yaml:"community,omitempty"`
			Documentation string `yaml:"documentation,omitempty"`
			Tracker       string `yaml:"tracker,omitempty"`
		} `yaml:"references,omitempty"`
		Profiles map[string]*struct {
			Description string   `yaml:"description,omitempty"`
			Rpms        []string `yaml:"rpms,omitempty"`
		} `yaml:"profiles,omitempty"`
		Profile map[string]*struct {
			Description string   `yaml:"description,omitempty"`
			Rpms        []string `yaml:"rpms,omitempty"`
		} `yaml:"profile,omitempty"`
		API struct {
			Rpms []string `yaml:"rpms,omitempty"`
		} `yaml:"api,omitempty"`
		Filter struct {
			Rpms []string `yaml:"rpms,omitempty"`
		} `yaml:"filter,omitempty"`
		Buildopts struct {
			Rpms struct {
				Macros    string   `yaml:"macros,omitempty"`
				Whitelist []string `yaml:"whitelist,omitempty"`
			} `yaml:"rpms,omitempty"`
			Arches []string `yaml:"arches,omitempty"`
		} `yaml:"buildopts,omitempty"`
		Components struct {
			Rpms map[string]*struct {
				Name          string   `yaml:"name,omitempty"`
				Rationale     string   `yaml:"rationale,omitempty"`
				Repository    string   `yaml:"repository,omitempty"`
				Cache         string   `yaml:"cache,omitempty"`
				Ref           string   `yaml:"ref,omitempty"`
				Buildonly     bool     `yaml:"buildonly,omitempty"`
				Buildroot     bool     `yaml:"buildroot,omitempty"`
				SrpmBuildroot bool     `yaml:"srpm-buildroot,omitempty"`
				Buildorder    int      `yaml:"buildorder,omitempty"`
				Arches        []string `yaml:"arches,omitempty"`
				Multilib      []string `yaml:"multilib,omitempty"`
			} `yaml:"rpms,omitempty"`
			Modules map[string]*struct {
				Rationale  string `yaml:"rationale,omitempty"`
				Repository string `yaml:"repository,omitempty"`
				Ref        string `yaml:"ref,omitempty"`
				Buildorder int    `yaml:"buildorder,omitempty"`
			} `yaml:"modules,omitempty"`
		} `yaml:"components,omitempty"`
		Artifacts struct {
			Rpms   []string `yaml:"rpms,omitempty"`
			RpmMap map[string]map[string]*struct {
				Name    string  `yaml:"name,omitempty"`
				Epoch   int     `yaml:"epoch,omitempty"`
				Version float64 `yaml:"version,omitempty"`
				Release string  `yaml:"release,omitempty"`
				Arch    string  `yaml:"arch,omitempty"`
				Nevra   string  `yaml:"nevra,omitempty"`
			} `yaml:"rpm-map,omitempty"`
		} `yaml:"artifacts,omitempty"`
	} `yaml:"data,omitempty"`
}

func Parse(input []byte) (*ModuleMd, error) {
	var ret ModuleMd
	err := yaml.Unmarshal(input, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}

func (m *ModuleMd) Marshal(fs billy.Filesystem, path string) error {
	bts, err := yaml.Marshal(m)
	if err != nil {
		return err
	}

	_ = fs.Remove(path)
	f, err := fs.Create(path)
	if err != nil {
		return err
	}
	_, err = f.Write(bts)
	if err != nil {
		return err
	}
	_ = f.Close()

	return nil
}
