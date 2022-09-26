# Everest Kiki

The purpose of the project is to address the exercise problem, as documented [here](https://github.com/sudevkk/everest_kiki/blob/main/docs/EverestEngineering_Coding_challenge__courier_service_.pdf)

## Contents
	

 - [Dev Build] (#development-build)
 - Documentation (#documentation)
 - [Project Structure](#project-structure)
	 - [cmd](#cmd)
	 - [internal](#internal)
		 - [package cargo](#package-cargo)
		 - [package offer](#package-offer)
		 - [package transport](#package-transport)
		 - [package trip](#package-trip)
 - [Other](#other)

## Development Build

TBA

## Documentation
		Can be autogeneratd with godoc
		Run,
		
> 		godoc -http=:6060

## Project Structure

List of packages/Modules.

### /cmd
Has two commands cost (/cmd/cost/cost.go) and time (/cmd/time/time.go). 
#### cost.go
dommand line tool to generate delivery cost for the given number of packages

usage:

> cost -no XX -basecost XX

**Parameters:-**

no : The number of packages to be considered
basecost: The basecost of delivery, same for all packages

The tool will next prompt and accepts further inputs as formatted below,

> PackageID Weight Distance Offercode

The above will need to be entered, repeatedly based on the original input

#### time.go
Command line tool to generate delivery cost and times for the given number of packages

usage:

> time -no XX -basecost XX

**Parameters:-**

no : The number of packages to be considered
basecost: The basecost of delivery, same for all packages

The tool will next prompt and accepts further inputs as formatted below,

> PackageID Weight Distance Offercode

The above will need to be entered, repeatedly based on the original input
Post that the inputs for the Delivery time calculation as below,

> NoOfVehicles MaxWeightPerVehicle MaxWeightPerVehicle

## /internal

Contains all the Project packages (Business)

#### package cargo
TBA
#### package offer 
TBA
#### package transport
TBA
#### package trip
TBA
