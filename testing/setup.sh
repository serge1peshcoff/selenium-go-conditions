mkdir testing
echo "Downloading selenium...."
wget "http://selenium-release.storage.googleapis.com/3.5/selenium-server-standalone-3.5.1.jar" -P testing -N

echo "Downloading geckodriver...."
wget https://github.com/mozilla/geckodriver/releases/download/v0.18.0/geckodriver-v0.18.0-linux64.tar.gz -P testing -N

echo "Unzipping geckodriver..."
tar zxf testing/geckodriver-v0.18.0-linux64.tar.gz -C testing/
chmod +x testing/geckodriver

echo "Copying geckodriver to /usr/local/bin..."
sudo cp testing/geckodriver /usr/local/bin