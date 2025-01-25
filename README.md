# Terminal GIFs API 📼

Use this API to present creative ways of your software commands.
Live in **http://terminalgifapi.com/api/v1/gif?commands=["cat README.md","Sleep 2s"]**.

# build image
docker build --rm -t vhspoc .

# debug container
docker run --rm -it -p 9001:80 -v .:/project -w /project vhspoc bash

# run server
go run cmd/server/main.go

## how to run w/ docker 🐳

```sh
# build image
make build-image

# debug container
make debug-container

# run server
go run cmd/server/main.go
```

From browser or [Bruno](./zarf/bruno/), open:

- http://terminalgifapi.com/api/v1/gif?commands=["Type \"echo 'Welcome to VHS!'\"","Enter","Type \"ls\"","Sleep 100ms","Enter","Sleep 2s"]
- http://localhost:9001/api/v1/gif?commands=["Type \"echo 'Welcome to VHS!'\"","Enter","Type \"ls\"","Sleep 100ms","Enter","Sleep 2s"]
- http://localhost:9001/api/v1/gif?commands=[
    "Type \"echo 'How to run locally w/ Docker'\"\n","Sleep 400ms","Enter","Sleep 200ms",
    "Type \"echo 'make build-image to build image'\"\n","Sleep 400ms","Enter","Sleep 200ms",
    "Type \"echo 'make debug-container to start a terminal from inside the container'\"\n","Sleep 400ms","Enter","Sleep 200ms",
    "Type \"echo 'go run cmd/server/main.go'\"\n","Sleep 400ms","Enter","Sleep 200ms",
    "Sleep 2s"]
- http://localhost:9001/api/v1/gif?commands=["Type \"cat README.md\"","Enter","Sleep 2s"]
- http://localhost:9001/api/v1/gif?commands=["Type \"less README.md\"","Enter","Sleep 2s"]
- http://localhost:9001/api/v1/gif?commands=["Type \"echo Hi\"","Enter","Sleep 2s"]

## Charmbracelets VHS

### how to run

Write instructions to `./output/demo.tape`, then run `docker run --rm -v $PWD:/vhs -w /vhs --name gifbuilder ghcr.io/charmbracelet/vhs output/demo.tape`.

## references

<!-- ![Anurag's GitHub stats](http://terminalgifapi.com/api/v1/gif?commands=["Type \"echo 'Welcome to VHS!'\"","Sleep 100ms","Enter","Sleep 100ms","Type \"ls -a\"","Sleep 100ms","Enter","Sleep 1s"])

![Anurag's GitHub stats](http://terminalgifapi.com/api/v1/mock) -->

<!-- http://terminalgifapi.com/api/v1/gif?commands=["Type \"echo 'Welcome to VHS!'\"","Sleep 100ms","Enter","Sleep 100ms","Type \"ls -a\"","Sleep 100ms","Enter","Sleep 1s"] -->
<!-- ![My GitHub Streak](http://github-readme-streak-stats.herokuapp.com?user=victorabarros) -->

- https://github.com/anuraghazra/github-readme-stats
- https://github.com/charmbracelet/vhs
- https://github.com/DenverCoder1/github-readme-streak-stats
- https://github.com/rahuldkjain/github-profile-readme-generator
- https://github.com/ryo-ma/github-profile-trophy

## [Licence](./LICENSE)

<!--
TODO

- improve dockerfile
  - create stage to copy code and build project
  - copy build to release fase
  - entrypoint to build
- backup https://github.com/charmbracelet/vhs/releases/download/v0.9.0/vhs_0.9.0_arm64.deb
- cache previous gifs
  - hash query to index it
  - store gifs on S3; or store locally and create a service to expire GIFs longers when directory is full?
  - persist hash -> GIF file path on redis (maybe better sqlite, because no needed of more service and it persist in disk)
- create homepage to introduce project

-->
