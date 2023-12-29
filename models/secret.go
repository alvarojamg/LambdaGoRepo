package models

type SecretRDSJson struct {
	UserName            string `json:"userName"`
	Password            string `json:"password"`
	Engine              string `json:"engine"`
	Host                string `json:"host"`
	Port                int    `json:"port"`
	DbClusterIdentifier string `json:"dbClusterIdentifier"`
}

type SiginUp struct {
	UserEmail string `json:"userEmail"`
	UserUUId  string `json:"userUUId"`
}
