package services

import (
	"github.com/kult0922/poker/backend/domain/analytics"
)

func (s *MyAppService) GetPreflopHandRange() map[string]int {

	return analytics.CalcPreflopRank(s.rankMap, s.rankMapFlush)
}
