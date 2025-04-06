package entities

type Question struct {
	ID      int
	QuizID  int
	Number  int
	Text    string
	Answers []string
	Right   int
}

type Quiz struct {
	User  User
	Theme string
}
