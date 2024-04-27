package system

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

type FTPService struct{}

var FTPServiceApp = new(FTPService)

func (f *FTPService) Upload(fileHeader *multipart.FileHeader) error {
	fmt.Println(fileHeader.Header, fileHeader.Filename)
	return nil
}

func (f *FTPService) UploadByFormParts(req *http.Request, saveDir string) error {
	mr, err := req.MultipartReader()

	if err != nil {
		return err
	}

	for {
		fp, err := mr.NextRawPart()
		var extStr string

		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		fileName := fp.FileName()
		if fileName == "" {
			continue
		}

		extIdx := strings.LastIndex(fileName, ".")
		if extIdx != -1 {
			extStr = fileName[extIdx:]
			fileName = fileName[0:extIdx]
		}

		if err := f.UploadFileByStream(fp, saveDir, fileName+"_"+time.Now().Format("2006_01_02_150405")+extStr); err != nil {
			return err
		}
	}
	return nil
}

func (f *FTPService) UploadFileByStream(r io.Reader, dir string, fileName string) (err error) {
	fullPath := path.Join(dir, fileName)

	if _, err = os.Stat(fullPath); err == nil {
		return fmt.Errorf("File is exist\n")
	}

	if err = os.MkdirAll(dir, os.ModeDir); err != nil {
		return err
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}

	bs := make([]byte, 1024)
	defer func() {
		_ = file.Close()
	}()
	for {
		num, err := r.Read(bs)
		if num > 0 {
			_, err = file.Write(bs[0:num])
			if err != nil {
				return err
			}
		}
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
	}

	return
}

func (f *FTPService) ReadRawFormPart(part *multipart.Part) error {
	bs := make([]byte, 1024)
	for {
		num, err := part.Read(bs)
		if num > 0 {
			fmt.Println("read ", part.FormName(), part.FileName(), num)
		}
		if err != nil {
			if err != io.EOF {
				return err
			}
			return nil
		}
	}
}
