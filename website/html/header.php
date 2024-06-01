
<?php 
include('constants.php');

$currentPage = $_SERVER['SCRIPT_NAME'];
$currentPage = substr($currentPage, 1);

include('helpers.php');
?>
<!doctype html>
<html lang='en'>
<head>
    <meta charset='UTF-8'>
    <link rel="shortcut icon" href="assets/img/gocure.ico"/>
    <meta name='viewport'
          content='width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0'>
    <meta http-equiv='X-UA-Compatible' content='ie=edge'>
    <title>Gocure</title>
    <link rel='preconnect' href='https://fonts.gstatic.com'>
    <link rel="stylesheet" href="assets/css/style.css">
    <link rel="stylesheet" href="assets/css/dismissible.css" />
    <script src="assets/js/dismissible.js"></script>
    <script src="assets/js/scripts.js"></script>
</head>
<body>
    <div id="loading"></div>
    <div id="dismissible-container"></div>
    <script src="assets/js/dismissibleInit.js"></script>

    <div class="header">
        <div class="header-title">
            <a href="index.php">
                <img src="assets/img/gocure_80.png" />
            </a>
        </div>
        <a href="htmlReport.php">
            <div class="header-block <?php if($currentPage == 'htmlReport.php'){echo 'active';} ?>">
                HTML Report
            </div>
        </a>
        <a href="embed.php">
            <div class="header-block <?php if($currentPage == 'embed.php'){echo 'active';} ?>">
                Embed Files
            </div>
        </a>
    </div>