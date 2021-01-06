package main

/*
//SearchPeople swapi struct
type SearchPeople struct {
	Count    int    `json:"count,omitempty"`
	Next     string `json:"next,omitempty"`
	Previous string `json:"previous,omitempty"`
	Results  []struct {
		Name      string   `json:"name,omitempty"`
		Height    string   `json:"height,omitempty"`
		Mass      string   `json:"mass,omitempty"`
		HairColor string   `json:"hair_color,omitempty"`
		SkinColor string   `json:"skin_color,omitempty"`
		EyeColor  string   `json:"eye_color,omitempty"`
		BirthYear string   `json:"birth_year,omitempty"`
		Gender    string   `json:"gender,omitempty"`
		Homeworld string   `json:"homeworld,omitempty"`
		Films     []string `json:"films,omitempty"`
		Species   []string `json:"species,omitempty"`
		Vehicles  []string `json:"vehicles,omitempty"`
		Starships []string `json:"starships,omitempty"`
		Created   string   `json:"created,omitempty"`
		Edited    string   `json:"edited,omitempty"`
		URL       string   `json:"url,omitempty"`
	} `json:"results,omitempty"`
} */

//People swapi struct
type People struct {
	Name      string   `json:"name,omitempty"`
	Height    string   `json:"height,omitempty"`
	Mass      string   `json:"mass,omitempty"`
	HairColor string   `json:"hair_color,omitempty"`
	SkinColor string   `json:"skin_color,omitempty"`
	EyeColor  string   `json:"eye_color,omitempty"`
	BirthYear string   `json:"birth_year,omitempty"`
	Gender    string   `json:"gender,omitempty"`
	Homeworld string   `json:"homeworld,omitempty"`
	Films     []string `json:"films,omitempty"`
	Species   []string `json:"species,omitempty"`
	Vehicles  []string `json:"vehicles,omitempty"`
	Starships []string `json:"starships,omitempty"`
	Created   string   `json:"created,omitempty"`
	Edited    string   `json:"edited,omitempty"`
	URL       string   `json:"url,omitempty"`
}

//Planets swapi struct
type Planets struct {
	Name           string   `json:"name,omitempty"`
	RotationPeriod string   `json:"rotation_period,omitempty"`
	OrbitalPeriod  string   `json:"orbital_Period,omitempty"`
	Diameter       string   `json:"diameter,omitempty"`
	Climate        string   `json:"climate,omitempty"`
	Gravity        string   `json:"gravity,omitempty"`
	Terrain        string   `json:"terrain,omitempty"`
	SurfaceWater   string   `json:"surface_water,omitempty"`
	Population     string   `json:"population,omitempty"`
	Residents      []string `json:"residents,omitempty"`
	Films          []string `json:"films,omitempty"`
	Created        string   `json:"created,omitempty"`
	Edited         string   `json:"edited,omitempty"`
	URL            string   `json:"url,omitempty"`
}

//Films swapi struct
type Films struct {
	Title        string   `json:"title,omitempty"`
	EpisodeID    string   `json:"episode_id,omitempty"`
	OpeningCrawl string   `json:"opening_crawl,omitempty"`
	Director     string   `json:"director,omitempty"`
	Producer     string   `json:"producer,omitempty"`
	ReleaseDate  string   `json:"release_date,omitempty"`
	Characters   []string `json:"characters,omitempty"`
	Planets      []string `json:"planets,omitempty"`
	Starships    []string `json:"starships,omitempty"`
	Vehicles     []string `json:"vehicles,omitempty"`
	Species      []string `json:"species,omitempty"`
	Created      string   `json:"created,omitempty"`
	Edited       string   `json:"edited,omitempty"`
	URL          string   `json:"url,omitempty"`
}

//Species swapi struct
type Species struct {
	Name            string   `json:"name,omitempty"`
	Classification  string   `json:"classification,omitempty"`
	Designation     string   `json:"designation,omitempty"`
	SkinColor       string   `json:"skin_color,omitempty"`
	HairColor       string   `json:"hair_color,omitempty"`
	EyeColor        string   `json:"eye_color,omitempty"`
	AverageLifespan string   `json:"average_lifespan,omitempty"`
	Homeworld       string   `json:"homeworld,omitempty"`
	Language        string   `json:"language,omitempty"`
	People          []string `json:"people,omitempty"`
	Films           []string `json:"films,omitempty"`
	Created         string   `json:"created,omitempty"`
	Edited          string   `json:"edited,omitempty"`
	URL             string   `json:"url,omitempty"`
}

//Vehicles swapi struct
type Vehicles struct {
	Name                 string   `json:"name,omitempty"`
	Model                string   `json:"model,omitempty"`
	Manufacturer         string   `json:"manufacturer,omitempty"`
	CostInCredits        string   `json:"cost_in_credits,omitempty"`
	Length               string   `json:"length,omitempty"`
	MaxAtmospheringSpeed string   `json:"max_atmosphering_speed,omitempty"`
	Crew                 string   `json:"crew,omitempty"`
	Passengers           string   `json:"passengers,omitempty"`
	CargoCapacity        string   `json:"cargo_capacity,omitempty"`
	Consumabes           string   `json:"consumables,omitempty"`
	VehicleClass         string   `json:"vehicle_class,omitempty"`
	Pilots               []string `json:"people,omitempty"`
	Films                []string `json:"films,omitempty"`
	Created              string   `json:"created,omitempty"`
	Edited               string   `json:"edited,omitempty"`
	URL                  string   `json:"url,omitempty"`
}

//Starships swapi struct
type Starships struct {
	Name                 string   `json:"name,omitempty"`
	Model                string   `json:"model,omitempty"`
	Manufacturer         string   `json:"manufacturer,omitempty"`
	CostInCredits        string   `json:"cost_in_credits,omitempty"`
	Length               string   `json:"length,omitempty"`
	MaxAtmospheringSpeed string   `json:"max_atmosphering_speed,omitempty"`
	Crew                 string   `json:"crew,omitempty"`
	Passengers           string   `json:"passengers,omitempty"`
	CargoCapacity        string   `json:"cargo_capacity,omitempty"`
	Consumabes           string   `json:"consumables,omitempty"`
	HyperdriveRating     string   `json:"hyperdrive_rating,omitempty"`
	Mglt                 string   `json:"MGLT,omitempty"`
	StarshipClass        string   `json:"starship_class,omitempty"`
	Pilots               []string `json:"people,omitempty"`
	Films                []string `json:"films,omitempty"`
	Created              string   `json:"created,omitempty"`
	Edited               string   `json:"edited,omitempty"`
	URL                  string   `json:"url,omitempty"`
}
