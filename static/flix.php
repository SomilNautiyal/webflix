<html>
<head></head>
<body bgcolor="grey"><center></b><br><br>
<?php
$email = filter_input(INPUT_POST, 'email');
$pass = filter_input(INPUT_POST, 'pass');

if (!empty($email)){
if (!empty($pass)){
$host = "localhost";
$dbusername = "root";
$dbpassword = "";
$dbname = "netflix";
// Create connection
$conn = new mysqli ($host, $dbusername, $dbpassword, $dbname);


if (mysqli_connect_error()){
die('Connect Error ('. mysqli_connect_errno() .') '
. mysqli_connect_error());
}
else{
$sql = "INSERT INTO subscribes (email, pass)values ('$email','$pass')";
if ($conn->query($sql)){
echo "New record is inserted sucessfully";
}
else{
echo "Error: ". $sql ."
". $conn->error;
}
$conn->close();
}
}
else{
echo "Password should not be empty";
die();
}
}
else{
echo "Email should not be empty";
die();
}
?>
</b><br><br>
<div class="btn-group">
	<button onclick="window.location.href = 'signin.html';">Previous</button>
	<button onclick="window.location.href = 'getstarted.html';">Next</button>
	<button onclick="window.location.href = 'webflix.html';">Home</button>
</div>
</center>
</body>
</html>