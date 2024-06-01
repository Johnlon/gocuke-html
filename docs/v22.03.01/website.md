# Website

> [!NOTE]
> To use this option, you have to start the docker container.
> 
> For more information about how to use Docker Gocure, go to [Docker](/v22.03.01/docker) section.

After start the Docker container, go to [http://localhost:8087/](http://localhost:8087/) to Embed or to generate HTML reports.

![](/_media/scr_website_01.png)

On the top of the page, you have to chose what do you want to do: Embed Files or hTML report.

![](/_media/scr_website_02.png)

## Embed Files

![](/_media/scr_website_03.png)

For Embed files, follow these instructions:
- On the ```Cucumber JSON Report``` card, select or drag and drop the Cucumber JSON report that you want to embed the files.
![](/_media/scr_website_04.png)

- On the ```Files to Embed``` card, select or drag and drop  the files you want to embed.
![](/_media/scr_website_05.png)

- On the ```Settings``` card, select the index for the Feature, Scenario, and the Step.
![](/_media/scr_website_06.png)

- Click on ```Embed Files``` button.
![](/_media/scr_website_07.png)

After click on **Embed Files** button, the ```Output files``` card will be shown below the Embed files button.
![](/_media/scr_website_08.png)

On this card you can view, download or delete a file.

### Deleting a file

Clicking the trash icon to the right of the output file name will delete only that file.
![](/_media/scr_website_09.png)

Clicking the trash icon to the right of the output files card will delete all output files.
![](/_media/scr_website_10.png)

## HTML Report

![](/_media/scr_website_11.png)

For HTML Report, follow these instructions:
- On the ```Cucumber JSON Report``` card, select or drag and drop the Cucumber JSON report that you want to use to generate a HTML report.
![](/_media/scr_website_12.png)

- On the ```Settings``` card, your have some options to chose:
  - Fill the report title.
  - Choose whether or not to merge files. (If true and you select more than one file to generate HTML report, Gocure will merge all the reports into a single HTML file, if not, an HTML file will be generated for each of the JSON reports)
  - Choose to ignore bad JSON files or not. (If true and you select more than one file to generate HTML report, Gocure will ignore if the JSON input does not have the correct format, if not, an error message will be shown if the JSON input does not have the correct format)
  - Choose to show embedded files or not. (If true, Gocure will show the embedded files, if not, Gocure will ignore these files)
![](/_media/scr_website_13.png)

- On the ```Metadata``` card, select all the Metadata information you want to use to generate an HTML report.
![](/_media/scr_website_14.png)

- Click on ```Generate HTML Report``` button.
![](/_media/scr_website_15.png)

After click on **Generate HTML Report** button, the ```Output files``` card will be shown below the Generate HTML Report button.
![](/_media/scr_website_16.png)

On this card you can view, download or delete a file.

### Deleting a file

Clicking the trash icon to the right of the output file name will delete only that file.
![](/_media/scr_website_17.png)

Clicking the trash icon to the right of the output files card will delete all output files.
![](/_media/scr_website_10.png)