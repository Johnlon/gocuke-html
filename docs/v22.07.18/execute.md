# Execute the binary

To use the binary, first you have to clone the project.

```shell
git clone git@gitlab.com:rodrigoodhin/gocure.git
```

After clone the project, you have to build the binary.

```shell
go build gocure.go
```

After build you can execute the binary.

```shell
./gocure -j /path/to/json/file.json
```

### Options

| Tag | Type | Description | Default |
| --- | --- | --- | --- |
| -h, --html | String | HTML option | |
| -e, --embedded | String | Embedded option | |

Common tags:

| Tag | Type | Description | Default |
| --- | --- | --- | --- |
| -j | String | Path of the Cucumber JSON report file | |

Tags for html option:

| Tag | Type | Description | Default |
| --- | --- | --- | --- |
| -f | String | Path of a directory with more than one Cucumber JSON report file<br>This option will be ignored if option "-j" were used | |
| -m | Bool | Merge all Cucumber JSON report files into only one HTML report<br>This option is only used if option "-f" were used | false |
| -i | Bool | Ignore bad Cucumber JSON report files<br>This option is only used if option "-f" were used | false |
| -o | String | Path of the HTML report output folder | |
| -t | String | Title of your report | GO Cucumber HTML Report |
| -s | Bool | Show embedded files | false |
| -AppVersion | String | Metadata - App Version | - |
| -TestEnvironment | String | Metadata - Test Environment | - |
| -Browser | String | Metadata - Browser | - |
| -Platform | String | Metadata - Platform | - |
| -Parallel | String | Metadata - Parallel | - |
| -Executed | String | Metadata - Executed | - |

Tags for embedded option:

| Tag | Type | Description | Default |
| --- | --- | --- | --- |
| -u | String | Output json path | |
| -l | String | Embedded file path | |
| -a | Int | Feature index | -1 |
| -c | Int | Scenario index | -1 |
| -p | Int | Step index | -1 |
