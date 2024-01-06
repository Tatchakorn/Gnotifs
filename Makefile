run:
	go run ./TronNotif/TSP/tsp.go
test:
	go test ./TronNotif/TSP/... -coverprofile=c.out
	go tool cover -html=c.out
fmt:
	go fmt ./TronNotif/TSP/...
runpy:
	python main.py
testpy:
	pytest --cache-clear --verbose -s
fmtpy:
	autopep8 . --recursive --in-place --pep8-passes 2000 --verbose