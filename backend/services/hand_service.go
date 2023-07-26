package services

import (
	"sort"

	"github.com/kult0922/go-react-blog/backend/models"
	"github.com/kult0922/go-react-blog/backend/util"
)

func (s *MyAppService) GetHandService(cards [5]models.Card) models.Hand {

	if IsRoyalFlush(cards) {
		return models.Hand{
			Cards: cards,
			Name:  "RoyalFlush",
		}
	}

	if IsStraightFlush(cards) {
		return models.Hand{
			Cards: cards,
			Name:  "StraightFlush",
		}
	}

	if IsQuads(cards) {
		return models.Hand{
			Cards: cards,
			Name:  "Quads",
		}
	}

	if IsFullHouse(cards) {
		return models.Hand{
			Cards: cards,
			Name:  "FullHouse",
		}
	}

	if IsFlush(cards) {
		return models.Hand{
			Cards: cards,
			Name:  "Flush",
		}
	}

	if IsStraight(cards) {
		return models.Hand{
			Cards: cards,
			Name:  "Straight",
		}
	}

	if IsTrips(cards) {
		return models.Hand{
			Cards: cards,
			Name:  "Trips",
		}
	}

	if IsTwoPair(cards) {
		return models.Hand{
			Cards: cards,
			Name:  "TwoPair",
		}
	}

	if IsPair(cards) {
		return models.Hand{
			Cards: cards,
			Name:  "Pair",
		}
	}

	return models.Hand{
		Cards: cards,
		Name:  "HighCard",
	}
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
