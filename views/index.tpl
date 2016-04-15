<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <title>首页</title>
    <link rel="stylesheet" href="//cdn.bootcss.com/bootstrap/3.3.5/css/bootstrap.min.css">
    <link rel="stylesheet" href="//cdn.bootcss.com/bootstrap/3.3.5/css/bootstrap-theme.min.css">
  </head>
  <body>
    <div class="container">
      <h1>你好，{{.Register.Name}}</h1>
      {{if .Status}}
      <div class="alert alert-success" role="alert">
        <h3>您的邮箱地址<mark>{{.Register.Email}}</mark>已验证成功！</h3>
      </div>
      {{else}}
      <div class="alert alert-warning" role="alert">
        <h3>邮件已发送到你的邮箱<mark>{{.Register.Email}}</mark>，请点击验证！</h3>
      </div>
      {{end}}
    </div>
    <script src="//cdn.bootcss.com/jquery/1.11.3/jquery.min.js"></script>
    <script src="//cdn.bootcss.com/bootstrap/3.3.5/js/bootstrap.min.js"></script>
  </body>
</html>