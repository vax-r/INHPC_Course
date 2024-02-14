# INHPC Courses
## Introduction
The system run as the courses managment system in INHPC.
## Setup
1. Download the source code with the following command.
```
$ git clone git@github.com:vax-r/INHPC_Course.git && cd INHPC_Course
```
2. Set up `.env` file and fill in the desired value.
```
$ cp .env.example .env
```
3. Build up the database and adminer
```
$ sudo make build
```
After this step, you should have 2 containers running in your system
4. Run the app to set up the backend server
```
$ sudo make dev-run
```
