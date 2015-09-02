# Golang VN Grok

This is an experiment project of Golang VN team.

## Prerequisite

Install and add these tools to your `$PATH`

- [Godep](https://github.com/tools/godep)
- [RethinkDB](http://rethinkdb.com/docs/install/)

## API Server

### Install dependencies

Clone this repository to `$GOPATH/src/github.com/golang-vietnam/grox` and run

```
make install
```

### Build

```
make build
```

### Start

```
$GOPATH/bin/grox-server
```

## Front-End Server

### Install dependencies

```
npm install
```

### Development

```
npm run dev
```

### Production

```
npm run build
npm start
```
