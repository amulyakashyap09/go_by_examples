package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const dirPath = "./records"

type User struct {
	id       int32
	name     string
	username string
	email    string
	address  *Address
	phone    string
	website  string
	company  *Company
}

type Company struct {
	name        string
	catchPhrase string
	bs          string
}

type Address struct {
	street  string
	suite   string
	city    string
	zipcode string
	geo     *Geo
}

type Geo struct {
	lat string
	lng string
}

func toJson(d interface{}) string {
	bytes, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return string(bytes)
}

func main() {

	for {
		//open the dir
		d, _ := os.Open(dirPath)
		files, _ := d.Readdir(-1) //return all files
		for _, fi := range files {
			filePath := dirPath + "/" + fi.Name() //creating file path
			f, _ := os.Open(filePath)
			data, _ := ioutil.ReadAll(f) //read all the data from file
			f.Close()
			os.Remove(filePath) //delete that file from data folder
			go func(data string) {
				x := toJson(data)
				fmt.Println(x)
			}(string(data))
		}
	}
}
