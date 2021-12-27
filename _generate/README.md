Hi, I'm Philipp and I like writing web servers, tinkering with my phone, automation and a bunch of other stuff.

I'm also creating a few Android apps; you can add [this F-Droid repo](https://github.com/xarantolus/fdroid) to your [F-Droid](https://f-droid.org/) client to download them (or just download the APK files).

<details>
  <summary>Click here to open a list of my projects or <a href="https://010.one/"><b>visit my website for a better overview</b></a></summary>

{{ range .Categories }}
#### {{.Name}}{{with .Description}}
{{.}}{{end}}
{{range .Repos}}
- {{with . | repo}}[**{{.Name}}**]({{.Link}}{{with .Title}} "{{.}}"{{end}}): {{.Desc | transform}}{{end}}{{end}}
{{end}}

So you've reached the end of this overview, but maybe you want to visit <a href="https://010.one/"><b>the web site</b></a> now?
 
</details>
