package demo

import (
	"bytes"
	"fmt"
	"html/template"
)

var Cart = []Item{
	{Name: "Coffee", Price: 3.5, Qty: 2},
	{Name: "Cloud Computing textbook", Price: 12, Qty: 1},
}

var pageTemplate = template.Must(template.New("page").Parse(`<!doctype html>
<html lang="en"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width,initial-scale=1">
<title>CI/CD Demo</title>
<style>
  :root{--accent:{{.Accent}}}
  *{box-sizing:border-box}
  body{margin:0;min-height:100vh;display:grid;place-items:center;background:#05070C;color:#F4F7FD;font-family:-apple-system,BlinkMacSystemFont,"Segoe UI",Roboto,sans-serif;padding:16px}
  main{width:min(560px,100%);text-align:center}
  h1{font-size:clamp(32px,7vw,56px);line-height:1.05;margin:0 0 12px;letter-spacing:-.02em}
  .sha{display:inline-block;font-family:ui-monospace,Menlo,monospace;padding:8px 16px;border-radius:999px;border:2px solid var(--accent);color:var(--accent);font-size:15px}
  .card{margin-top:28px;text-align:left;background:#0E1420;border:1px solid #1C2536;border-radius:20px;padding:22px}
  ul{list-style:none;margin:0;padding:0}
  li{display:flex;justify-content:space-between;padding:8px 0;color:#A3AEC4}
  .total{display:flex;justify-content:space-between;border-top:1px solid #1C2536;margin-top:10px;padding-top:14px;font-size:20px}
  .total b{color:var(--accent)}
  small{display:block;margin-top:22px;color:#5F6C86;font-family:ui-monospace,Menlo,monospace}
</style></head>
<body><main>
  <h1>{{.Headline}}</h1>
  <span class="sha">version {{.Version}}</span>
  <div class="card">
    <ul>{{range .Items}}<li><span>{{.Qty}} × {{.Name}}</span><b>${{.Subtotal}}</b></li>{{end}}</ul>
    <div class="total"><span>Total with STUDENT20</span><b>${{.Total}}</b></div>
  </div>
  <small>deployed {{.DeployedAt}}</small>
</main></body></html>`))

func RenderPage(version, deployedAt string) ([]byte, error) {
	if len(version) > 7 {
		version = version[:7]
	}
	type row struct {
		Name     string
		Qty      int
		Subtotal string
	}
	rows := make([]row, 0, len(Cart))
	for _, item := range Cart {
		rows = append(rows, row{item.Name, item.Qty, fmt.Sprintf("%.2f", item.Price*float64(item.Qty))})
	}
	data := struct {
		Headline   string
		Accent     string
		Version    string
		DeployedAt string
		Items      []row
		Total      string
	}{Headline, Accent, version, deployedAt, rows, fmt.Sprintf("%.2f", CalculateTotal(Cart, "STUDENT20"))}
	var output bytes.Buffer
	err := pageTemplate.Execute(&output, data)
	return output.Bytes(), err
}
