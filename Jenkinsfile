pipeline {
    agent any
    stages {
        stage('Clone Repository') {
            steps {
                echo 'Cloning repository...'
                git branch: 'main', url: 'https://github.com/ErikaRguez/API-Rest.git'
            }
        }
        stage('Test') {
            steps {
                echo 'Running tests...'
                sh 'go test ./...'
            }
        }
        stage('Deploy to Kubernetes') {
            steps {
                echo 'Deploying to kubernetes...'
                sh 'kubectl apply -f k8s/deployment.yaml'
                sh 'kubectl apply -f k8s/service.yaml'
                sh 'kubectl apply -f k8s/ingress.yaml'
            }
        }
    }
}