package home

const (
	minAttempts = 3
	maxAttempts = 10
)

var attemptsCounter = 3

func increaseAttempts() {
	if attemptsCounter < maxAttempts {
		attemptsCounter++
	}
}

func decreaseAttempts() {
	if attemptsCounter > minAttempts {
		attemptsCounter--
	}
}
