<?php

if (isset($_GET["rcf"]) && htmlspecialchars($_GET["rcf"]) != '' &&
    isset($_GET["rct"]) && htmlspecialchars($_GET["rct"]) == 'embed') {
    RemoveCacheFile($GLOBALS['embedOutputPath'],htmlspecialchars($_GET["rcf"]));
} else if (isset($_GET["rc"]) && htmlspecialchars($_GET["rc"]) == 'all' &&
            isset($_GET["rct"]) && htmlspecialchars($_GET["rct"]) == 'embed') {
    $output = RemoveCache($GLOBALS['embedOutputPath']);
    if ($output != '') {
        ?><script>dismissible.error("error removing file '<?php echo $file ?>' from directory '<?php echo $folder ?>'.");</script><?php
    } else {
        ?><script>dismissible.success("Files '<?php echo htmlspecialchars($_GET["rcf"]) ?>' successfully removed!");</script><?php
    }
}

if(isset($_POST['submit'])) {

    $featureIndex = $_POST['featureIndex'];
    $scenarioIndex = $_POST['scenarioIndex'];
    $stepIndex = $_POST['stepIndex'];

    $reportFile = $GLOBALS['embedReportPath'].basename($_FILES['fileToUpload']['name']);

    if((basename($_FILES['fileToUpload']['name']) != "") && (basename($_FILES['filesToUpload']['name'][0]) != "")) {

        RemoveCache($GLOBALS['embedReportPath']);
        RemoveCache($GLOBALS['embedUploadPath']);

        $filename = basename($_FILES['fileToUpload']['name']);
        $fileSize = $_FILES['fileToUpload']['size'];
        $fileType = $_FILES['fileToUpload']['type'];
        $fileTmpName = $_FILES['fileToUpload']['tmp_name'];
        $targetFile = $GLOBALS['embedReportPath'].$filename;
        UploadFile($filename, $fileSize, $fileType, $fileTmpName, $targetFile, false);
        
        for ($i=0; $i<count($_FILES['filesToUpload']['name']); $i++) {
            $filename = basename($_FILES['filesToUpload']['name'][$i]);
            $fileSize = $_FILES['filesToUpload']['size'][$i];
            $fileType = $_FILES['filesToUpload']['type'][$i];
            $fileTmpName = $_FILES['filesToUpload']['tmp_name'][$i];
            $targetFile = $GLOBALS['embedUploadPath'].$filename;
            UploadFile($filename, $fileSize, $fileType, $fileTmpName, $targetFile, true);
        }
        
        if (sizeof($msgs) == 0) {
            $t = date("Ymd_His.").gettimeofday()["usec"];
            $t = str_replace(".","_",$t);

            $time = date('Ymd_his.u', time());
            $gocure_exec_cmd = './gocure -e -j '.$reportFile.' -u '.$GLOBALS['embedOutputPath'].'embedded_report_'.$t.'.json';
 
            if ($featureIndex != '') {
                $gocure_exec_cmd .= ' -a '.$featureIndex;
            }
            if ($scenarioIndex != '') {
                $gocure_exec_cmd .= ' -c '.$scenarioIndex;
            }
            if ($stepIndex != '') {
                $gocure_exec_cmd .= ' -p '.$stepIndex;
            }

            if ($handle = opendir($GLOBALS['embedUploadPath'])) {
                while (false !== ($file = readdir($handle))) {
                    if ($file != "." && $file != ".." && $file != ".gitkeep") {
                        $gocure_exec_cmd .= ' -l "'.$GLOBALS['embedUploadPath'].$file.'"';
                    }
                }
                closedir($handle);
            }

            $output = shell_exec('cd / && '.$gocure_exec_cmd.' 2>&1');
            if ($output != '') {
                ?><script>dismissible.error("errors executing Gocure.");</script><?php
            } else {
                ?><script>dismissible.success('Files successfully embedded!');</script><?php
            }

            RemoveCache($GLOBALS['embedReportPath']);
            RemoveCache($GLOBALS['embedUploadPath']);
        } else { 
            $errors = '';
            for ($i=0; $i<count($msgs); $i++) {
                $errors .= $msgs[$i].'\n';
            } 
            ?><script>dismissible.error("<?php echo $errors; ?>");</script><?php
        }
    }
}

if ($handle = opendir($GLOBALS['embedOutputPath'])) {
    $files = array();
    $thelist = '';
    while (false !== ($file = readdir($handle))) {
        array_push($files, $file);
    }
    closedir($handle);
    rsort($files);
    foreach ($files as &$file) {
        if ($file != "." && $file != ".." && $file != ".gitkeep") {
            $thelist .= '<div class="reports"><div class="reports-left">'.$file.'</div><div class="reports-right"><a href="'.$GLOBALS['embedOutputPathLink'].$file.'" target="_Blank" class="tooltip top" data-tooltip="View"><span class="material-icons" >visibility</span></a><a href="'.$GLOBALS['embedOutputPathLink'].$file.'" download class="tooltip top" data-tooltip="Download"><span class="material-icons">cloud_download</span></a><a href="htmlReport.php?embeddedFile='.$file.'" class=" tooltip top" data-tooltip="Generate HTML Report" onclick="return showLoading()"><span class="material-icons">html</span></a><a href="embed.php?rcf='.$file.'&rct=embed" class=" tooltip top" data-tooltip="Delete"><span class="material-icons">delete</span></a></div></div>';
        }
    }

    if($thelist != ''){
        ?><div class="card" id="card-output-files">
        <div class="card-title">Output files <a href="embed.php?rc=all&rct=embed" class=" tooltip top" data-tooltip="Delete all" onclick="return showLoading()"><span class="material-icons">delete</span></a><a href="downloadAll.php?t=json" target="_Blank" class="tooltip top" data-tooltip="Download All"><span class="material-icons">cloud_download</span></a></div>
        <div class="card-body"><?php echo $thelist; ?></div>
        </div><?php
    }
}
?>