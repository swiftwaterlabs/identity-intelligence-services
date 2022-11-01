resource "aws_secretsmanager_secret" "directories_table" {
  name = "${local.service_name}_directories_table"
}

resource "aws_secretsmanager_secret_version" "directories_table" {
  secret_id     = aws_secretsmanager_secret.directories_table.id
  secret_string = aws_dynamodb_table.directories.name
}

resource "aws_secretsmanager_secret" "ingestion_queue" {
  name = "${local.service_name}_ingestion_queue"
}

resource "aws_secretsmanager_secret_version" "ingestion_queue" {
  secret_id     = aws_secretsmanager_secret.ingestion_queue.id
  secret_string = aws_sqs_queue.signal_ingestion.name
}