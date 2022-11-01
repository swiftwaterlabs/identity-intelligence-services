locals {
  service_name = "identity_intelligence_${var.environment}"
  signal_bucket_name = replace("${local.service_name}data","_","")
}