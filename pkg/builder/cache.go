package builder

import (
	"os"
	"path/filepath"
)

func (b *Builder) compareModTime(source, object string) (bool, error) {
	sourceStat, err := os.Stat(source)
	if err != nil {
		return false, err
	}

	objectStat, err := os.Stat(object)
	if err != nil {
		return false, err
	}

	return sourceStat.ModTime().After(objectStat.ModTime()), nil
}

func (b *Builder) needRebuild(source, object string) (bool, error) {
	if !b.exists(object) {
		return true, nil
	}

	compare, err := b.compareModTime(source, object)
	if err != nil {
		return true, err
	}

	return compare, nil
}

func (b *Builder) needRebuildObject(source string) (bool, error) {
	return b.needRebuild(source, b.toObject(source))
}

func (b *Builder) needRebuildExe() (bool, error) {
	exe := filepath.Join("build", b.cfg.Package.Name)
	for _, v := range b.objects {
		need, err := b.needRebuild(v, exe)
		if err != nil {
			return true, err
		}

		if need {
			return true, nil
		}
	}

	return false, nil
}
