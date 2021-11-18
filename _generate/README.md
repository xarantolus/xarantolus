Hi, I'm Philipp and I like writing web servers, tinkering with my phone, automation and a bunch of other stuff.

You might be interested in my Android apps, which you can install by adding this [this F-Droid repo](https://github.com/xarantolus/fdroid) to your [F-Droid](https://f-droid.org/) client (or just downloading the APK file).

I also have [a blog](https://blog.010.one) where I sometimes write about technical stuff that interests me :)

<details>
  <summary>Click here to open a list of my projects or <a href="https://010.one/"><b>visit my website for a better overview</b></a></summary>

{{ range .Categories }}
#### {{.Name}}{{with .Description}}
{{.}}{{end}}
{{range .Repos}}
- {{with . | repo}}[{{.Name}}]({{.Link}}{{with .Title}} "{{.}}"{{end}}): {{.Desc | transform}}{{end}}{{end}}
{{end}}

So you've reached the end of this overview, but maybe you want to visit <a href="https://010.one/"><b>the web site</b></a> now?
 
</details>
