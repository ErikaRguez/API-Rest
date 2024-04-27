pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                sh 'go build -o myapp'
            }
        }
        stage('Test') {
            steps {
                sh 'go test ./...'
            }
        }
        stage('Deploy') {
            steps {
                sh 'docker build -t myapp .'
                sh 'docker-compose up -d'
            }
        }
    }
}
