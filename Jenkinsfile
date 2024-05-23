pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                echo 'Cloning repository...'
                git branch: 'main', url: 'https://github.com/ErikaRguez/API-Rest.git'
                sh 'go build -o appi'
            }
        }
        stage('Test') {
            steps {
                echo 'Running tests...'
                sh 'go test ./...'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Building Docker image...'
                sh 'docker build -t myapp .'
                echo 'Starting Docker container...'
                sh 'docker-compose up -d'
            }
        }
    }
}
