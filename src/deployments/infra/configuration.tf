resource "aws_secretsmanager_secret" "directories_table" {
  name = "${local.service_name}_directories_table"
}

resource "aws_secretsmanager_secret_version" "directories_table" {
  secret_id     = aws_secretsmanager_secret.directories_table.id
  secret_string = aws_dynamodb_table.directories.name
}