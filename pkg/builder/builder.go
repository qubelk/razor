package builder

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"razor/pkg/config"
	"strings"
	"sync"

	"golang.org/x/sync/errgroup"
)

type Builder struct {
	cfg     *config.Config
	mtx     sync.Mutex
	sources []string
	objects []string
}

// Source must contains only .c files
func New(cfg *config.Config, sources []string) *Builder {
	return &Builder{
		cfg:     cfg,
		mtx:     sync.Mutex{},
		sources: sources,
		objects: make([]string, 0, len(sources)),
	}
}

func (b *Builder) exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func (b *Builder) createOutpath(path string) (string, error) {
	trimedPath := strings.TrimSuffix(path, ".c")
	outpath := filepath.Join("build", trimedPath+".o")

	dir := filepath.Dir(outpath)
	if !b.exists(dir) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return "", err
		}
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
	if len(out) > 0 {
		fmt.Print(string(out))
	}

	if err != nil {
		fmt.Print(err)
		return err
	}

	b.objects = append(b.objects, cmd[len(cmd)-1])

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

func (b *Builder) buildExecutable() error {
	exePath := filepath.Join("build", b.cfg.Package.Name)
	var cmd []string

	cmd = append(cmd, b.objects...)
	cmd = append(cmd, "-o", exePath)

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

func (b *Builder) Build() error {
	err := b.buildObjects()
	if err != nil {
		return err
	}

	return b.buildExecutable()
}
