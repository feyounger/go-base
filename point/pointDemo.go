package main

func main() {
	done, _ := bbb()
	done()
}

func bbb() (done func(), _ error) {
	return func() {
		print("bbb: surprise!")
		done()
	}, nil
}
