package models

type RegisterUserInput struct {
	Nama     string `json:"nama" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

type ReceivedData struct {
	First AntaresDetail `json:"m2m:cin"`
}

type AntaresDetail struct {
	Rn  string `json:"rn"`
	Ty  int    `json:"ty"`
	Ri  string `json:"ri"`
	Pi  string `json:"pi"`
	Ct  string `json:"ct"`
	Lt  string `json:"lt"`
	St  int    `json:"st"`
	Cnf string `json:"cnf"`
	Cs  int    `json:"cs"`
	Con string `json:"con"`
}

type SensorData struct {
	Kapasitas     int    `json:"kapasitas"`
	Status_device string `json:"status"`
}

type ObjectAntares5 struct {
	Rn  string `json:"rn"`
	Ty  int    `json:"ty"`
	Ri  string `json:"ri"`
	Pi  string `json:"pi"`
	Ct  string `json:"ct"`
	Lt  string `json:"lt"`
	St  int    `json:"st"`
	Cnf string `json:"cnf"`
	Cs  int    `json:"cs"`
	Con string `json:"con"`
}

type ObjectAntares4 struct {
	M2m_cin ObjectAntares5 `json:"m2m:cin"`
}

type ObjectAntares3 struct {
	M2m_rep ObjectAntares4 `json:"m2m:rep"`
	M2m_rss int            `json:"m2m:rss"`
}

type ObjectAntares2 struct {
	M2m_nev ObjectAntares3 `json:"m2m:nev"`
	M2m_sud bool           `json:"m2m:sud"`
	M2m_sur string         `json:"m2m:sur"`
}

type ObjectAntares1 struct {
	First ObjectAntares2 `json:"m2m:sgn"`
}
