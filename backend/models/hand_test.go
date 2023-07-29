package models_test

import (
	"testing"

	"github.com/kult0922/go-react-blog/backend/models"
)

func TestHandName(t *testing.T) {
	tests := []struct {
		name string
		args [5]models.Card
		want string
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
			want: "RoyalFlush",
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
			want: "StraightFlush",
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
			want: "Quads",
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
			want: "FullHouse",
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
			want: "Flush",
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
			want: "Straight",
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
			want: "Straight",
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
			want: "Trips",
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
			want: "TwoPair",
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
			want: "Pair",
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
			want: "HighCard",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := models.HandName(tt.args); got != tt.want {
				t.Errorf("GetHandService() = %v, want %v", got, tt.want)
			}
		})
	}

}
