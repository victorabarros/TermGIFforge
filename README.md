# TermGIFforge 👾

On-the-Fly Terminal GIF Effortlessly Generation
Elevate your documentation with **TermGIFforge**, the API that transforms your terminal commands into polished, animated GIFs.

<p align="center">
  <img src="http://terminalgifapi.com/api/v1/gif?commands=[%20%22Set%20FontSize%2050%22,%20%22Set%20TypingSpeed%2075ms%22,%20%22Type%20\%22echo%20\%22%22,%20%22Set%20TypingSpeed%20500ms%22,%20%22Type%20\%22%27YEY\%22%22,%20%22Set%20TypingSpeed%2075ms%22,%20%22Type%20\%22!!!!%27\%22%22,%20%22Sleep%20100ms%22,%20%22Enter%22,%20%22Sleep%202s%22]"/>
  <img src="http://terminalgifapi.com/api/v1/gif?commands=[%22Type%20\%22echo%20%27The%20Magic%20Happens%20Here%27\%22%22,%22Enter%22,%22Sleep%202s%22]"/>
</p>

Use this API to present creative ways of your software commands.

## Instructions

**TermGIFforge** uses the [VHS](https://github.com/charmbracelet/vhs) to generate the GIFs, [here you'll find the command references](https://github.com/charmbracelet/vhs?tab=readme-ov-file#vhs-command-reference) to generate your own GIFs.

## How to run

### Locally w/ docker 🐳

<p align="center">
  <img src="http://terminalgifapi.com/api/v1/gif?commands=[%20%22Type%20\%22echo%20%27How%20to%20run%20locally%20w/%20Docker%27\%22\n%22,%22Sleep%20400ms%22,%22Enter%22,%22Sleep%20200ms%22,%20%22Type%20\%22echo%20%27make%20build-image%20to%20build%20image%27\%22\n%22,%22Sleep%20400ms%22,%22Enter%22,%22Sleep%20200ms%22,%20%22Type%20\%22echo%20%27make%20debug-container%20to%20start%20a%20terminal%20from%20inside%20the%20container%27\%22\n%22,%22Sleep%20400ms%22,%22Enter%22,%22Sleep%20200ms%22,%20%22Type%20\%22echo%20%27go%20run%20cmd/server/main.go%27\%22\n%22,%22Sleep%20400ms%22,%22Enter%22,%22Sleep%20200ms%22,%20%22Sleep%202s%22]"/>
</p>

Than, from browser or [Bruno](./zarf/bruno/), open:

- http://localhost:9001/api/v1/gif?commands=["Type \"echo Hi\"","Enter","Sleep 2s"]

## Troubleshooting

- How to force image (cache) update in README.rst on GitHub
  - run `curl -X PURGE {url of cached badge image}` and refresh the page. [link](https://stackoverflow.com/questions/26898052/how-to-force-image-cache-update-in-readme-rst-on-github)

## References

- https://github.com/charmbracelet/vhs
<!-- - https://github.com/anuraghazra/github-readme-stats
- https://github.com/DenverCoder1/github-readme-streak-stats
- https://github.com/rahuldkjain/github-profile-readme-generator
- https://github.com/ryo-ma/github-profile-trophy -->

## Licence

[MIT](./LICENSE)

<!--
TODO

- list of requirements to publish idea
  - create worker to exclude oldest GIFs
  - Improve README

- cmds to build image using mac or linux
- create homepage to introduce project
- improve dockerfile
  - create stage to copy code and build project
  - copy build to release fase
  - entrypoint to build
- backup https://github.com/charmbracelet/vhs/releases/download/v0.9.0/vhs_0.9.0_arm64.deb

- write article/post
- post on https://x.com/i/communities/1685641800449462272 and of show HN

-->
