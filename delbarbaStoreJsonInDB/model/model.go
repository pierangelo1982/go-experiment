package model

type Credenziali struct {
	Token string
}

type Customer struct {
	ID                        int    `json:"id"`
	DataInserimento           string `json:"data_inserimento"`
	Nome                      string `json:"nome"`
	Cognome                   string `json:"cognome"`
	IDGommista                string `json:"id_gommista"`
	Indirizzo                 string `json:"indirizzo"`
	Citta                     string `json:"citta"`
	Provincia                 string `json:"provincia"`
	Cap                       string `json:"cap"`
	Telefono                  string `json:"telefono"`
	Email                     string `json:"email"`
	Cellulare                 string `json:"cellulare"`
	Note                      string `json:"note"`
	IsApp                     string `json:"is_app"`
	DataScadenzaPatente       string `json:"data_scadenza_patente"`
	DataRegistrazioneApp      string `json:"data_registrazione_app"`
	DataConfermaRegistrazione string `json:"data_conferma_registrazione_app"`
	DataUltimoLogin           string `json:"data_ultimo_login_app"`
	//Veicoli                   []string `json:"veicoli"`
}

type Veicoli struct {
	ID               int
	IDCliente        int
	Marca            string
	Modello          string
	PercorrenzaAnnua int64
	KmAttuali        int64
	Targa            string
	TipoVeicolo      int
	IsApp            bool
	CanBeEdited      bool
	CanBeDeleted     bool
}
