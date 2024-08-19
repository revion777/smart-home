****Smart home****

For Windows Users:
1. Set the environment variable %APP_PATH%:
This variable should point to the root directory of your application.
Command:
cmd
Copy code
set APP_PATH=C:\Path\To\Your\App
Replace C:\Path\To\Your\App with the actual path to your application's root directory.

2. Navigate to %APP_PATH% and configure AWS:
Open a command prompt, navigate to the path stored in %APP_PATH%, and run AWS configuration.
Commands:
cmd
Copy code
cd %APP_PATH%
aws configure
Follow the prompts to enter your AWS Access Key ID, Secret Access Key, region, and output format.

3. Run Serverless Framework (sls) and choose the existing app option:
Command:
cmd
Copy code
sls
When prompted, choose the option to work with an existing app.

4. Run build.bat to build the project:
This script will compile your code and prepare the necessary files for deployment.
Command:
cmd
Copy code
build.bat

5. Deploy the application:
You can either run deploy.bat or use the Serverless Framework (sls deploy) to deploy your application to AWS.
Command:
cmd
Copy code
deploy.bat
Or alternatively:
cmd
Copy code
sls deploy
