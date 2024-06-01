<?php

$title = '';
$merge = true;
$ignore = true;
$showEmbed = true;
$appVersion = '';
$testEnvironment = '';
$browser = '';
$platform = '';
$parallel = '';
$executed = '';

if (isset($_GET["rcf"]) && htmlspecialchars($_GET["rcf"]) != '' &&
    isset($_GET["rct"]) && htmlspecialchars($_GET["rct"]) == 'htmlReport') {
    RemoveCacheFile($GLOBALS['htmlReportOutputPath'],htmlspecialchars($_GET["rcf"]));
} else if (isset($_GET["rc"]) && htmlspecialchars($_GET["rc"]) == 'all' &&
        isset($_GET["rct"]) && htmlspecialchars($_GET["rct"]) == 'htmlReport') {
    $output = RemoveCache($GLOBALS['htmlReportOutputPath']);
    if ($output != '') {
        ?><script>dismissible.error("error removing files.");</script><?php
    } else {
        ?><script>dismissible.success("Files successfully removed!");</script><?php
    }
} else if (isset($_GET["embeddedFile"]) && htmlspecialchars($_GET["embeddedFile"]) != '') {
    RunGocure('-j',$GLOBALS['embedOutputPath'].htmlspecialchars($_GET["embeddedFile"]));
}  

if(isset($_POST['submit'])) {

    $GLOBALS['title'] = $_POST['title'];
    $GLOBALS['merge'] = filter_input(INPUT_POST, 'merge', FILTER_SANITIZE_STRING);
    $GLOBALS['ignore'] = filter_input(INPUT_POST, 'ignore', FILTER_SANITIZE_STRING);
    $GLOBALS['showEmbed'] = filter_input(INPUT_POST, 'showEmbed', FILTER_SANITIZE_STRING);
    $GLOBALS['appVersion'] = $_POST['appVersion'];
    $GLOBALS['testEnvironment'] = $_POST['testEnvironment'];
    $GLOBALS['browser'] = $_POST['browser'];
    $GLOBALS['platform'] = $_POST['platform'];
    $GLOBALS['parallel'] = $_POST['parallel'];
    $GLOBALS['executed'] = $_POST['executed'];

    if(basename($_FILES['filesToUpload']['name'][0]) != "") {

        RemoveCache($GLOBALS['htmlReportUploadPath']);
        
        for ($i=0; $i<count($_FILES['filesToUpload']['name']); $i++) {
            $filename = basename($_FILES['filesToUpload']['name'][$i]);
            $fileSize = $_FILES['filesToUpload']['size'][$i];
            $fileType = $_FILES['filesToUpload']['type'][$i];
            $fileTmpName = $_FILES['filesToUpload']['tmp_name'][$i];
            $targetFile = $GLOBALS['htmlReportUploadPath'].$filename;
            UploadFile($filename, $fileSize, $fileType, $fileTmpName, $targetFile, false);
        }
        
        if (sizeof($msgs) == 0) {
           RunGocure('-f',$GLOBALS['htmlReportUploadPath']);
        } else {
            $errors = '';
            for ($i=0; $i<count($msgs); $i++) {
                $errors .= $msgs[$i].'\n';
            } 
            ?><script>dismissible.error("<?php echo $errors; ?>");</script><?php
        }
    }
}

if ($handle = opendir($GLOBALS['htmlReportOutputPath'])) {
    $files = array();
    $thelist = '';
    while (false !== ($file = readdir($handle))) {
        array_push($files, $file);
    }
    closedir($handle);
    rsort($files);
    foreach ($files as &$file) {
        if ($file != "." && $file != ".." && $file != ".gitkeep") {
            $thelist .= '<div class="reports"><div class="reports-left">'.$file.'</div><div class="reports-right"><a href="'.$GLOBALS['htmlReportOutputPathLink'].$file.'" target="_Blank" class="tooltip top" data-tooltip="View"><span class="material-icons" >visibility</span></a><a href="'.$GLOBALS['htmlReportOutputPathLink'].$file.'" download class="tooltip top" data-tooltip="Download"><span class="material-icons">cloud_download</span></a><a href="htmlReport.php?rcf='.$file.'&rct=htmlReport" class=" tooltip top" data-tooltip="Delete"><span class="material-icons">delete</span></a></div></div>';
        }
    }

    if($thelist != ''){
        ?><div class="card" id="card-output-files">
        <div class="card-title">Output files <a href="htmlReport.php?rc=all&rct=htmlReport" class=" tooltip top" data-tooltip="Delete all" onclick="return showLoading()"><span class="material-icons">delete</span></a><a href="downloadAll.php?t=html" target="_Blank" class="tooltip top" data-tooltip="Download All"><span class="material-icons">cloud_download</span></a></div>
        <div class="card-body"><?php echo $thelist; ?></div>
        </div><?php
    }
}

function RunGocure($param, $execPath) {
    $gocure_exec_cmd = './gocure -h '.$param.' '.$execPath.' -o '.$GLOBALS['htmlReportOutputPath'];
    if ($GLOBALS['merge']) {
        $gocure_exec_cmd .= ' -m';
    }
    if ($GLOBALS['ignore']) {
        $gocure_exec_cmd .= ' -i';
    }
    if ($GLOBALS['showEmbed']) {
        $gocure_exec_cmd .= ' -s';
    }
    if ($GLOBALS['title'] != '') {
        $gocure_exec_cmd .= ' -t "'.$title.'"';
    }
    if ($GLOBALS['appVersion'] != '') {
        $gocure_exec_cmd .= ' -AppVersion "'.$appVersion.'"';
    }
    if ($GLOBALS['testEnvironment'] != '') {
        $gocure_exec_cmd .= ' -TestEnvironment "'.$testEnvironment.'"';
    }
    if ($GLOBALS['browser'] != '') {
        $gocure_exec_cmd .= ' -Browser "'.$browser.'"';
    }
    if ($GLOBALS['platform'] != '') {
        $gocure_exec_cmd .= ' -Platform "'.$platform.'"';
    }
    if ($GLOBALS['parallel'] != '') {
        $gocure_exec_cmd .= ' -Parallel "'.$parallel.'"';
    }
    if ($GLOBALS['executed'] != '') {
        $gocure_exec_cmd .= ' -Executed "'.$executed.'"';
    }

    $output = shell_exec('cd / && '.$gocure_exec_cmd.' 2>&1');
    if ($output != '') {
        ?><script>dismissible.error("errors executing Gocure.");</script><?php
    } else {
        ?><script>dismissible.success('HTML report successfully generated!');</script><?php
    }

    RemoveCache($GLOBALS['htmlReportUploadPath']);
}
?>