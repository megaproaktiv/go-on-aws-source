package lambdasummary

//build cloud local
import (
	"bytes"
	"fmt"
	"html/template"
)

type FunctionListPageData struct {
	PageTitle string
	Summaries []*LambdaSummary
}

type LambdaSummary struct {
	Account        string
	Region         string
	RuntimeCounter map[string]int
}

func Render(InfoPageData FunctionListPageData) string {

	tmpl, err := template.ParseFiles("result.html")
	if err != nil {
		fmt.Println("Got an error retrieving users:")
		fmt.Println(err)
		return "Error"
	}

	var body bytes.Buffer

	err = tmpl.Execute(&body, InfoPageData)
	if err != nil {
		fmt.Println("executing template:", err)
	}
	return body.String()
}
