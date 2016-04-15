<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>注册</title>
		<link rel="stylesheet" href="//cdn.bootcss.com/bootstrap/3.3.5/css/bootstrap.min.css">
		<link rel="stylesheet" href="//cdn.bootcss.com/bootstrap/3.3.5/css/bootstrap-theme.min.css">
	</head>
	<body>
		<div class="container">
			<h1>邮箱注册</h1>
			<form method="post" action="/register">
				<fieldset>
					<div class="form-group">
						<label for="Name">昵称</label>
						<input type="text" name="Name" class="form-control" placeholder="请输入昵称">
					</div>
					<div class="form-group">
						<label for="Email">邮箱</label>
						<input type="email" name="Email" class="form-control" placeholder="请输入邮箱地址">
					</div>
					<button type="submit" class="btn btn-success">注册</button>
				</fieldset>
			</form>
		</div>
		<script src="//cdn.bootcss.com/jquery/1.11.3/jquery.min.js"></script>
		<script src="//cdn.bootcss.com/bootstrap/3.3.5/js/bootstrap.min.js"></script>
	</body>
</html>