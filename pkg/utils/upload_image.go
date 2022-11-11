package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

func UploadImage(c *fiber.Ctx) (string, error) {
	method := "POST"
	file, err := c.FormFile("file")
	if err != nil {
		return "", err
	}
	buffer, err := file.Open()

	if err != nil {
		return "", err
	}
	defer buffer.Close()

	client := &http.Client{}
	req, err := http.NewRequest(method, os.Getenv("UPLOAD_URL"), buffer)
	if err != nil {
		return "", err
	}

	println("UPLOAD ENV : ", os.Getenv("UPLOAD_URL"))

	req.Header.Set("Content-Type", "multipart/form-data")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", errors.New(res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	myString := string(body)
	fmt.Println(myString)
	// data := &model.TransCodeResponse{}
	// json.Unmarshal([]byte(myString), data)
	// if data.JobId == "" {
	// 	return "", errors.New(fmt.Sprint("can't transcode"))
	// }
	return myString, nil
}
