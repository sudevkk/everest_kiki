
# Everest Kiki

The purpose of the project is to address the exercise problem, as documented [here](https://github.com/sudevkk/everest_kiki/blob/main/docs/EverestEngineering_Coding_challenge__courier_service_.pdf)

## Contents
	

 - [Dev Build] (#dev-build)
 - [Project Structure](#project-structure)
 - [Other](#other)

## Dev Build

TBA

## Project Structure

List of packages/Modules.

### cmd (/cmd)
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


