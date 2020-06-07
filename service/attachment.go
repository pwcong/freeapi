package service

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"freeapi/model"
	"freeapi/utils"
)

type AttachmentService struct {
	Base *BaseService
}

func (ctx *AttachmentService) SaveAttachment(file *multipart.FileHeader) (model.Attachment, error) {
	src, err := file.Open()
	if err != nil {
		return model.Attachment{}, err
	}

	defer src.Close()

	buffer, err := ioutil.ReadAll(src)
	if err != nil {
		return model.Attachment{}, err
	}

	h := md5.New()
	_, err = h.Write(buffer)
	if err != nil {
		return model.Attachment{}, err
	}

	symbol := hex.EncodeToString(h.Sum(nil))

	var attachment model.Attachment

	db := ctx.Base.DB
	now := time.Now()

	notFound := db.Where("symbol = ?", symbol).First(&attachment).RecordNotFound()
	if notFound {
		attachment = model.Attachment{
			Symbol:   symbol,
			Filename: file.Filename,
			Year:     now.Format("2006"),
			Month:    now.Format("01"),
			Date:     now.Format("02"),
			ExtName:  utils.GetExtension(file.Filename),
		}

		rootDir := utils.GetRootDir()

		dir := filepath.Join(rootDir, "public/attachment/"+attachment.Year+"/"+attachment.Month+"/"+attachment.Date)

		err = os.MkdirAll(dir, 0666)
		if err != nil {
			return model.Attachment{}, err
		}

		err = ioutil.WriteFile(filepath.Join(dir, attachment.Symbol+"."+attachment.ExtName), buffer, 0666)
		if err != nil {
			return model.Attachment{}, err
		}

		if dbc := db.Create(&attachment); dbc.Error != nil {
			return model.Attachment{}, dbc.Error
		}
	}

	return attachment, nil
}
