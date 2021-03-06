package main

import (
	"fmt"

	"github.com/StratoAPI/Interface/plugins"

	"github.com/StratoAPI/Plugin-Skeleton/skeleton"
)

type SkeletonPlugin string

func (SkeletonPlugin) Name() string {
	return "Skeleton"
}

func (SkeletonPlugin) Entrypoint() {
	plugins.GetRegistry().RegisterFacade("skeleton", &skeleton.SkeletonFacade{})
	plugins.GetRegistry().RegisterStorage("skeleton", &skeleton.SkeletonStorage{})
	plugins.GetRegistry().RegisterFilter("skeleton", &skeleton.SkeletonFilter{})
	plugins.GetRegistry().RegisterMiddleware("skeleton", &skeleton.SkeletonMiddleware{})
	plugins.GetRegistry().AssociateFilter("skeleton", "skeleton")
	fmt.Println("Skeleton Plugin Initialized")
}

var CorePlugin SkeletonPlugin
