pipeline {
    agent any
    stages {
        stage("Build Docker Image") {
            steps {
                script {
                    def imageName = 'my-golang-web-app'
                    docker.build(imageName, '-f Dockerfile .')
                }
            }
        }
        stage("Run Docker Container") {
            steps {
                script {
                    def containerName = 'my-golang-web-app-container'
                    docker.image('my-golang-web-app').run('--name ' + containerName + ' -p 8088:8088')
                }
            }
        }
    }
}