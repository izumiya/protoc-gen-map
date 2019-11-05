# Contributing

Thank you for contributing to protoc-gen-map.

Please follow the code of conduct when interacting with the protoc-gen-map project.

## Pull Request Process

The recommended PR process is as follows.

1. Please first discuss the change you wish to make via issue.
2. Upgrade with `go get -u github.com/jackskj/protoc-gen-map`
3. cd `$GOPATH/src/github.com/jackskj/protoc-gen-map`
4. Reflect your changes in example files as well as the README. 
5. Update example and test .proto, .pb.go and .pg.map.go files when needed.
   You can use 'make gen' defined in the Makefile.
   Please ensure to include those changes in the merge request.
6. Update BUILD files, as well as dependencies in bazel directory if needed.
   You can use gazelle functions (gazelle, repos, and fix) from the Makefile. 
7. Make sure that the build succedes with `go build` or `make build` 
   and all tests pass with `make test`
8. Make sure to follow [Effective Go](https://golang.org/doc/effective_go.html)  
   as well as [Go Code Review Comments](https://golang.org/wiki/CodeReviewComments)

## Code of Conduct


Please do not make aggressive, harassing or condescending comments based on someone's
age, body size, disability, ethnicity, gender identity and expression, level of experience,
nationality, personal appearance, race, religion, or sexual identity and
orientation.

Please do not troll, or make insulting/derogatory comments  pr personal/political attacks
