package swapistructs

//People swapi struct
type People struct {
	Name      string       `json:"name,omitempty"`
	Height    string       `json:"height,omitempty"`
	Mass      string       `json:"mass,omitempty"`
	HairColor string       `json:"hair_color,omitempty"`
	SkinColor string       `json:"skin_color,omitempty"`
	EyeColor  string       `json:"eye_color,omitempty"`
	BirthYear string       `json:"birth_year,omitempty"`
	Gender    string       `json:"gender,omitempty"`
	Homeworld *Planets     `json:"homeworld,omitempty"`
	Films     *[]Films     `json:"films,omitempty"`
	Species   *[]Species   `json:"species,omitempty"`
	Vehicles  *[]Vehicles  `json:"vehicles,omitempty"`
	Starships *[]Starships `json:"starships,omitempty"`
	Created   string       `json:"created,omitempty"`
	Edited    string       `json:"edited,omitempty"`
	URL       string       `json:"url,omitempty"`
}

//Planets swapi struct
type Planets struct {
	Name           string    `json:"name,omitempty"`
	RotationPeriod int       `json:"rotation_period,omitempty"`
	OrbitalPeriod  int       `json:"orbital_Period,omitempty"`
	Diameter       int       `json:"diameter,omitempty"`
	Climate        string    `json:"climate,omitempty"`
	Gravity        string    `json:"gravity,omitempty"`
	Terrain        string    `json:"terrain,omitempty"`
	SurfaceWater   int       `json:"surface_water,omitempty"`
	Population     int       `json:"population,omitempty"`
	Residents      *[]People `json:"people,omitempty"`
	Films          *[]Films  `json:"films,omitempty"`
	Created        string    `json:"created,omitempty"`
	Edited         string    `json:"edited,omitempty"`
	URL            string    `json:"url,omitempty"`
}

//Films swapi struct
type Films struct {
	Title        string       `json:"title,omitempty"`
	EpisodeID    int          `json:"episode_id,omitempty"`
	OpeningCrawl string       `json:"opening_crawl,omitempty"`
	Director     string       `json:"director,omitempty"`
	Producer     string       `json:"producer,omitempty"`
	ReleaseDate  string       `json:"release_date,omitempty"`
	Characters   *[]People    `json:"characters,omitempty"`
	Planets      *[]Planets   `json:"planets,omitempty"`
	Starships    *[]Starships `json:"starships,omitempty"`
	Vehicles     *[]Vehicles  `json:"vehicles,omitempty"`
	Species      *[]Species   `json:"species,omitempty"`
	Created      string       `json:"created,omitempty"`
	Edited       string       `json:"edited,omitempty"`
	URL          string       `json:"url,omitempty"`
}

//Species swapi struct
type Species struct {
	Name            string    `json:"name,omitempty"`
	Classification  string    `json:"classification,omitempty"`
	Designation     string    `json:"designation,omitempty"`
	SkinColor       string    `json:"skin_color,omitempty"`
	HairColor       string    `json:"hair_color,omitempty"`
	EyeColor        string    `json:"eye_color,omitempty"`
	AverageLifespan int       `json:"average_lifespan,omitempty"`
	Homeworld       *Planets  `json:"homeworld,omitempty"`
	Language        string    `json:"language,omitempty"`
	People          *[]People `json:"people,omitempty"`
	Films           *[]Films  `json:"films,omitempty"`
	Created         string    `json:"created,omitempty"`
	Edited          string    `json:"edited,omitempty"`
	URL             string    `json:"url,omitempty"`
}

//Vehicles swapi struct
type Vehicles struct {
	Name                 string    `json:"name,omitempty"`
	Model                string    `json:"model,omitempty"`
	Manufacturer         string    `json:"manufacturer,omitempty"`
	CostInCredits        string    `json:"cost_in_credits,omitempty"`
	Length               float32   `json:"length,omitempty"`
	MaxAtmospheringSpeed int       `json:"max_atmosphering_speed,omitempty"`
	Crew                 int       `json:"crew,omitempty"`
	Passengers           int       `json:"passengers,omitempty"`
	CargoCapacity        int       `json:"cargo_capacity,omitempty"`
	Consumabes           string    `json:"consumables,omitempty"`
	VehicleClass         string    `json:"vehicle_class,omitempty"`
	Pilots               *[]People `json:"people,omitempty"`
	Films                *[]Films  `json:"films,omitempty"`
	Created              string    `json:"created,omitempty"`
	Edited               string    `json:"edited,omitempty"`
	URL                  string    `json:"url,omitempty"`
}

//Starships swapi struct
type Starships struct {
	Name                 string    `json:"name,omitempty"`
	Model                string    `json:"model,omitempty"`
	Manufacturer         string    `json:"manufacturer,omitempty"`
	CostInCredits        string    `json:"cost_in_credits,omitempty"`
	Length               float32   `json:"length,omitempty"`
	MaxAtmospheringSpeed int       `json:"max_atmosphering_speed,omitempty"`
	Crew                 int       `json:"crew,omitempty"`
	Passengers           int       `json:"passengers,omitempty"`
	CargoCapacity        int       `json:"cargo_capacity,omitempty"`
	Consumabes           string    `json:"consumables,omitempty"`
	HyperdriveRating     float32   `json:"hyperdrive_rating,omitempty"`
	Mglt                 int       `json:"MGLT,omitempty"`
	StarshipClass        string    `json:"starship_class,omitempty"`
	Pilots               *[]People `json:"people,omitempty"`
	Films                *[]Films  `json:"films,omitempty"`
	Created              string    `json:"created,omitempty"`
	Edited               string    `json:"edited,omitempty"`
	URL                  string    `json:"url,omitempty"`
}
