package wordMapper

import (
	"../dataStructure"
	"../utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func getMappedPage(page *dataStructure.StemmedPageJson) dataStructure.PageElement {
	var mappedText = make(map[string]float64)

	for _, rev := range page.Revision {
		for _, word := range rev.Text {
			if _, ok := mappedText[word]; ok {
				mappedText[word] += 1
			} else {
				mappedText[word] = 1
			}
		}
	}
	return dataStructure.PageElement{PageId: page.PageID, Word: mappedText}
}

func WordMapperByPage(resultDir string) {
	fileList := utils.FilesInDir(resultDir, ".json", "S")
	nFile := len(fileList)

	for i, file := range fileList {
		fmt.Printf("\rOn %d/%d", i, nFile)

		jsonFile, err := os.Open(file)
		// if we os.Open returns an error then handle it
		if err != nil {
			panic(err)
		}
		// defer the closing of our jsonFile so that we can parse it later on

		// read our opened xmlFile as a byte array.
		byteValue, _ := ioutil.ReadAll(jsonFile)

		_ = jsonFile.Close()

		var page dataStructure.StemmedPageJson

		// we unmarshal our byteArray which contains our
		// jsonFile's content into 'users' which we defined above
		_ = json.Unmarshal(byteValue, &page)

		mappedPage := getMappedPage(&page)
		_ = os.Remove(file)
		if len(mappedPage.Word) > 0 {
			utils.WriteMappedPage(resultDir, &mappedPage)
		}
	}
}