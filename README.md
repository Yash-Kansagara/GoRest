| [![Quality gate](https://sonarcloud.io/api/project_badges/quality_gate?project=Yash-Kansagara_GoRest)](https://sonarcloud.io/summary/new_code?id=Yash-Kansagara_GoRest)  | [![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=Yash-Kansagara_GoRest&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=Yash-Kansagara_GoRest) [![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=Yash-Kansagara_GoRest&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=Yash-Kansagara_GoRest) [![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=Yash-Kansagara_GoRest&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=Yash-Kansagara_GoRest) [![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=Yash-Kansagara_GoRest&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=Yash-Kansagara_GoRest)  |
| ------------ | ------------ |
|  [![Go Report Card](https://goreportcard.com/badge/github.com/Yash-Kansagara/GoRest)](https://goreportcard.com/report/github.com/Yash-Kansagara/GoRest) |

# REST API starter code for *GO*
- Uses MySQL DB as storage
- http.ServeMux for routing
- rate limiter using [time](https://pkg.go.dev/golang.org/x/time@v0.14.0/rate "time")
- middleware: CORS, field validation,
- no external library used for validation, only reflection
    - suggested: 
		- https://pkg.go.dev/github.com/go-playground/validator/v10
		- https://pkg.go.dev/github.com/Azure/go-autorest/autorest/validation