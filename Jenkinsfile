pipeline {
    agent any // 在任何可用的 Jenkins 节点上运行

    environment {
        // 定义需要的环境变量
        DOCKER_IMAGE_NAME = 'my-golang-web-app' // 镜像名称
        DOCKERFILE = 'Dockerfile' // Dockerfile 文件名
        EC2_INSTANCE_IP = 'ec2-3-14-250-133.us-east-2.compute.amazonaws.com' // EC2 实例的公共 IP 地址
        SSH_CREDENTIALS_ID = '3b1469a0-15f6-4804-8e0b-b0f72c007d4e'
    }

    stages {
        stage('Checkout') {
            steps {
                // 从代码仓库中检出代码
                checkout scm
            }
        }

        stage('Build Docker Image') {
            steps {
                // 使用 Dockerfile 构建 Docker 镜像
                script {
                    sh "docker build -t ${DOCKER_IMAGE_NAME} -f ${DOCKERFILE} ."
                }
            }
        }
        stage('Test EC2 Connection') {
            steps {
                sshagent(credentials: [SSH_CREDENTIALS_ID]) {
                    sh '''
                        ssh -o StrictHostKeyChecking=no ec2-user@${EC2_INSTANCE_IP}
                        touch test.txt
                        exit
                        '''
                }
            }
        }
        stage('Transfer Docker Image to EC2') {
            steps {
                sshagent(credentials: [SSH_CREDENTIALS_ID]) {
                    sh "scp ${DOCKER_IMAGE_NAME} ec2-user@${EC2_INSTANCE_IP}:/home/ec2-user/"

                }
            }
        }

        stage('Run Docker Image on EC2') {
            steps {
                sshagent(credentials: [SSH_CREDENTIALS_ID]) {
                    sh "ssh ec2-user@${EC2_INSTANCE_IP} 'docker load -i /home/ec2-user/${DOCKER_IMAGE_NAME} && docker run -d -p 8088:8088 ${DOCKER_IMAGE_NAME}'"
                }
            }
        }
    }

    post {
        always {
            // 清理：在 Pipeline 完成后删除构建时生成的 Docker 镜像（可选）
            script {
                sh "docker rmi ${DOCKER_IMAGE_NAME}"
            }
        }
    }
}
