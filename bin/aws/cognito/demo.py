from dotenv import load_dotenv
import os
import hmac
import hashlib
import base64
import boto3
from botocore.exceptions import BotoCoreError

def register_user(client, user_pool_id, client_id, username, password, email):
    """Register a new user."""
    try:
        # Try to get the user
        user = client.admin_get_user(
            UserPoolId=user_pool_id,
            Username=username
        )
    except client.exceptions.UserNotFoundException:
        # If the user does not exist, sign them up
        response = client.sign_up(
            ClientId=client_id,
            Username=username,
            Password=password,
            UserAttributes=[
                {
                    'Name': 'email',
                    'Value': email
                },
            ]
        )

        client.admin_confirm_sign_up(
            UserPoolId=user_pool_id,
            Username=username
        )

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
    username = ''  # Replace with the actual username
    password = ''  # Replace with the actual password
    email = ''  # Replace with the actual email

    # Create Cognito IDP client
    client = boto3.client('cognito-idp', region_name=region)

    # Register user
    register_user(client, user_pool_id, client_id, username, password, email)

    # Sign in the user
    auth_response = sign_in_user(client, client_id, username, password)

    # Get the JWT
    id_token = auth_response['AuthenticationResult']['IdToken']
    access_token = auth_response['AuthenticationResult']['AccessToken']
    refresh_token = auth_response['AuthenticationResult']['RefreshToken']

    print('ID Token:', id_token)
    print('Access Token:', access_token)
    print('Refresh Token:', refresh_token)

if __name__ == "__main__":
    main()