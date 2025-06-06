package words

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"
	"sync"
	"unicode"
)

var lock sync.Mutex
var instance *Store

//go:embed words.txt
var defaultWords []byte

func GetStore() *Store {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &Store{}
			instance.readDefaultWords()
		}
	}
	return instance
}

type Store struct {
	words           []string
	length          int
	availableLength map[int]bool
	maxLength       int
	minLength       int
}

func (s *Store) readDefaultWords() {
	br := bytes.NewReader(defaultWords)
	r := bufio.NewReader(br)
	s.readReader(r)
}

func (s *Store) ReadFromFile(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("could not get list of words: %v", err)
	}
	defer func() {
		_ = f.Close()
	}()

	r := bufio.NewReader(f)

	s.readReader(r)

	if len(s.words) < 1 {
		return fmt.Errorf("the word list is empty")
	}
	return nil
}

func (s *Store) readReader(reader *bufio.Reader) {
	var err error
	var word string
	s.minLength = math.MaxInt
	s.maxLength = math.MinInt
	s.availableLength = make(map[int]bool)
	s.words = make([]string, 0)

	for err == nil {
		word, err = reader.ReadString('\n')
		word = strings.Trim(word, " \n")
		invalid := false

		for _, v := range word {
			if !unicode.IsLetter(v) {
				invalid = true
			}
		}
		length := len(word)
		if length <= 0 || invalid {
			continue
		}

		if length < s.minLength {
			s.minLength = length
		}
		if length > s.maxLength {
			s.maxLength = length
		}
		s.availableLength[length] = true
		word = strings.ToLower(word)
		s.words = append(s.words, word)
	}
	s.length = s.minLength

}

func (s *Store) CurrentLength() int {
	return s.length
}

func (s *Store) IncreaseLength() {
	if s.length == s.maxLength {
		return
	}

	s.length++
	for !s.availableLength[s.length] && s.length != s.maxLength {
		s.length++
	}
}

func (s *Store) DecreaseLength() {
	if s.length == s.minLength {
		return
	}

	s.length--
	for !s.availableLength[s.length] && s.length != s.minLength {
		s.length--
	}
}

func (s *Store) GetWord() string {
	// filtered := slices.DeleteFunc(s.words, func(val string) bool { return len(val) != s.length })
	filtered := make([]string, 0, len(s.words)/2)

	for _, v := range s.words {
		if len(v) == s.length {
			filtered = append(filtered, v)
		}
	}

	random := rand.Intn(len(filtered))

	return filtered[random]
}
