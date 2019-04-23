def label = "docker-${UUID.randomUUID().toString()}"

podTemplate(label: label, inheritFrom: 'base') {
  node(label) {
    stage('Checkout Repository') {
      container('base') {
        checkout scm
      }
    }

    stage('Login to Dockerhub') {
      withCredentials([usernamePassword(credentialsId: 'DockerHubAccessYardstick', usernameVariable: 'USER', passwordVariable: 'PASS')]) {
        container('base') {
          sh "docker login --username ${USER} --password ${PASS}"
        }
      }
    }

    stage('Build Docker image') {
      container('base') {
        sh "docker build -t yardstick/bivac:${BRANCH_NAME} ."
        sh "docker push yardstick/bivac:${BRANCH_NAME}"
      }
    }
  }
}