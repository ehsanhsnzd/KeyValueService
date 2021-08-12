package Services

import (
	"biges/Repository"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"time"
)

// SaveData store json struct data in a file
func SaveData(data *Repository.Data) {

	content, _ := json.MarshalIndent(data.Content, "", " ")
	fmt.Println(string(content))

	timestamp:= strconv.Itoa(int(time.Now().Unix()))
	// writing json to file in tmp directory
	err := ioutil.WriteFile("tmp/"+timestamp+"-db.txt", content, 0644)

	if err != nil {
		log.Fatal(err)
	}

}
