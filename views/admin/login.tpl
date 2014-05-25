<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<title>用户登录</title>
<link href="/static/css/login.css" rel="stylesheet" type="text/css" />
<script type="text/javascript" src="/static/js/jquery-2.1.1.js"></script>
</head>
<body>
    <div id="login">
	     <div id="top">
		      <div id="top_left"><img src="/static/img/login_03.gif" /></div>
			  <div id="top_center"></div>
		 </div>
		 <div id="center">
		      <div id="center_left"></div>
			  <div id="center_middle">
			       <div id="login_user">用 户
			         <input type="text" name="username" />
			       </div>
				   <div class="login_input">密   码
				     <input type="password" name="password" />
				   </div>
				   <div id="btn"><a href="javascript:void(0)" id="nowLogin">登录</a><a href="/admin/public/register">注册</a></div>
			  </div>
			  <div id="center_right"></div>		 
		 </div>
		 <div id="down">
		      <div id="down_left">
			      <div id="inf">
                       <span class="inf_text">版本信息</span>
					   <span class="copyright">V信公众号管理平台</span>
			      </div>
			  </div>
			  <div id="down_center"></div>		 
		 </div>
	</div>
</body>
</html>
<script type="text/javascript">
	$('#nowLogin').click(function(){
		var param = new Object();
		param.username = $('input[name="username"]').val();
		param.password = $('input[name="password"]').val();
		if(param.username.length <= 0){
			alert('请输入用户名！')
		}else if(param.password.length <= 0){
			alert('请输入密码！')
		}else{
			$.post('/admin/public/CheckLogin',param,function(data){
				alert(data.res);
			})
		}
		return false ;
	})	
</script>