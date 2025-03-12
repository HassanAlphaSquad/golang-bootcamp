// module github.com/HassanAlphaSquad/golang-bootcamp/week_2/main

// go 1.23.6

// module github.com/HassanAlphaSquad/golang-bootcamp/hellogo

// go 1.23.6

// module github.com/HassanAlphaSquad/golang-bootcamp/mystrings

// go 1.23.6


// replace github.com/HassanAlphaSquad/golang-bootcamp/mystrings v0.0.0 => ../mystrings

// require (
// 	github.com/HassanAlphaSquad/golang-bootcamp/week_2/hellogo v0.0.0
// )

module example.com/username/hellogo

go 1.23.0

replace github.com/HassanAlphaSquad/golang-bootcamp/mystrings v0.0.0 => ../mystrings

require (
	example.com/username/mystrings v0.0.0
)