module local/tmp/mod-app

go 1.19

require (
	local/tmp/mod1 v0.0.0
	local/tmp/mod2 v0.0.0
)

require github.com/google/uuid v1.6.0 // indirect

replace (
	local/tmp/mod1 => ../mod1
	local/tmp/mod2 => ../mod2
)
