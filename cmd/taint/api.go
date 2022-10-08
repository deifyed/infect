package taint

import (
	"fmt"

	"github.com/deifyed/infect/pkg/storage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

func RunE(fs *afero.Afero) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		targetPath := args[0]

		trackedPath, err := storage.Get(fs, targetPath)
		if err != nil {
			return fmt.Errorf("retrieving tracked path: %w", err)
		}

		trackedPath.Taint = true

		err = storage.Add(fs, trackedPath)
		if err != nil {
			return fmt.Errorf("adding tracked path: %w", err)
		}

		return nil
	}
}
