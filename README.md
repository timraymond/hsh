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

Coming Soon
-----------

- The ability to switch the target application without exiting the shell
- Run system commands without exiting the shell
