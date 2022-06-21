pipeline {
    agent {
        docker { image 'golang:1.17.11-alpine'}
        }
    environment {
        DISABLE_AUTH = 'true'
        DB_ENGINE    = 'sqlite'
    }
    stages {
        stage('build') {
            steps {
                sh 'go version'
                sh 'go env'
            }
        }
        stage("printenv"){
            sh 'printenv'
        }
    }
}
