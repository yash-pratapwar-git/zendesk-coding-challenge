
# ZENDESK-CODING CHALLENGE INTERN 2022

This is a solution for the coding challenge provided by ZenDesk for the summer 2022 internship.

## Requirements
1. Go (golang) - This can be downloaded from ```https://go.dev/doc/install```
2. Please follow the steps for your OS mentioned on ```https://go.dev/doc/install``` 

## Setting up GO environment in your OS
### (Preferred OS is windows - because of the simplicity of installation)
1. WINDOWS - Just install Go from the downloaded .msi file (from the above step)
2. LINUX, MAC - Need to set environment paths mentioned here ```https://go.dev/doc/install```

## Clone this repository
1. Open your home directory and create new folder with the name ```go```
2. Inside ```go``` folder, create new folder ```github.com```
3. Open ```github.com``` folder and clone this repository there.

## Usage
1. Open the cloned repository and go to ```/constants``` directory.
2. In ```constants.go``` file, update the values of UserName and Password fields.
3. Open the terminal at ```/go/github.com/zendesk-coding-challenge``` level
4. Enter the following command to install dependencies: ```go get```
5. Enter the following command to execute the program: ```go run main.go```

## Testing
1. Open the terminal at ```/go/github.com/zendesk-coding-challenge/services``` level.
2. Enter the following command to run the unit tests: ```go test```
3. The last line of the terminal output should start with an "ok" message if all the test cases are passed.   
4. ```services_test.go``` file contains all the unit tests.

## Outputs
![Main Menu](https://github.com/yash-pratapwar-git/zendesk-coding-challenge/blob/main/docs/result_1.png)
![List All Tickets](https://github.com/yash-pratapwar-git/zendesk-coding-challenge/blob/main/docs/result_2.png)
![View Single Ticket](https://github.com/yash-pratapwar-git/zendesk-coding-challenge/blob/main/docs/result_3.png)