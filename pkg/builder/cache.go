package builder

import (
	"fmt"
	"os"
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

func (b *Builder) needRebuild(source string) (bool, error) {
	object := b.toObject(source)
	if !b.exists(object) {
		fmt.Println("object not exists, rebuild")
		return true, nil
	}

	compare, err := b.compareModTime(source, object)
	if err != nil {
		return true, err
	}

	return compare, nil
}
