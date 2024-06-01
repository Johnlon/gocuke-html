# REST API

To use the API, first you have to clone the project.

```shell
git clone git@gitlab.com:rodrigoodhin/gocure.git

```

After clone the project, go to the Gocure folder.

```shell
cd gocure
```

And start the API.

```shell
go run api/api.go
```

The API will start at [http://localhost:7087/](http://localhost:7087/)

![](/_media/scr_api.png)

> [!NOTE]
> Is it also possible to use the API starting the docker container.
> For more information about how to use Docker Gocure, go to [Docker](/v22.07.18/docker) section

## Swagger

With API running, you can go to [http://localhost:7087/swagger/index.html](http://localhost:7087/swagger/index.html) to check the Swagger Documentation for the API.

![](/_media/scr_api_swagger.png)

## Endpoints

To work with your reports files, you will have four endpoints:
- /embed/toFeature
- /embed/toScenario
- /embed/toStep
- /html/generate

### Embed to Feature

**Method**: POST<br>
**Path**: /embed/toFeature<br >
**Body**:
```json
{
  "inputJsonPath": "report.json",
  "outputJsonPath": "newReport.json",
  "files": [
    "video.mp4",
    "text.txt",
    "https://i.ibb.co/LpRkTqf/gocure-small.png"
  ],
  "featureIndex": 0
}
```

### Embed to Scenario

**Method**: POST<br>
**Path**: /embed/toScenario<br >
**Body**:
```json
{
  "inputJsonPath": "report.json",
  "outputJsonPath": "newReport.json",
  "files": [
    "video.mp4",
    "text.txt",
    "https://i.ibb.co/LpRkTqf/gocure-small.png"
  ],
  "featureIndex": 0,
  "scenarioIndex": 0
}
```

### Embed to Step

**Method**: POST<br>
**Path**: /embed/toStep<br >
**Body**:
```json
{
  "inputJsonPath": "report.json",
  "outputJsonPath": "newReport.json",
  "files": [
    "video.mp4",
    "text.txt",
    "https://i.ibb.co/LpRkTqf/gocure-small.png"
  ],
  "featureIndex": 0,
  "scenarioIndex": 0,
  "stepIndex": 2
}
```

#### Embed Options

To embed files, these are the options:

| Name | Type | Description | Default |
| --- | --- | --- | --- |
| inputJsonPath | String | Path of the Cucumber JSON report file | |
| outputJsonPath | String | Output json path | |
| files | []String | Paths of the files to be embedded | |
| featureIndex | Int | Index of the Feature | |
| scenarioIndex | Int | Index of the Scenario | |
| stepIndex | Int | Index of the Step | |

### HTML Generate

**Method**: POST<br>
**Path**: /html/generate<br >
**Body**:
```json
{
  "mergeFiles": true,
  "ignoreBadJsonFiles": true,
  "outputHtmlFolder": "example/",
  "title": "",
  "showEmbeddedFiles": true,
  "inputJsonPath": "report.json",
  "inputFolderPath": "",
  "metadata": {
    "appVersion": "0.8.7",
    "testEnvironment": "development",
    "browser": "Google Chrome",
    "platform": "Linux",
    "parallel": "Scenarios",
    "executed": "Remote"
  }
}
```

#### HTML Options

| Option | Type | Description | Default |
| --- | --- | --- | --- |
| inputJsonPath | String | Path of the Cucumber JSON report file | |
| inputFolderPath | String | Path of a directory with more than one Cucumber JSON report file<br>This option will be ignored if option "InputJsonPath" were used | |
| mergeFiles | Bool | Merge all Cucumber JSON report files into only one HTML report<br>This option is only used if option "InputJsonPath" were used | false |
| ignoreBadJsonFiles | Bool | Ignore bad Cucumber JSON report files<br>This option is only used if option "InputJsonPath" were used | false |
| outputHtmlFolder | String | Path of the HTML report output folder | |
| title | String | Title of your report | GO Cucumber HTML Report |
| showEmbeddedFiles | Bool | Show embedded files | false |
| metadata.AppVersion | String | Metadata - App Version | - |
| metadata.TestEnvironment | String | Metadata - Test Environment | - |
| metadata.Browser | String | Metadata - Browser | - |
| metadata.Platform | String | Metadata - Platform | - |
| metadata.Parallel | String | Metadata - Parallel | - |
| metadata.Executed | String | Metadata - Executed | - |