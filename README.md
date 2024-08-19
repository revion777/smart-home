# SMART-HOME

## For Windows Users:

1. **Set the environment variable `%APP_PATH%`:**
   - This variable should point to the root directory of the application.
   - **Command:**
     ```cmd
     set APP_PATH=C:\Path\To\Your\App
     ```
     Replace `C:\Path\To\Your\App` with the actual path to the application's root directory.

2. **Navigate to `%APP_PATH%` and configure AWS:**
   - Open a command prompt, navigate to the path stored in `%APP_PATH%`, and run AWS configuration.
   - **Commands:**
     ```cmd
     cd %APP_PATH%
     aws configure
     ```
   - Follow the prompts to enter your AWS Access Key ID, Secret Access Key, region, and output format.

3. **Run `tests.bat` to test the project:**
    - This script will test a code before the deployment.
    - **Command:**
      ```cmd
      tests.bat
      ```

4. **Run Serverless Framework (`sls`) and choose the existing app option:**
   - **Command:**
     ```cmd
     sls
     ```
   - When prompted, choose the option to work with an existing app.

5. **Run `build.bat` to build the project:**
   - This script will compile your code and prepare the necessary files for deployment.
   - **Command:**
     ```cmd
     build.bat
     ```

6. **Deploy the application:**
   - run `deploy.bat` or use the Serverless Framework (`sls deploy`) to deploy the application to AWS.
   - **Command:**
     ```cmd
     deploy.bat
     ```
   - Or alternatively:
     ```cmd
     sls deploy
     ```
