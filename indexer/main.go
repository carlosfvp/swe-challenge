package main

import (
	"bufio"
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/mail"
	"os"
	"strings"
)

func list_all_files(folder string) []string {
	files, err := os.ReadDir(folder)

	if err != nil {
		return nil
	}

	file_len := 0

	for _, f := range files {
		if !f.IsDir() {
			file_len++
		}
	}

	arr_files := make([]string, file_len)
	arr_index := 0

	for _, f := range files {
		if f.IsDir() {
			arr_inner_files := list_all_files(folder + "/" + f.Name())
			arr_files = append(arr_files, arr_inner_files...)
		} else {
			if strings.Contains(f.Name(), ".DS_Store") {
				continue
			}
			//fmt.Printf("%v\n", folder+"/"+f.Name())
			arr_files[arr_index] = folder + "/" + f.Name()
			arr_index++
		}
	}

	return arr_files
}

func digest_mail(mailInfo *mail.Message) map[string]string {
	mailMap := make(map[string]string)

	header := mailInfo.Header

	for key, value := range header {
		mailMap[key] = value[0]
	}

	body, err := io.ReadAll(mailInfo.Body)
	if err != nil {
		log.Fatal(err)
	}

	mailMap["Body"] = string(body[:])

	return mailMap
}

func push_index(client *http.Client, data map[string]string) bool {
	const user = "admin"
	const pass = "Complexpass#123"
	const authString = user + ":" + pass
	authBytes := make([]byte, len(authString))
	copy(authBytes, authString)
	encodedCredentials := b64.StdEncoding.EncodeToString(authBytes)

	const baseURL = "http://localhost:4080/api/enron_mails/_doc"

	jsonString, _ := json.Marshal(data)

	req, err := http.NewRequest("PUT", baseURL, bytes.NewReader(jsonString))

	if err != nil {
		log.Println("Could not make http.NewRequest")
		return false
	}

	req.Header.Add("Content-type", "application/json")
	req.Header.Add("Authorization", "Basic "+encodedCredentials)

	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return false
	}

	defer resp.Body.Close()

	return resp.StatusCode == 200
}

func main() {
	filePaths := list_all_files("/Users/carlos/Downloads/enron_mail_20110402/maildir")

	tr := &http.Transport{
		//MaxIdleConns:        10,
		//MaxIdleConnsPerHost: 10,
	}
	client := &http.Client{Transport: tr}

	for _, filePath := range filePaths {
		fileHandler, err := os.Open(filePath)
		if err != nil {
			continue
		}

		log.Print(filePath)

		fileReader := bufio.NewReader(fileHandler)
		m, err := mail.ReadMessage(fileReader)
		if err != nil {
			fileHandler.Close()
			log.Println("Error parsing!")
			continue
		}

		mailFields := digest_mail(m)

		result := push_index(client, mailFields)

		if !result {
			log.Println("Error pushing!")
		}
	}
}
