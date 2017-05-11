<?php
  //globals and includes
  include './../Go/src/Credentials/credentials.php';

?>
<?php
  //main

  // Create connection
  $conn = new mysqli($server, $user, $password, $database);

  // Check connection
  if ($conn->connect_error) {
    die("Connection failed: " . $conn->connect_error);
  }

  $emailAddress = $argv[1];

  $title = "Welcome to SkyTracker!";

  $sql = "SELECT * FROM Users WHERE UserEmailAddress = '$emailAddress';";
  $result = mysqli_query($conn, $sql);
  $result = mysqli_fetch_assoc($result);

  $salt = $result['UserSalt'];

  $emailBody = "";
  $emailBody .= "<h3>Welcome!</h3>";
  $emailBody .= "<p> It appears that you have signed up to receive flight price reports from SkyTracker, if this is not you, then please feel free to <a href = 'http://www.skytracker.co?unsubscribe&email=$emailAddress&token=$salt'>unsubscribe</a> or check us out at <a href = 'http://www.skytracker.co'>SkyTracker</a> and decide for yourself if you want to join! <br><br>If you <b>did</b> mean to sign up, please do us a favour and add us as a contact so we can get our reports straight to your inbox!<br><br>Thanks!</p>";

  echo $emailBody . "\n";

  // exec("echo \"$emailBody\" > \"testfile.html\"");
  exec("echo \"$emailBody\" | mail -a \"From: SkyTracker <no-reply@skytracker.co>\" -a \"Content-type: text/html\" -s \"$title\" $emailAddress");

  $conn->close();

?>
