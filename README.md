# Terraform Provider for Google Cloud Translate API

This is a *very* rudimentary Terraform Provider for Google
Cloud Translate API.

It pretty much does one thing:

- Take in text.
- Translate to target language.

Since it's a Cloud Translate API job, it pretty much
just runs once.

## Caveats

- You cannot "delete" a translation job.

- You cannot "update" a translation job, just run a new
  one with new output.

- There are no acceptance tests right now, so use at your
  own risk.

## Build the Plugin

Run:

```shell
make plugin
```

This will build the plugin and load it into your home directory.