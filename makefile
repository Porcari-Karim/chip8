

build:
	go build -o ./bin/chip8.exe ./cmd/chip8/main.go


run:
	go run ./cmd/chip8/main.go


run_executable:
	./bin/chip8.exe

test_samples:
	go run ./test/random_testing/$(FILE)