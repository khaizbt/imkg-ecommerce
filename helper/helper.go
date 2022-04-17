package helper

import (
	"bytes"
	"crypto/rand"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"

	"text/template"

	"gopkg.in/gomail.v2"
)

func GenerateNumber(length int) (string, error) {
	var chars = "0123456789"
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "0", err
	}

	charsLength := len(chars)
	for i := 0; i < length; i++ {
		buffer[i] = chars[int(buffer[i])%charsLength]
	}

	return string(buffer), nil
}

func UploadFile(alias, path, fileType string, file multipart.File, handler *multipart.FileHeader) (string, error) {
	defer file.Close()

	dir, err := os.Getwd()

	if err != nil {
		return "", err
	}

	fileExt := filepath.Ext(handler.Filename)

	if fileType == "image" {
		if fileExt != ".jpg" && fileExt != ".png" && fileExt != ".jpeg" {
			return "", errors.New("format file invalid")
		}
	}

	//TODO File Extension Document dll

	filename := handler.Filename

	if alias != "" {
		filename = fmt.Sprintf("%s%s", alias, fileExt)
	}

	if path != "" {
		filename = fmt.Sprintf("%s%s%s", path, "/", filename)
	}

	_, err = os.Stat(path)

	if err == nil {

		os.Mkdir("Storage/"+path, os.ModePerm)
	}

	fileLocation := filepath.Join(dir, "Storage", filename)

	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		return "", err
	}

	defer targetFile.Close()

	if _, err := io.Copy(targetFile, file); err != nil {
		return "", err
	}

	filePath := "Storage/" + filename

	return filePath, nil

}

func SendEmail(to, subject, template_name string, message interface{}) {
	body, err := renderEmailTemplate(template_name, message)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	mail := gomail.NewMessage()

	mail.SetHeader("From", os.Getenv("MAIL_FROM"))

	mail.SetHeader("To", to)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", body)
	mailPort, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
	d := gomail.NewDialer(os.Getenv("MAIL_HOST"), mailPort, os.Getenv("MAIL_USERNAME"), os.Getenv("MAIL_PASSWORD"))

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(mail); err != nil {
		fmt.Println(err)
		panic(err)
	}

}

func renderEmailTemplate(template_string string, body interface{}) (string, error) {
	template_html, err := template.ParseFiles(template_string)

	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)

	if err = template_html.Execute(buf, body); err != nil {
		return "", err
	}

	return buf.String(), nil

}
