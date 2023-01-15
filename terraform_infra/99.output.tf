output "api_gateway_execution_arn" {
  value = aws_api_gateway_rest_api.example.execution_arn
}
output "api_url" {
  value = aws_api_gateway_deployment.api_deployment.invoke_url
}