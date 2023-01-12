package main

import (
	"context"
	"fmt"
	"os"

	"github.com/magefile/mage/mg"
)

type Bootstrap mg.Namespace

func (Bootstrap) Bevvy(ctx context.Context, toPath string) error {
	if _, err := os.Stat(toPath); err != nil {
		return fmt.Errorf("cannot locate destination path %s", toPath)
	}

	return nil
}
