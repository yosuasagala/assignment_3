<html>
    <head>
        <title>Login</title>
    </head>

    <body>
    <form action="/edit" method="post">
        <div>
        <table>
            <tr>
            <td>ID</td>
            <td>:</td> 
            <td><input type="text" name="uid" value="{{.Users.ID}}"><br></td>
            </tr>
            <tr>
            <td>Username</td>
            <td>:</td> 
            <td><input type="text" name="username" value="{{.Users.Username}}"><br></td>
            </tr>  
            <tr>
            <td>Password</td>
            <td>:</td> 
            <td><input type="text" name="password" value="{{.Users.Password}}"><br></td>
            </tr>
            <tr>
            <tr>
            <td><input type="submit" value="Submit" ></td>
            </tr>
        </table>
        </div>
        </form>
    </body>

</html>