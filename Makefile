DC = docker-compose

all: start

init: # init
	@echo "Initalize enviornment"
	cp -f ${PWD}/.env.dist ${PWD}/.env
	@echo "Remember changes .env content to your settings"

start: # deploy project
	@echo "Starting project..."
	${DC} up -d

develop: # deploy project with program log
	@echo "Starting project with program logging..."
	${DC} up --build

stop: # stop project
	@echo "Stopping project..."
	${DC} down

remove: # remove project
	@echo "Removing project..."
	${DC} down --remove-orphans
	
redeploy: # redeploy project
	@echo "Redeploying project..."
	${DC} down --remove-orphans --volumes
	${DC} up --build -d