data "archive_file" "signal_persistance_lambda_zip" {
  type        = "zip"
  source_file = "../../cmd/lambda-directoryobjectreceiver/main"
  output_path = "main.zip"
}

resource "aws_lambda_function" "signal_persistance" {
  function_name = "${local.service_name}_signal_persistance"

  role = aws_iam_role.lambda_exec.arn

  filename          = data.archive_file.signal_persistance_lambda_zip.output_path
  handler           = "main"
  source_code_hash  = filebase64sha256(data.archive_file.signal_persistance_lambda_zip.output_path)
  runtime           = "go1.x"

  environment {
    variables = {
      aws_region = var.aws_region
      directoryobject_blobstore = aws_s3_bucket.signal_data.id
    }
  }
  
}

resource "aws_cloudwatch_log_group" "signal_persistance" {
  name = "/aws/lambda/${aws_lambda_function.signal_persistance.function_name}"

  retention_in_days = 30
}

resource "aws_lambda_event_source_mapping" "signal_persistance" {
  event_source_arn = aws_sqs_queue.signal_ingestion.arn
  function_name    = aws_lambda_function.signal_persistance.arn
}