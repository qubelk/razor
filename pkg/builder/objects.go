package builder

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"golang.org/x/sync/errgroup"
)

func (b *Builder) toObject(path string) string {
	trimedPath := strings.TrimSuffix(path, ".c")
	outpath := filepath.Join("build", trimedPath+".o")

	return outpath
}

func (b *Builder) createOutpath(path string) (string, error) {
	outpath := b.toObject(path)

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

	return nil
}

func (b *Builder) buildObjects() error {
	var eg errgroup.Group

	for _, v := range b.sources {
		b.objects = append(b.objects, b.toObject(v))
		eg.Go(func() error {
			need, err := b.needRebuild(v)
			if err != nil {
				return err
			}

			if !need {
				fmt.Println("no need rebuild")
				return nil
			}

			if err := b.executeBuild(v); err != nil {
				return fmt.Errorf("failed to build object file for %s: %w", v, err)
			}

			return nil
		})
	}

	return eg.Wait()
}
