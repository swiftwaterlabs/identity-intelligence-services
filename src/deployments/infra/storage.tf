resource "aws_s3_bucket" "signal_data" {
  bucket = local.signal_bucket_name

}

resource "aws_s3_bucket_acl" "signal_data" {
  bucket = aws_s3_bucket.signal_data.id
  acl    = "private"
}