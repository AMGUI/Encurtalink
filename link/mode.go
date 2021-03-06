package link

import (
	//"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

//Datamode retorna um link alteraado
func Datamode(Url string) ModeJason {
	var dadoRetornoBanco ModeJason
	modo := make([]ModeJason, 0)
	jsonFile, err := os.Open(`urls.json`)

	if err != nil {

		fmt.Println(err)
	}

	//Aqui o arquivo é convertido para uma variável array de bytes, através do pacote "io/ioutil"
	byteValueJSON, _ := ioutil.ReadAll(jsonFile)

	//fmt.Println(bytes.NewBuffer(byteValueJSON))
	json.Unmarshal(byteValueJSON, &modo)

	//	newMode, _ := json.Marshal(AddnovoLink(modo, Url))
	estruturaDB := AddnovoLinkV2(modo, Url)
	//result, err := json.Marshal(AddnovoLinkV2(modo, Url))
	result, err := json.Marshal(estruturaDB)
	//fmt.Println(estruturaDB[len(estruturaDB)-1])

	if err != nil {

		fmt.Println(err)
	}

	ioutil.WriteFile("urls.json", result, 0644)
	//fmt.Println(AddnovoLink(modo, Url))

	defer jsonFile.Close()
	dadoRetornoBanco = estruturaDB[len(estruturaDB)-1]
	return dadoRetornoBanco

}

func AddnovoLinkV2(modejunior []ModeJason, url string) []ModeJason {
	var novoindice = len(modejunior) + 1
	return append(modejunior, ModeJason{
		ID:        novoindice,
		URL:       url,
		LinkMenor: "AppLink-" + strconv.Itoa(novoindice),
	})

}

func AddnovoLink(mNewstruct ModeJason, url string) ModeJason {
	mNewstruct.ID += 1
	mNewstruct.URL = url
	mNewstruct.LinkMenor = "APPLINK" + string(mNewstruct.ID)
	return mNewstruct
}
