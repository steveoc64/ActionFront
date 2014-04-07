all: ActionFrontOne

ActionFrontOne: ActionFrontOne.go
	go build -o ActionFrontOne ActionFrontOne.go
