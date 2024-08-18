package main

//go:generate go run . -o ./pkg/

import (
	"runtime/debug"

	"github.com/diamondburned/gotk4-sourceview/gensourceview"
	"github.com/diamondburned/gotk4/gir/cmd/gir-generate/gendata"
	"github.com/diamondburned/gotk4/gir/cmd/gir-generate/genmain"
	"github.com/diamondburned/gotk4/gir/girgen"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

var gomodModule = func() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		panic("cannot read build info")
	}
	return info.Main.Path + "/pkg"
}()

var preprocessors = []types.Preprocessor{}

var postprocessors = map[string][]girgen.Postprocessor{}

var Data = genmain.Overlay(
	gendata.Main,
	gensourceview.Data,
	genmain.Data{
		Module: gomodModule,
		Packages: []genmain.Package{
			{
				Name:       "libspelling-1",
				Namespaces: []string{"Spelling-1"},
			},
		},
		PkgGenerated: []string{"spelling"},
		PkgExceptions: []string{
			"go.mod",
			"go.sum",
			"LICENSE",
		},
		Preprocessors:  preprocessors,
		Postprocessors: postprocessors,
	},
)

func main() {
	genmain.Run(Data)
}
