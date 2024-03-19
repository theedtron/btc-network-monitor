## btc-network-monitor

<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
  </ol>
</details>


## Getting Started

### Prerequisites

Before this project can be run locally you will need to install [Go](https://golang.org/doc/install)

### Installation

To utilize the project, the following needs to be done:
1. Clone this repository
2. Install the dependencies, using the following command:
```
go mod tidy
```

## Usage

1. To run migration, use the following command:
```
make up
```

2. To run the project locally, use the following command:
```
make run
```

3. To check the health status, use Postman to make a GET request to the following URL:
```
http://localhost:{port}/api/v1/healthcheck
```
