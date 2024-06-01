<?php 
include 'header.php'; 
?>

<form action="embed.php" method="post" enctype="multipart/form-data" onsubmit="return showLoading()">

    <div class="card">
        <div class="card-title">
                Cucumber JSON Report
        </div>
        <div class="card-body">
            <div class="upload-area-embed">
                <input type="file" name="fileToUpload" id="fileToUpload" class="custom-file-input" title="">
            </div>
            <div class="upload-info" id="uploadInfoEmbed">
                <div class="upload-info-content" id="uploadContentEmbed">
                </div>
            </div>
        </div>
    </div>

    <div class="card">
        <div class="card-title">
                Files to Embed
        </div>
        <div class="card-body">
            <div class="upload-area-embed-multiple">
                <input type="file" name="filesToUpload[]" id="filesToUpload" multiple class="custom-file-input" title="">
            </div>
            <div class="upload-info" id="uploadInfoEmbedMultiple">
                <div class="upload-info-content" id="uploadContentEmbedMultiple">
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
                    <input type="number" name="featureIndex" class="form-field" placeholder="Feature Index" id="featureIndex" min="0"/>
                    <label for="featureIndex" class="form-label">Feature Index</label>
                </div>

                <div class="form-group-small field">
                    <input type="number" name="scenarioIndex" class="form-field" placeholder="Scenario Index" id="scenarioIndex" min="0" />
                    <label for="scenarioIndex" class="form-label">Scenario Index</label>
                </div>

                <div class="form-group-small field">
                    <input type="number" name="stepIndex" class="form-field" placeholder="Step Index" id="stepIndex" />
                    <label for="stepIndex" class="form-label">Step Index</label>
                </div>

            </div>
        </div>
    </div>

    <input type="submit" value="Embed Files" name="submit" id="submit">

</form>

<script src="assets/js/embedUpload.js"></script>

<?php
include('embedUpload.php');
include('footer.php');
?>