package swapi

// swapi People struct
type people struct {
	Name      string       `json:"name,omitempty"`
	Height    int          `json:"height,omitempty"`
	Mass      int          `json:"mass,omitempty"`
	HairColor string       `json:"hair_color,omitempty"`
	SkinColor string       `json:"skin_color,omitempty"`
	EyeColor  string       `json:"eye_color,omitempty"`
	BirthYear string       `json:"birth_year,omitempty"`
	Gender    string       `json:"gender,omitempty"`
	Homeworld *planets     `json:"homeworld,omitempty"`
	Films     *[]films     `json:"films,omitempty"`
	Species   *[]species   `json:"species,omitempty"`
	Vehicles  *[]vehicles  `json:"vehicles,omitempty"`
	Starships *[]starships `json:"starships,omitempty"`
	Created   string       `json:"created,omitempty"`
	Edited    string       `json:"edited,omitempty"`
	URL       string       `json:"url,omitempty"`
}

//Planets swapi struct
type planets struct {
	Name           string    `json:"name,omitempty"`
	RotationPeriod int       `json:"rotation_period,omitempty"`
	OrbitalPeriod  int       `json:"orbital_Period,omitempty"`
	Diameter       int       `json:"diameter,omitempty"`
	Climate        string    `json:"climate,omitempty"`
	Gravity        string    `json:"gravity,omitempty"`
	Terrain        string    `json:"terrain,omitempty"`
	SurfaceWater   int       `json:"surface_water,omitempty"`
	Population     int       `json:"population,omitempty"`
	Residents      *[]people `json:"people,omitempty"`
	Films          *[]films  `json:"films,omitempty"`
	Created        string    `json:"created,omitempty"`
	Edited         string    `json:"edited,omitempty"`
	URL            string    `json:"url,omitempty"`
}

//films swapi struct
type films struct {
	Title        string       `json:"title,omitempty"`
	EpisodeID    int          `json:"episode_id,omitempty"`
	OpeningCrawl string       `json:"opening_crawl,omitempty"`
	Director     string       `json:"director,omitempty"`
	Producer     string       `json:"producer,omitempty"`
	ReleaseDate  string       `json:"release_date,omitempty"`
	Characters   *[]people    `json:"characters,omitempty"`
	Planets      *[]planets   `json:"planets,omitempty"`
	Starships    *[]starships `json:"starships,omitempty"`
	Vehicles     *[]vehicles  `json:"vehicles,omitempty"`
	Species      *[]species   `json:"species,omitempty"`
	Created      string       `json:"created,omitempty"`
	Edited       string       `json:"edited,omitempty"`
	URL          string       `json:"url,omitempty"`
}

//species swapi struct
type species struct {
	Name            string    `json:"name,omitempty"`
	Classification  string    `json:"classification,omitempty"`
	Designation     string    `json:"designation,omitempty"`
	SkinColor       string    `json:"skin_color,omitempty"`
	HairColor       string    `json:"hair_color,omitempty"`
	EyeColor        string    `json:"eye_color,omitempty"`
	AverageLifespan int       `json:"average_lifespan,omitempty"`
	Homeworld       *planets  `json:"homeworld,omitempty"`
	Language        string    `json:"language,omitempty"`
	People          *[]people `json:"people,omitempty"`
	Films           *[]films  `json:"films,omitempty"`
	Created         string    `json:"created,omitempty"`
	Edited          string    `json:"edited,omitempty"`
	URL             string    `json:"url,omitempty"`
}

//vehicles swapi struct
type vehicles struct {
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
	Pilots               *[]people `json:"people,omitempty"`
	Films                *[]films  `json:"films,omitempty"`
	Created              string    `json:"created,omitempty"`
	Edited               string    `json:"edited,omitempty"`
	URL                  string    `json:"url,omitempty"`
}

//starships swapi struct
type starships struct {
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
	Pilots               *[]people `json:"people,omitempty"`
	Films                *[]films  `json:"films,omitempty"`
	Created              string    `json:"created,omitempty"`
	Edited               string    `json:"edited,omitempty"`
	URL                  string    `json:"url,omitempty"`
}
