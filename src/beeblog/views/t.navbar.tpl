{{define "navbar"}}
<div>
 <a class ="navbar-brand" href="/">myblog</a>
        <ul class="nav navbar-nav">
          <li {{if .IsHome}}class="active"{{end}}><a href="/">home</a></li>
          <li {{if .IsCategory}}class="active"{{end}}><a href="/category">category</a></li>
          <li {{if .IsTopic}}class="active"{{end}}><a href="/topic">topic</a></li>
        </ul>
</div>

<div class="pull-right">
    <ul class = "nav navbar-nav">
    {{if .IsLogin}}
    <li><a href="/login?exit=true">Logout</a></li>
    {{else}}
    <li><a href="/login">Login</a></li>
    {{end}}
    </ul>
</div>
{{end}}
