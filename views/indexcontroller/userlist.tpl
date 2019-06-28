{{template "public/header.tpl" .}}
<div style="width:1024px;margin:0 auto; border:#f00 solid 1px;">
  {{ if .collegelist}}
      <ul>
        <li>编号</li>
        <li></li>
        <li></li>
        <li></li>
        <li></li>
        <li></li>
        <li></li>
        <li></li>
        <li></li>
        <li>操作</li>
      </ul>
      {{range $ind, $college := .collegelist}}
         <ul>
            <li>{{$college.Id}}</li>
            <li>{{$college.CollegeName}}</li>
            <li>{{$college.CollegeName}}</li>
            <li>{{$college.CollegeName}}</li>
            <li>{{$college.CollegeName}}</li>
            <li>{{$college.CollegeName}}</li>
            <li>{{$college.CollegeName}}</li>
            <li>{{$college.CollegeName}}</li>
            <li>{{$college.CollegeName}}</li>
            <li>{{$college.CollegeName}}</li>
          </ul>
      {{end}}
  {{else}}
  目前没有任何数据！
  {{end}}

</div>
{{template "public/footer.tpl" .}}
