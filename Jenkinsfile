pipeline {
    agent any
    stages {
        stage("api_auto_test") {
            steps{
              bat 'go run main.go'
            }
        }
    }
}