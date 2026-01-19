# TermGIFforge 👾

[![Go Report Card](https://goreportcard.com/badge/github.com/victorabarros/termgifforge)](https://goreportcard.com/report/github.com/victorabarros/termgifforge) ![GitHub License](https://img.shields.io/github/license/victorabarros/TermGIFforge)

**Turn your terminal recordings into captivating GIFs effortlessly.**
Ideal for tutorials, documentation, and sharing terminal workflows.

<p align="center">
  <img src="http://terminalgifapi.com/api/v1/gif/39323864-6331-3838-3761-333566623361"/>
</p>

Use this API to present creative ways of your software commands.

## Table of Contents

- [TermGIFforge 👾](#termgifforge-)
  - [Table of Contents](#table-of-contents)
  - [Instructions](#instructions)
  - [How to run](#how-to-run)
    - [Locally w/ docker 🐳](#locally-w-docker-)
  - [Troubleshooting](#troubleshooting)
  - [Contributing](#contributing)
  - [Support](#support)
  - [References](#references)

## Instructions

**TermGIFforge** uses the [VHS](https://github.com/charmbracelet/vhs) to generate the GIFs, [here you'll find the command references](https://github.com/charmbracelet/vhs?tab=readme-ov-file#vhs-command-reference) to generate your own GIFs.
Try:

- Having the commands on query param, like:
  - `http://terminalgifapi.com/api/v1/gif?commands=["Type \"echo 'The Magic Happens Here'\"","Enter","Sleep 2s"]`
  - `http://terminalgifapi.com/api/v1/gif?commands=[
      "Set FontSize 50",
      "Set TypingSpeed 75ms",
      "Type \"echo \"",
      "Set TypingSpeed 500ms",
      "Type \"'YEY\"",
      "Set TypingSpeed 75ms",
      "Type \"!!!'\"",
      "Sleep 100ms",
      "Enter",
      "Sleep 2s"]`
  - `http://terminalgifapi.com/api/v1/gif?commands=["Type \"echo 'Welcome to VHS!'\"","Enter","Type \"ls\"","Sleep 100ms","Enter","Sleep 2s"]`
- Having the commands on the requesst body, like:
  - `curl -X POST http://terminalgifapi.com/api/v1/gif --header 'content-type: application/json' --data '{"commands": ["Type \"echo '\''Welcome t00 V0HS!'\''\"","Enter","Type \"ls\"","Sleep 100ms","Enter","Sleep 2s"]}'`
- Using the GIF id on the route param:
  - `http://terminalgifapi/api/v1/gif/33613662-3838-3532-3532-313636613838`

<!--
- http://terminalgifapi.com/api/v1/gif?commands=[
    "Type \"echo 'How to run locally w/ Docker'\"\n","Sleep 400ms","Enter","Sleep 200ms",
    "Type \"echo 'make build-image to build image'\"\n","Sleep 400ms","Enter","Sleep 200ms",
    "Type \"echo 'make debug-container to start a terminal from inside the container'\"\n","Sleep 400ms","Enter","Sleep 200ms",
    "Type \"echo 'go run cmd/server/main.go'\"\n","Sleep 400ms","Enter","Sleep 200ms",
    "Sleep 2s"]
- http://terminalgifapi.com/api/v1/gif?commands=["Type \"echo 'Welcome to VHS!'\"","Sleep 100ms","Enter","Sleep 2s"]

- http://localhost:9001/api/v1/gif?commands=[
  "Set FontSize 25",
  "Set Height 800",
  "Set TypingSpeed 80ms",
  "Hide",
  "Type \"bash btc_logo.sh\"",
  "Enter",
  "Show",
  "Sleep 5s"] -->

## How to run

### Locally w/ docker 🐳

<!--
make build-image
make debug-container
go run cmd/server/main.go
 -->

<p align="center">
  <img src="http://terminalgifapi.com/api/v1/gif/62643236-3530-3762-6262-366563626532"/>
</p>

Then, from browser or [Bruno](./zarf/bruno/), open:

- [http://localhost:9001/api/v1/gif?commands=\["Type \\"echo Hi\\"","Enter","Sleep 2s"\]](http://localhost:9001/api/v1/gif?commands=["Type%20\"echo%20Hi\"","Enter","Sleep%202s"])

## Troubleshooting

- How to encode the URL to use in the HTML tag?
  - Simply enter your URL to the browser and it'll parse it to you.
- How to force image (cache) update in README.rst on GitHub?
  - You have to delete the cache by run `curl -X PURGE {url of cached badge image}` and refresh the page [(reference)](https://stackoverflow.com/questions/26898052/how-to-force-image-cache-update-in-readme-rst-on-github).

## Contributing

Contributions are welcome! Feel free to open issues or fork the repo and submit pull requests to enhance the project.

## Support

<p>
  <img src="http://terminalgifapi.com/api/v1/gif?commands=[%20%22Set%20FontSize%2060%22,%20%22Type%20\%22I%20am%20happy%20to%20be%20honored%20with%20your%20support!%20S2S2\%22%22,%20%22Sleep%202s%22%20]" height="200px"/>
  <br/>
  <br/>
  <a href="https://victor.barros.engineer/support" target="_blank">
    <img src="https://bitcoin.org/img/icons/logotop.svg?1671880122" height="40px">
  </a>
  <br/>
  <a href="https://www.buymeacoffee.com/victorbarros" target="_blank">
    <img src="https://cdn.buymeacoffee.com/buttons/v2/default-yellow.png" height="45px">
  </a>
</p>

## References

- https://github.com/charmbracelet/vhs

---

<p align="center">
  <br/>
  Made in Brazil
  <br/>
  <img src="https://github.com/victorabarros/ura-bot/blob/main/assets/BrazilFlag.png" height="30px"/>
</p>

<!--
## how to run in production

```
# install docker
curl -fsSL https://raw.githubusercontent.com/victorabarros/victorabarros/master/scripts/install_docker_ubuntu.sh | bash

# install make
apt install -y make

# clone project
git clone https://github.com/victorabarros/TermGIFforge.git

cd TermGIFforge

cp .env.example .env.production

# write variables to .env.production

make build-image
make compile ENV_FILE=.env.production
make run-app ENV_FILE=.env.production PORT=80
```
 -->

<!--
- https://github.com/anuraghazra/github-readme-stats
- https://github.com/DenverCoder1/github-readme-streak-stats
- https://github.com/rahuldkjain/github-profile-readme-generator
- https://github.com/ryo-ma/github-profile-trophy
- [![Star History Chart](https://api.star-history.com/svg?repos=getumbrel/umbrel&type=Date)](https://star-history.com/#getumbrel/umbrel&Date)
-->

<!--
TODO

- honeybadger https://app.honeybadger.io/projects/130099/install/go
- codeclimate
- sonar
- improve dockerfile
  - cmds to build image using mac or linux
    - receive OS as arg and select script to install vhs
  - create stage with shared volume and build project
  - copy build to release fase
  - entrypoint to run builded
- implement workers to limit GIF processing at same time
- better readme https://github.com/Azure-Samples/deepseek-azure-javascript?tab=readme-ov-file#deepseek-on-azure---javascript-demos
- write article/post
- post on
  - ask friends to give star
  - https://x.com/i/communities/1685641800449462272,
  - gopher discord,
  - ~gopher slack~ https://gophers.slack.com/archives/C8VFRARPY/p1738249701304709
  - "show HN" (https://news.ycombinator.com/show)...
  - msg: "If you love the beauty of what charmbracelet have being doing for CLI tools, you'll love this.
You no longer needs to install VHS to create your GIFs, I just released the TermGIFforge, the VHS as an API!"

-->
