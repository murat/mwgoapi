test:
	go test ./... -coverprofile=cover.out
	curl -Ls https://coverage.codacy.com/get.sh -o codacy.sh && \
	bash ./codacy.sh report -s --force-coverage-parser go -r cover.out -t ${CODACY_PROJECT_TOKEN}
