package models

import (
	"sort"

	"github.com/kult0922/go-react-blog/backend/util"
)

func HandName(cards [5]Card) string {
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

func IsRoyalFlush(cards [5]Card) bool {
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

func IsStraightFlush(cards [5]Card) bool {
	if IsFlush(cards) && IsStraight(cards) {
		return true
	}

	return false
}

func IsQuads(cards [5]Card) bool {
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

func IsFullHouse(cards [5]Card) bool {
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

func IsFlush(cards [5]Card) bool {
	firstSuit := cards[0].Suit
	return util.All(func(c Card) bool { return c.Suit == firstSuit }, cards[:])
}

func IsStraight(cards [5]Card) bool {
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

func IsTrips(cards [5]Card) bool {
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

func IsTwoPair(cards [5]Card) bool {
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

func IsPair(cards [5]Card) bool {
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
