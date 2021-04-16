package link

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//Datamode retorna um link alteraado
func Datamode(Url string) {

	modo := ModeJason{}
	jsonFile, err := os.Open(`url.json`)

	if err != nil {

		fmt.Println(err)
	}

	//Aqui o arquivo é convertido para uma variável array de bytes, através do pacote "io/ioutil"
	byteValueJSON, _ := ioutil.ReadAll(jsonFile)

	//fmt.Println(bytes.NewBuffer(byteValueJSON))
	json.Unmarshal(byteValueJSON, &modo)

	newMode, _ := json.Marshal(AddnovoLink(modo, Url))
	ioutil.WriteFile("url.json", newMode, 0644)
	//fmt.Println(AddnovoLink(modo, Url))

	//fmt.Println(modo)

	defer jsonFile.Close()

}

func AddnovoLink(mNewstruct ModeJason, url string) ModeJason {
	mNewstruct.ID += 1
	mNewstruct.URL = url
	mNewstruct.LinkMenor = "APPLINK" + string(mNewstruct.ID)
	return mNewstruct
}
