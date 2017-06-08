package extres

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// APPENVDEV et. al. are constants describing the environment where
// the app is running. It is set via the config.json file
const (
	APPENVDEV  = 0
	APPENVPROD = 1
	APPENVQA   = 2
)

// ExternalResources is a type defining several resources
// that are used accross the Accord suite.
type ExternalResources struct {
	Env         int    `json:"Env"`
	Dbuser      string `json:"Dbuser"`
	Dbname      string `json:"Dbname"`
	Dbpass      string `json:"Dbpass"`
	Dbhost      string `json:"Dbhost"`
	Dbport      int    `json:"Dbport"`
	Dbtype      string `json:"Dbtype"`
	SMTPHost    string `json:"SmtpHost"`
	SMTPPort    int    `json:"SmtpPort"`
	SMTPLogin   string `json:"SmtpLogin"`
	SMTPPass    string `json:"SmtpPass"`
	RRDbuser    string `json:"RRDbuser"`
	RRDbname    string `json:"RRDbname"`
	RRDbpass    string `json:"RRDbpass"`
	RRDbhost    string `json:"RRDbhost"`
	RRDbport    int    `json:"RRDbport"`
	RRDbtype    string `json:"RRDbtype"`
	MojoDbuser  string `json:"MojoDbuser"`
	MojoDbname  string `json:"MojoDbname"`
	MojoDbpass  string `json:"MojoDbpass"`
	MojoDbhost  string `json:"MojoDbhost"`
	MojoDbport  int    `json:"MojoDbport"`
	MojoDbtype  string `json:"MojoDbtype"`
	MojoWebAddr string `json:"MojoWebAddr"`
	Timezone    string `json:"Timezone"`
}

// ReadConfig will read the configuration file "config.json" if
// it exists in the current directory
//=======================================================================================
func ReadConfig(fname string, a *ExternalResources) error {
	var err error
	a.Timezone = "GMT" // default, overridden by value in config.json
	if _, err = os.Stat(fname); err != nil {
		return err
	}

	content, err := ioutil.ReadFile(fname)
	if err != nil {
		return err
	}
	err = json.Unmarshal(content, a)
	// fmt.Printf("After unmarshal: a = %#v\n", *a)
	return err
}

// GetSQLOpenString builds the string to use for opening an sql database.
// Input string is the name of the database:  "accord" for phonebook, "rentroll" for RentRoll
// Returns:  a string to pass to sql.Open()
//=======================================================================================
func GetSQLOpenString(dbname string, a *ExternalResources) string {
	s := ""
	switch strings.ToLower(dbname) {
	case "accord":
		switch a.Env {
		case APPENVDEV: //development
			s = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True",
				a.Dbuser, a.Dbpass, dbname)
		case APPENVPROD: //production
			s = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True",
				a.Dbuser, a.Dbpass, a.Dbhost, a.Dbport, dbname)
		default:
			fmt.Printf("Unhandled configuration environment: %d\n", a.Env)
			os.Exit(1)
		}
	case "rentroll":
		switch a.Env {
		case APPENVDEV: //dev
			s = fmt.Sprintf("%s:@/%s?charset=utf8&parseTime=True", a.RRDbuser, dbname)
		case APPENVPROD: //production
			s = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True",
				a.RRDbuser, a.RRDbpass, a.RRDbhost, a.RRDbport, dbname)
		default:
			fmt.Printf("Unhandled configuration environment: %d\n", a.Env)
			os.Exit(1)
		}
	case "mojo":
		switch a.Env {
		case APPENVDEV: //dev
			s = fmt.Sprintf("%s:@/%s?charset=utf8&parseTime=True", a.MojoDbuser, dbname)
		case APPENVPROD: //production
			s = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True",
				a.MojoDbuser, a.MojoDbpass, a.MojoDbhost, a.MojoDbport, dbname)
		default:
			fmt.Printf("Unhandled configuration environment: %d\n", a.Env)
			os.Exit(1)
		}
	default:
		s = fmt.Sprintf("%s:@/%s?charset=utf8&parseTime=True", a.Dbuser, dbname)
	}
	return s
}
