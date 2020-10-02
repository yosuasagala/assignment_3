<html>
    <head>
		<title>Login</title>
		<meta charset= "UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.4.1/css/bootstrap.min.css">
        <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.4.1/js/bootstrap.min.js"></script>
    </head>

    <body>
		<div class="container">
			<div class="row">
				<div class="col-sm">
			<form action="/input" method="post">
				<div>
				<table>
					<div class="form-group mx-auto" >
						<label >Username</label>
						<input type="text" class="form-control w-25 p-3" id="username" >
					  </div>
					  <div class="form-group">
						<label >Password</label>
						<input type="password" class="form-control w-25 p-3" id="password">
					  </div>
                      <td>
					  <button type="submit" class="btn btn-primary" value="Login">Log In</button>
                      </td>
				</form>
				</div>
				</div>
		  </div>
        
    </body>

</html>