// Create log group for repo-service API Lambda.
resource "aws_cloudwatch_log_group" "repo_service_api_lambda_log_group" {
  name              = "/aws/lambda/${aws_lambda_function.repo_service_api_lambda.function_name}"
  retention_in_days = 30
  tags              = local.common_tags
}

// Repo SERVICE API GATEWAY
resource "aws_cloudwatch_log_group" "repo_service_gateway_log_group" {
  name = "${var.environment_name}/${var.service_name}/repo-service-api-gateway"

  retention_in_days = 30
}
