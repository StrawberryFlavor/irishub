# iriscli record

## Description

Record allows you to submit, query and download your records on chain.

## Usage

```shell
iriscli record [command]
```

## Record Introduction

1. Any users can initiate a record request. It will cost you some tokens. If there’s no record of the data on the existing chains, the request will be completed successfully and the relevant metadata will be recorded onchain. And you will be returned a record ID to confirm your ownership of the data.
2. If any others initiate a record request for the same data, the request will be directly rejected and it will hint that the relevant record data has already existed.
3. Any users can search/download on chain based on the record ID.
4. At present, the maximum amount of stored data at most 1K Bytes. In the future, the dynamic adjustment of parameters will be implemented in conjunction with the governance module.

## Available Commands

| Name                    | Description                                                   |
| ------------------------| --------------------------------------------------------------|
| [query](query.md)       | Query specified record                                        |
| [download](download.md) | Download related data with unique record ID to specified file |
| [submit](submit.md)     | Submit a new record                                           |

## Flags

| Name, shorthand | Default | Description     | Required |
| --------------- | ------- | --------------- | -------- |
| --help, -h      |         | help for record |          |
