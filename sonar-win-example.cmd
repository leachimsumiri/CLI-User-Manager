# docker run --rm -e SONAR_HOST_URL="http://localhost:9000" -v "C:/Projekte/sde22-asd-exercise:/usr/src" sonarsource/sonar-scanner-cli
sonar-scanner.bat -D"sonar.projectKey=at.pyr.asd-exercise" -D"sonar.sources=." -D"sonar.host.url=http://localhost:9000" -D"sonar.login=0837e7c2a216184bad8805ca49ae96e67edd8578"
pause