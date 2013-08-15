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
	failed := false
	for {
		card, ok := <-feed
		if !ok {
			break
		}

		if card != testData[i] {
			t.Log("Data Test Failed", i, "| Card:", card, "Data:", testData[i])
			t.Fail()
			failed = true
		}
		i++
	}

	if failed {
		t.Fail()
	}
}
