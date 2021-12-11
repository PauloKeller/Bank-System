`run tests: go test ./... -coverprofile cover.out`
`generate cover profile: go test ./... -coverprofile cover.out`
`open profile as html: go tool corver -html=cover.out`