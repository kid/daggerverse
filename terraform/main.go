package main

import (
	"github.com/kid/daggerverse/terraform/internal/dagger"
)

type Terraform struct{}

func (m *Terraform) Bar(
	// +optional
	// +defaultPath="/"
	// +ignore=[".git", ".devenv", ".direnv"]
	source *dagger.Directory,
) *dagger.Directory {
	return source
}
