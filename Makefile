all: gde1

gde1: gameDataEditor.go gamedatadb/gamedatadb.go
	go build -o gde1 gameDataEditor.go
