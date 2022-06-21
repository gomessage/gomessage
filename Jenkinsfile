Jenkinsfile (Declarative Pipeline)
pipeline {
    agent { docker 'golang:1.17.11-alpine' }
    stages {
        stage('build') {
            steps {
                sh 'go version'
                sh 'go env'
            }
        }
    }
}
