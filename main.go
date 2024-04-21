package main

import (
	"errors"
	"fmt"
	"time"
)

type TravelHandler struct {
	CityService *CityService
	RoadService *RoadService
}

func NewTravelHandler(cityService *CityService, roadService *RoadService) *TravelHandler {
	return &TravelHandler{CityService: cityService, RoadService: roadService}
}

type Menu int
type Model int

const (
	MainMenu  Menu  = 0
	Help      Menu  = 1
	Add       Menu  = 2
	Delete    Menu  = 3
	Path      Menu  = 4
	Exit      Menu  = 5
	cityModel Model = 1
	roadModel Model = 2
)

type City struct {
	Id   uint
	Name string
}
type Road struct {
	Id            uint
	Name          string
	From          uint
	To            uint
	Through       []uint
	SpeedLimit    uint
	Length        uint
	BiDirectional bool
	City          *City
}

func main() {
	inputHandler()
}

func inputHandler() {

	handler := InitHandler()

	var req Menu
	for {
		switch req {
		case MainMenu:
			handler.MainMenuHandler()
		case Help:
			handler.HelpHandler()
		case Add:
			handler.AddModelHandler()
		case Delete:
			handler.DeleteModelHandler()
			req = MainMenu
		case Path:
			handler.PathHandler()
		case Exit:
			return
		default:
			handler.InvalidInputHandler()
		}
		GetMenu(&req)
	}
}

func InitHandler() *TravelHandler {
	//init repositories
	cityRepo := NewMapCityRepo()
	roadRepo := NewMapRoadRepo()

	// init services
	cityService := NewCityService(cityRepo)
	roadService := NewRoadService(roadRepo)

	// init handler
	handler := NewTravelHandler(cityService, roadService)
	return handler
}

func (h *TravelHandler) MainMenuHandler() {
	fmt.Println("Main Menu - Select an action:")
	fmt.Println("1. Help")
	fmt.Println("2. Add")
	fmt.Println("3. Delete")
	fmt.Println("4. Path")
	fmt.Println("5. Exit")
}

func GetMenu(input *Menu) {
	_, err := fmt.Scanln(input)
	if err != nil {
		return
	}
}

func (h *TravelHandler) AddModelHandler() {
	model := selectModel()
	for {
		switch Model(model) {
		case cityModel:

			h.CityService.CreateCity()
		case roadModel:
			h.RoadService.CreateRoad()
		}

		var input Menu
		GetMenu(&input)
		if input == 2 {
			h.MainMenuHandler()
			return
		}
	}
}
func printSuccess(item string, id uint) {
	fmt.Printf(`
%s with id=%v added!
Select your next action:
1. Add another %s
2. Main Menu
`, item, id, item)
}

func selectModel() int {
	var m int
	fmt.Println(`Select model:
1. City
2. Road`)
	fmt.Scanln(&m)
	return m
}

func (h *TravelHandler) HelpHandler() {
	fmt.Println("Select a number from shown menu and enter. For example 1 is for help")
	h.MainMenuHandler()
}

func (h *TravelHandler) DeleteModelHandler() {
	model := selectModel()

	switch Model(model) {
	case cityModel:
		h.CityService.DeleteCity()
	case roadModel:
		h.RoadService.DeleteRoad()
	}
}

func (h *TravelHandler) PathHandler() *Menu {
	var (
		SourceCityId, DestinationCityId uint
	)
	fmt.Scanf("%v:%v", &SourceCityId, &DestinationCityId)

	SourceCityName, _ := h.CityService.GetCity(SourceCityId)
	DestinationCityName, _ := h.CityService.GetCity(DestinationCityId)

	AllRoads := h.RoadService.GetAllRoads()

	roadName, takeTime := findThePath(SourceCityId, DestinationCityId, AllRoads)

	fmt.Printf("%s:%s via Road %s: Takes %v", SourceCityName, DestinationCityName, roadName, takeTime)

	return nil
}

