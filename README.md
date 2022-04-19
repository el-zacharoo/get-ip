# get-ip
## This app allows you to recieve the IPv4 Address

* Run `go get` to install depenedencies
`go run .` to start the app

* To make a build on `docker` run `docker build --tag ip .`

* To run on docker `docker run -d -p 8080:8080 ip`

* It should run on `localhost:8080`

# Setup Heroku 

## Install the Heroku CLI


* To start you need install Heroku CLI `brew tap heroku/brew && brew install heroku`

* Login into heroku from your terminal with `heroku login`

* Create a new application using either Heroku dashboard - from the 'New' tab OR
  in your terminal run command `heroku create -a <app-name>` . If you dont specify an app name, one will be assigned.
  This will return the heroku url address for your app and a git url.

* `heroku git:remote -a <app-name>` | This command adds a git remote to the app repo

* The default stack that Heroku assigns is 'heroku-20'. You can check this under the settings on the
  dashboard. We need to set the stack to 'container'. Run this command `heroku stack:set container`
  This will also automatically set the Framework type to Container as well.

* Setup a enviroment variable on Heroku dashboard. Settings -> Config Vars -> "PORT" : "8080" 
  This is for the first initialization.

* Then run `git push heroku main` to trigger deployment from your local machine to Heroku remote.
  
* Once this has finished building, you will find your URL for the container under the settings tab.

* heroku.yaml | Add this file to your code repo. This tells Heroku to build a Docker container according to the
  Docker file located in the root directory of the application.

* https://devcenter.heroku.com/articles/heroku-cli-commands | Heroku CLI commands