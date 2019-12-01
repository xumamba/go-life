package spider

import (
	"log"
	"testing"
)

func TestTranslation(t *testing.T) {
	aimContent, srcLang := Translation("en", "你好")
	log.Printf("aimContent:%s,srcLang:%s", aimContent, srcLang)
}
