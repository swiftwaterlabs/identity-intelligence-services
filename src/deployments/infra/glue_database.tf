resource "aws_glue_catalog_database" "object_db" {
  name = local.service_name
}

resource "aws_glue_catalog_table" "user" {
  name          = "user"
  database_name = aws_glue_catalog_database.object_db.name

  table_type = "EXTERNAL_TABLE"

  parameters = {
    EXTERNAL              = "TRUE"
  }

  storage_descriptor {
    location      = "s3://${local.signal_bucket_name}/user"
    input_format  = "org.apache.hadoop.mapred.TextInputFormat"
    output_format = "org.apache.hadoop.hive.ql.io.HiveIgnoreKeyTextOutputFormat"

    ser_de_info {
      name                  = "json"
      serialization_library = "org.apache.hive.hcatalog.data.JsonSerDe"

      parameters = {
        "serialization.format" = 1
      }
    }

    columns {
      name = "id"
      type = "string"
    }

    columns {
      name = "directory"
      type = "string"
    }

    columns {
      name = "name"
      type = "string"
    }

    columns {
      name = "location"
      type = "string"
    }

    columns {
      name = "type"
      type = "string"
    }

    columns {
      name = "objecttype"
      type = "string"
    }

    columns {
      name = "upn"
      type = "string"
    }

    columns {
      name = "givenname"
      type = "string"
    }

    columns {
      name = "surname"
      type = "string"
    }

    columns {
      name = "email"
      type = "string"
    }

    columns {
      name = "manager"
      type = "string"
    }

    columns {
      name = "company"
      type = "string"
    }

    columns {
      name = "department"
      type = "string"
    }

    columns {
      name = "title"
      type = "string"
    }
  }
}

esource "aws_glue_catalog_table" "group" {
  name          = "group"
  database_name = aws_glue_catalog_database.object_db.name

  table_type = "EXTERNAL_TABLE"

  parameters = {
    EXTERNAL              = "TRUE"
  }

  storage_descriptor {
    location      = "s3://${local.signal_bucket_name}/group"
    input_format  = "org.apache.hadoop.mapred.TextInputFormat"
    output_format = "org.apache.hadoop.hive.ql.io.HiveIgnoreKeyTextOutputFormat"

    ser_de_info {
      name                  = "json"
      serialization_library = "org.apache.hive.hcatalog.data.JsonSerDe"

      parameters = {
        "serialization.format" = 1
      }
    }

    columns {
      name = "id"
      type = "string"
    }

    columns {
      name = "directory"
      type = "string"
    }

    columns {
      name = "name"
      type = "string"
    }

    columns {
      name = "location"
      type = "string"
    }

    columns {
      name = "type"
      type = "string"
    }

    columns {
      name = "objecttype"
      type = "string"
    }
  }
}