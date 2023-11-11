# SliverX

Customized version of Sliver C2 Framework.

- How to compile after debuggin:

1) First run the "make" command in the main SliverX folder
2) Then run each of these customized commands to build binaries sliver-server and sliver-client, the ones that "make" builds do not work.

Server binary: 

go build -ldflags "-X github.com/bishopfox/sliver/client/version.Version=1.0.0 -X github.com/bishopfox/sliver/client/version.CompiledAt=November -X github.com/bishopfox/sliver/client/version.GithubReleasesURL=github.com -X github.com/bishopfox/sliver/client/version.GitCommit=HostFileChanger -X github.com/bishopfox/sliver/client/version.GitDirty=RedTeam" -tags osusergo,netgo,go_sqlite,server ./server/main.go

Client binary:
go build -ldflags "-X github.com/bishopfox/sliver/client/version.Version=1.0.0 -X github.com/bishopfox/sliver/client/version.CompiledAt=November -X github.com/bishopfox/sliver/client/version.GithubReleasesURL=github.com -X github.com/bishopfox/sliver/client/version.GitCommit=HostFileChanger -X github.com/bishopfox/sliver/client/version.GitDirty=RedTeam" -tags osusergo,netgo,go_sqlite,server ./client/main.go

Version Addon Description
----------------
## SliverX v1.0.0

Customized version of Sliver.

Addons:

1) Added an implant windows command "modifyhostsfile" which takes as parameters a domain and an ip.
Example:

modifyhostsfile new -d www.newdomain.com -i 1.2.3.4 -t 40000
