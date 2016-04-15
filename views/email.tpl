<!DOCTYPE html>
<html lang="zh-CN">
	<head>
		<meta charset="UTF-8">
		<title>邮件验证</title>
		<style type="text/css">
		.alert {
		padding: 15px;
		margin-bottom: 20px;
		border: 1px solid transparent;
		border-radius: 4px;
		}
		.alert-info {
		color: #31708f;
		background-color: #d9edf7;
		border-color: #bce8f1;
		}
		</style>
	</head>
	<body>
		<h1>欢迎您注册Beego的邮件测试范例</h1>
		<div class="alert alert-info">
			请点击链接激活邮箱：
			<a href="{{.Link}}" class="alert-link">{{.Link}}</a>
		</div>
	</body>
</html>