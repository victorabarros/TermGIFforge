# Terminal GIFs API 📼

Use this API to present creative ways of your software commands.
Live in **http://terminalgifapi.com/api/v1/gif?commands=["Type \"echo 'Welcome to VHS!'\"","Enter","Type \"ls\"","Sleep 100ms","Enter","Sleep 2s"]**.

## how to run w/ docker 🐳

```sh
# build image
docker build --rm -t vhspoc .

# debug container
docker run --rm -it -p 9001:80 -v .:/project -w /project vhspoc bash

# run server
go run cmd/server/main.go
```

From browser or [Bruno](./zarf/bruno/), open **http://localhost:9001/api/v1/gif?commands=["Type \"echo 'Welcome to VHS!'\"","Enter","Type \"ls\"","Sleep 100ms","Enter","Sleep 2s"]**.

## Charmbracelets VHS

### how to run

Write instructions to `./demo.tape`, then run `docker run --rm -v $PWD:/vhs -w /vhs --name gifbuilder ghcr.io/charmbracelet/vhs demo.tape`.

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

- cache previous gifs
  - hash query to index it
  - store gifs on S3; or store locally and create a service to expire GIFs longers when directory is full?
  - persist hash -> GIF file path on redis (maybe better sqlite, because no needed of more service and it persist in disk)
- understand why execution takes longer w/ time; temporary work around; reset aplication every X minutes
- create homepage to introduce project

-->
