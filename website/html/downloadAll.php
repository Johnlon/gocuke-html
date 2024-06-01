<?php
    include('constants.php');

    if (isset($_GET["t"]) && htmlspecialchars($_GET["t"]) == 'html') {
        DownloadAll($GLOBALS['htmlReportOutputPath'], "htmlReports");
    } else if (isset($_GET["t"]) && htmlspecialchars($_GET["t"]) == 'json') {
        DownloadAll($GLOBALS['embedOutputPath'], "jsonReports");
    }   

    function DownloadAll($inputPath, $filePrefix){
    
        $zip = new ZipArchive;
    
        $fileName = $filePrefix.date("_Ymd_his", mktime(11, 14, 54, 8, 12, 2014)).'.zip';
        $filePath = $GLOBALS['reportZipPath'].$fileName;
    
        if ($zip->open($filePath,ZIPARCHIVE::CREATE) == true) {
    
            if ($handle = opendir($inputPath)) {
                $files = array();
                while (false !== ($file = readdir($handle))) {
                    array_push($files, $file);
                }
                closedir($handle);
                rsort($files);

                foreach ($files as &$file) {
                    if ($file != "." && $file != ".." && $file != ".gitkeep") {
                        $zip->addFile($inputPath.$file , $file);
                    }
                }
            }
    
            $zip->close();
    
           header("Content-type: application/zip"); 
           header("Content-Disposition: attachment; filename=".$fileName);
           header("Content-length: " . filesize($filePath));
           header("Pragma: no-cache"); 
           header("Expires: 0"); 
           readfile($filePath);
           unlink($filePath);
    
       }
    
    }