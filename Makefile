templ:
	@templ generate
run:
	@npx tailwindcss -i ./resources/global.css -o ./resources/output.css
	@templ generate
	@go run ./cmd/main.go