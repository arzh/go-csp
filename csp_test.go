package main

import "testing"

func DataReader(data []string) chan string {
	cr := make(chan string)

	go func() {
		for _, card := range data {
			cr <- card
		}
		close(cr)
	}()

	return cr
}

func TestCopy(t *testing.T) {
	data, testData := csp_getTestData()

	cr := DataReader(data)

	feed := Assemble(Copy(Disassemble(cr)))

	i := 0
	for {
		card, ok := <-feed
		if !ok {
			if len(testData) > i {
				t.Error("expecting more data")
			}
			t.Log("Channel Closed")
			break
		}

		if len(testData) <= i {
			if len(card) > 0 {
				t.Error("Recived more data then expected: Card:", card)
				break
			} else if len(card) == 0 {
				t.Log("Recived empty card")
				continue
			} else {
				t.Log("Channel not closed but ran out of tests")
				continue
			}
		}

		// Find the difference and make a string to help point it out
		if card != testData[i] {
			t.Log("Data Test Failed", i)

			t.Log("Card len:", len(card), "Test len:", len(testData[i]))

			diff := ""
			for j, _ := range testData[i] {
				if i < len(card) {
					if testData[i][j] != card[j] {
						diff += "^"
						break
					} else {
						diff += " "
					}
				} else {
					diff += "^"
				}
			}

			t.Errorf("Diff found at %d|\n\tCard:`%s`\n\tDiff: %s\n\tData:`%s`", len(diff)-1, card, diff, testData[i])

			break
		}
		i++
	}

}
