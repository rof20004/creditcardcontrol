build:
	mkdir -p functions
	cd lambda/go && \
	go get ./... && \
	go build -o ../../functions/create_card card/create_card/create_card.go && \
	go build -o ../../functions/delete_card card/delete_card/delete_card.go