package drive

import (
	"fmt"
	"io"
)

type DeleteArgs struct {
	Out       io.Writer
	ID        string
	Recursive bool
}

func (self *Drive) Delete(args DeleteArgs) error {
	f, err := self.service.Files.Get(args.ID).Fields("name", "mimeType").Do()
	if err != nil {
		return fmt.Errorf("failed to get file: %s", err.Error())
	}

	if isDir(f) && !args.Recursive {
		return fmt.Errorf("'%s' is a directory, use the 'recursive' flag to delete directories", f.Name)
	}

	err = self.service.Files.Delete(args.ID).Do()
	if err != nil {
		return fmt.Errorf("failed to delete file: %s", err.Error())
	}

	fmt.Fprintf(args.Out, "Deleted '%s'\n", f.Name)
	return nil
}

func (self *Drive) deleteFile(fileID string) error {
	err := self.service.Files.Delete(fileID).Do()
	if err != nil {
		return fmt.Errorf("failed to delete file: %s", err.Error())
	}
	return nil
}
