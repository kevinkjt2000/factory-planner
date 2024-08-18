# factory-planner

This project is still under heavy construction, so expect that most of this
documentation is subject to change.

Minecraft GregTech New Horizons (GTNH) is a complex modpack that has 284,724
recipes as of version 2.6.1. This project aims to make visualizing chains of
dependent recipes easier. Imagine searching for a recipe, clicking a button
next to desired results to add those to a visualization section. Once
visualized, those recipes can be moved around and connected with arrows as you
see fit by drag-n-drop. Now also imagine that this program is also capable of
calculating how many of each machine is required to maintain certain amounts of
output items at any point in the chain of recipes. Yes, this is a lofty project
idea, but already portions of this exist scattered around GitHub.

## Database

To set up the database, I used https://github.com/D-Cysteine/nesql-exporter
configured to output to postgres. There's a few tweaks I did to set config for
that mod before letting it dump to docker-compose managed postgres that I
started with:

```
docker compose up -d
```

## Usage

I do not expect people to get anything useful from this project during this
super early state, but here are some instructions. I selected Taskfile for
automation of common tasks. This is optional for you to install, but I
recommend it so that you can shorten commands and have automatic rebuilds.
These are the tasks I spawn during development:

```
task --watch --verbose run
task --watch --verbose test
task db
task db-log
```
