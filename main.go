package main

import (
	"fmt"
	"time"
)

type InputReq struct {
	Input Action
	Model Model
	City  City
	Road  Road
}

type Action int
type Model int

const (
	MainMenu  Action = 0
	Help      Action = 1
	Add       Action = 2
	Delete    Action = 3
	Path      Action = 4
	Exit      Action = 5
	cityModel Model  = 1
	roadModel Model  = 2

	initiateMenu = `Main Menu - Select an action: 
1. Help
2. Add
3. Delete
4. Path
5. Exit`
	invalidInputMsg = "Invalid input. Please enter 1 for more info."
	helpMsg         = "Select a number from shown menu and enter. For example 1 is for help"
)

type City struct {
	Name string
}
type Road struct {
	Name          string
	From          int
	To            int
	Through       []int
	SpeedLimit    int
	Length        int
	BiDirectional bool
}

var CityData = make(map[int]City)

var RoadData = make(map[int]Road)

func main() {
	CityData[21] = City{Name: "Tehran"}
	CityData[251] = City{Name: "Qom"}
	CityData[361] = City{Name: "Kashan"}

	RoadData[1] = Road{
		Name:          "T-K",
		From:          21,
		To:            361,
		Through:       []int{21, 251},
		SpeedLimit:    80,
		Length:        600,
		BiDirectional: false,
	}
	var req InputReq
	req.Input = 0
	for {
		switch req.Input {
		case MainMenu:
			fmt.Println(initiateMenu)
			fmt.Scanln(&req.Input)
		case Help:
			req.HelpFunc()
		case Add:
			req.AddFunc()
		case Delete:
			req.DeleteFunc()
			req.Input = 0
		case Path:
			req.PathFunc()
		case Exit:
			return
		default:
			fmt.Println(invalidInputMsg)
			fmt.Scanln(&req.Input)
		}
	}
}

func (input *InputReq) AddFunc() *Action {
	var (
		id   int
		item string
	)
	model := input.selectModel()

	for {
		fmt.Println("id=?")
		fmt.Scanln(&id)
		fmt.Println("name=?")
		switch model {
		case cityModel:
			item = "City"
			fmt.Scanln(&input.City.Name)
			CityData[id] = City{
				Name: input.City.Name,
			}
			fmt.Println("city is :", CityData[id])
		case roadModel:
			item = "Road"
			fmt.Scanln(&input.Road.Name)
			fmt.Println("from?")
			fmt.Scanln(&input.Road.From)
			fmt.Println("to?")
			fmt.Scanln(&input.Road.To)
			fmt.Println("through?")
			fmt.Scanln(&input.Road.Through)
			fmt.Println("speed_limit?")
			fmt.Scanln(&input.Road.SpeedLimit)
			fmt.Println("length?")
			fmt.Scanln(&input.Road.Length)
			fmt.Println("bi_directional?")
			var biDirectional int
			fmt.Scanln(&biDirectional)
			if biDirectional == 1 {
				input.Road.BiDirectional = true
			} else {
				input.Road.BiDirectional = false
			}
			RoadData[id] = Road{
				Name:          input.Road.Name,
				From:          input.Road.From,
				To:            input.Road.To,
				Through:       input.Road.Through,
				SpeedLimit:    input.Road.SpeedLimit,
				Length:        input.Road.Length,
				BiDirectional: input.Road.BiDirectional,
			}
		}
		fmt.Printf(`
%s with id=%v added!
Select your next action:
1. Add another %s
2. Main Menu
`, item, id, item)
		fmt.Scanln(&input.Input)
		if input.Input == 2 {
			input.Input = 0
			return &input.Input
		}
	}
}

func (input *InputReq) selectModel() Model {
	fmt.Println(`Select model:
1. City
2. Road`)
	fmt.Scanln(&input.Model)
	return input.Model
}

func (input *InputReq) HelpFunc() *Action {
	fmt.Println(helpMsg)
	fmt.Println(initiateMenu)
	fmt.Scanln(&input.Input)
	return &input.Input
}

func (input *InputReq) DeleteFunc() {
	model := input.selectModel()
	var (
		id   int
		item string
	)
	fmt.Println("id=?")
	switch model {
	case cityModel:
		item = "City"
		fmt.Scanln(&id)
		if _, ok := CityData[id]; ok {
			delete(CityData, id)
			fmt.Println("%s :%v deleted!", item, id)
		} else {
			fmt.Printf("%s with id %v not found!", item, id)
		}
	case roadModel:
		item = "Road"
		fmt.Scanln(&id)
		if _, ok := RoadData[id]; ok {
			delete(RoadData, id)
			fmt.Printf("%s :%v deleted!", item, id)
		} else {
			fmt.Printf("%s with id %v not found!", item, id)
		}
	}
}

func (input *InputReq) PathFunc() *Action {
	var (
		SourceCityId, DestinationCityId int
		RoadName                        string
		TakeTime                        time.Duration
	)
	fmt.Scanf("%v:%v", &SourceCityId, &DestinationCityId)
	SourceCityName, _ := CityData[SourceCityId]
	DestinationCityName, _ := CityData[DestinationCityId]
	for _, value := range RoadData {
		var linkedList []int
		linkedList = append(linkedList, value.From)
		linkedList = append(linkedList, value.Through...)
		linkedList = append(linkedList, value.To)
		found1 := false
		found2 := false
		for i, city := range linkedList {
			if !value.BiDirectional {
				if city == SourceCityId {
					if linkedList[i+1] == DestinationCityId {
						RoadName = value.Name
						TakeTime = time.Duration(value.Length / value.SpeedLimit)
						break
					}
				} else {
					i++
				}
			} else {
				if city == SourceCityId {
					found1 = true
				} else if city == DestinationCityId {
					found2 = true
				}
				if found1 && found2 {
					RoadName = value.Name
					TakeTime = time.Duration(value.Length / value.SpeedLimit)
				}
			}
		}
	}

	fmt.Printf("%s:%s via Road %s: Takes %v", SourceCityName, DestinationCityName, RoadName, TakeTime)

	return nil
}
