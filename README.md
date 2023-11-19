# sonarqube-cli
Example of using the CLI to consume the [sonarcloud API](https://sonarcloud.io/web_api/)
``````
$ sonarqube -h
Usage:
  sonarqube [command]

Available Commands:
  completion      Generate the autocompletion script for the specified shell
  help            Help about any command
  project         Manager Sonar Project
  quality-gate    Associate a project to a quality gate.
  quality-profile Associate a project with a quality profile.
  user-token      Generate a user access token.

Flags:
  -x, --debug     Debugmode (default: false)
  -h, --help      help for sonarqube
  -v, --version   version for sonarqube

Use "sonarqube [command] --help" for more information about a command.

``````
I'm using [Cobra](https://github.com/spf13/cobra)



## To start using sonarqube-cli
```
git clone https://github.com/evertonvps/sonarqube-cli
cd sonarqube-cli
make deploy
```

## Roadmap 
- [x] Create a project.
* [x] Associate a project to a quality gate.
* [x] Associate a project with a quality profile.
* [x] Generate a user access token.
* [ ] Revoke a user access token.
* [ ] Set tags on a project..
* [ ] Generate badge for project's measure as an SVG.
* [ ] Generate badge for project's quality gate as an SVG.
* [ ] Delete a non-main branch of a project.

