from dotenv import load_dotenv
import os
import boto3
from botocore.exceptions import BotoCoreError

def sign_in_user(client, client_id, username, password):
    """Sign in the user and return the authentication response."""
    auth_response = client.initiate_auth(
        ClientId=client_id,
        AuthFlow='USER_PASSWORD_AUTH',
        AuthParameters={
            'USERNAME': username,
            'PASSWORD': password
        }
    )
    return auth_response

def main():
    """Main function."""
    # Load environment variables from .env file
    load_dotenv()

    # Read environment variables
    region = os.getenv('AWS_REGION')
    user_pool_id = os.getenv('AWS_COGNITO_USER_POOL_ID')
    client_id = os.getenv('AWS_COGNITO_APP_CLIENT_ID')
    username = os.getenv('AWS_COGNITO_API_USER_EMAIL')
    password = os.getenv('AWS_COGNITO_API_PASSWORD')

    # Create Cognito IDP client
    client = boto3.client('cognito-idp', region_name=region)

    # Sign in the user
    auth_response = sign_in_user(client, client_id, username, password)

    # Get the JWT
    jwt_token = auth_response['AuthenticationResult']['IdToken']

    print('JWT Token:', jwt_token)

if __name__ == "__main__":
    main()
