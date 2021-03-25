package extres

// External Resource keeper

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
// that are used across the Accord suite.
type ExternalResources struct {
	Env            int    `json:"Env"`
	AuthNHost      string `json:"AuthNHost"`
	AuthNType      string `json:"AuthNType"`
	AuthNPort      int    `json:"AuthNPort"`
	Dbuser         string `json:"Dbuser"`
	Dbname         string `json:"Dbname"`
	Dbpass         string `json:"Dbpass"`
	Dbhost         string `json:"Dbhost"`
	Dbport         int    `json:"Dbport"`
	Dbtype         string `json:"Dbtype"`
	SMTPHost       string `json:"SmtpHost"`
	SMTPPort       int    `json:"SmtpPort"`
	SMTPLogin      string `json:"SmtpLogin"`
	SMTPPass       string `json:"SmtpPass"`
	RRDbuser       string `json:"RRDbuser"`
	RRDbname       string `json:"RRDbname"`
	RRDbpass       string `json:"RRDbpass"`
	RRDbhost       string `json:"RRDbhost"`
	RRDbport       int    `json:"RRDbport"`
	RRDbtype       string `json:"RRDbtype"`
	MojoDbuser     string `json:"MojoDbuser"`
	MojoDbname     string `json:"MojoDbname"`
	MojoDbpass     string `json:"MojoDbpass"`
	MojoDbhost     string `json:"MojoDbhost"`
	MojoDbport     int    `json:"MojoDbport"`
	MojoDbtype     string `json:"MojoDbtype"`
	MojoWebAddr    string `json:"MojoWebAddr"`
	WREISDbuser    string `json:"WREISDbuser"`
	WREISDbname    string `json:"WREISDbname"`
	WREISDbpass    string `json:"WREISDbpass"`
	WREISDbhost    string `json:"WREISDbhost"`
	WREISDbport    int    `json:"WREISDbport"`
	WREISDbtype    string `json:"WREISDbtype"`
	Timezone       string `json:"Timezone"`       // see $GOROOT/lib/time/zoneinfo.zip, or try: tar tvf $GOROOT/lib/time/zoneinfo.zip
	SessionTimeout int    `json:"SessionTimeout"` // session timeout in minutes
	RootHandler    string `json:"RootHandler"`    // Handler for web app where url filepath is "/"
	Tester1Name    string `json:"Tester1Name"`    // username for tester 1
	Tester1Pass    string `json:"Tester1Pass"`    // password for tester 1
	Tester2Name    string `json:"Tester2Name"`    // username for tester 2
	Tester2Pass    string `json:"Tester2Pass"`    // password for tester 2
	RepoUser       string `json:"RepoUser"`       // for Artifactory
	RepoPass       string `json:"RepoPass"`       // for Artifactory
	RepoURL        string `json:"RepoURL"`        // for Artifactory
	S3Region       string `json:"S3Region"`       //
	S3BucketHost   string `json:"S3BucketHost"`   // S3 Bucket host. It is for storing profile images. e.g., https://s3.ap-south-1.amazonaws.com
	S3BucketName   string `json:"S3BucketName"`   // S3 Bucket name. e.g., upload-profile-image
	S3BucketKeyID  string `json:"S3BucketKeyID"`  // Access key id to use AWS S3 service
	S3BucketKey    string `json:"S3BucketKey"`    // Secret key to use AWS S3 service
	CryptoKey      string `json:"CryptoKey"`      // key for encryption/decription, must be 32 chars long
	MapKey         string `json:"MapKey"`         // key for using google map
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
	case "receipts":
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
	case "wreis":
		switch a.Env {
		case APPENVDEV: //dev
			s = fmt.Sprintf("%s:@/%s?charset=utf8&parseTime=True", a.WREISDbuser, dbname)
		case APPENVPROD: //production
			s = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True",
				a.WREISDbuser, a.WREISDbpass, a.WREISDbhost, a.WREISDbport, dbname)
		default:
			fmt.Printf("Unhandled configuration environment: %d\n", a.Env)
			os.Exit(1)
		}
	default:
		fmt.Printf("extres.GetSQLOpenString: db %s is not recognized, a restrictive login string is returned\n", dbname)
		s = fmt.Sprintf("%s:@/%s?charset=utf8&parseTime=True", a.Dbuser, dbname)
	}
	return s
}
