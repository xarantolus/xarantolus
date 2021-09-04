Hi, I'm Philipp and I like writing web servers, tinkering with my phone, automation and a bunch of other stuff. You might also be interested in [my blog](https://xarantolus.github.io/blog/) where I sometimes write about technical stuff. Anyways, take a look around :)


<details>
  <summary>Open this or <a href="https://xarantolus.github.io/">visit my website to see an overview of my projects</a></summary>

{{ range .Categories }}
#### {{.Name}}{{with .Description}}
{{.}}{{end}}
{{range .Repos}}
- {{with . | repo}}[{{.Name}}]({{.Link}}{{with .Title}} "{{.}}"{{end}}): {{.Desc | transform}}{{end}}{{end}}
{{end}}

</details>
