package main

import (
	"log"
	"regexp"
	"strings"
)

func main() {
	log.Println(GetAvatarCharacter("测测a"))
}

var (
	avatarPattern1 = regexp.MustCompile(`^[\x{0800}-\x{4e00}\x{AC00}-\x{D7A3}\x{3130}-\x{318F}\x{4e00}-\x{9fa5}]+$`)
	avatarPattern2 = regexp.MustCompile(`\s+`)
)

// GetAvatarCharacter ...
func GetAvatarCharacter(character string) string {
	// 纯东亚三国语言取后2位，带空格的名字取前2字的首字母，其它取前2字符
	if matched := avatarPattern1.MatchString(character); matched {
		character = avatarPattern2.ReplaceAllString(character, "")
		runes := []rune(character)
		start := len(runes) - 2
		if start < 0 {
			start = 0
		}
		character = string(runes[start:])
		log.Println(character, len(character))
	} else if matched := avatarPattern2.MatchString(character); matched {
		character = strings.ToUpper(character)
		arr := []string{}
		for _, i := range avatarPattern2.Split(character, -1) {
			arr = append(arr, i[0:1])
		}
		character = strings.Join(arr, "")
		if len(character) > 2 {
			character = character[0:2]
		}
	} else {
		character = strings.ToUpper(character)
		runes := []rune(character)
		if len(runes) > 2 {
			character = string(runes[0:2])
		}
	}

	return character
}