func findThePath(SourceCityId, DestinationCityId uint, allRoads []*Road) (string, time.Duration) {
	var (
		RoadName string
		TakeTime time.Duration
	)

	for _, road := range allRoads {
		var cities []uint
		cities = append(cities, road.From)
		cities = append(cities, road.Through...)
		cities = append(cities, road.To)

		found1 := false
		found2 := false

		for i, city := range cities {
			if !road.BiDirectional {
				if city == SourceCityId {
					if cities[i+1] == DestinationCityId {
						RoadName = road.Name
						TakeTime = time.Duration(road.Length / road.SpeedLimit)
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
					RoadName = road.Name
					TakeTime = time.Duration(road.Length / road.SpeedLimit)
				}
			}
		}
	}
	return RoadName, TakeTime

}

func (h *TravelHandler) InvalidInputHandler() {
	fmt.Println("Invalid input. Please enter 1 for more info.")
}

type IRoadRepository interface {
	getAll() []*Road
	create(road Road) error
	delete(id uint) error
}

type ICityRepository interface {
	get(id uint) (*City, error)
	create(city City) error
	delete(id uint) error
}

type MapRoadRepo struct {
	repo map[uint]Road
}

func NewMapRoadRepo() *MapRoadRepo {
	return &MapRoadRepo{repo: make(map[uint]Road)}
}

func (m *MapRoadRepo) getAll() []*Road {
	var roadList []*Road
	for _, road := range m.repo {
		roadList = append(roadList, &road)
	}
	return roadList
}

func (m MapRoadRepo) create(road Road) error {
	m.repo[road.Id] = road
	return nil
}

func (m MapRoadRepo) delete(id uint) error {
	if _, ok := m.repo[id]; ok {
		delete(m.repo, id)
		return nil
	}
	return errors.New("")
}

type RoadService struct {
	repo IRoadRepository
}

func NewRoadService(repo IRoadRepository) *RoadService {
	return &RoadService{repo: repo}
}

func (s *RoadService) CreateRoad() {
	var road Road
	fmt.Println("id=?")
	fmt.Scanln(&road.Id)
	fmt.Scanln(&road.Name)
	fmt.Println("from?")
	fmt.Scanln(&road.From)
	fmt.Println("to?")
	fmt.Scanln(&road.To)
	fmt.Println("through?")
	fmt.Scanln(&road.Through)
	fmt.Println("speed_limit?")
	fmt.Scanln(&road.SpeedLimit)
	fmt.Println("length?")
	fmt.Scanln(&road.Length)
	fmt.Println("bi_directional?")
	var biDirectional int
	fmt.Scanln(&biDirectional)
	if biDirectional == 1 {
		road.BiDirectional = true
	} else {
		road.BiDirectional = false
	}
	s.repo.create(road)
	printSuccess("Road", road.Id)

}

func (s *RoadService) DeleteRoad() {
	var id uint
	fmt.Scanln(&id)
	if err := s.DeleteRoad; err != nil {
		fmt.Printf("Road with id %v not found!", id)
	} else {
		fmt.Printf("Road :%v deleted!", id)
	}
}

func (s *RoadService) GetAllRoads() []*Road {
	return s.repo.getAll()
}

type MapCityRepo struct {
	repo map[uint]City
}

func NewMapCityRepo() *MapCityRepo {
	cityMap := make(map[uint]City)
	return &MapCityRepo{repo: cityMap}
}

func (m MapCityRepo) get(id uint) (*City, error) {
	city, ok := m.repo[id]
	if !ok {
		return nil, errors.New("")
	}
	return &city, nil
}

func (m MapCityRepo) create(city City) error {
	m.repo[city.Id] = city
	return nil
}

func (m MapCityRepo) delete(id uint) error {
	if _, ok := m.repo[id]; ok {
		delete(m.repo, id)
		return nil
	}
	return errors.New("")
}

type CityService struct {
	repo ICityRepository
}

func NewCityService(repo ICityRepository) *CityService {
	return &CityService{repo: repo}
}

func (s *CityService) CreateCity() {
	var city City
	fmt.Println("id=?")
	fmt.Scanln(&city.Id)
	fmt.Println("name=?")
	fmt.Scanln(&city.Name)
	s.repo.create(city)

	printSuccess("City", city.Id)
}
func (s *CityService) DeleteCity() {
	var id uint
	fmt.Println("id?")
	fmt.Scanln(&id)
	err := s.repo.delete(id)
	if err != nil {
		fmt.Printf("City with id %v not found!", id)
	} else {
		fmt.Printf("City : %v deleted!", id)
	}
}

func (s *CityService) GetCity(id uint) (*City, error) {
	city, err := s.repo.get(id)
	if err != nil {
		return nil, err
	}
	return city, nil
}
