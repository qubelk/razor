package builder

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"razor/pkg/config"
	"strings"
)

type Builder struct {
	cfg     *config.Config
	sources []string
}

func New(cfg *config.Config, sources []string) *Builder {
	return &Builder{
		cfg:     cfg,
		sources: sources,
	}
}

func (b *Builder) createOutpath(path string) (string, error) {
	trimedPath := strings.TrimSuffix(path, ".c")
	outpath := filepath.Join("build", trimedPath+".o")

	if err := os.MkdirAll(outpath, 0755); err != nil {
		return "", err
	}

	return outpath, nil
}

func (b *Builder) buildCommand(path string) ([]string, error) {
	outpath, err := b.createOutpath(path)
	if err != nil {
		return nil, err
	}

	cmd := []string{
		"-I.",
		fmt.Sprintf("-std=c%d", b.cfg.Conf.Standard),
		"-c",
		path,
		"-o",
		outpath,
	}

	return cmd, nil
}

func (b *Builder) executeBuild(path string) error {
	cmd, err := b.buildCommand(path)
	if err != nil {
		return err
	}

	out, err := exec.Command(b.cfg.Conf.Compiler, cmd...).CombinedOutput()

	if err != nil {
		fmt.Println(err)
		return err
	}

	if len(out) > 0 {
		fmt.Println(string(out))
	}

	return nil
}

func (b *Builder) Build() error {
	for _, v := range b.sources {
		b.executeBuild(v)
	}

	return nil
}
