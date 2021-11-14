Hi, I'm Philipp and I like writing web servers, tinkering with my phone, automation and a bunch of other stuff.

You might be interested in [my blog](https://blog.010.one) where I sometimes write about technical stuff, or my Android apps. You can install them by adding [this F-Droid repo](https://github.com/xarantolus/fdroid) to your [F-Droid](https://f-droid.org/) client.

Anyways, take a look around :)

<details>
  <summary>Open this or <a href="https://010.one/">visit my website to see an overview of my projects</a></summary>

{{ range .Categories }}
#### {{.Name}}{{with .Description}}
{{.}}{{end}}
{{range .Repos}}
- {{with . | repo}}[{{.Name}}]({{.Link}}{{with .Title}} "{{.}}"{{end}}): {{.Desc | transform}}{{end}}{{end}}
{{end}}

</details>
