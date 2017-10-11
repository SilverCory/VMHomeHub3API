package vmapi

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const URL = "http://192.168.100.1/"

var (
	ErrorBadPassword = errors.New("incorrect password")
	ErrorLockedOut   = errors.New("too many failed attempts to log in")
	ErrorSingleUser  = errors.New("a user is already logged in")
	ErrorUnsupported = errors.New("the interface or router being used isn't supported")
)

type Instance struct {
	Timeout  uint
	OnLAN    bool
	password string
}

func New(password string) (*Instance, error) {

	ret := &Instance{
		Timeout:  30,
		password: password,
	}

	if err := ret.Login(); err != nil {
		return nil, err
	}

	return ret, nil

}

func (i *Instance) Login() error {

	doc, err := goquery.NewDocument(URL)
	if err != nil {
		return err
	}

	passwordInput := doc.Find("#password")
	value, exists := passwordInput.Attr("name")
	if !exists {
		return ErrorUnsupported
	}

	loginResponse, err := http.PostForm(URL+"cgi-bin/VmLoginCgi", url.Values{value: {i.password}})
	if err != nil {
		return err
	}

	doc, err = goquery.NewDocumentFromResponse(loginResponse)
	if err != nil {
		return err
	}

	documentLines := strings.Split(doc.Text(), "\n")
	for _, line := range documentLines {

		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "var res=\"") {

			res, err := strconv.Atoi(string(line[9]))
			if err != nil {
				continue
			}

			switch res {
			case 1:
				return ErrorBadPassword
			case 3:
				return ErrorLockedOut
			}
			continue

		} else if strings.HasPrefix(line, "var defpass=\"") {

			defPass, err := strconv.Atoi(string(line[14]))
			if err != nil {
				continue
			}

			if defPass == 1 {
				fmt.Println("You are using the default password. THIS ISN'T SAFE!!!")
				fmt.Println("You are using the default password. THIS ISN'T SAFE!!!")
				fmt.Println("You are using the default password. THIS ISN'T SAFE!!!")
				fmt.Println("You are using the default password. THIS ISN'T SAFE!!!")
			}

			continue

		} else if strings.HasPrefix(line, "var singleUser=\"") {

			singleUser, err := strconv.Atoi(string(line[16]))
			if err != nil {
				continue
			}

			if singleUser == 1 {
				return ErrorSingleUser
			}

			continue
		} else if strings.HasPrefix(line, "var lanAccess=\"") {

			lanAccess, err := strconv.Atoi(string(line[15]))
			if err != nil {
				continue
			}

			i.OnLAN = lanAccess == 1
			continue
		}
	}

	return nil

}

func (i *Instance) Close() {
	http.Get(URL + "/VmLogout2.html")
}
