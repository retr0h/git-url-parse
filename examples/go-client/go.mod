module example.com/client

go 1.22.2

replace github.com/retr0h/git-url-parse => ../../../git-url-parse/

require (
	github.com/lmittmann/tint v1.0.4
	github.com/retr0h/git-url-parse v0.0.0-00010101000000-000000000000
)

require github.com/chainguard-dev/git-urls v1.0.2 // indirect
