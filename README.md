# 🐘 Elephant
Docker orchestrator project 🐘.

I'm using the Docker SDK for this. 

## The concept

Elephants are big. And solid. And they're great.

So they can carry several containers.
For that, they need to eat a docker-compose like file. But it can be named the way you want (even something like `cool.yml` it's up to you).

You also gotta tell them to start and carry them, using the command `walk`.

Elephants are so so great that there are other commands they understand :)


## Commands

- `walk`: Start a stack using a docker-compose file. The CLI will also ask you to name your elephant, this is important.
- `ls`: List the elephants and the containers they carry.
- `stomp`: Your elephant will drop and stomp on all the containers he carries (in a short way, it stops the containers associated to your elephant).

## Examples

You can either use `go run main.go`, or `./elephant` if you previously run `go build -o elephant`.

- `elephant walk ~/garbageDir/cool.yml` or `elephant walk ./cool.yml`
- `elephant stomp myStack`
- `elephant ls`

I'd recommend to set your GOPATH env variable and then run `go install applinh/elephant` so you can directly use the `elephant` command anywhere.

## 🐘❤️

Feel free to git clone this, or open any pull request to contribute.

Much luv <3

Special thanks to [BoltDB](https://github.com/boltdb/bolt) and also [Docker](https://pkg.go.dev/github.com/docker/docker) for their work.


