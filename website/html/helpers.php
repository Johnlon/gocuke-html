<?php

$msgs = [];

function UploadFile($filename, $fileSize, $fileType, $fileTmpName, $targetFile, $anyFile) {
    $uploadOk = 1;
    
    // Check if file already exists
    if (file_exists($targetFile)) {
        array_push($GLOBALS['msgs'], 'File \''.$filename.'\' already exists.');
        $uploadOk = 0;
    }
    
    // Check file size
    if ($fileSize > 50000000) {
        array_push($GLOBALS['msgs'], 'File \''.$filename.'\' is too large.');
        $uploadOk = 0;
    }
    
    // Allow certain file formats
    if(!$anyFile && $fileType != 'application/json' ) {
        array_push($GLOBALS['msgs'], 'File \''.$filename.'\' have a wrong type. Only JSON files are allowed.');
        $uploadOk = 0;
    }
    
    // Check if $uploadOk is set to 1, upload the file
    if ($uploadOk == 1) {
        if (!move_uploaded_file($fileTmpName, $targetFile)) {
            array_push($GLOBALS['msgs'], 'Error uploading the file \''.$filename.'\'.');
        }
    }

}

function RemoveCache($folder){
    $output = '';
    if ($handle = opendir($folder)) {
        while (false !== ($file = readdir($handle))) {
            if ($file != "." && $file != ".." && $file != ".gitkeep") {
                $output .= shell_exec('cd '.$folder.' && rm '.$file.' 2>&1');
            }
        }
        closedir($handle);
    }
    return $output;
}

function RemoveCacheFile($folder, $filename){
    if ($handle = opendir($folder)) {
        while (false !== ($file = readdir($handle))) {
            if ($file != "." && $file != ".." && $file != ".gitkeep" && $file == $filename) {
                $output = shell_exec('cd '.$folder.' && rm '.$file.' 2>&1');
                if ($output != '') {
                    ?><script>dismissible.error("error removing file '<?php echo $file ?>' from directory.");</script><?php
                } else {
                    ?><script>dismissible.success("File '<?php echo htmlspecialchars($_GET["rcf"]) ?>' successfully removed!");</script><?php
                }
            }
        }
        closedir($handle);
    }
}

?>