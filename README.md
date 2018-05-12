# Elastic Operator Toolkit

Provide an easy (albeit simplified) access to some common tasks performed while using Elasticsearch.

- [Configuration](#configuration)
- [Commands](#commands)
  - [config](#config)
  - [template](#template)
  - [mapping](#mapping)

## Configuration

Copy `config.yml.dist` as `config.yml`, then edit it to suit your needs.

## Commands

- [config](#config)
- [template](#template)
- [mapping](#mapping)

### config

Use to display current configuration (contents of `config.yml` or the specified configuration loaded via `-c` flag)

```bash
$ elasticop config
{
    "clusters": [
        {
            "name": "cluster_1",
            "description": "Cluster 1",
            "url": "http://localhost:9200"
        },
        {
            "name": "cluster_2",
            "description": "Cluster 2",
            "url": "http://localhost:9200",
            "username": "does_not_exit",
            "password": "not_a_password"
        }
    ]
}
```

### template

TBD

### mapping

TBD
