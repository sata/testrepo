// +build mage

package main

import (
	"github.com/sata/build/pkg/build"
)

func Build() {
	build.Build("cmd")
}
