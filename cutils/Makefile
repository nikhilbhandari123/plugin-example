NAME=myplugin.so


.PHONY: build
build: 
	go build  -buildmode=plugin -o $(NAME) .	

.PHONY: clean
clean:
	rm -rf $(NAME)
