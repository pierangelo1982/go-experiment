package model

type Credenziali struct {
	Token string
}

type Customer struct {
	ID              string `json:"id"`
	DataInserimento string `json:"data_inserimento"`
	Nome            string `json:"nome"`
	Cognome         string `json:"cognome"`
	IDGommista      string `json:"id_gommista"`
	Indirizzo       string `json:"indirizzo"`
	Citta           string `json:"citta"`
	Provincia       string `json:"provincia"`
	Cap             string `json:"cap"`
	Telefono        string `json:"telefono"`
	Email           string `json:"email"`
	Cellulare       string `json:"cellulare"`
	Note            string `json:"note"`
	//IsApp                     bool     `json:"is_app"`
	IsApp                     string    `json:"is_app"`
	DataScadenzaPatente       string    `json:"data_scadenza_patente"`
	DataRegistrazioneApp      string    `json:"data_registrazione_app"`
	DataConfermaRegistrazione string    `json:"data_conferma_registrazione_app"`
	DataUltimoLogin           string    `json:"data_ultimo_login_app"`
	Veicoli                   []Veicoli `json:"veicoli"`
}

type Veicoli struct {
	ID               string `json:"id,int"`
	IDCliente        string `json:"id_cliente"`
	Marca            string `json:"marca"`
	Modello          string `json:"veicolo"`
	PercorrenzaAnnua string `json:"percorrenza_annua, int"`
	KmAttuali        string `json:"km_attuali, int"`
	Targa            string `json:"targa"`
	TipoVeicolo      string `json:"tipo_veicolo"`
	IsApp            string `json:"is_app"`
}
