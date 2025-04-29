output "repo_service_api_lambda_arn" {
  value = aws_lambda_function.repo_service_api_lambda.arn
}

output "repo_service_api_lambda_invoke_arn" {
  value = aws_lambda_function.repo_service_api_lambda.invoke_arn
}

output "repo_service_api_lambda_function_name" {
  value = aws_lambda_function.repo_service_api_lambda.function_name
}
