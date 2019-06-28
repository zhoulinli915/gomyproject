{{template "public/header.tpl" .}}
<div style="width:1024px;margin:0 auto; border:#f00 solid 1px;">
    <form action="{{.action}}"  method="post">
        <div>用户名：<input type="text" name="username" /></div>
        <div>密码：<input type="password" name="pwd" /></div>
        <div><input type="submit" name="submit" value="提 交" /></div>

    </for>
</div>
{{template "public/footer.tpl" .}}
