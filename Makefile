build:
	mkdir -p functions
	cd lambda/go
	go get ./...
	cd -
	go build -o functions/create_card lambda/go/card/create_card/create_card.go
	go build -o functions/delete_card lambda/go/card/delete_card/delete_card.go