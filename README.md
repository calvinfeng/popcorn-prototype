# Movie Gopher
Coming soon! This shall replace my Consilium project. 

## Setup
### Dependency Management with `dep`
First of all, get `dep` for dependency management

```
go get -u github.com/golang/dep/cmd/dep
```

If you are using Mac OS X then it's even easier

```
brew install dep 
brew upgrade dep
```

I fucking love Homebrew on Mac. It has everything!

### Database
I am going to use PostgreSQL for this project, so let's create one. The superuser on my computer is `cfeng` so I will
use that to create a database named `goyak_development`

If you don't have a role or wish to create a separate role for this project, then just do the following
```
$ psql postgres
postgres=# create role <name> superuser login;
```

Example:
```
$ psql postgres
postgres=# create role cfeng superuser login;
```

Then create a database named `movie_gopher_development`
```
$ psql postgres
postgres=# create database movie_gopher_development owner=cfeng;
```

Then just exit with `\q`

Actually just in case you don't remember the password to your `ROLE`, do the following
```
postgres=# alter user <your_username> with password <whatever you like>
```

I did mine with
```
postgres=# alter user cfeng with password "cfeng";
```