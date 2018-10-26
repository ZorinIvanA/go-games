package fias

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

//GetFias returns list of fias records
func GetFias(fiasURL string, fiasEnvelope string) (Envelope, error) {
	client := &http.Client{}
	sRequestContent := fiasEnvelope
	requestContent := []byte(sRequestContent)
	req, err := http.NewRequest("POST", fiasURL, bytes.NewBuffer(requestContent))
	if err != nil {
		return Envelope{}, err
	}

	req.Header.Add("Content-Type", "application/soap+xml; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		return Envelope{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return Envelope{}, errors.New("Error Respose " + resp.Status)
	}

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Envelope{}, err
	}

	envelope := Envelope{}
	lenght := len(xml.Header)
	str1 := string(contents)
	str := str1[lenght-1 : len(str1)]
	//str := "<soap:Envelope><soap:Body><GetAllDownloadFileInfoResponse><GetAllDownloadFileInfoResult><DownloadFileInfo></DownloadFileInfo><DownloadFileInfo></DownloadFileInfo></GetAllDownloadFileInfoResult></GetAllDownloadFileInfoResponse></soap:Body></soap:Envelope>"
	//bts := bytes.NewReader([]byte(str))
	bts := []byte(str)

	fmt.Println("result is: " + str)

	// decoder := xml.NewDecoder(bts)
	// errUml := decoder.Decode(envelope)

	errUml := xml.Unmarshal(bts, &envelope)

	if err != nil {
		fmt.Println(errUml)
		return Envelope{}, err
	}
	// for _, downloadInfo := range envelope.Body.GetAllDownloadFileInfoResponse.GetAllDownloadFileInfoResult {
	// 	fmt.Printf("Url: %s Version: %d\n", downloadInfo.FiasCompleteDbfURL, downloadInfo.VersionID)
	// }

	return envelope, nil
}

func generateRequestContent() string {
	return ""
}
