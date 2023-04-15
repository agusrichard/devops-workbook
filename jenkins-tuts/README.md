# Jenkins Tutorial

<br />

## List of Contents:

### 1. [Learn Jenkins! Complete Jenkins Course - Zero to Hero](#content-1)

<br />

---

## Contents:
## [Learn Jenkins! Complete Jenkins Course - Zero to Hero](https://www.youtube.com/watch?v=6YZvp2GwT0A) <span id="content-1"></span>


### Jenkins Infrasturcture
- Master server:
  - Control pipelines
  - Schedule builds
- Agents/minions:
  - Perform the build


### Agent Types
- Pernament agent:
  - Dedicated servers for running jobs
  - Standalone server like Linux and Windows
- Cloud agent
  - Ephemeral/dynamic agents spun up on demand

### Build Types
- Freestyle build: Pretty much like scripting
- Pipelines: Composed by stages or steps

### Setting up Jenkins using Docker
- Dockerfile
```dockerfile
FROM jenkins/jenkins:2.332.3-jdk11
USER root
RUN apt-get update && apt-get install -y lsb-release
RUN curl -fsSLo /usr/share/keyrings/docker-archive-keyring.asc \
  https://download.docker.com/linux/debian/gpg
RUN echo "deb [arch=$(dpkg --print-architecture) \
  signed-by=/usr/share/keyrings/docker-archive-keyring.asc] \
  https://download.docker.com/linux/debian \
  $(lsb_release -cs) stable" > /etc/apt/sources.list.d/docker.list
RUN apt-get update && apt-get install -y docker-ce-cli
USER jenkins
RUN jenkins-plugin-cli --plugins "blueocean:1.25.3 docker-workflow:1.28"
```
- Build the image:
```shell
docker build -t myjenkins-blueocean:2.332.3-1 .
```
- Create network:
```shell
docker network create jenkins
```
- Run the container:
```shell
docker run --name jenkins-blueocean --restart=on-failure --detach \
  --network jenkins --env DOCKER_HOST=tcp://docker:2376 \
  --env DOCKER_CERT_PATH=/certs/client --env DOCKER_TLS_VERIFY=1 \
  --publish 8080:8080 --publish 50000:50000 \
  --volume jenkins-data:/var/jenkins_home \
  --volume jenkins-docker-certs:/certs/client:ro \
  myjenkins-blueocean:2.332.3-1
```
- Get the password:
```shell
docker exec jenkins-blueocean cat /var/jenkins_home/secrets/initialAdminPassword
```


**[⬆ back to top](#list-of-contents)**

<br />

---

## References:
- https://www.youtube.com/watch?v=6YZvp2GwT0A
- https://www.jenkins.io/doc/pipeline/tour/getting-started/
- https://devopscube.com/jenkins-2-tutorials-getting-started-guide/
- https://github.com/devopsjourney1/jenkins-101
