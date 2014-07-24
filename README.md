hsh
===

Tired of typing out `heroku #{thing_I_want_to_do} -a #{my_long_app_name}` ? `hsh` solves this for you! Just start the shell with the name of the application that you would like to execute your commands against:

```
➜  hsh git:(master) hsh elixir-docset
elixir-docset > apps:info
=== elixir-docset
Addons:        loggly:mole
Git URL:       git@heroku.com:elixir-docset.git
Owner Email:   xtjraymond.x@gmail.com
Region:        us
Repo Size:     1M
Slug Size:     1M
Stack:         cedar
Web URL:       http://elixir-docset.herokuapp.com/
elixir-docset > restart
Restarting dynos... done
elixir-docset >
```
To exit the shell type `:exit`

Shell specific commands are invoked by prepending a colon character, an approach inspired by the fabulous [gitsh](https://github.com/thoughtbot/gitsh).

Installing
----------

The hope is that this will be installable via package managers in the future. For now, the steps are:

1. Setup a working golang dev environment. I won't write how to do this as it's been covered extensively in other places. Here is the official Golang [install page](http://golang.org/doc/install).
2. `go get github.com/timraymond/hsh`
3. `cd $(echo $GOPATH)/src/github.com/timraymond/hsh`
4. `go get`
5. If you're on OS X, you may have to install GNU readline through Homebrew. Simply `brew install readline` and take note of the post install documentation on the CPPFLAGS and LDFLAGS.
6. Change the paths under the LDFLAGS and CFLAGS vars under `$GOPATH/src/github.com/shavac/readline/readline.go` to match the ones that came from Homebrew
7. Go back to `$GOPATH/src/github.com/timraymond` and run `go build`
8. `go install` to have it be linked system-wide, and you're done!

Coming Soon
-----------

- Like I said, the install steps above suck, so I'd like to make this as easy as `brew install hsh` or `apt-get install hsh`
- The ability to switch the target application without exiting the shell
- Run system commands without exiting the shell
