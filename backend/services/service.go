package services

import (
	"database/sql"
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
)

// 1. サービス構造体を定義
type MyAppService struct {
	// 2. フィールドに sql.DB 型を含める
	db           *sql.DB
	rankMap      map[int]int
	rankMapFlush map[int]int
}

// コンストラクタの定義
func NewMyAppService(db *sql.DB) *MyAppService {
	file, err := os.Open("data/hand.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	primeMap := map[string]int{
		"2": 2,
		"3": 3,
		"4": 5,
		"5": 7,
		"6": 11,
		"7": 13,
		"8": 17,
		"9": 19,
		"T": 23,
		"J": 29,
		"Q": 31,
		"K": 37,
		"A": 41,
	}

	rankMap := make(map[int]int)
	rankMapFlush := make(map[int]int)

	for _, fields := range records {
		hand := fields[5]
		// arr A A A 2 3
		arr := strings.Split(hand, " ")

		hash := 1
		for _, v := range arr {
			prime, _ := primeMap[v]
			hash = hash * prime
			rank, _ := strconv.Atoi(fields[0])

			handName := fields[6]
			if (handName == "SF") || (handName == "F") {
				rankMapFlush[hash] = rank
			} else {
				rankMap[hash] = rank
			}

		}
	}

	return &MyAppService{
		db:           db,
		rankMap:      rankMap,
		rankMapFlush: rankMapFlush,
	}
}
