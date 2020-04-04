package install

import (
	"io"
	"os"

	"github.com/google/go-github/v30/github"
)

// DefaultBinDir is the default location for binary files
const DefaultBinDir = "/usr/local/bin/"

// Application handles binary assets
func Application(asset *github.ReleaseAsset) error {
	filename, err := downloadFile(asset.GetBrowserDownloadURL(), asset.GetName())
	if err != nil {
		return err
	}

	from, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer from.Close()

	to, err := os.OpenFile(DefaultBinDir+asset.GetName(), os.O_RDWR|os.O_CREATE, 0744)
	if err != nil {
		return err
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	return err
}
