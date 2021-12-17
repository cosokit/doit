//go:build !darwin
// +build !darwin

package main

import (
	"fmt"
	"github.com/cosokit/doit/pkg/comm/utils"
)

var ConfigureFile = fmt.Sprintf("%s/doit.yaml", utils.GetExeDir())
var ResourcesPath = utils.GetExeDir()
