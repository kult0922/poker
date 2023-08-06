package hand

import (
	"sort"

	"github.com/kult0922/poker/backend/domain/models"
	"github.com/kult0922/poker/backend/util"
)

func SevenCardsHandRank(rankMap map[int]int, rankMapFlush map[int]int, cards [7]models.Card) int {
	len := 7

	var handRank = 999999

	for i := 0; i < len-4; i++ {
		for j := i + 1; j < len-3; j++ {
			for k := j + 1; k < len-2; k++ {
				for l := k + 1; l < len-1; l++ {
					for m := l + 1; m < len; m++ {
						hand := [5]models.Card{
							cards[i],
							cards[j],
							cards[k],
							cards[l],
							cards[m],
						}
						HandName := HandName(hand)
						handRank = util.Min(handRank, HandRank(rankMap, rankMapFlush, hand, HandName))
					}
				}
			}
		}
	}

	return handRank
}

func HandRank(rankMap map[int]int, rankMapFlush map[int]int, cards [5]models.Card, handName string) int {
	sort.Slice(cards[:], func(i, j int) bool {
		return cards[i].Rank < cards[j].Rank
	})

	primeMap := map[int]int{
		2:  2,
		3:  3,
		4:  5,
		5:  7,
		6:  11,
		7:  13,
		8:  17,
		9:  19,
		10: 23,
		11: 29,
		12: 31,
		13: 37,
		1:  41,
	}

	hash := primeMap[cards[0].Rank] * primeMap[cards[1].Rank] * primeMap[cards[2].Rank] *
		primeMap[cards[3].Rank] * primeMap[cards[4].Rank]

	switch handName {
	case "RoyalFlush":
		return rankMapFlush[hash]
	case "StraightFlush":
		return rankMapFlush[hash]
	case "Flush":
		return rankMapFlush[hash]
	default:
		return rankMap[hash]
	}
}

func HandName(cards [5]models.Card) string {
	if IsRoyalFlush(cards) {
		return "RoyalFlush"
	}

	if IsStraightFlush(cards) {
		return "StraightFlush"
	}

	if IsQuads(cards) {
		return "Quads"
	}

	if IsFullHouse(cards) {
		return "FullHouse"
	}

	if IsFlush(cards) {
		return "Flush"
	}

	if IsStraight(cards) {
		return "Straight"
	}

	if IsTrips(cards) {
		return "Trips"
	}

	if IsTwoPair(cards) {
		return "TwoPair"
	}

	if IsPair(cards) {
		return "Pair"
	}

	return "HighCard"
}

func IsRoyalFlush(cards [5]models.Card) bool {
	sort.Slice(cards[:], func(i, j int) bool {
		return cards[i].Rank < cards[j].Rank
	})

	if cards[0].Rank != 1 || cards[1].Rank != 10 || cards[2].Rank != 11 ||
		cards[3].Rank != 12 || cards[4].Rank != 13 {
		return false
	}

	if !IsFlush(cards) {
		return false
	}

	return true
}

func IsStraightFlush(cards [5]models.Card) bool {
	if IsFlush(cards) && IsStraight(cards) {
		return true
	}

	return false
}

func IsQuads(cards [5]models.Card) bool {
	sort.Slice(cards[:], func(i, j int) bool {
		return cards[i].Rank < cards[j].Rank
	})

	if cards[0].Rank == cards[1].Rank && cards[1].Rank == cards[2].Rank &&
		cards[2].Rank == cards[3].Rank {
		return true
	}

	if cards[1].Rank == cards[2].Rank && cards[2].Rank == cards[3].Rank &&
		cards[3].Rank == cards[4].Rank {
		return true
	}

	return false
}

func IsFullHouse(cards [5]models.Card) bool {
	sort.Slice(cards[:], func(i, j int) bool {
		return cards[i].Rank < cards[j].Rank
	})

	if cards[0].Rank == cards[1].Rank && cards[1].Rank == cards[2].Rank &&
		cards[3].Rank == cards[4].Rank {
		return true
	}

	if cards[0].Rank == cards[1].Rank && cards[2].Rank == cards[3].Rank &&
		cards[3].Rank == cards[4].Rank {
		return true
	}

	return false
}

func IsFlush(cards [5]models.Card) bool {
	firstSuit := cards[0].Suit
	return util.All(func(c models.Card) bool { return c.Suit == firstSuit }, cards[:])
}

func IsStraight(cards [5]models.Card) bool {
	sort.Slice(cards[:], func(i, j int) bool {
		return cards[i].Rank < cards[j].Rank
	})

	// エースを14として使うストレート
	if cards[0].Rank == 1 && cards[1].Rank == 10 && cards[2].Rank == 11 &&
		cards[3].Rank == 12 && cards[4].Rank == 13 {
		return true
	}

	for i := 0; i < len(cards)-1; i++ {
		if cards[i].Rank+1 != cards[i+1].Rank {
			return false
		}
	}

	return true
}

func IsTrips(cards [5]models.Card) bool {
	sort.Slice(cards[:], func(i, j int) bool {
		return cards[i].Rank < cards[j].Rank
	})

	if cards[0].Rank == cards[1].Rank && cards[1].Rank == cards[2].Rank {
		return true
	}

	if cards[1].Rank == cards[2].Rank && cards[2].Rank == cards[3].Rank {
		return true
	}

	if cards[2].Rank == cards[3].Rank && cards[3].Rank == cards[4].Rank {
		return true
	}

	return false
}

func IsTwoPair(cards [5]models.Card) bool {
	sort.Slice(cards[:], func(i, j int) bool {
		return cards[i].Rank < cards[j].Rank
	})

	if cards[0].Rank == cards[1].Rank && cards[2].Rank == cards[3].Rank {
		return true
	}

	if cards[0].Rank == cards[1].Rank && cards[3].Rank == cards[4].Rank {
		return true
	}

	if cards[1].Rank == cards[2].Rank && cards[3].Rank == cards[4].Rank {
		return true
	}

	return false
}

func IsPair(cards [5]models.Card) bool {
	sort.Slice(cards[:], func(i, j int) bool {
		return cards[i].Rank < cards[j].Rank
	})

	if cards[0].Rank == cards[1].Rank {
		return true
	}

	if cards[1].Rank == cards[2].Rank {
		return true
	}

	if cards[2].Rank == cards[3].Rank {
		return true
	}

	if cards[3].Rank == cards[4].Rank {
		return true
	}

	return false
}
