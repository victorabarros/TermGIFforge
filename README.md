# TermGIFforge 👾

On-the-Fly Terminal GIF Effortlessly Generation
Elevate your documentation with TermGIFforge, the API that transforms your terminal commands into polished, animated GIFs.

<p align="center">
  <img src="http://terminalgifapi.com/api/v1/gif?commands=[%20%22Set%20FontSize%2050%22,%20%22Set%20TypingSpeed%2075ms%22,%20%22Type%20\%22echo%20\%22%22,%20%22Set%20TypingSpeed%20500ms%22,%20%22Type%20\%22%27YEY\%22%22,%20%22Set%20TypingSpeed%2075ms%22,%20%22Type%20\%22!!!%27\%22%22,%20%22Sleep%20100ms%22,%20%22Enter%22,%20%22Sleep%202s%22]"/>
</p>

Use this API to present creative ways of your software commands.
<!-- Live in <link>http://terminalgifapi.com/api/v1/gif?commands=["Type \"echo 'The Magic Happens Here'\"","Enter","Sleep 2s"]</link>

<a
    href="http://terminalgifapi.com/api/v1/gif?commands=[%22Type%20%5C%22echo%20%27The%20Magic%20Happens%20Here%27%5C%22%22%2C%22Enter%22%2C%22Sleep%202s%22]">
    http://terminalgifapi.com/api/v1/gif?commands=["Type \"echo 'The Magic Happens Here'\"","Enter","Sleep 2s"]
</a> -->

## Instructions

TermGIFforge uses the [VHS](https://github.com/charmbracelet/vhs) to generate the GIFs, [here you'll find the commands references](https://github.com/charmbracelet/vhs?tab=readme-ov-file#vhs-command-reference) to generate your own GIFs.

## How to run

### Locally w/ docker 🐳

<p align="center">
  <img src="http://terminalgifapi.com/api/v1/gif?commands=[%20%22Type%20\%22echo%20%27How%20to%20run%20locally%20w/%20Docker%27\%22\n%22,%22Sleep%20400ms%22,%22Enter%22,%22Sleep%20200ms%22,%20%22Type%20\%22echo%20%27make%20build-image%20to%20build%20image%27\%22\n%22,%22Sleep%20400ms%22,%22Enter%22,%22Sleep%20200ms%22,%20%22Type%20\%22echo%20%27make%20debug-container%20to%20start%20a%20terminal%20from%20inside%20the%20container%27\%22\n%22,%22Sleep%20400ms%22,%22Enter%22,%22Sleep%20200ms%22,%20%22Type%20\%22echo%20%27go%20run%20cmd/server/main.go%27\%22\n%22,%22Sleep%20400ms%22,%22Enter%22,%22Sleep%20200ms%22,%20%22Sleep%202s%22]"/>
</p>

<!--
http://terminalgifapi.com/api/v1/gif?commands=[
    "Type \"# build image run: make build-image\"",
    "Sleep 100ms",
    "Enter",
    "Type \"# debug container sharing volume run: make debug-container\"",
    "Sleep 100ms",
    "Enter",
    "Type \"# start server from inside container run: go run cmd/server/main.go\"",
    "Sleep 100ms",
    "Enter",
    "Sleep 2s"]
 -->

<!--
http://terminalgifapi.com/api/v1/gif?commands=[
"Set TypingSpeed 75ms",
"Set FontSize 22",
"Set Width 1300",
"Set Height 650",
"Type \"neofetch\"",
"Sleep 500ms",
"Enter",
"Sleep 2s",
"Type \"Welcome to VHS!\"",
"Sleep 1",
"Space",
"Type \"A tool for generating terminal GIFs from code.\"",
"Sleep 5s"
]

 -->
From browser or [Bruno](./zarf/bruno/), open:

<!-- - http://terminalgifapi.com/api/v1/gif?commands=["Type \"echo 'The Magic Happens Here'\"","Enter","Sleep 2s"]
- http://terminalgifapi.com/api/v1/gif?commands=[
    "Set FontSize 50",
    "Set TypingSpeed 75ms",
    "Type \"echo \"",
    "Set TypingSpeed 500ms",
    "Type \"'YEY\"",
    "Set TypingSpeed 75ms",
    "Type \"!!!'\"",
    "Sleep 100ms",
    "Enter",
    "Sleep 2s"]
- http://terminalgifapi.com/api/v1/gif?commands=["Type \"echo 'Welcome to VHS!'\"","Sleep 100ms","Enter","Sleep 2s"]
- http://terminalgifapi.com/api/v1/gif?commands=["Type \"echo 'Welcome to VHS!'\"","Enter","Type \"ls\"","Sleep 100ms","Enter","Sleep 2s"]
- http://localhost:9001/api/v1/gif?commands=["Type \"echo 'Welcome to VHS!'\"","Sleep 100ms","Enter","Sleep 2s"]
- http://localhost:9001/api/v1/gif?commands=[
    "Type \"echo 'How to run locally w/ Docker'\"\n","Sleep 400ms","Enter","Sleep 200ms",
    "Type \"echo 'make build-image to build image'\"\n","Sleep 400ms","Enter","Sleep 200ms",
    "Type \"echo 'make debug-container to start a terminal from inside the container'\"\n","Sleep 400ms","Enter","Sleep 200ms",
    "Type \"echo 'go run cmd/server/main.go'\"\n","Sleep 400ms","Enter","Sleep 200ms",
    "Sleep 2s"]
- http://localhost:9001/api/v1/gif?commands=["Type \"cat README.md\"","Enter","Sleep 2s"]
- http://localhost:9001/api/v1/gif?commands=["Type \"less README.md\"","Enter","Sleep 2s"] -->
- http://localhost:9001/api/v1/gif?commands=["Type \"echo Hi\"","Enter","Sleep 2s"]

## References

<!-- ![Anurag's GitHub stats](http://terminalgifapi.com/api/v1/gif?commands=["Type \"echo 'Welcome to VHS!'\"","Sleep 100ms","Enter","Sleep 100ms","Type \"ls -a\"","Sleep 100ms","Enter","Sleep 1s"])

![Anurag's GitHub stats](http://terminalgifapi.com/api/v1/mock) -->

<!-- http://terminalgifapi.com/api/v1/gif?commands=["Type \"echo 'Welcome to VHS!'\"","Sleep 100ms","Enter","Sleep 100ms","Type \"ls -a\"","Sleep 100ms","Enter","Sleep 1s"] -->
<!-- ![My GitHub Streak](http://github-readme-streak-stats.herokuapp.com?user=victorabarros) -->

- https://github.com/charmbracelet/vhs
- https://github.com/anuraghazra/github-readme-stats
- https://github.com/DenverCoder1/github-readme-streak-stats
- https://github.com/rahuldkjain/github-profile-readme-generator
- https://github.com/ryo-ma/github-profile-trophy

## [Licence](./LICENSE)

<!--
TODO

- list of requirements to publish idea
  - create worker to exclude oldest GIFs
  - Improve README
    - use fancier namming prompt to improve name https://github.com/f/awesome-chatgpt-prompts?tab=readme-ov-file#act-as-a-fancy-title-generator
    - add especial thanks to Charmbracelets VHS

- cmds to build image using mac or linux
- create homepage to introduce project
- improve dockerfile
  - create stage to copy code and build project
  - copy build to release fase
  - entrypoint to build
- backup https://github.com/charmbracelet/vhs/releases/download/v0.9.0/vhs_0.9.0_arm64.deb

- use prompt to help w/ doc and post
- write article/post
- post on https://x.com/i/communities/1685641800449462272 and of show HN

-->
