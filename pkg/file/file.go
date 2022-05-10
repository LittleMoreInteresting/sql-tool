package file

import (
	"io"
	"os"
)



func CheckSavePath(dst string) bool  {
	_,err := os.Stat(dst)
	return os.IsNotExist(err)
}


func CheckPermission(dst string) bool {
	_,err := os.Stat(dst)
	return os.IsPermission(err)
}

func CreateSavePath(dst string) error {
	err := os.MkdirAll(dst, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func CreateWriter(dst string) (io.Writer,error) {

	out,err := os.Create(dst)
	if err != nil {
		return nil,err
	}
	return out,nil
}