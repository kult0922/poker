package services_test

import (
	"testing"

	"github.com/kult0922/go-react-blog/backend/models"
	"github.com/kult0922/go-react-blog/backend/services"
)

func TestGetHandService(t *testing.T) {
	tests := []struct {
		name string
		args [5]models.Card
		want models.Hand
	}{
		{
			name: "RoyalFlush",
			args: [5]models.Card{
				{Suit: "Spade", Rank: 1},
				{Suit: "Spade", Rank: 13},
				{Suit: "Spade", Rank: 12},
				{Suit: "Spade", Rank: 11},
				{Suit: "Spade", Rank: 10},
			},
			want: models.Hand{
				Cards: [5]models.Card{
					{Suit: "Spade", Rank: 1},
					{Suit: "Spade", Rank: 13},
					{Suit: "Spade", Rank: 12},
					{Suit: "Spade", Rank: 11},
					{Suit: "Spade", Rank: 10},
				},
				Name: "RoyalFlush",
			},
		},
		{
			name: "StraightFlush",
			args: [5]models.Card{
				{Suit: "Spade", Rank: 1},
				{Suit: "Spade", Rank: 2},
				{Suit: "Spade", Rank: 3},
				{Suit: "Spade", Rank: 4},
				{Suit: "Spade", Rank: 5},
			},
			want: models.Hand{
				Cards: [5]models.Card{
					{Suit: "Spade", Rank: 1},
					{Suit: "Spade", Rank: 2},
					{Suit: "Spade", Rank: 3},
					{Suit: "Spade", Rank: 4},
					{Suit: "Spade", Rank: 5},
				},
				Name: "StraightFlush",
			},
		},
		{
			name: "Quads",
			args: [5]models.Card{
				{Suit: "Spade", Rank: 1},
				{Suit: "Heart", Rank: 1},
				{Suit: "Diamond", Rank: 1},
				{Suit: "Club", Rank: 1},
				{Suit: "Spade", Rank: 2},
			},
			want: models.Hand{
				Cards: [5]models.Card{
					{Suit: "Spade", Rank: 1},
					{Suit: "Heart", Rank: 1},
					{Suit: "Diamond", Rank: 1},
					{Suit: "Club", Rank: 1},
					{Suit: "Spade", Rank: 2},
				},
				Name: "Quads",
			},
		},
		{
			name: "FullHouse",
			args: [5]models.Card{
				{Suit: "Spade", Rank: 1},
				{Suit: "Heart", Rank: 1},
				{Suit: "Diamond", Rank: 1},
				{Suit: "Spade", Rank: 2},
				{Suit: "Heart", Rank: 2},
			},
			want: models.Hand{
				Cards: [5]models.Card{
					{Suit: "Spade", Rank: 1},
					{Suit: "Heart", Rank: 1},
					{Suit: "Diamond", Rank: 1},
					{Suit: "Spade", Rank: 2},
					{Suit: "Heart", Rank: 2},
				},
				Name: "FullHouse",
			},
		},
		{
			name: "Flush",
			args: [5]models.Card{
				{Suit: "Spade", Rank: 1},
				{Suit: "Spade", Rank: 2},
				{Suit: "Spade", Rank: 3},
				{Suit: "Spade", Rank: 4},
				{Suit: "Spade", Rank: 6},
			},
			want: models.Hand{
				Cards: [5]models.Card{
					{Suit: "Spade", Rank: 1},
					{Suit: "Spade", Rank: 2},
					{Suit: "Spade", Rank: 3},
					{Suit: "Spade", Rank: 4},
					{Suit: "Spade", Rank: 6},
				},
				Name: "Flush",
			},
		},
		{
			name: "Straight",
			args: [5]models.Card{
				{Suit: "Spade", Rank: 1},
				{Suit: "Heart", Rank: 2},
				{Suit: "Spade", Rank: 3},
				{Suit: "Spade", Rank: 4},
				{Suit: "Spade", Rank: 5},
			},
			want: models.Hand{
				Cards: [5]models.Card{
					{Suit: "Spade", Rank: 1},
					{Suit: "Heart", Rank: 2},
					{Suit: "Spade", Rank: 3},
					{Suit: "Spade", Rank: 4},
					{Suit: "Spade", Rank: 5},
				},
				Name: "Straight",
			},
		},
		{
			name: "Straight with A",
			args: [5]models.Card{
				{Suit: "Spade", Rank: 1},
				{Suit: "Heart", Rank: 10},
				{Suit: "Spade", Rank: 11},
				{Suit: "Spade", Rank: 12},
				{Suit: "Spade", Rank: 13},
			},
			want: models.Hand{
				Cards: [5]models.Card{
					{Suit: "Spade", Rank: 1},
					{Suit: "Heart", Rank: 10},
					{Suit: "Spade", Rank: 11},
					{Suit: "Spade", Rank: 12},
					{Suit: "Spade", Rank: 13},
				},
				Name: "Straight",
			},
		},
		{
			name: "Trips",
			args: [5]models.Card{
				{Suit: "Spade", Rank: 1},
				{Suit: "Heart", Rank: 1},
				{Suit: "Club", Rank: 1},
				{Suit: "Spade", Rank: 4},
				{Suit: "Spade", Rank: 5},
			},
			want: models.Hand{
				Cards: [5]models.Card{
					{Suit: "Spade", Rank: 1},
					{Suit: "Heart", Rank: 1},
					{Suit: "Club", Rank: 1},
					{Suit: "Spade", Rank: 4},
					{Suit: "Spade", Rank: 5},
				},
				Name: "Trips",
			},
		},
		{
			name: "TwoPair",
			args: [5]models.Card{
				{Suit: "Spade", Rank: 1},
				{Suit: "Heart", Rank: 1},
				{Suit: "Spade", Rank: 2},
				{Suit: "Heart", Rank: 2},
				{Suit: "Spade", Rank: 3},
			},
			want: models.Hand{
				Cards: [5]models.Card{
					{Suit: "Spade", Rank: 1},
					{Suit: "Heart", Rank: 1},
					{Suit: "Spade", Rank: 2},
					{Suit: "Heart", Rank: 2},
					{Suit: "Spade", Rank: 3},
				},
				Name: "TwoPair",
			},
		},
		{
			name: "Pair",
			args: [5]models.Card{
				{Suit: "Spade", Rank: 1},
				{Suit: "Heart", Rank: 1},
				{Suit: "Spade", Rank: 2},
				{Suit: "Heart", Rank: 3},
				{Suit: "Spade", Rank: 4},
			},
			want: models.Hand{
				Cards: [5]models.Card{
					{Suit: "Spade", Rank: 1},
					{Suit: "Heart", Rank: 1},
					{Suit: "Spade", Rank: 2},
					{Suit: "Heart", Rank: 3},
					{Suit: "Spade", Rank: 4},
				},
				Name: "Pair",
			},
		},
		{
			name: "HighCard",
			args: [5]models.Card{
				{Suit: "Spade", Rank: 1},
				{Suit: "Heart", Rank: 2},
				{Suit: "Spade", Rank: 3},
				{Suit: "Heart", Rank: 4},
				{Suit: "Spade", Rank: 6},
			},
			want: models.Hand{
				Cards: [5]models.Card{
					{Suit: "Spade", Rank: 1},
					{Suit: "Heart", Rank: 2},
					{Suit: "Spade", Rank: 3},
					{Suit: "Heart", Rank: 4},
					{Suit: "Spade", Rank: 6},
				},
				Name: "HighCard",
			},
		},
	}

	s := services.NewMyAppService(nil)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := s.GetHandService(tt.args); got != tt.want {
				t.Errorf("GetHandService() = %v, want %v", got, tt.want)
			}
		})
	}

}
