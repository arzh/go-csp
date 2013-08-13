package main

import "unicode/utf8"

func Copy(west chan rune) (east chan rune) {
	east = make(chan rune, 3)
	
	go func() {
		for {
			c, ok := <-west
			if !ok {
				close(east)
				return // Break out when west is closed
			}
			east <- c
		}
	}()
	
	return
}

func Squash(west chan rune) (east chan rune) {
	east = make(chan rune, 3)
	go csp_squash(west, east)
	return
}

func csp_squash(west, east chan rune) {
	for {
		c, ok := <-west
		if !ok {
			close(east)
			return
		}

		if string(c) == "*" {
			c2, ok := <-west
			if !ok {
				east <- c
				close(east)
				return
			}

			if string(c2) == "*" {
				r, _ := utf8.DecodeRuneInString("^")
				east <- r
				continue
			}
		}

		east <- c
	}
}

func Assemble(e chan rune) chan string {
	lp := make(chan string)

	go func(){
		s := ""

		for {
			c, ok := <-e
			if !ok {
				lp <- s
				close(lp)
				return
			} else {
				s += string(c)
				if len(s) == 125 {
					s += " "
					lp <- s
					s = ""
				}
			}
		}
	}()

	return lp
}

func Disassemble(cr chan string) chan rune {
	c := make(chan rune, 3)

	go func(){
		for {
			card, ok := <-cr
			if !ok {
				close(c)
				return
			}

			for _, char := range card {
				c <- char
			}

		}
	}()

	return c
}