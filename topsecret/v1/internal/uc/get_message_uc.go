package uc

import "strings"

func GetMessage(messages ...[]string) (msg string) {
	sanitizedMms, majorCount := sanitizeMss(messages)
	completedMms := completeEmptySpaces(sanitizedMms, majorCount)
	var mss = make([]string, majorCount)

	for i := 0; i < majorCount; i++ {
		for j := range completedMms {
			if completedMms[j][i] != "" && !inArray(mss, completedMms[j][i]) {
				mss[i] = completedMms[j][i]
			}
		}
	}

	return strings.Join(mss, " ")
}

func sanitizeMss(m [][]string) ([][]string, int) {
	var mssStrings [][]string
	majorCount := 0

	for _, row := range m {
		joinRow := strings.Join(row, " ")
		rowClear := strings.TrimLeft(joinRow, " ")
		rowClear = strings.ReplaceAll(rowClear, " ", ",")
		mssStrings = append(mssStrings, strings.Split(rowClear, ","))
		lenCurrentRow := len(strings.Split(rowClear, ","))

		if lenCurrentRow > majorCount {
			majorCount = lenCurrentRow
		}
	}

	return mssStrings, majorCount
}

func completeEmptySpaces(m [][]string, majorCount int) [][]string {
	var completedSpaces []string

	for i, row := range m {
		if len(row) < majorCount {
			for j := 0; j < majorCount-len(row); j++ {
				completedSpaces = append(completedSpaces, "")
			}

			m[i] = append(completedSpaces, row...)
		}
	}

	return m
}

func inArray(array []string, text string) bool {
	for _, element := range array {
		if element == text {
			return true
		}
	}
	return false
}
