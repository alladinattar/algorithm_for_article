package w_parser

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"strings"
)

type dataUnit struct {
	tag    string
	parent string
	text   string
	level  int
	index  int
	//childs map[string] [] obj
}

type dataUnitWithoutText struct {
	Tag    string
	Parent string
	Level  int
	Index  int
	//childs map[string] [] obj
}

type textUnit struct {
	tag  string
	text string
}

type counterHelper struct {
	insideWhichTag []string
	Depth          int
	Id             int
}

func main() {
	str := "<w id='#ABCA;' ><c id='AE;' >Плательщик</c> : <v id='AE;' qqw='293:fio;293:pB;293:pIDo;'>ТФОМС Санкт-Петербург</v> : <c id='AEAB;' >Тип плательщика</c> : <v id='AEAB;' qqw='293:tpl;'>ОМС</v> ; <c id='AEAC;' >Тарифный план</c> {ОМС}</w><z id='AE;' qqw='293:fio;293:pB;293:pIDo;'>ТФОМС Санкт-Петербург</z>"
	texts, tags := GetDataContentAndTagSet(str)

	fmt.Println(texts)
	fmt.Print(tags)

}

func GetTagContentByIndex(strToParse string, contentIndex int) (string, error) {
	arrayToCheck := parse(strToParse)
	if contentIndex < len(arrayToCheck) {
		return arrayToCheck[contentIndex].text, nil
	} else {
		return "", errors.New("Index out of range")
	}
}

func getTokenCount(strToParse string) int {
	z := getTokenizer(strToParse)
	tagsLen := 0
	for {
		tt1 := z.Next()
		if tt1 == html.ErrorToken {
			break
		}
		tagsLen++
	}
	return tagsLen
}

func getTokenizer(strToParse string) *html.Tokenizer {
	r := strings.NewReader(strToParse)
	z := html.NewTokenizer(r)
	return z
}

func parse(strToParse string) []dataUnit {
	z := getTokenizer(strToParse)
	tagsCount := getTokenCount(strToParse)
	var tags []dataUnit
	counter := counterHelper{}
	var textUnitArray []textUnit
	var contentTemp string
	var content string

	for j := 0; j < tagsCount; j++ {
		currentTokenizer := z.Next()
		contentTemp = string(z.Text())
		if len(contentTemp) > 0 {
			content = contentTemp
		}
		token := z.Token()
		tagName := token.Data
		if currentTokenizer == html.StartTagToken {
			counter.Depth++
			counter.insideWhichTag = append(counter.insideWhichTag, tagName)
			var dataUnit0 dataUnit
			if counter.Depth == 1 {
				dataUnit0 = dataUnit{tag: tagName, parent: "None", level: counter.Depth, index: counter.Id}
			}
			if counter.Depth > 1 {
				var parent string
				for i := len(tags) - 1; i >= 0; i-- {
					if tags[i].level == counter.Depth-1 {
						parent = tags[i].tag
						break
					}
				}
				dataUnit0 = dataUnit{tag: tagName, parent: parent, level: counter.Depth, index: counter.Id}
			}
			tags = append(tags, dataUnit0)
			counter.Id++
		}
		if currentTokenizer == html.EndTagToken {
			if len(counter.insideWhichTag) != 0 && len(tags) != 0 {
				if counter.insideWhichTag[len(counter.insideWhichTag)-1] != "w" {
					tags[len(tags)-1].text = content
				}
			}
			counter.Depth--
			if (len(counter.insideWhichTag)) > 0 {
				counter.insideWhichTag = remove(counter.insideWhichTag)
			}
		}
		if currentTokenizer == html.TextToken && len(counter.insideWhichTag) > 0 {
			textUnitArray = append(textUnitArray, textUnit{tag: counter.insideWhichTag[len(counter.insideWhichTag)-1], text: content})
		}
	}
	return tags
}

func assembleTagsArrayWithoutText(strToBuildArray string) []dataUnitWithoutText {
	tagsLen := getTokenCount(strToBuildArray)
	z := getTokenizer(strToBuildArray)
	var tags []dataUnitWithoutText
	counter := counterHelper{}
	for j := 0; j < tagsLen; j++ {
		currentTokenizer := z.Next()
		token := z.Token()
		data := token.Data

		if currentTokenizer == html.StartTagToken {
			counter.Depth++
			counter.insideWhichTag = append(counter.insideWhichTag, data)
			var dataUnit0 dataUnitWithoutText

			if counter.Depth == 1 {
				dataUnit0 = dataUnitWithoutText{Tag: data, Parent: "None", Level: counter.Depth, Index: counter.Id}
			}
			if counter.Depth > 1 {
				var parent string
				for i := len(tags) - 1; i >= 0; i-- {
					if tags[i].Level == counter.Depth-1 {
						parent = tags[i].Tag
						break
					}
				}
				dataUnit0 = dataUnitWithoutText{Tag: data, Parent: parent, Level: counter.Depth, Index: counter.Id}
			}
			tags = append(tags, dataUnit0)
			counter.Id++
		}
		if currentTokenizer == html.EndTagToken {
			counter.Depth--
			counter.insideWhichTag = remove(counter.insideWhichTag)
		}
	}
	return tags
}

func ParseStrWithoutText(strToParse string) []string {
	tags := assembleTagsArrayWithoutText(strToParse)
	var unitArray []string
	for i := 0; i < len(tags); i++ {
		jsonUnit, _ := json.Marshal(tags[i])
		unitArray = append(unitArray, string(jsonUnit))
	}
	return unitArray
}

func remove(s []string) []string {
	return s[:len(s)-1]
}

func GetDataContentAndTagSet(strToParse string) ([]string, []string) {
	doc, err := html.Parse(strings.NewReader(strToParse))
	if err != nil {
		log.Fatal(err)
	}

	var tagStr string
	var tags []string
	var texts []string
	var f func(*html.Node)

	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data != "html" && n.Data != "head" && n.Data != "body" {
			tags = append(tags, n.Data)
			tagStr = n.Data
		}
		if n.Type == html.TextNode {
			texts = append(texts, n.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return texts, tags
}
