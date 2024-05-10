// Terraform Code to configure Cognito User Pool and Client
// This code creates a Cognito User Pool, a Cognito User Pool Client, and a Cognito User Pool Domain

// Variables
variable "aws_region" {
  type = string
  description = "The region in which the resources will be created"
  default = "us-east-1"
}

// Providers
provider "aws" {
  region     = var.aws_region
}

// Resources
resource "aws_cognito_user_pool" "user_pool" {
  name = "GoBlitzUserPool"

  username_attributes = ["email"]
  auto_verified_attributes = ["email"]
  password_policy {
    minimum_length = 6
  }

  verification_message_template {
    default_email_option = "CONFIRM_WITH_CODE"
    email_subject = "Account Confirmation"
    email_message = "Your confirmation code is {####}"
  }

  schema {
    attribute_data_type      = "String"
    developer_only_attribute = false
    mutable                  = true
    name                     = "email"
    required                 = true

    string_attribute_constraints {
      min_length = 1
      max_length = 256
    }
  }
}

resource "aws_cognito_user_pool_client" "client" {
  name = "GoBlitzClient"

  user_pool_id = aws_cognito_user_pool.user_pool.id
  generate_secret = false
  refresh_token_validity = 90
  prevent_user_existence_errors = "ENABLED"

  allowed_oauth_flows = ["code", "implicit"]
  allowed_oauth_scopes = ["email", "openid"]
  callback_urls = ["http://localhost:8000/"]  // Replace with your actual callback URL
  allowed_oauth_flows_user_pool_client = true
  supported_identity_providers = ["COGNITO"]
}

resource "aws_cognito_user_pool_domain" "cognito-domain" {
  domain       = "goblitz"
  user_pool_id = "${aws_cognito_user_pool.user_pool.id}"
}

output "user_pool_id" {
  value = aws_cognito_user_pool.user_pool.id
}

output "client_id" {
  value = aws_cognito_user_pool_client.client.id
}
