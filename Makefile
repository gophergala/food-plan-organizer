download:
	curl -L https://www.ars.usda.gov/SP2UserFiles/Place/12354500/Data/SR27/dnload/sr27asc.zip -o sr27asc.zip
	unzip sr27asc.zip -d data
	rm sr27asc.zip

database:
	go run cmd/parse/main.go -data.dir data -database default.db

server:
	go run cmd/serve/main.go -database default.db

atom:
	/Applications/Atom-Shell.app/Contents/MacOS/Atom .