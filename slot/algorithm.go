package slot

import (

)

type (
	// saved in db
	CardOddsInDb struct {
		line int
		odd int
		ifFree bool
		ifBonus bool
		specificCards map[int]map[int]int
		/*
		// example for specificcards
		{
			{1: 10, 2: 10, 1: 0, 2: 10, 4: 10},
			{2: 11, 3: 11, 1: 0, 4:, 11},
			{4: 14, 4: 14, 1: 0}
		}
		*/
	}

	SingleLineOdd struct {
		card int
		odd int
		cards []int
	}

	CardType struct {
		line int
		odd int
		info []CardInfo
	}

	CardInfo struct {
		cardNum int
		info []int
	}
)

func init() {

}

func CreateLine() {
	rowCount, colCounr, lineCount := 4, 5, 50

	odds := map[int]map[int]int{
		10: map[int]int{3: 35, 4: 130, 5: 400}, // 10
		11: map[int]int{3: 40, 4: 150, 5: 450}, // J
		12: map[int]int{3: 50, 4: 200, 5: 600}, // Q
		13: map[int]int{3: 70, 4: 280, 5: 850}, // K
		14: map[int]int{3: 80, 4: 300, 5: 1000}, // A
		15: map[int]int{3: 100, 4: 400, 5: 1200}, // wine bottle
		16: map[int]int{3: 150, 4: 600, 5: 1800}, // knife
		17: map[int]int{3: 240, 4: 900, 5: 2700}, // commpass
		18: map[int]int{3: 450, 4:1800, 5: 5400} // anchors
	}

	lines := map[int][]{
		1: {1, 1, 1, 1, 1},
		2: {2, 2, 2, 2, 2},
		3: {3, 3, 3, 3, 3},
		4: {4, 4, 4, 4, 4},
		5: {1, 1, 2, 1, 1},
		6: {2, 2, 1, 2, 2},
		7: {3, 3, 4, 3, 3},
		8: {4, 4, 3, 4, 4},
		9: {1, 1, 3, 1, 1},
		10: (2, 2, 4, 2, 2),
		11: {3, 3, 1, 3, 3},
		12: {4, 4, 2, 4, 4},
		13: {1, 1, 4, 1, 1},
		14: {2, 2, 3, 2, 2},
		15: {3, 3, 2, 3, 3},
		16: {4, 4, 1, 4, 4},
		17: {1, 2, 1, 2, 1},
		18: {2, 1, 2, 1, 2},
		19: {3, 4, 3, 4, 3},
		20: {4, 3, 4, 3, 4},
		21: {1, 2, 2, 2, 1},
		22: {2, 1, 1, 1, 2},
		23: {3, 4, 4, 4, 3},
		24: {4, 3, 3, 3, 4},
		25: {1, 2, 3, 2, 1},
		26: {2, 1, 4, 1, 2},
		27: {3, 4, 1, 4, 3},
		28: {4, 3, 2, 3, 4},
		29: {1, 3, 1, 3, 1},
		30: {2, 4, 2, 4, 2},
		31: {3, 1, 3, 1, 3},
		32: {4, 2, 4, 2, 4},
		33: {1, 3, 2, 3, 1},
		34: {2, 4, 3, 4, 2},
		35: {3, 1, 2, 1, 3},
		36: {4, 2, 3, 2, 4},
		37: {1, 3, 3, 3, 1},
		38: {2, 4, 4, 4, 2},
		39: {3, 1, 1, 1, 3},
		40: {4, 2, 2, 2, 4},
		41: {1, 4, 1, 4, 1},
		42: {2, 3, 2, 3, 2},
		43: {3, 2, 3, 2, 3},
		44: {4, 1, 4, 1, 4},
		45: {1, 4, 3, 4, 1},
		46: {2, 3, 4, 3, 2},
		47: {3, 2, 1, 2, 3},
		48: {4, 1, 2, 1, 4},
		49: {1, 4, 3, 2, 1},
		50: {4, 1, 2, 3, 4}
	}

	lineNums := map[int][]{
		1: {3, 4, 5}
		2: {3, 4, 5}
		3: {3, 4, 5}
		4: {3, 4, 5}
		5: {3, 4, 5}
		6: {3, 4, 5}
		7: {3, 4, 5}
		8: {3, 4, 5}
		9: {3, 4, 5}
		10: {3, 4, 5}
		11: {3, 4, 5}
		12: {3, 4, 5}
		13: {3, 4, 5}
		14: {3, 4, 5}
		15: {3, 4, 5}
		16: {3, 4, 5}
		17: {3, 4, 5}
		18: {3, 4, 5}
		19: {3, 4, 5}
		20: {3, 4, 5}
		21: {3, 4, 5}
		22: {3, 4, 5}
		23: {3, 4, 5}
		24: {3, 4, 5}
		25: {3, 4, 5}
		26: {3, 4, 5}
		27: {3, 4, 5}
		28: {3, 4, 5}
		29: {3, 4, 5}
		30: {3, 4, 5}
		31: {3, 4, 5}
		32: {3, 4, 5}
		33: {3, 4, 5}
		34: {3, 4, 5}
		35: {3, 4, 5}
		36: {3, 4, 5}
		37: {3, 4, 5}
		38: {3, 4, 5}
		39: {3, 4, 5}
		40: {3, 4, 5}
		41: {3, 4, 5}
		42: {3, 4, 5}
		43: {3, 4, 5}
		44: {3, 4, 5}
		45: {3, 4, 5}
		46: {3, 4, 5}
		47: {3, 4, 5}
		48: {3, 4, 5}
		49: {3, 4, 5}
		50: {3, 4, 5}
	}

	// First: stripped of all possible odds of winning.
	var singleLineOdds []SingleLineOdd{}

	for k, v := range odds {
		for k2, v2 := range v {
			var cards []int
			for i := 0; i < k2; i++ {
				cards = append(cards, k)
			}
			singleLineOdd := SingleLineOdd{
				card: k,
				odd: v2,
				cards: cards
			}

			singleLineOdds = append(singleLineOdds, singleLineOdd)
		}
	}

	// Second: each line to generate the respective odds combination
	totalLineOdds := make(map[int]map[int]TotalLineOdd)
	for key, val := range lines {
		for k, v := range singleLineOdds {
			secCards := make(map[int]map[int]int)
			for _, v2 := range v.cards {
				tmp := map[int]int{
					v: v2
				}

				secCards = append(secCards, tmp)
			}

			totalLineOdd := TotalLineOdd{
				line: key,
				single: v,
				cards: secCards
			}

			totalLineOdds = append(totalLineOdds[key], totalLineOdd)
		}
	}

	// Third: combination of 50 lines
	// Get the totaL num of odds
	totalOddsNum := 0
	for _, v := range odds {
		tmp := len(v)
		totalOddsNum += tmp
	}

	for i := 0; i < lineCount; i++ {
		tmpLineOdd := getCardOddsFromGivenLines(i, totalLineOdds)
	}
}

func getCardOddsFromGivenLines(i int, ) () {

}

// n total line num
// m need line num
func getNeedLinesFromGivenLines(n, m int) {
	result := make([][]int, 0, mathCombination(n, m))

	indexs := make([]int, n)
	for i := 0; i < n; i++ {
		if i < m {
			indexs[i] = 1
		} else {
			indexs[i] = 0
		}
	}

	result = addTo(result, indexs)
}

func mathCombination(n, m int) int {
	return mathFactorial(n) / (mathFactorial(n - m) * mathFactorial(m))
}

func mathFactorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}

	return result
}

// arr map[int]int{3, 4, 5} the number of sheets can be winning
// n int numerical calculation
func getCombinationFromGivenMapAndNumber(arr map[int]int, n int) map[int][]int {
	retArr := make(map[int][]int)
	tmpArr := make([][]int)
	retArr[1], tmpArr[0] = []int{1}, 1

	for i := 2; i <= n; i++ {
		retArr[i] := []int{
			i
		}

		for _, val := range tmpArr {
			val = append(val, i)
			retArr[i] = append(retArr[i], val)
			tmpArr = append(tmpArr, val)
		}

		tmpArr = append(tmpArr, []int{i})
	}

	return retArr
}

func newFunc() {

}
