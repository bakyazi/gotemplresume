templ:
	@templ generate
run:
	@npx tailwindcss -i ./resources/global.css -o ./styles/output.css
	@templ generate
	@go run ./cmd/main.go