package command

import (
	"fmt"
	"log"

	"github.com/igungor/tlbot"
)

func init() {
	register(cmdImg)
}

var cmdImg = &Command{
	Name:      "img",
	ShortLine: "resim filan ara",
	Run:       runImg,
}

func runImg(b *tlbot.Bot, msg *tlbot.Message) {
	args := msg.Args()
	if len(args) == 0 {
		term := randChoice(imgExamples)
		txt := fmt.Sprintf("ne resmi aramak istiyorsun? örneğin: */img %s*", term)
		err := b.SendMessage(msg.Chat.ID, txt, tlbot.ModeMarkdown, false, nil)
		if err != nil {
			log.Printf("Error while sending message: %v\n", err)
		}
		return
	}

	u, err := searchImage(args...)
	if err != nil {
		log.Printf("Error while searching image. Err: %v\n", err)
		if err == errImageSearchQuotaExceeded {
			b.SendMessage(msg.Chat.ID, `¯\_(ツ)_/¯`, tlbot.ModeNone, false, nil)
		}
		return
	}

	photo := tlbot.Photo{File: tlbot.File{FileURL: u}}
	err = b.SendPhoto(msg.Chat.ID, photo, "", nil)
	if err != nil {
		log.Printf("Error while sending photo: %v\n", err)
		return
	}
}

var imgExamples = []string{
	"burdur",
	"kapadokya",
}
