package erring

var ErrGettingClientsFromDB = NewAppError("client:could-not-retrieve-clients", "error retrieving clients from database")
