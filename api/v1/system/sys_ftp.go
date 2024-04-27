package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

type FTPApi struct{}

func (f *FTPApi) Upload(c *gin.Context) {
	//file, err := c.FormFile("file")
	err := ftpService.UploadByFormParts(c.Request, "./tmp")
	if err != nil {
		c.Error(err)
		return
	}
	c.String(http.StatusOK, "ok")
	//if err = ftpService.Upload(file); err != nil {
	//	c.Error(err)
	//	return
	//}
	//c.JSON(http.StatusOK, gin.H{
	//	"fileName": file.Filename,
	//	"size":     file.Size,
	//})
}

func saveFile(fp *multipart.Part) string {
	var errors = ""
	dst, _ := os.OpenFile("./"+fp.FileName(),
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644) // 创建文件（本演示忽略错误）
	defer dst.Sync()          // 同步到磁盘
	defer dst.Close()         // 关闭文件
	buf := make([]byte, 1024) // 创建一个buffer用于读取http上传的文件
	fmt.Println("start saveing file ", fp.FileName())
	for {
		n, err := fp.Read(buf) // 循环读取文件内容
		if n > 0 {
			_, err = dst.Write(buf[0:n]) // 将buffer的内容写入新创建的文件
			if err != nil {
				errors += err.Error() + "\r\n"
				continue
			}
		}
		if err != nil {
			if err != io.EOF { // 读取遇到非文件读取结束的异常
				errors += err.Error() + "\r\n"
			}
			break
		}
	}
	fmt.Println("end saveing file ", fp.FileName())
	return errors
}
