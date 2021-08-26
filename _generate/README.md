Hi, I'm Philipp and I like writing web servers, tinkering with my phone, building robots (and electronics in general!), automation and a bunch of other stuff. You might also be interested in my [blog](https://xarantolus.github.io/blog/) where I sometimes write guides. Anyways, take a look around :)

{{ range .Categories }}
#### {{.Name}}{{with .Description}}
{{.}}{{end}}
{{range .Repos}}
- {{with . | repo}}[{{.Name}}]({{.Link}}{{with .Title}} "{{.}}"{{end}}): {{.Desc | transform}}{{end}}{{end}}
{{end}}
