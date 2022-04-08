package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"thesis_work/pkg/config"
	"thesis_work/pkg/forms"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	flash := app.session.PopString(r, "flash")
	app.render(w, r, "home.page.tmpl", &templateData{
		Flash: flash,
	})
}

func (app *application) redirect(w http.ResponseWriter, r *http.Request) {
	flash := app.session.PopString(r, "flash")
	app.render(w, r, "dimash.page.tmpl", &templateData{
		Flash: flash,
	})
}

func (app *application) ipCheckerForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "ip-checker.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) ipChecker(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	form := forms.New(r.PostForm)
	ip := form.Get("ip")

	if net.ParseIP(ip) == nil {
		app.render(w, r, "error.page.tmpl", &templateData{Form: form})
		return
	}

	http.Redirect(w, r, "/ip-info?ip="+ip, http.StatusSeeOther)
	return
}

func (app *application) ipInfo(w http.ResponseWriter, r *http.Request) {
	ip := r.URL.Query().Get("ip")

	//-------
	url := "https://www.virustotal.com/api/v3/search?query=" + ip

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")

	req.Header.Add("x-apikey", "b264d918456da4d250bf936dbd1df2f84fbbc45a122598fc39fa2848b30e21a7")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	//fmt.Println(res)

	JSONobj := &config.Config{}
	if err := json.Unmarshal(body, JSONobj); err != nil {
		fmt.Printf("JSON unmarshal error: %v\n", err)
		return
	}
	//fmt.Println(JSONobj.Links)
	//fmt.Println(JSONobj.Data[0].ID)
	//fmt.Println(JSONobj.Data[0].Type)
	//fmt.Println(JSONobj.Data[0].Links.Self)
	//fmt.Println(JSONobj.Data[0].Attributes.AsOwner)

	//-------
	flash := app.session.PopString(r, "flash")
	app.render(w, r, "ip-info.page.tmpl", &templateData{
		Flash:  flash,
		Config: JSONobj,
		//Data:   JSONobj.Data,
	})

}
