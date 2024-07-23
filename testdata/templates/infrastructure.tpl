{{- define "infra_1" }}
{
    {{ template "id1" }},
    "aliases": [],
    "name": "my-big-query",
    "type": "Database",
    "owner": {
      {{ template "id1" }}
    },
    "locked": true,
    "ownerLocked": false,
    "data": {
      "name": "my-big-query",
      "engine": "BigQuery",
      "replica": false,
      "endpoint": "https://google.com"
    }
}
{{end}}
{{- define "infra_2" }}
{
    {{ template "id2" }},
    "aliases": [
      "arn:aws:ec2:ca-central-1:XXXXXXXXXX:vpc/vpc-XXXXXXXXXX"
    ],
    "name": "vpc-XXXXXXXXXX",
    "type": "Network",
    "owner": null,
    "locked": false,
    "ownerLocked": false,
    "data": {
      "name": "vpc-XXXXXXXXXX",
      "zone": "ca-central-1",
      "subnets": [],
      "ipv4_cidr": "172.31.0.0/16",
      "is_default": true,
      "external_id": "arn:aws:ec2:ca-central-1:XXXXXXXXXX:vpc/vpc-XXXXXXXXXX",
      "nat_gateway": null,
      "internet_gateway": null,
      "dns_support_enabled": true,
      "dns_hostnames_enabled": true
    }
}
{{end}}
{{- define "infra_3" }}
{
    {{ template "id3" }},
    "aliases": [
      "arn:aws:elasticache:us-east-1:XXXXXXXXXX:cluster:production-demo"
    ],
    "name": "production-demo",
    "type": "Cache",
    "owner": {
      "teamAlias": "platform",
      {{ template "id1" }}
    },
    "ownerLocked": true,
    "data": {
      "name": "production-demo",
      "zone": "us-east-1a",
      "engine": "redis",
      "endpoint": "",
      "external_id": "arn:aws:elasticache:us-east-1:XXXXXXXXXX:cluster:production-demo",
      "instance_type": "cache.t3.micro",
      "engine_version": "6.2.6",
      "maintenance_window": "thu:05:30-thu:06:30"
    }
}
{{end}}
{{- define "infra_schema_1" }}
{
"type": "Database",
"schema": {
  "type": "object",
  "additionalProperties": false,
  "required": [
    "name"
  ],
  "properties": {
    "external_id": {
      "type": "string"
    },
    "name": {
      "type": "string"
    },
    "dns_address": {
      "type": [
        "string",
        "null"
      ]
    },
    "engine": {
      "type": "string"
    },
    "engine_version": {
      "type": "string"
    },
    "endpoint": {
      "type": [
        "string",
        "null"
      ]
    },
    "multi_zone_availability_enabled": {
      "type": [
        "bool",
        "null"
      ]
    },
    "publicly_accessible": {
      "type": [
        "bool",
        "null"
      ]
    },
    "port": {
      "type": [
        "number",
        "null"
      ]
    },
    "vpc": {
      "type": "string"
    },
    "zone": {
      "type": "string"
    },
    "instance_type": {
      "type": "string"
    },
    "storage_iops": {
      "type": "object",
      "required": [
        "value",
        "unit"
      ],
      "properties": {
        "value": {
          "type": [
            "number",
            "null"
          ]
        },
        "unit": {
          "type": "string"
        }
      }
    },
    "storage_type": {
      "type": "string"
    },
    "storage_size": {
      "type": "object",
      "required": [
        "value",
        "unit"
      ],
      "properties": {
        "value": {
          "type": [
            "number",
            "null"
          ]
        },
        "unit": {
          "type": "string"
        }
      }
    },
    "storage_encrypted": {
      "type": [
        "bool",
        "null"
      ]
    },
    "replica": {
      "type": [
        "bool"
      ]
    },
    "replica_source": {
      "type": [
        "string",
        "null"
      ]
    },
    "deletion_protection": {
      "type": [
        "bool",
        "null"
      ]
    },
    "maintenance_window": {
      "type": "string"
    },
    "table_id": {
      "type": "string"
    },
    "creation_date": {
      "type": [
        "date",
        "null"
      ]
    },
    "table_size": {
      "type": "object",
      "required": [
        "value",
        "unit"
      ],
      "properties": {
        "value": {
          "type": [
            "number",
            "null"
          ]
        },
        "unit": {
          "type": "string"
        }
      }
    },
    "item_count": {
      "type": "number"
    },
    "read_capacity_units": {
      "type": "number"
    },
    "write_capacity_units": {
      "type": "number"
    },
    "billing_mode": {
      "type": [
        "string",
        "null"
      ]
    }
  }
}
}
{{end}}
{{- define "infra_schema_2" }}
{
"type": "Compute",
"schema": {
  "type": "object",
  "additionalProperties": false,
  "required": [
    "name"
  ],
  "properties": {
    "external_id": {
      "type": "string"
    },
    "name": {
      "type": "string"
    },
    "instance_id": {
      "type": "string"
    },
    "image_id": {
      "type": "string"
    },
    "instance_type": {
      "type": "string"
    },
    "zone": {
      "type": "string"
    },
    "vpc": {
      "type": "string"
    },
    "launch_time": {
      "type": [
        "date",
        "null"
      ]
    },
    "platform_details": {
      "type": "string"
    },
    "public_ip_address": {
      "type": "string"
    },
    "ipv_6_address": {
      "type": [
        "string",
        "null"
      ]
    },
    "volume_list": {
      "type": "array",
      "items": {
        "type": "object",
        "properties": {
          "type": {
            "type": "string"
          },
          "storage_size": {
            "type": "object",
            "required": [
              "value",
              "unit"
            ],
            "properties": {
              "value": {
                "type": [
                  "number",
                  "null"
                ]
              },
              "unit": {
                "type": "string"
              }
            }
          },
          "zone": {
            "type": "string"
          },
          "storage_encrypted": {
            "type": "bool"
          },
          "iops": {
            "type": "number"
          }
        }
      }
    }
  }
}
}
{{end}}
{{- define "infra_schema_3" }}
{
"type": "Queue",
"schema": {
  "type": "object",
  "additionalProperties": false,
  "required": [
    "name"
  ],
  "properties": {
    "external_id": {
      "type": "string"
    },
    "name": {
      "type": "string"
    },
    "zone": {
      "type": "string"
    },
    "queue_type": {
      "type": "string"
    },
    "encryption_enabled": {
      "type": "bool"
    },
    "creation_date": {
      "type": [
        "date",
        "null"
      ]
    },
    "message_retention_period": {
      "type": [
        "object",
        "null"
      ],
      "required": [
        "value",
        "unit"
      ],
      "properties": {
        "value": {
          "type": [
            "number",
            "null"
          ]
        },
        "unit": {
          "type": "string"
        }
      }
    },
    "visibility_timeout": {
      "type": [
        "object",
        "null"
      ],
      "required": [
        "value",
        "unit"
      ],
      "properties": {
        "value": {
          "type": [
            "number",
            "null"
          ]
        },
        "unit": {
          "type": "string"
        }
      }
    }
  }
}
}
{{end}}
