# Angular Project Builder
A CLI tool which builds Angular projects to your specification. Perfect for organisations.
```powershell
Usage: angular-gen -d <targetDir> -n <appName> -l <libraries>

Example: angular-gen -d C:\Users\Kai\Desktop -n MyApp -l ng-bootstrap,ngx-charts
```
## How to use
The latest build is included in the **dist** folder for those who wish to run the application as-is.
If you plan for widespread usage of this application, you can add the path to the executable file to your environment variables to avoid step 2.
1. Open a command prompt.
2. Navigate to the **dist** folder.
3. Execute **angular-gen -help** for an in-depth description on flags.

## Developer notes
The **types.go** file includes a variety of structs which are mapped to the structure of the default **angular.json**.
This makes it easier to set default configurations for your Angular projects.

```go
addLibrary(projectDir string, libName string)
```
The **addLibrary** function can be used to add additional third-party libraries via the **ng add** command.
If your application requires scripts or styles from the library, add the relative path to the respective array in **updater.go**.

## Contributions
Simply fork the project and create a pull request. Include a thorough description of your changes and it will be reviewed.