package builder

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"razor/pkg/config"
)

type Builder struct {
	cfg     *config.Config
	sources []string
	objects []string
}

// Source must contains only .c files
func New(cfg *config.Config, sources []string) *Builder {
	return &Builder{
		cfg:     cfg,
		sources: sources,
		objects: make([]string, 0, len(sources)),
	}
}

func (b *Builder) exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func (b *Builder) buildExeCommand() []string {
	outpath := filepath.Join("build", b.cfg.Package.Name)
	var cmd []string

	cmd = append(cmd, b.objects...)
	cmd = append(cmd, "-o", outpath)

	return cmd
}

func (b *Builder) buildExecutable() error {
	need, err := b.needRebuildExe()
	if err != nil {
		return err
	}

	if !need {
		return nil
	}

	cmd := b.buildExeCommand()

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

func (b *Builder) createVersionFile() error {
	if !b.exists("build") {
		if err := os.MkdirAll("build", 0755); err != nil {
			return err
		}
	}

	versionPath := filepath.Join("build", ".VERSION")

	f, err := os.Create(versionPath)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	version := scanner.Text()
	if b.cfg.Package.Version == version {
		return nil
	}

	_, err = f.WriteString(b.cfg.Package.Version)
	return err
}

func (b *Builder) Build() error {
	if err := b.buildObjects(); err != nil {
		return err
	}

	if err := b.buildExecutable(); err != nil {
		return err
	}

	if err := b.createVersionFile(); err != nil {
		return err
	}

	return nil
}
