package testutils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jbactad/grpc-vs-rest/pb"
)

type Request struct {
	Path   string
	Random *pb.Random
}

func Get(client http.Client, path string, output interface{}) error {
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		log.Println("error creating request ", err)
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		log.Println("error executing request ", err)
		return err
	}

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("error reading response body ", err)
		return err
	}

	err = json.Unmarshal(bytes, output)
	if err != nil {
		log.Println("error unmarshaling response ", err)
		return err
	}

	return nil
}

func Post(client http.Client, path string, input interface{}, output interface{}) error {
	data, err := json.Marshal(input)
	if err != nil {
		log.Println("error marshalling input ", err)
		return err
	}
	body := bytes.NewBuffer(data)

	req, err := http.NewRequest("POST", path, body)
	if err != nil {
		log.Println("error creating request ", err)
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		log.Println("error executing request ", err)
		return err
	}

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("error reading response body ", err)
		return err
	}

	err = json.Unmarshal(bytes, output)
	if err != nil {
		log.Println("error unmarshaling response ", err)
		return err
	}

	return nil
}
