package helpers

import (
	"os/exec"
	"path/filepath"
)

type GitHelper struct{}

func (g *GitHelper) cloneTag(repoUrl, tag, basePath, name string) error {
	finalPath := filepath.Join(basePath, "source", tag, name)
	cmd := exec.Command("git", "clone", "--branch", tag, "--depth", "1", repoUrl, finalPath)
	return cmd.Run()
}

func (g *GitHelper) CloneBackend(tag, basePath string) error {
	repoUrl := "https://github.com/qonsensus/backend.git"
	return g.cloneTag(repoUrl, tag, basePath, "backend")
}

func (g *GitHelper) CloneFrontend(tag, basePath string) error {
	repoUrl := "https://github.com/qonsensus/frontend.git"
	return g.cloneTag(repoUrl, tag, basePath, "frontend")
}
