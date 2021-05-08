properties([pipelineTriggers([githubPush()])])
node {
    git url: 'https://github.com/Ehco1996/ehco',
    branch: 'jenkins'
}

pipeline {
    agent any
    tools {
        go 'go1.16'
    }
    environment {
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    stages {
        stage('checkout branch') {
            steps {
                echo 'checkout ...'
                git changelog: false, poll: false, url: 'https://github.com/Ehco1996/ehco'
            }
        }

        stage('Pre Test') {
            steps {
                echo 'Installing dependencies'
                sh 'go version'
                sh 'go env -w GOPROXY=https://goproxy.cn,direct'
                sh 'make ensure'
            }
        }

        stage('Test') {
            steps {
                withEnv(["PATH+GO=${GOPATH}/bin"]) {
                    echo 'Running vetting'
                    sh 'make test'
                }
            }
        }

        stage('Build') {
            steps {
                echo 'Compiling and building'
                sh 'make build'
            }
        }
    }
}
