package words

import (
	"bufio"
	"bytes"
	_ "embed"
	"hangman-tui/pkg/boot"
	"math"
	"math/rand"
	"os"
	"slices"
	"strings"
	"sync"
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

func (s *Store) ReadFromFile(file string) {
	f, err := os.Open(file)
	if err != nil {
		boot.Shutdown("Failed to get words: " + err.Error())
	}
	defer func() {
		_ = f.Close()
	}()

	r := bufio.NewReader(f)

	s.readReader(r)
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

		length := len(word)
		if length <= 0 {
			continue
		}
		if length < s.minLength {
			s.minLength = length
		}
		if length > s.maxLength {
			s.maxLength = length
		}
		s.availableLength[length] = true
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
	filtered := slices.DeleteFunc(s.words, func(val string) bool { return len(val) != s.length })

	random := rand.Intn(len(filtered))

	return filtered[random]
}
