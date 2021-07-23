package product

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path"
	"text/template"
	"time"
)

type ProductReportFilter struct {
	NameFilter         string `json:"productName"`
	ManufacturerFilter string `json:"manufacturer"`
	SKUFilter          string `json:"sku"`
}

func handleProductReport(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var productFilter ProductReportFilter
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &productFilter)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		products, err := searchForProductData(productFilter)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		t := template.New("report.gotmpl").Funcs(template.FuncMap{"mod": func(i, j int) bool { return i%j == 0 }})
		t, err = t.ParseFiles(path.Join("templates", "report.gotmpl"))

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var tmpl bytes.Buffer

		if len(products) > 0 {
			err = t.Execute(&tmpl, products)
		} else {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		rdr := bytes.NewReader(tmpl.Bytes())
		w.Header().Set("Content-Disposition", "Attachement")
		http.ServeContent(w, r, "report.html", time.Now(), rdr)
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
