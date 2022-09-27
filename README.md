# Everest Kiki

The purpose of the project is to address the exercise problem, as documented [here](https://github.com/sudevkk/everest_kiki/blob/main/docs/EverestEngineering_Coding_challenge__courier_service_.pdf)

## Contents
	

 - [Build & Tests](#build)
 - [Documentation](#documentation)
 - [Project Structure](#project-structure)
	 - [cmd (Commands) ](#cmd)
	 - [internal (Project Packages. Business)](#internal)
		 - [package cargo](#1.-package-cargo)
		 - [package offer](#2.-package-offer)
		 - [package transport](#3.-package-transport)
		 - [package trip](#4.-package-trip)
 - [Other](#other)

## Build and Tests

TBA

(Makefile is added but incomplete. do _go run_/_build_ directly)

Tests: Example tests at _/cmd/cost_ and _/cms/time_ 

## Documentation
Can be auto generatd with godoc
		
> 	godoc -http=:6060

## Project Structure

List of packages/Modules.

## /cmd
Has two commands cost (/cmd/cost/cost.go) and time (/cmd/time/time.go). 
#### 1. cost.go
dommand line tool to generate delivery cost for the given number of packages

***usage:***

> cost -no XX -basecost XX

***Parameters:-***

no : The number of packages to be considered
basecost: The basecost of delivery, same for all packages

The tool will next prompt and accepts further inputs as formatted below,

> PackageID Weight Distance Offercode

The above will need to be entered, repeatedly based on the original input

#### 2. time.go
Command line tool to generate delivery cost and times for the given number of packages

***usage:***

> time -no XX -basecost XX

***Parameters:-***

no : The number of packages to be considered
basecost: The basecost of delivery, same for all packages

The tool will next prompt and accepts further inputs as formatted below,

> PackageID Weight Distance Offercode

The above will need to be entered, repeatedly based on the original input
Post that the inputs for the Delivery time calculation as below,

> NoOfVehicles MaxWeightPerVehicle MaxWeightPerVehicle

## /internal

Contains all the Project packages (Business)

#### 1. package cargo
##### Overview 
The package deals with the Freight collection and the Freight content itself This is the Raw package before it converts to the Actual Consignment
#### Index
[type Box](#type-box)

[func NewBox(id string, weight float64) *Box](#function-newbox)

###### *Type Box*
A Physical Package that needs to be delivered. Package seems to be ideal name, not using to avoid confusion Can contain Package's Physical data and the data collected from the Sender (Addresses, Received Date)

    type Box struct {
	    ID     [string](http://localhost:6060/pkg/builtin/#string)  // Package ID
	    Weight [float64](http://localhost:6060/pkg/builtin/#float64) // Package Weight in KGs
		}

###### *Function newBox*

    func NewBox(id string, weight float64) *Box

Returns a New [cargo.Box]

#### 2. package offer 
##### Overview 
Package deals with the Offers/Discount Coupons etc.
A Global for holding the available list of Offercodes and details as a Key-Value pair is used. Assuming only initialized once during init and only reads post that, (Thus, though [offers] is not threadsafe, is not a problem here).
*TODO:*
*The actual implementation will involve persisting the Offer/Coupon data in DB etc. And also since this information will be something that will get fetched frequently (On each consignment addition), to reduce the actual DB reads this can be cached. Possible possible caching strategy: Use a Redis or some in-memory store to save the Coupon/Offer data. On application bootstrap the DB can be read and all the info can be Written to Redis (Offer code as Key). Will need a Mechanism/Service to check and invalidate codes from Redis as needed based on diff criteria (Code Expiry, Max allowed usage exhausted for a code etc)*

*For simplicity the persistence/caching is not done here, instead a simple in memory list of Maps holds the list of Codes that will be used directly.*

#### Index
[func  DiscountByCode(offerCode string, d float64, weight float64, amount *money.Money) (*money.Money, bool)]()

[func  DiscountPercByCode(offerCode string, d float64, weight float64) (discountPerc float64, isValidOffercode bool)]()

###### Function DiscountByCode

    func DiscountByCode(offerCode string, d float64, weight float64, amount *money.Money) (*money.Money, bool)

Returns a applicable Discount amount (Assuming only Dollar as currency for now), for the supplied offer code Returns 0, false if no qualifying offer code was found TODO: Make this currency flexible.

###### func DiscountPercByCode

       func DiscountPercByCode(offerCode string, d float64, weight float64) (discountPerc float64, isValidOffercode bool)

Returns a Discount Percentage, for the supplied offer code Returns 0, false if no qualifying offer code was found
#### 3. package transport

##### Overview 
Deals with [transport.Consignment] Delivery cost Calculations Could be part of Box package, separated to manage different Costing structures that might be possible over time eg. Festive Rate, Corporate/Partner plans etc. better to Isolate Box from the actual Costing scheme Uses Strategy pattern to accommodate different costing strategies

The Package deals with the Action of actual Transportation (Costing, Delivery Time Estimation, Consignments generation and Planning etc)

#### Index

 - [type  Consignment]()
	 - [func  NewConsignment(b *cargo.Box) *Consignment]() 	
	 - [func (c *Consignment) CalcCost()]() 	
	 - [func (c *Consignment) SetCostingStratergy(coster CostingStrategy)]()
	 - [func (c  *Consignment) SetDistance(distance float64)]() 	    
	 - [func (c *Consignment) SetOffercode(offercode string)]()
 - [type  CostingStrategy]() 
 - [type  RoadTransport]() [type StandardCosting]() 
 - [type  TransportMode]() 
 - [func (s StandardCosting) Calc(weight float64, distance float64) (cost *money.Money, discount
   *money.Money)]()

###### type  Consignment
A Consignment that gets Transported One to One Mapping with a [cargo.Box] [cargo.Box] gets converted to a Consignment with Distance, Coster, Offer etc applied on it. ([cargo.Box] could have raw info like From and To addresses etc, which on converting to a [transport.Consignement]. will be mapped to the Actual Distance, or the Weight gets converted to the Volumetric Weight etc)

    type Consignment struct {
        Box       *[cargo](http://localhost:6060/pkg/github.com/sudevkk/everest_kiki/internal/cargo/).[Box](http://localhost:6060/pkg/github.com/sudevkk/everest_kiki/internal/cargo/#Box)   // The [carg.Box] that is part of this consignment
        Distance  [float64](http://localhost:6060/pkg/builtin/#float64)      // The actual distance for the shipment
        Offercode [string](http://localhost:6060/pkg/builtin/#string)       // Offercode, if any
        Cost      *money.[Money](http://localhost:6060/pkg/github.com/sudevkk/everest_kiki/internal/transport/#Money) // The calculated cost for the shipment based on the selected [transport.Coster]
        Discount  *money.[Money](http://localhost:6060/pkg/github.com/sudevkk/everest_kiki/internal/transport/#Money) // The calculated discount based on the supplied Offercode, if any. 0 by default
        Coster    [CostingStrategy](http://localhost:6060/pkg/github.com/sudevkk/everest_kiki/internal/transport/#CostingStrategy)
        Mode      [TransportMode](http://localhost:6060/pkg/github.com/sudevkk/everest_kiki/internal/transport/#TransportMode)
    }

###### type costingStratergy 

TBA


#### 4. package trip
##### Overview 
Package deals with the Trip (Delivery of selected consignments) Manages the Scheduling etc

##### Index

 - [type  Vehicle]()
 - [type  FleetVehicle]()
 - [type  Fleet]()
	 - [func  NewFleet(vehicles []Vehicle) Fleet]()    
	 - [func (f Fleet) 	   AddVehicle(v Vehicle)]()
	 - [func (f Fleet) GetAvailableVehicleNode(weight float64) (*rbt.Node, bool)]()
 - [type  Trip]()
	 - [func  New(v []Vehicle) Trip]() 	
	 - [func (t Trip) RunSchedule(consignments []transport.Consignment)]()

###### type  Vehicle

    // The Vehicle entity
    
    type  Vehicle  struct {
    
    ID int
    
    MaxLoadWeight float64  // The maximum weight in KGs the vehicle can carry
    
    Speed float64  // Speed KM/HR (Assuming constant speed)
    
    }
###### type  FleetVehicle
A [trip.Vehicle] that is added to a [trip.Trip]

    type  FleetVehicle  struct {
    
    Vehicle
    
    CurrentLoadWeight float64
    
    NextAvailableFrom time.Time
    
    consignments []transport.Consignment
    
    }
###### type  Fleet
The collection of assets used for a [trip.Trip]. 

Vehicles are saved as an [RB tree](https://en.wikipedia.org/wiki/Red%E2%80%93black_tree) for simplicity

    type  Fleet  struct {
    
    Vehicles *rbt.Tree // List of available Vehicles in the Fleet (Stored as a RB Tree)
    
    }
###### type  Trip
Represents a Trip / delivery route. The scheduling and other actions happens on a [trip.Trip]

    type  Trip  struct {
    
    fleet Fleet
    
    startTime time.Time // Only Used for the simulation time for schedule
    
    }

###### func RunSchedule
Runs a scheduling (simulated) for a given list of consignements on the current [trip.Trip]
The consignments are sorted on Weight and Distance.
The scheduler utilises the RB tree of vehicle nodes to do the assignment in priority
The complexity (n logn * n). This can be likely optimised further and not attempted.

