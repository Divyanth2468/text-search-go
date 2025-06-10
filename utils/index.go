package utils

import (
	"encoding/json"
	"os"
)

type Index map[string][]int

func (idx Index) SaveToFile(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", " ")
	return enc.Encode(idx)
}

func LoadIndexFromFile(path string) (Index, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	var idx Index
	dec := json.NewDecoder(f)
	err = dec.Decode(&idx)
	return idx, err
}

func (idx Index) Add(docs []Document) {
	for _, doc := range docs {
		for _, token := range analyze(doc.Text) {
			ids := idx[token]
			if ids != nil && ids[len(ids)-1] == doc.ID {
				// Don't add ID twice
				continue
			}
			idx[token] = append(ids, doc.ID)
		}
	}
}

func Intersection(a []int, b []int) []int {
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}

	r := make([]int, 0, maxLen)
	var i, j int

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		} else {
			r = append(r, a[i])
			i++
			j++
		}
	}
	return r
}

func (idx Index) Search(text string) []int {
	var r []int

	for _, token := range analyze(text) {
		if ids, ok := idx[token]; ok {
			if r == nil {
				r = ids
			} else {
				r = Intersection(r, ids)
			}
		} else {
			// Token doesn't exist
			return nil
		}
	}
	return r
}
