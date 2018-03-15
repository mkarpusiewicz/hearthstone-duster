package services

import (
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/mkarpusiewicz/hearthstone-duster/types"
)

const urlFormat string = "https://www.hearthpwn.com/members/%s/collection"

// GetUserCards - Scrape hearthpwn for user cards
func GetUserCards(userName string) []types.MyCard {
	//doc, err := goquery.NewDocument(fmt.Sprintf(urlFormat, userName))
	file, _ := os.Open("_tools/hearthpwn.html")
	defer file.Close()

	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		log.Fatal(err)
	}

	var cards []types.MyCard

	doc.Find(".card-image-item.owns-card").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("data-card-name")
		gold, _ := s.Attr("data-is-gold")
		count, _ := s.Find(".inline-card-count").Attr("data-card-count")

		countInt, _ := strconv.Atoi(count)
		goldBool, _ := strconv.ParseBool(gold)

		cards = append(cards, types.MyCard{Name: name, Count: countInt, IsGold: goldBool})

		sort.Slice(cards, func(i, j int) bool {
			return cards[i].Name < cards[j].Name
		})
	})

	return cards
}
