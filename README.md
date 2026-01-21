# TermGIFforge 👾

[![Go Report Card](https://goreportcard.com/badge/github.com/victorabarros/termgifforge)](https://goreportcard.com/report/github.com/victorabarros/termgifforge) ![GitHub License](https://img.shields.io/github/license/victorabarros/TermGIFforge)

**Turn your terminal recordings into captivating GIFs effortlessly.**
Ideal for tutorials, documentation, and sharing terminal workflows.

<p align="center">
  <img src="http://terminalgifapi.com/api/v1/gif?commands=[%20%22Set%20FontSize%2050%22,%20%22Set%20TypingSpeed%2075ms%22,%20%22Type%20\%22echo%20\%22%22,%20%22Set%20TypingSpeed%20500ms%22,%20%22Type%20\%22%27YEY\%22%22,%20%22Set%20TypingSpeed%2075ms%22,%20%22Type%20\%22!!!!%27\%22%22,%20%22Sleep%20100ms%22,%20%22Enter%22,%20%22Sleep%202s%22]"/>
</p>

<!-- curl -fsSL https://raw.githubusercontent.com/victorabarros/victorabarros/master/scripts/btc_logo.sh | bash -->


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

http://terminalgifapi.com/api/v1/gif?commands=[
"Set FontSize 25",
  "Set Height 900",
  "Set Width 700",
"Sleep 1s",
"Hide",
"Type \"curl -fsSL https://raw.githubusercontent.com/victorabarros/victorabarros/master/scripts/btc_logo.sh | bash\"",
"Show",
"Sleep 1s",
"Enter",
"Sleep 10s"
]

- http://localhost:9001/api/v1/gif?commands=[
  "Set FontSize 25",
  "Set Height 900",
  "Set Width 700",
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
  <img src="http://terminalgifapi.com/api/v1/gif?commands=[%20%22Type%20\%22echo%20How%20to%20run%20locally%20w/%20Docker\%22\n%22,%22Sleep%20400ms%22,%22Enter%22,%22Sleep%20200ms%22,%20%22Type%20\%22echo%20make%20build-image%20to%20build%20image\%22\n%22,%22Sleep%20400ms%22,%22Enter%22,%22Sleep%20200ms%22,%20%22Type%20\%22echo%20make%20debug-container%20to%20start%20a%20terminal%20from%20inside%20the%20container\%22\n%22,%22Sleep%20400ms%22,%22Enter%22,%22Sleep%20200ms%22,%20%22Type%20\%22echo%20%27go%20run%20cmd/server/main.go%27\%22\n%22,%22Sleep%20400ms%22,%22Enter%22,%22Sleep%20200ms%22,%20%22Sleep%202s%22]"/>

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
  <img src="http://terminalgifapi.com/api/v1/gif?commands=[%20%22Set%20FontSize%2025%22,%20%22Set%20Height%20900%22,%20%22Set%20Width%20700%22,%20%22Sleep%201s%22,%20%22Hide%22,%20%22Type%20\%22curl%20-fsSL%20https://raw.githubusercontent.com/victorabarros/victorabarros/master/scripts/btc_logo.sh%20|%20bash\%22%22,%20%22Show%22,%20%22Sleep%201s%22,%20%22Enter%22,%20%22Sleep%2010s%22%20]" height="600px"/>
  <br/>
  <br/>
  <a href="https://victor.barros.engineer/support" target="_blank">
    <img src="https://bitcoin.org/img/icons/logotop.svg?1671880122" height="40px">
  </a>
  <br/>
  <a href="https://www.buymeacoffee.com/victorbarros" target="_blank">
    <img src="https://cdn.buymeacoffee.com/buttons/v2/default-yellow.png" height="45px">
  </a>
  <br/>
  <img src="http://terminalgifapi.com/api/v1/gif?commands=[%20%22Set%20FontSize%2060%22,%20%22Type%20\%22I%20am%20happy%20to%20be%20honored%20with%20your%20support!%20S2S2\%22%22,%20%22Sleep%202s%22%20]" height="200px"/>
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
