pipeline {
    agent any // 在任何可用的 Jenkins 节点上运行

    environment {
        // 定义需要的环境变量
        DOCKER_IMAGE_NAME = 'my-golang-web-app' // 镜像名称
        DOCKERFILE = 'Dockerfile' // Dockerfile 文件名
        GITHUB_REPO = "https://github.com/Kassaking7/Mota.git"
        EC2_INSTANCE_IP = 'ec2-3-14-250-133.us-east-2.compute.amazonaws.com' // EC2 实例的公共 IP 地址
        SSH_CREDENTIALS_ID = '3b1469a0-15f6-4804-8e0b-b0f72c007d4e'
    }

    stages {

        stage('Cloning to EC2') {
            steps {
                sshagent(credentials: [SSH_CREDENTIALS_ID]) {
                    sh "ssh -o StrictHostKeyChecking=no ec2-user@${EC2_INSTANCE_IP} 'git clone ${GITHUB_REPO}'"
                }
            }
        }
        stage('Build Docker Image') {
            steps {
                // 使用 Dockerfile 构建 Docker 镜像
                script {
                    sh "ssh -o StrictHostKeyChecking=no ec2-user@${EC2_INSTANCE_IP} 'docker build -t ${DOCKER_IMAGE_NAME} .'"
                }
            }
        }

        stage('Run Docker Image') {
            steps {
                // 使用 Dockerfile 构建 Docker 镜像
                script {
                    sh "ssh -o StrictHostKeyChecking=no ec2-user@${EC2_INSTANCE_IP} 'docker run -p 8088:8088 ${DOCKER_IMAGE_NAME}'"
                }
            }
        }
    }

}
