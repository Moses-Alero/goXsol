start:
	./templ generate --watch --proxy="http://localhost:3000" --cmd="go run ."

dev:
	./templ generate
	go run .
