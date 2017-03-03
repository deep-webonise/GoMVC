package mafiaDB

import "time"

type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo

type User struct {
	Id                    int
	Name                  string
	Email                 string
	Password              string
	Created_timestamp     time.Time
	Games_played          int
	Games_won_as_mafia    int
	Games_won_notas_mafia int
}

type CharacterType struct {
	Id                int
	Name              string
	CanKill           int
	CanSave           int
	CanDetect         int
	CanVote           int
	Created_timestamp time.Time
}

type Round struct {
	Id   int
	Name string
}

type HostedGames struct {
	Id                int
	Name              string
	Creator_user_id   int
	Character_list    string
	Game_status       int
	Created_timestamp time.Time
	End_timestamp     time.Time
}

type JoinedPlayers struct {
	Id              int
	Hosted_game_id  int
	User_id         int
	Character_id    int
	Joined_datetime time.Time
	Is_alive        int
	Vote_against    int
	Action_against  int
	Extra_column_1  string
	Extra_column_2  string
}

type GameData struct {
	Id                       int
	Hosted_game              int
	Round_number             int
	Round_status             int
	Turn_user_id             int
	Winner_user_id           int
	Game_data_extra_column_1 string
	Game_data_extra_column_2 string
}

// connect to db using standard Go database/sql API
// use whatever database/sql driver you wish
