
This is at the artifact accompanying the PLDI'22 paper titled "A Study of Real-world Data Races in Golang."

The artifact includes the sample data race examples shown in Section 4 "Observations of Data Races in Go." 

The data race examples are associated with a Listing number used in the paper. A few additional data race examples are also included, which are mentioned but not shown as a listing in the paper.

You can exercise the artifact in one of the following ways:
A. Using a docker
B. Without using a docker.

# A. Using the docker
Follow these steps.
1. `git clone https://github.com/mkr-plse/gorace-examples.git`
2. `cd gorace-examples`
3. `docker build -t golang_datarace_artifact`
4. `docker run -it golang_datarace_artifact /bin/bash`
5. `bash run.sh` # this will run each data race example.
6. [Optional] You may investigate the *.log files present which show the found data race logs.
    Each log file contains the call stacks where the data race is found.
    You can open up the corresponding .go file to see the source code of the data race.
    You may investigate `run.bash` file to add more data race examples.
4. `exit` # exit from the docker.

# B. Without using the docker
Follow these steps.
1. Install Go `go1.15.5` matching you distribution from here: https://go.dev/dl/
2. `git clone https://github.com/mkr-plse/gorace-examples.git`
3. `cd gorace-examples`
4. `bash run.sh` # this will run each data race example.
5.  [Optional] You may investigate the *.log files present which show the found data race logs.
    Each log file contains the call stacks where the data race is found.
    You can open up the corresponding .go file to see the source code of the data race.
    You may investigate `run.bash` file to add more data race examples.

