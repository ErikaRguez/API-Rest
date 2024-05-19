pipeline {
    agent any

    environment {
        DOCKER_IMAGE = "ivonnerodriguez0818/api-rest:latest"
    }

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
                echo 'Updating Kubernetes manifests...'
                script {
                    sh """
                    sed -i 's|image: .*|image: ${DOCKER_IMAGE}|' kubernetes/deployment.yaml
                    """
                }
                echo 'Deploying to Kubernetes...'
                sh 'kubectl apply -f kubernetes/deployment.yaml'
                sh 'kubectl apply -f kubernetes/service.yaml'
                sh 'kubectl apply -f kubernetes/ingress.yaml'
            }
        }
    }
}
