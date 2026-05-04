# French Business Search Module For Go

[![Go Version](https://img.shields.io/badge/Go-1.23%2B-blue)](https://go.dev/)
[![Test Status](https://github.com/alex-cos/entapi/actions/workflows/test.yml/badge.svg)](https://github.com/alex-cos/entapi/actions/workflows/test.yml)
[![Lint Status](https://github.com/alex-cos/entapi/actions/workflows/lint.yml/badge.svg)](https://github.com/alex-cos/entapi/actions/workflows/lint.yml)
[![License](https://img.shields.io/badge/License-MIT-green)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/alex-cos/entapi)](https://goreportcard.com/report/github.com/alex-cos/entapi)

The Business Search API allows everyone to find a company, an association, or a public service in France. It offers a large number of search criteria, in particular the name, the address and the managers or elected officials.

Please consult the reference API documentation at:
[https://recherche-entreprises.api.gouv.fr/docs](https://recherche-entreprises.api.gouv.fr/docs/#section/Bienvenue-sur-la-documentation-interactive-d'API-Recherche-d'entreprises-!)

## Data accessible in the API

Since the API is completely open access, it has limitations. Thus, the following are not accessible in the API:

- the predecessors and successors of an establishment
- non-disseminable companies
- companies that have been refused registration with the RCS

**Warning**: this API does not allow access to the complete data of the Sirene database, but only to search for a company, by its name or address.

## Request Limit

The server accepts a maximum of 7 requests per second. Beyond this, a 429 code is returned indicating that the call volume has been exceeded.

## Install

With Go installed, you can install with command line interface:

```bash
  go get github.com/alex-cos/entapi
```
