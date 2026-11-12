package builder

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"razor/pkg/config"
	"strings"

	"golang.org/x/sync/errgroup"
)

type Builder struct {
	cfg     *config.Config
	sources []string
}

// Source must contains only .c files
func New(cfg *config.Config, sources []string) *Builder {
	return &Builder{
		cfg:     cfg,
		sources: sources,
	}
}

func (b *Builder) createOutpath(path string) (string, error) {
	trimedPath := strings.TrimSuffix(path, ".c")
	outpath := filepath.Join("build", trimedPath+".o")

	if err := os.MkdirAll(filepath.Dir(outpath), 0755); err != nil {
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
		"-I%s", b.cfg.Conf.Include,
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
	if len(out) > 0 {
		fmt.Print(string(out))
	}

	if err != nil {
		fmt.Print(err)
		return err
	}

	return nil
}

func (b *Builder) buildObjects() error {
	var eg errgroup.Group

	for _, v := range b.sources {
		eg.Go(func() error {
			if err := b.executeBuild(v); err != nil {
				return fmt.Errorf("failed to build object file for %s: %w", v, err)
			}

			return nil
		})
	}

	return eg.Wait()
}

func (b *Builder) Build() error {
	return b.buildObjects()
}
