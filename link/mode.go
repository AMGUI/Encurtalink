package link

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type modeloconvesor struct {
	id        string `json:"id"`
	url       string `json:"url"`
	linkMenor string `Json:"linkMenor`
}

//Datamode retorna um link alteraado
func Datamode() {

	mode := modeloconvesor{}
	jsonFile, err := os.Open(`url.json`)

	if err != nil {

		fmt.Println(err)
	}

	//Aqui o arquivo é convertido para uma variável array de bytes, através do pacote "io/ioutil"
	byteValueJSON, _ := ioutil.ReadAll(jsonFile)

	fmt.Println(bytes.NewBuffer(byteValueJSON))
	json.Unmarshal(byteValueJSON, &mode)
	fmt.Println(mode)

	defer jsonFile.Close()

}
