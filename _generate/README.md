Hi, I'm Philipp and I like writing web servers, tinkering with my phone, building robots (and electronics in general!), automation and a bunch of other stuff. You might also be interested in my [blog](https://xarantolus.github.io/blog/) where I sometimes write guides. Anyways, take a look around :)


<details>
  <summary>Open this or visit <a href="https://xarantolus.github.io/">my website</a> to see an overview of my projects</summary>

{{ range .Categories }}
#### {{.Name}}{{with .Description}}
{{.}}{{end}}
{{range .Repos}}
- {{with . | repo}}[{{.Name}}]({{.Link}}{{with .Title}} "{{.}}"{{end}}): {{.Desc | transform}}{{end}}{{end}}
{{end}}

</details>
