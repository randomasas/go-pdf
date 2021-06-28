package gopdf

import (
	"os"
)

func Exist(localFile string) (bool, error) {
	_, err := os.Stat(localFile)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
