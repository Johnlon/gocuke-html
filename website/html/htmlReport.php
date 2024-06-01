<?php 
include('header.php'); 
?>

<form action="htmlReport.php" method="post" enctype="multipart/form-data" onsubmit="return showLoading()">

    <div class="card">
        <div class="card-title">
                Cucumber JSON Reports
        </div>
        <div class="card-body">
            <div class="upload-area">
                <input type="file" name="filesToUpload[]" id="filesToUpload" multiple class="custom-file-input" title="">
            </div>
            <div class="upload-info" id="uploadInfo">
                <div class="upload-info-content" id="uploadContent">
                </div>
            </div>
        </div>
    </div>

    <div class="card">
        <div class="card-title">
                Settings
        </div>
        <div class="card-body">
            <div class="inputs-section">

                <div class="form-group-small field">
                    <input type="checkbox" name="merge" id="merge" checked>
                    <label for="merge" class="lbl-checkbox">Merge files into one report ?</label><br>
                </div>

                <div class="form-group-small field">
                    <input type="checkbox" name="ignore" id="ignore" checked>
                    <label for="ignore" class="lbl-checkbox">Ignore bad json files ?</label><br>
                </div>

                <div class="form-group-small field">
                    <input type="checkbox" name="showEmbed" id="showEmbed" checked>
                    <label for="showEmbed" class="lbl-checkbox">Show embedded files ?</label><br>
                </div>

                <div class="form-group-full field">
                    <input type="text" name="title" class="form-field" placeholder="Report Title" id="title" />
                    <label for="title" class="form-label">Report Title</label>
                </div>

            </div>
        </div>
    </div>

    <div class="card">
        <div class="card-title">
                Metadata
        </div>
        <div class="card-body">
            <div class="inputs-section">

                <div class="form-group field">
                    <input type="text" name="appVersion" class="form-field" placeholder="App Version" id="appVersion" />
                    <label for="appVersion" class="form-label">App Version</label>
                </div>

                <div class="form-group field">
                    <input type="text" name="testEnvironment" class="form-field" placeholder="Test Environment" id="testEnvironment" />
                    <label for="testEnvironment" class="form-label">Test Environment</label>
                </div>

                <div class="form-group field">
                    <input type="text" name="browser" class="form-field" placeholder="Browser" id="browser" />
                    <label for="browser" class="form-label">Browser</label>
                </div>

                <div class="form-group field">
                    <input type="text" name="platform" class="form-field" placeholder="Platform" id="platform" />
                    <label for="platform" class="form-label">Platform</label>
                </div>

                <div class="form-group field">
                    <input type="text" name="parallel" class="form-field" placeholder="Parallel" id="parallel" />
                    <label for="parallel" class="form-label">Parallel</label>
                </div>

                <div class="form-group field">
                    <input type="text" name="executed" class="form-field" placeholder="Executed" id="executed" />
                    <label for="executed" class="form-label">Executed</label>
                </div>

            </div>
        </div>
    </div>

    <input type="submit" value="Generate HTML Report" name="submit" id="submit">
</form>

<script src="assets/js/htmlReportUpload.js"></script>

<?php
include('htmlReportUpload.php');
include('footer.php');
?>