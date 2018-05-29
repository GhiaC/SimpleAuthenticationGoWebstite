package Models

type PageVariables struct {
	PageTitle string
	Answer    string
}

type HomePageVariables struct {
	Date        string
	Time        string
	LoginStatus string
}

type StatusPageVariables struct {
	Users [] User
}