package tokenizer

import (
	"embed"
	"newsSearchEngine/engine/utils"
	"strings"

	"github.com/wangbin/jiebago"
)

var (
	//go:embed dict/*.txt
	dictionaryFS embed.FS
)

type Tokenizer struct {
	seg jiebago.Segmenter
}

func NewTokenizer(dictionaryPath string) *Tokenizer {
	file, err := dictionaryFS.Open("dict/dict.txt")
	if err != nil {
		panic(err)
	}
	utils.ReleaseAssets(file, dictionaryPath)

	tokenizer := &Tokenizer{}

	err = tokenizer.seg.LoadDictionary(dictionaryPath)
	if err != nil {
		panic(err)
	}

	return tokenizer
}

func (t *Tokenizer) Cut(text string) []string {
	//all lower
	text = strings.ToLower(text)
	//process text
	text = utils.RemovePunctuation(text)
	text = utils.RemoveSpace(text)

	var wordMap = make(map[string]struct{})

	resultChan := t.seg.CutForSearch(text, true)
	var wordsSlice []string
	for {
		w, ok := <-resultChan
		if !ok {
			break
		}
		_, found := wordMap[w]
		if !found {
			//remove duplicate words
			wordMap[w] = struct{}{}
			wordsSlice = append(wordsSlice, w)
		}
	}

	return wordsSlice
}
