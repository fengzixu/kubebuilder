/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package file

// Input is the input for scaffolding a file
type Input struct {
	// Path is the file to write
	Path string

	// IfExistsAction determines what to do if the file exists
	IfExistsAction IfExistsAction

	// TemplateBody is the template body to execute
	TemplateBody string

	// Boilerplate is the contents of a Boilerplate go header file
	Boilerplate string

	// Domain is the domain for the APIs
	Domain string

	// Repo is the go project package
	Repo string

	// MultiGroup is the multi-group boolean from the PROJECT file
	MultiGroup bool
}

// Domain allows a domain to be set on an object
type Domain interface {
	// SetDomain sets the domain
	SetDomain(string)
}

// SetDomain sets the domain
func (i *Input) SetDomain(d string) {
	if i.Domain == "" {
		i.Domain = d
	}
}

// Repo allows a repo to be set on an object
type Repo interface {
	// SetRepo sets the repo
	SetRepo(string)
}

// SetRepo sets the repo
func (i *Input) SetRepo(r string) {
	if i.Repo == "" {
		i.Repo = r
	}
}

// Boilerplate allows boilerplate text to be set on an object
type Boilerplate interface {
	// SetBoilerplate sets the boilerplate text
	SetBoilerplate(string)
}

// SetBoilerplate sets the boilerplate text
func (i *Input) SetBoilerplate(b string) {
	if i.Boilerplate == "" {
		i.Boilerplate = b
	}
}

// MultiGroup allows the project version to be set on an object
type MultiGroup interface {
	// SetVersion sets the project version
	SetMultiGroup(value bool)
}

// SetVersion sets the MultiGroup value
func (i *Input) SetMultiGroup(v bool) {
	i.MultiGroup = v
}

// Template is a scaffoldable file template
type Template interface {
	// GetInput returns the Input for creating a scaffold file
	GetInput() (Input, error)
}

// RequiresValidation is a file that requires validation
type RequiresValidation interface {
	Template
	// Validate returns true if the template has valid values
	Validate() error
}
