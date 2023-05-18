package entities

type LoginBasicInfo struct {
	Id             int    `json:"id"`
	Username       string `json:"username"`
	ChangePassword bool   `json:"change_password"`
	Rol            string `json:"rol"`
}

type LoginData struct {
	DNI      string `json:"dni"`
	Password string `json:"password"`
}

type RegisterData struct {
	DNI         string `json:"dni"`
	Name        string `json:"name"`
	LastName    string `json:"last_name"`
	SeniorityID string `json:"seniority_id"`
	AreaID      string `json:"area_id"`
}

type UserRegisterError struct {
	User RegisterData `json:"user"`
	Err  error        `json:"error"`
}

type UpdatePassword struct {
	Id          int    `json:"id"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

var Names = []string{
	"Calamar",
	"Foca",
	"Delfín",
	"Pez",
	"León",
	"Cachalote",
	"Ballena",
	"Anguila",
	"Medusa",
	"Tiburón",
	"Sardinas",
	"Gamba",
	"Manatí",
	"Trucha",
	"Pulpo",
	"Caracol",
	"Sapo",
	"Dragón",
	"Tortuga",
	"Langosta",
	"Carpa",
	"Loro",
	"Atún",
	"Cerdo",
	"Salmón",
	"Almeja",
	"Coral",
	"Piraña",
	"Pingüino",
	"Bacalao",
	"Oso",
}

var Lastnames = []string{
	"Feroz",
	"Veloz",
	"Escurridizo",
	"Curioso",
	"Astuto",
	"Elegante",
	"Intrépido",
	"Travieso",
	"Sigiloso",
	"Misterioso",
	"Juguetón",
}
