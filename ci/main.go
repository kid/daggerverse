package main

import (
	"github.com/kid/daggerverse/ci/internal/dagger"
)

type Ci struct{}

func (m *Ci) Foo(
	// +optional
	// +defaultPath="/"
	// +ignore=[".git", ".devenv", ".direnv"]
	source *dagger.Directory,
) *dagger.Directory {
	return source
}

func (m *Ci) Bar() *dagger.Directory {
	return dag.Terraform().Bar()
}
