# TODO

## tech debt

## code improvement

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

## prod improvement

- store GIF assets on S3
- create DB to store the commands asked; this way can re-create in case of lost
- install `vim` to the container
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
