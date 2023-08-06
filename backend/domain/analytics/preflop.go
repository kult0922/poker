package analytics

import (
	"github.com/kult0922/poker/backend/domain/hand"
	"github.com/kult0922/poker/backend/domain/models"
	"github.com/kult0922/poker/backend/util"
)

func CalcPreflopRank(rankMap map[int]int, rankMapFlush map[int]int) map[string]int {
	preflopHandRange := make(map[string]int)

	for i := 1; i <= 13; i++ {
		for j := 1; j <= 13; j++ {
			startingHand := [2]models.Card{
				{Suit: "unknown", Rank: i},
				{Suit: "unknown", Rank: j},
			}

			if startingHand[0].Greater(startingHand[1]) {
				startingHand[0].Suit = "spade"
				startingHand[1].Suit = "spade"
				preflopHandRange[startingHand[0].RankString()+startingHand[1].RankString()+"s"] =
					CalcStartingHandRank(rankMap, rankMapFlush, startingHand)
			}

			if startingHand[1].Greater(startingHand[0]) {
				startingHand[0].Suit = "spade"
				startingHand[1].Suit = "club"
				preflopHandRange[startingHand[1].RankString()+startingHand[0].RankString()+"o"] =
					CalcStartingHandRank(rankMap, rankMapFlush, startingHand)
			}

			if startingHand[0].Equal(startingHand[1]) {
				startingHand[0].Suit = "spade"
				startingHand[1].Suit = "club"
				preflopHandRange[startingHand[0].RankString()+startingHand[1].RankString()] =
					CalcStartingHandRank(rankMap, rankMapFlush, startingHand)
			}

		}
	}

	return preflopHandRange
}

func CalcStartingHandRank(rankMap map[int]int, rankMapFlush map[int]int, startingHand [2]models.Card) int {
	stock := models.GetStock()

	// 2枚のカードを除外
	for i, card := range stock {
		if card == startingHand[0] {
			stock = util.Remove(i, stock)
		}
	}
	for i, card := range stock {
		if card == startingHand[1] {
			stock = util.Remove(i, stock)
		}
	}

	totalRank := 0

	len := len(stock)

	for i := 0; i < len-4; i++ {
		for j := i + 1; j < len-3; j++ {
			for k := j + 1; k < len-2; k++ {
				for l := k + 1; l < len-1; l++ {
					for m := l + 1; m < len; m++ {
						sevenCards := [7]models.Card{
							stock[i],
							stock[j],
							stock[k],
							stock[l],
							stock[m],
							startingHand[0],
							startingHand[1],
						}
						totalRank += hand.SevenCardsHandRank(rankMap, rankMapFlush, sevenCards)
					}
				}
			}
		}
	}

	return totalRank
}
