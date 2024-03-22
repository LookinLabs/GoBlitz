# Example Go Web service

Example Go Web Service that serves static files and exposes a simple API path.

## Build and Deploy

To build and deploy the image to Amazon Elastic Container Registry (ECR), follow these steps:

1. Install the AWS Command Line Interface (CLI) if you haven't already.

2. Authenticate the Docker CLI to your ECR registry by running the following command:

    ```shell
    aws ecr get-login-password --region <your-region> | docker login --username AWS --password-stdin <your-account-id>.dkr.ecr.<your-region>.amazonaws.com
    ```

    Replace `<your-region>` with the AWS region where your ECR registry is located, and `<your-account-id>` with your AWS account ID.

3. Build the Docker image by running the following command:

    ```shell
    docker build -t <your-image-name> .
    
    # EXAMPLE
    docker build -t go-web-service:v1.0 .
    ```

    Replace `<your-image-name>` with the desired name for your Docker image.

4. Tag the Docker image with the ECR repository URI by running the following command:

    ```shell
    docker tag <your-image-name> <your-account-id>.dkr.ecr.<your-region>.amazonaws.com/<your-repository-name>:<your-tag>

    # EXAMPLE
    docker tag go-web-service:v1.0 1234567890.dkr.ecr.eu-west-1.amazonaws.com/aws-app-runner:v1.0
    ```

    Replace `<your-repository-name>` with the name of your ECR repository, and `<your-tag>` with the desired tag for your Docker image.

5. Push the Docker image to ECR by running the following command:

    ```shell
    docker push 1234567890.dkr.ecr.eu-west-1.amazonaws.com/aws-app-runner:v1.0
    ```

    Replace `<your-repository-name>` with the name of your ECR repository, and `<your-tag>` with the tag you used in the previous step.

6. Your Docker image is now available in ECR and can be used for deployment.
